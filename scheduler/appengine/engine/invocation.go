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
	"strings"
	"time"

	"golang.org/x/net/context"

	"go.chromium.org/gae/service/datastore"

	"go.chromium.org/luci/common/auth/identity"
	"go.chromium.org/luci/common/clock"
	"go.chromium.org/luci/common/data/rand/mathrand"
	"go.chromium.org/luci/common/errors"
	"go.chromium.org/luci/common/retry/transient"

	"go.chromium.org/luci/scheduler/appengine/task"
)

// errInvocationIDConflict is returned by generateInvocationID.
var errInvocationIDConflict = errors.New("could not find available invocationID", transient.Tag)

const (
	// debugLogSizeLimit is how many bytes the invocation debug log can be before
	// it gets trimmed. See 'trimDebugLog'. The debug log isn't supposed to be
	// huge.
	debugLogSizeLimit = 200000

	// debugLogTailLines is how many last log lines to keep when trimming the log.
	debugLogTailLines = 100
)

// Jan 1 2015, in UTC.
var invocationIDEpoch time.Time

func init() {
	var err error
	invocationIDEpoch, err = time.Parse(time.RFC822, "01 Jan 15 00:00 UTC")
	if err != nil {
		panic(err)
	}
}

// generateInvocationID is called within a transaction to pick a new Invocation
// ID and ensure it isn't taken yet.
//
// 'parent', if not nil, defines an entity group for the new Invocation entity.
// It can be nil, in which case the Invocation entity will be in its own entity
// group.
//
// Format of the invocation ID:
//   - highest order bit set to 0 to keep the value positive.
//   - next 43 bits set to negated time since some predefined epoch, in ms.
//   - next 16 bits are generated by math.Rand
//   - next 4 bits set to 0. They indicate ID format.
//
// Makes one attempt at allocating an ID. If it fails (should be extremely
// rare), the entire transaction should be retried. We do it to avoid
// unnecessarily hitting multiple entity groups from a single transaction.
//
// Returns only transient errors.
func generateInvocationID(c context.Context, parent *datastore.Key) (int64, error) {
	// See http://play.golang.org/p/POpQzpT4Up.
	invTs := int64(clock.Now(c).UTC().Sub(invocationIDEpoch) / time.Millisecond)
	invTs = ^invTs & 8796093022207 // 0b111....1, 42 bits (clear highest bit)
	invTs = invTs << 20

	randSuffix := mathrand.Int63n(c, 65536)
	invID := invTs | (randSuffix << 4)
	exists, err := datastore.Exists(c, datastore.NewKey(c, "Invocation", "", invID, parent))
	if err != nil {
		return 0, transient.Tag.Apply(err)
	}
	if !exists.All() {
		return invID, nil
	}

	return 0, errInvocationIDConflict
}

// Invocation entity stores single attempt to run a job. Its parent entity
// is corresponding Job, its ID is generated based on time.
type Invocation struct {
	_kind  string                `gae:"$kind,Invocation"`
	_extra datastore.PropertyMap `gae:"-,extra"`

	// ID is identifier of this particular attempt to run a job. Multiple attempts
	// to start an invocation result in multiple entities with different IDs, but
	// with same InvocationNonce.
	ID int64 `gae:"$id"`

	// JobKey is the key of parent Job entity.
	//
	// For v1 Invocation entities it is set to the corresponding job key. For v2
	// entities it is always nil.
	JobKey *datastore.Key `gae:"$parent"`

	// JobID is '<ProjectID>/<JobName>' string of a parent job.
	//
	// For v1 Invocation entities it is "". For v2 entities it is set when the
	// invocation is created and never changes.
	JobID string `gae:",noindex"`

	// IndexedJobID is '<ProjectID>/<JobName>' string of a parent job, but it is
	// set only for finished invocations.
	//
	// It is used to make the invocations appear in the listings of finished
	// invocations.
	//
	// We can't use JobID field for this since the invocation launch procedure can
	// potentially generate orphaned "garbage" invocations in some edge cases (if
	// Invocation transaction lands, but separate Job transaction doesn't). They
	// are harmless, but we don't want them to show up in listings.
	//
	// Used only for v2 entities.
	IndexedJobID string

	// Started is time when this invocation was created.
	Started time.Time `gae:",noindex"`

	// Finished is time when this invocation transitioned to a terminal state.
	Finished time.Time `gae:",noindex"`

	// InvocationNonce identifies a request to start a job, produced by
	// StateMachine.
	//
	// TODO(vadimsh): Remove in v2.
	InvocationNonce int64

	// TriggeredBy is identity of whoever triggered the invocation, if it was
	// triggered via ForceInvocation ("Run now" button).
	//
	// Empty identity string if it was triggered by the service itself.
	TriggeredBy identity.Identity

	// IncomingTriggers is a list of triggers that the invocation consumed.
	//
	// They are popped from job's pending triggers set when the invocation
	// starts.
	IncomingTriggers []task.Trigger `gae:",noindex"`

	// OutgoingTriggers is a list of triggers that the invocation produced.
	//
	// They are fanned out into pending trigger sets of corresponding triggered
	// jobs (specified by TriggeredJobIDs).
	OutgoingTriggers []task.Trigger `gae:",noindex"`

	// Revision is revision number of config.cfg when this invocation was created.
	// For informational purpose.
	Revision string `gae:",noindex"`

	// RevisionURL is URL to human readable page with config file at
	// an appropriate revision. For informational purpose.
	RevisionURL string `gae:",noindex"`

	// Task is the job payload for this invocation in binary serialized form.
	// For informational purpose. See Catalog.UnmarshalTask().
	Task []byte `gae:",noindex"`

	// TriggeredJobIDs is a list of jobIDs of jobs which this job triggers.
	// The list is sorted and without duplicates.
	TriggeredJobIDs []string `gae:",noindex"`

	// DebugLog is short free form text log with debug messages.
	DebugLog string `gae:",noindex"`

	// RetryCount is 0 on a first attempt to launch the task. Increased with each
	// retry. For informational purposes.
	RetryCount int64 `gae:",noindex"`

	// Status is current status of the invocation (e.g. "RUNNING"), see the enum.
	Status task.Status

	// ViewURL is optional URL to a human readable page with task status, e.g.
	// Swarming task page. Populated by corresponding TaskManager.
	ViewURL string `gae:",noindex"`

	// TaskData is a storage where TaskManager can keep task-specific state
	// between calls.
	TaskData []byte `gae:",noindex"`

	// MutationsCount is used for simple compare-and-swap transaction control.
	//
	// It is incremented on each change to the entity.
	MutationsCount int64 `gae:",noindex"`
}

