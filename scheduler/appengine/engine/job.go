// Copyright 2017 The LUCI Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package engine

import (
	"bytes"
	"fmt"
	"hash/fnv"
	"strconv"
	"strings"
	"time"

	"github.com/golang/protobuf/proto"
	"golang.org/x/net/context"

	"go.chromium.org/gae/service/datastore"
	"go.chromium.org/luci/common/errors"
	"go.chromium.org/luci/common/logging"

	"go.chromium.org/luci/scheduler/appengine/acl"
	"go.chromium.org/luci/scheduler/appengine/catalog"
	"go.chromium.org/luci/scheduler/appengine/engine/dsset"
	"go.chromium.org/luci/scheduler/appengine/internal"
	"go.chromium.org/luci/scheduler/appengine/schedule"
)

// Job stores the last known definition of a scheduler job, as well as its
// current state. Root entity, its kind is "Job".
type Job struct {
	_kind  string                `gae:"$kind,Job"`
	_extra datastore.PropertyMap `gae:"-,extra"`

	// cachedSchedule and cachedScheduleErr are used by ParseSchedule().
	cachedSchedule    *schedule.Schedule `gae:"-"`
	cachedScheduleErr error              `gae:"-"`

	// JobID is '<ProjectID>/<JobName>' string. JobName is unique with a project,
	// but not globally. JobID is unique globally.
	JobID string `gae:"$id"`

	// ProjectID exists for indexing. It matches <projectID> portion of JobID.
	ProjectID string

	// Flavor describes what category of jobs this is, see the enum.
	Flavor catalog.JobFlavor `gae:",noindex"`

	// Enabled is false if the job was disabled or removed from config.
	//
	// Disabled jobs do not show up in UI at all (they are still kept in the
	// datastore though, for audit purposes).
	Enabled bool

	// Paused is true if job's schedule is ignored and job can only be started
	// manually via "Run now" button.
	Paused bool `gae:",noindex"`

	// Revision is last seen job definition revision.
	Revision string `gae:",noindex"`

	// RevisionURL is URL to human readable page with config file at
	// an appropriate revision.
	RevisionURL string `gae:",noindex"`

	// Schedule is the job's schedule in regular cron expression format.
	Schedule string `gae:",noindex"`

	// Task is the job's payload in serialized form. Opaque from the point of view
	// of the engine. See Catalog.UnmarshalTask().
	Task []byte `gae:",noindex"`

	// TriggeredJobIDs is a list of jobIDs of jobs which this job triggers.
	// The list is sorted and without duplicates.
	TriggeredJobIDs []string `gae:",noindex"`

	// ACLs are the latest ACLs applied to Job and all its invocations.
	Acls acl.GrantsByRole `gae:",noindex"`

	// State is the job's state machine state, see StateMachine.
	//
	// Used by v1 jobs only.
	State JobState

	// ActiveInvocations is ordered set of active invocation IDs.
	//
	// It contains IDs of pending, running or recently finished invocations,
	// the most recent at the end.
	//
	// Used by v2 jobs only.
	ActiveInvocations []int64 `gae:",noindex"`
}

// JobName returns name of this Job as defined its project's config.
//
// This is "<name>"" part extracted from "<project>/<name>" job ID.
func (e *Job) JobName() string {
	chunks := strings.Split(e.JobID, "/")
	return chunks[1]
}

// EffectiveSchedule returns schedule string to use for the job, considering its
// Paused field.
//
// Paused jobs always use "triggered" schedule.
func (e *Job) EffectiveSchedule() string {
	if e.Paused {
		return "triggered"
	}
	return e.Schedule
}

// ParseSchedule returns *Schedule object, parsing e.Schedule field.
//
// If job is paused e.Schedule field is ignored and "triggered" schedule is
// returned instead.
func (e *Job) ParseSchedule() (*schedule.Schedule, error) {
	if e.cachedSchedule == nil && e.cachedScheduleErr == nil {
		hash := fnv.New64()
		hash.Write([]byte(e.JobID))
		seed := hash.Sum64()
		e.cachedSchedule, e.cachedScheduleErr = schedule.Parse(e.EffectiveSchedule(), seed)
		if e.cachedSchedule == nil && e.cachedScheduleErr == nil {
			panic("no schedule and no error")
		}
	}
	return e.cachedSchedule, e.cachedScheduleErr
}

// IsEqual returns true iff 'e' is equal to 'other'.
func (e *Job) IsEqual(other *Job) bool {
	return e == other || (e.JobID == other.JobID &&
		e.ProjectID == other.ProjectID &&
		e.Flavor == other.Flavor &&
		e.Enabled == other.Enabled &&
		e.Paused == other.Paused &&
		e.Revision == other.Revision &&
		e.RevisionURL == other.RevisionURL &&
		e.Schedule == other.Schedule &&
		e.Acls.Equal(&other.Acls) &&
		bytes.Equal(e.Task, other.Task) &&
		equalSortedLists(e.TriggeredJobIDs, other.TriggeredJobIDs) &&
		e.State.Equal(&other.State) &&
		equalInt64Lists(e.ActiveInvocations, other.ActiveInvocations))
}

