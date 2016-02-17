// Copyright 2015 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package store

import (
	"fmt"
	"reflect"
	"sync"
	"time"

	"github.com/luci/luci-go/common/clock"
	"github.com/luci/luci-go/common/tsmon/distribution"
	"github.com/luci/luci-go/common/tsmon/field"
	"github.com/luci/luci-go/common/tsmon/target"
	"github.com/luci/luci-go/common/tsmon/types"
	"golang.org/x/net/context"
)

type inMemoryStore struct {
	defaultTarget types.Target
	data          map[string]*metricData
	lock          sync.RWMutex
}

type cellKey struct {
	fieldValuesHash, targetHash uint64
}

type metricData struct {
	types.MetricInfo

	cells map[cellKey][]*types.CellData
	lock  sync.Mutex
}

func (m *metricData) get(fieldVals []interface{}, t types.Target, resetTime time.Time) (*types.CellData, error) {
	fieldVals, err := field.Canonicalize(m.Fields, fieldVals)
	if err != nil {
		return nil, err
	}

	targetHash := t.Hash()

	key := cellKey{field.Hash(fieldVals), targetHash}
	cells, ok := m.cells[key]
	if ok {
		for _, cell := range cells {
			if reflect.DeepEqual(fieldVals, cell.FieldVals) &&
				reflect.DeepEqual(t, cell.Target) {
				return cell, nil
			}
		}
	}

	cell := &types.CellData{fieldVals, t, resetTime, nil}
	m.cells[key] = append(cells, cell)
	return cell, nil
}

// NewInMemory creates a new metric store that holds metric data in this
// process' memory.
func NewInMemory(defaultTarget types.Target) Store {
	return &inMemoryStore{
		defaultTarget: defaultTarget,
		data:          map[string]*metricData{},
	}
}

// Register does nothing.
func (s *inMemoryStore) Register(m types.Metric) {}

// Unregister removes the metric (along with all its values) from the store.
func (s *inMemoryStore) Unregister(h types.Metric) {
	s.lock.Lock()
	defer s.lock.Unlock()

	delete(s.data, h.Info().Name)
}

func (s *inMemoryStore) getOrCreateData(m types.Metric) *metricData {
	s.lock.RLock()
	d, ok := s.data[m.Info().Name]
	s.lock.RUnlock()
	if ok {
		return d
	}

	s.lock.Lock()
	defer s.lock.Unlock()

	// Check again in case another goroutine got the lock before us.
	if d, ok := s.data[m.Info().Name]; ok {
		return d
	}

	d = &metricData{
		MetricInfo: m.Info(),
		cells:      map[cellKey][]*types.CellData{},
	}

	s.data[m.Info().Name] = d
	return d
}

func (s *inMemoryStore) DefaultTarget() types.Target {
	return s.defaultTarget
}

// Get returns the value for a given metric cell.
func (s *inMemoryStore) Get(ctx context.Context, h types.Metric, resetTime time.Time, fieldVals []interface{}) (value interface{}, err error) {
	if resetTime.IsZero() {
		resetTime = clock.Now(ctx)
	}

	m := s.getOrCreateData(h)
	m.lock.Lock()
	defer m.lock.Unlock()

	c, err := m.get(fieldVals, target.GetWithDefault(ctx, s.defaultTarget), resetTime)
	if err != nil {
		return nil, err
	}

	return c.Value, nil
}

// Set writes the value into the given metric cell.
func (s *inMemoryStore) Set(ctx context.Context, h types.Metric, resetTime time.Time, fieldVals []interface{}, value interface{}) error {
	if resetTime.IsZero() {
		resetTime = clock.Now(ctx)
	}
	return s.set(h, resetTime, fieldVals, target.GetWithDefault(ctx, s.defaultTarget), value)
}

func (s *inMemoryStore) set(h types.Metric, resetTime time.Time, fieldVals []interface{}, t types.Target, value interface{}) error {
	m := s.getOrCreateData(h)
	m.lock.Lock()
	defer m.lock.Unlock()

	c, err := m.get(fieldVals, t, resetTime)
	if err != nil {
		return err
	}

	c.Value = value
	return nil
}

// Incr increments the value in a given metric cell by the given delta.
func (s *inMemoryStore) Incr(ctx context.Context, h types.Metric, resetTime time.Time, fieldVals []interface{}, delta interface{}) error {
	if resetTime.IsZero() {
		resetTime = clock.Now(ctx)
	}
	return s.incr(h, resetTime, fieldVals, target.GetWithDefault(ctx, s.defaultTarget), delta)
}

func (s *inMemoryStore) incr(h types.Metric, resetTime time.Time, fieldVals []interface{}, t types.Target, delta interface{}) error {
	m := s.getOrCreateData(h)
	m.lock.Lock()
	defer m.lock.Unlock()

	c, err := m.get(fieldVals, t, resetTime)
	if err != nil {
		return err
	}

	switch m.ValueType {
	case types.CumulativeDistributionType:
		d, ok := delta.(float64)
		if !ok {
			return fmt.Errorf("Incr got a delta of unsupported type (%v)", delta)
		}
		v, ok := c.Value.(*distribution.Distribution)
		if !ok {
			v = distribution.New(h.(types.DistributionMetric).Bucketer())
			c.Value = v
		}
		v.Add(float64(d))

	case types.CumulativeIntType:
		if v, ok := c.Value.(int64); ok {
			c.Value = v + delta.(int64)
		} else {
			c.Value = delta.(int64)
		}

	case types.CumulativeFloatType:
		if v, ok := c.Value.(float64); ok {
			c.Value = v + delta.(float64)
		} else {
			c.Value = delta.(float64)
		}

	default:
		return fmt.Errorf("attempted to increment non-cumulative metric %s by %v",
			m.Name, delta)
	}

	return nil
}

func (s *inMemoryStore) ModifyMulti(ctx context.Context, mods []Modification) error {
	contextTarget := target.GetWithDefault(ctx, s.defaultTarget)

	for _, m := range mods {
		resetTime := m.ResetTime
		if resetTime.IsZero() {
			resetTime = clock.Now(ctx)
		}

		t := m.Target
		if t == nil {
			t = contextTarget
		}

		if m.SetValue != nil {
			if err := s.set(m.Metric, resetTime, m.FieldVals, t, m.SetValue); err != nil {
				return err
			}
		} else if m.IncrDelta != nil {
			if err := s.incr(m.Metric, resetTime, m.FieldVals, t, m.IncrDelta); err != nil {
				return err
			}
		}
	}
	return nil
}

// ResetForUnittest resets all metric values without unregistering any metrics.
// Useful for unit tests.
func (s *inMemoryStore) ResetForUnittest() {
	s.lock.Lock()
	defer s.lock.Unlock()

	for _, m := range s.data {
		m.lock.Lock()
		m.cells = make(map[cellKey][]*types.CellData)
		m.lock.Unlock()
	}
}

// GetAll efficiently returns all cells in the store.
func (s *inMemoryStore) GetAll(ctx context.Context) []types.Cell {
	s.lock.Lock()
	defer s.lock.Unlock()

	ret := []types.Cell{}
	for _, m := range s.data {
		m.lock.Lock()
		for _, cells := range m.cells {
			for _, cell := range cells {
				ret = append(ret, types.Cell{m.MetricInfo, *cell})
			}
		}
		m.lock.Unlock()
	}
	return ret
}