// jobID returns ID of the job the invocation belongs too.
//
// Supports both v1 and v2 invocations.
func (e *Invocation) jobID() string {
	if e.JobKey != nil {
		return e.JobKey.StringID() // for v1
	}
	return e.JobID // for v2
}

// isEqual returns true iff 'e' is equal to 'other'
func (e *Invocation) isEqual(other *Invocation) bool {
	return e == other || (e.ID == other.ID &&
		e.MutationsCount == other.MutationsCount && // compare it first, it changes most often
		(e.JobKey == other.JobKey || e.JobKey.Equal(other.JobKey)) &&
		e.JobID == other.JobID &&
		e.IndexedJobID == other.IndexedJobID &&
		e.Started.Equal(other.Started) &&
		e.Finished.Equal(other.Finished) &&
		e.InvocationNonce == other.InvocationNonce &&
		e.TriggeredBy == other.TriggeredBy &&
		equalTriggerLists(e.IncomingTriggers, other.IncomingTriggers) &&
		equalTriggerLists(e.OutgoingTriggers, other.OutgoingTriggers) &&
		e.Revision == other.Revision &&
		e.RevisionURL == other.RevisionURL &&
		bytes.Equal(e.Task, other.Task) &&
		equalSortedLists(e.TriggeredJobIDs, other.TriggeredJobIDs) &&
		e.DebugLog == other.DebugLog &&
		e.RetryCount == other.RetryCount &&
		e.Status == other.Status &&
		e.ViewURL == other.ViewURL &&
		bytes.Equal(e.TaskData, other.TaskData))
}

// debugLog appends a line to DebugLog field.
func (e *Invocation) debugLog(c context.Context, format string, args ...interface{}) {
	debugLog(c, &e.DebugLog, format, args...)
}

// trimDebugLog makes sure DebugLog field doesn't exceed limits.
//
// It cuts the middle of the log. We need to do this to keep the entity small
// enough to fit the datastore limits.
func (e *Invocation) trimDebugLog() {
	if len(e.DebugLog) <= debugLogSizeLimit {
		return
	}

	const cutMsg = "--- the log has been cut here ---"
	giveUp := func() {
		e.DebugLog = e.DebugLog[:debugLogSizeLimit-len(cutMsg)-2] + "\n" + cutMsg + "\n"
	}

	// We take last debugLogTailLines lines of log and move them "up", so that
	// the total log size is less than debugLogSizeLimit. We then put a line with
	// the message that some log lines have been cut. If these operations are not
	// possible (e.g. we have some giant lines or something), we give up and just
	// cut the end of the log.

	// Find debugLogTailLines-th "\n" from the end, e.DebugLog[tailStart:] is the
	// log tail.
	tailStart := len(e.DebugLog)
	for i := 0; i < debugLogTailLines; i++ {
		tailStart = strings.LastIndex(e.DebugLog[:tailStart-1], "\n")
		if tailStart <= 0 {
			giveUp()
			return
		}
	}
	tailStart++

	// Figure out how many bytes of head we can keep to make trimmed log small
	// enough.
	tailLen := len(e.DebugLog) - tailStart + len(cutMsg) + 1
	headSize := debugLogSizeLimit - tailLen
	if headSize <= 0 {
		giveUp()
		return
	}

	// Find last "\n" in the head.
	headEnd := strings.LastIndex(e.DebugLog[:headSize], "\n")
	if headEnd <= 0 {
		giveUp()
		return
	}

	// We want to keep 50 lines of the head no matter what.
	headLines := strings.Count(e.DebugLog[:headEnd], "\n")
	if headLines < 50 {
		giveUp()
		return
	}

	// Remove duplicated 'cutMsg' lines. They may appear if 'debugLog' (followed
	// by 'trimDebugLog') is called on already trimmed log multiple times.
	lines := strings.Split(e.DebugLog[:headEnd], "\n")
	lines = append(lines, cutMsg)
	lines = append(lines, strings.Split(e.DebugLog[tailStart:], "\n")...)
	trimmed := make([]byte, 0, debugLogSizeLimit)
	trimmed = append(trimmed, lines[0]...)
	for i := 1; i < len(lines); i++ {
		if !(lines[i-1] == cutMsg && lines[i] == cutMsg) {
			trimmed = append(trimmed, '\n')
			trimmed = append(trimmed, lines[i]...)
		}
	}
	e.DebugLog = string(trimmed)
}