// MatchesDefinition returns true if job definition in the entity matches the
// one specified by catalog.Definition struct.
func (e *Job) MatchesDefinition(def catalog.Definition) bool {
	return e.JobID == def.JobID &&
		e.Flavor == def.Flavor &&
		e.Schedule == def.Schedule &&
		e.Acls.Equal(&def.Acls) &&
		bytes.Equal(e.Task, def.Task) &&
		equalSortedLists(e.TriggeredJobIDs, def.TriggeredJobIDs)
}

// CheckRole returns nil if the caller has the given role or ErrNoPermission
// otherwise.
//
// May also return transient errors.
func (e *Job) CheckRole(c context.Context, role acl.Role) error {
	switch yep, err := e.Acls.CallerHasRole(logging.SetField(c, "JobID", e.JobID), role); {
	case err != nil:
		return err
	case !yep:
		return ErrNoPermission
	}
	return nil
}

////////////////////////////////////////////////////////////////////////////////
// v2 stuff.

// recentlyFinishedSet is a set with IDs of all recently finished invocations.
//
// This is an accumulator of IDs to remove from ActiveInvocations list next
// time we run a triage for the corresponding Job.
//
// Invocation IDs are serialized with fmt.Sprintf("%d").
func recentlyFinishedSet(c context.Context, jobID string) *invocationIDSet {
	return &invocationIDSet{
		Set: dsset.Set{
			ID:              "finished:" + jobID,
			ShardCount:      8,
			TombstonesRoot:  datastore.KeyForObj(c, &Job{JobID: jobID}),
			TombstonesDelay: 30 * time.Minute,
		},
	}
}

// invocationIDSet is a dsset.Set that stores invocation IDs.
type invocationIDSet struct {
	dsset.Set
}

// Add adds a bunch of invocation IDs to the set.
func (s *invocationIDSet) Add(c context.Context, ids []int64) error {
	items := make([]dsset.Item, len(ids))
	for i, id := range ids {
		items[i].ID = fmt.Sprintf("%d", id)
	}
	return s.Set.Add(c, items)
}

// ItemToInvID takes a dsset.Item and returns invocation ID stored there or 0 if
// it's malformed.
func (s *invocationIDSet) ItemToInvID(i *dsset.Item) int64 {
	id, _ := strconv.ParseInt(i.ID, 10, 64)
	return id
}

// pendingTriggersSet is a set of not yet consumed triggers for the job.
//
// This is incoming triggers. They are processed in the triage procedure,
// resulting in new invocations.
func pendingTriggersSet(c context.Context, jobID string) *triggersSet {
	return &triggersSet{
		Set: dsset.Set{
			ID:              "triggers:" + jobID,
			ShardCount:      8,
			TombstonesRoot:  datastore.KeyForObj(c, &Job{JobID: jobID}),
			TombstonesDelay: 30 * time.Minute,
		},
	}
}

// triggersSet is a dsset.Set that stores internal.Trigger protos.
type triggersSet struct {
	dsset.Set
}

// Add adds triggers to the set.
func (s *triggersSet) Add(c context.Context, triggers []*internal.Trigger) error {
	items := make([]dsset.Item, 0, len(triggers))
	for _, t := range triggers {
		blob, err := proto.Marshal(t)
		if err != nil {
			return fmt.Errorf("failed to marshal proto - %s", err)
		}
		items = append(items, dsset.Item{
			ID:    t.Id,
			Value: blob,
		})
	}
	return s.Set.Add(c, items)
}

// Triggers returns all triggers in the set, in no particular order.
//
// Returns the original dsset listing (that can be used later with BeginPop or
// CleanupGarbage), as well as the actual deserialized set of triggers.
func (s *triggersSet) Triggers(c context.Context) (*dsset.Listing, []*internal.Trigger, error) {
	l, err := s.Set.List(c)
	if err != nil {
		return nil, nil, err
	}
	out := make([]*internal.Trigger, len(l.Items))
	for i, item := range l.Items {
		out[i] = &internal.Trigger{}
		if err := proto.Unmarshal(item.Value, out[i]); err != nil {
			return nil, nil, errors.Annotate(err, "failed to unmarshal trigger").Err()
		}
		if out[i].Id != item.ID {
			return nil, nil, fmt.Errorf("trigger ID in the body (%q) doesn't match item ID %q", out[i].Id, item.ID)
		}
	}
	return l, out, nil
}