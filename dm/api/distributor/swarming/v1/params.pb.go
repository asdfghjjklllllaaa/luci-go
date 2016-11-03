// Code generated by protoc-gen-go.
// source: github.com/luci/luci-go/dm/api/distributor/swarming/v1/params.proto
// DO NOT EDIT!

package swarmingV1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/luci/luci-go/common/proto/google"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Parameters represents the set of swarming parameters that the Swarming v1
// distributor can interpret for use with a Swarming v1 compatible service.
type Parameters struct {
	Scheduling *Parameters_Scheduling `protobuf:"bytes,1,opt,name=scheduling" json:"scheduling,omitempty"`
	Meta       *Parameters_Meta       `protobuf:"bytes,2,opt,name=meta" json:"meta,omitempty"`
	Job        *Parameters_Job        `protobuf:"bytes,3,opt,name=job" json:"job,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

func (m *Parameters) GetScheduling() *Parameters_Scheduling {
	if m != nil {
		return m.Scheduling
	}
	return nil
}

func (m *Parameters) GetMeta() *Parameters_Meta {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *Parameters) GetJob() *Parameters_Job {
	if m != nil {
		return m.Job
	}
	return nil
}

// These parameters affect how the Executions for this Quest are scheduled.
type Parameters_Scheduling struct {
	// Priority adjusts the scheduling preference for Executions. The higher the
	// priority number, the longer it will take to schedule, and vice versa (ala
	// `nice`).
	//
	// A value of 0 (default) corresponds to a default priority (currently 100).
	// Other values correspond directly to the swarming task Priority.
	//
	// THIS MEANS THAT PRIORITY 0 IS NOT PERMITTED HERE! If you want 0, use
	// 1 instead.
	//
	// This must be <= 255.
	Priority uint32 `protobuf:"varint,1,opt,name=priority" json:"priority,omitempty"`
	// These specifiy the profile of the machine to use for Execuions of this
	// quest. These can indicate OS, number of cores, amount of ram, GPU type,
	// pool, etc. See the swarming service instance for available dimensions.
	Dimensions map[string]string `protobuf:"bytes,2,rep,name=dimensions" json:"dimensions,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// These dimensions will be snapshotted from the first execution of each
	// attempt, and will be subsequently re-used for all following executions of
	// that attempt.
	//
	// The most-specific value for these dimensions will be taken for tasks
	// where a given dimension has multiple values.
	SnapshotDimensions []string `protobuf:"bytes,3,rep,name=snapshot_dimensions,json=snapshotDimensions" json:"snapshot_dimensions,omitempty"`
	// This indicates the maximum amount of time that an Execution may run
	// without emitting IO on stdout/err. 0 means 'no timeout'.
	IoTimeout *google_protobuf.Duration `protobuf:"bytes,4,opt,name=io_timeout,json=ioTimeout" json:"io_timeout,omitempty"`
}

func (m *Parameters_Scheduling) Reset()                    { *m = Parameters_Scheduling{} }
func (m *Parameters_Scheduling) String() string            { return proto.CompactTextString(m) }
func (*Parameters_Scheduling) ProtoMessage()               {}
func (*Parameters_Scheduling) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0, 0} }

func (m *Parameters_Scheduling) GetDimensions() map[string]string {
	if m != nil {
		return m.Dimensions
	}
	return nil
}

func (m *Parameters_Scheduling) GetIoTimeout() *google_protobuf.Duration {
	if m != nil {
		return m.IoTimeout
	}
	return nil
}

type Parameters_Meta struct {
	// The 'human readable' name prefix for Executions of this quest. DM will
	// automatically prepend this to the execution ID. So if this was "cool
	// job", the swarming task name would be
	//   "cool job / <quest_ID>|<attempt>|<execution>"
	NamePrefix string `protobuf:"bytes,1,opt,name=name_prefix,json=namePrefix" json:"name_prefix,omitempty"`
}

func (m *Parameters_Meta) Reset()                    { *m = Parameters_Meta{} }
func (m *Parameters_Meta) String() string            { return proto.CompactTextString(m) }
func (*Parameters_Meta) ProtoMessage()               {}
func (*Parameters_Meta) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0, 1} }

type Parameters_Job struct {
	Inputs *Parameters_Job_Inputs `protobuf:"bytes,1,opt,name=inputs" json:"inputs,omitempty"`
	// This is the "argv" to run with this job. This includes substitution
	// paramters defined by swarming's run_isolated.py script:
	//   https://github.com/luci/luci-py/blob/master/client/run_isolated.py
	//
	// Additionally, DM provides the following substitutions:
	//   ${DM.PREVIOUS.EXECUTION.STATE:PATH} - the path to a JSONPB encoding of
	//     the swarming_v1.Result from the previous Execution of this Attempt.
	//   ${DM.QUEST.DATA.DESC:PATH} - the path to the
	//     dm.Quest.Data.Desc JSONPB for the quest that this execution is part
	//     of.
	//   ${DM.HOST} - the hostname to use to access DM's pRPC API.
	//
	// DM also provides a JSONPB encoded dm.Execution.Auth via the LUCI_CONTEXT
	// swarming.secret_bytes value. See:
	//   https://github.com/luci/luci-py/blob/master/client/LUCI_CONTEXT.md
	//
	// Command MUST be specified; specifying a command in any of the isolated
	// inputs WILL NOT DO ANYTHING.
	Command []string `protobuf:"bytes,2,rep,name=command" json:"command,omitempty"`
	// Environment variables for the executions.
	Env map[string]string `protobuf:"bytes,3,rep,name=env" json:"env,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *Parameters_Job) Reset()                    { *m = Parameters_Job{} }
func (m *Parameters_Job) String() string            { return proto.CompactTextString(m) }
func (*Parameters_Job) ProtoMessage()               {}
func (*Parameters_Job) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0, 2} }

func (m *Parameters_Job) GetInputs() *Parameters_Job_Inputs {
	if m != nil {
		return m.Inputs
	}
	return nil
}

func (m *Parameters_Job) GetEnv() map[string]string {
	if m != nil {
		return m.Env
	}
	return nil
}

type Parameters_Job_Inputs struct {
	// 0 or more isolated IDs that will be 'included' together into the final
	// job .isolated sent to swarming.
	//
	// The "server" value must either be omitted, or equal the isolate server
	// defined by this distributor's config (the `isolate.host` field,
	// prepended with "https://").
	Isolated []*IsolatedRef `protobuf:"bytes,1,rep,name=isolated" json:"isolated,omitempty"`
	// CIPD packages to use for the job. These specs may contain templated
	// parameters for package names or non-instance_ids for the package
	// versions. The first successful execution for each attempt will
	// resolve+snapshot all package names and versions. These package names
	// and versions will be used for all subsequent executions of that
	// attempt.
	Cipd *CipdSpec `protobuf:"bytes,2,opt,name=cipd" json:"cipd,omitempty"`
}

func (m *Parameters_Job_Inputs) Reset()                    { *m = Parameters_Job_Inputs{} }
func (m *Parameters_Job_Inputs) String() string            { return proto.CompactTextString(m) }
func (*Parameters_Job_Inputs) ProtoMessage()               {}
func (*Parameters_Job_Inputs) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0, 2, 0} }

func (m *Parameters_Job_Inputs) GetIsolated() []*IsolatedRef {
	if m != nil {
		return m.Isolated
	}
	return nil
}

func (m *Parameters_Job_Inputs) GetCipd() *CipdSpec {
	if m != nil {
		return m.Cipd
	}
	return nil
}

func init() {
	proto.RegisterType((*Parameters)(nil), "swarmingV1.Parameters")
	proto.RegisterType((*Parameters_Scheduling)(nil), "swarmingV1.Parameters.Scheduling")
	proto.RegisterType((*Parameters_Meta)(nil), "swarmingV1.Parameters.Meta")
	proto.RegisterType((*Parameters_Job)(nil), "swarmingV1.Parameters.Job")
	proto.RegisterType((*Parameters_Job_Inputs)(nil), "swarmingV1.Parameters.Job.Inputs")
}

func init() {
	proto.RegisterFile("github.com/luci/luci-go/dm/api/distributor/swarming/v1/params.proto", fileDescriptor3)
}

var fileDescriptor3 = []byte{
	// 515 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x93, 0x51, 0x6b, 0x13, 0x41,
	0x10, 0xc7, 0x49, 0x2e, 0xc6, 0x66, 0x82, 0x28, 0x6b, 0xc1, 0xf3, 0x04, 0x8d, 0xfa, 0x60, 0x1e,
	0xf4, 0x96, 0xa4, 0x28, 0x55, 0xf0, 0xa1, 0xb4, 0x05, 0x1b, 0x10, 0xea, 0x56, 0x7c, 0x0d, 0x7b,
	0x77, 0x9b, 0xcb, 0x6a, 0x6e, 0x67, 0xd9, 0xdd, 0x8b, 0xf6, 0x03, 0xf9, 0x05, 0xc4, 0x0f, 0x28,
	0xb7, 0x77, 0x49, 0x0e, 0xa1, 0x16, 0xfb, 0x12, 0x32, 0x3b, 0xbf, 0xff, 0x7f, 0x76, 0x66, 0xe7,
	0xe0, 0x38, 0x97, 0x6e, 0x59, 0x26, 0x71, 0x8a, 0x05, 0x5d, 0x95, 0xa9, 0xf4, 0x3f, 0xaf, 0x72,
	0xa4, 0x59, 0x41, 0xb9, 0x96, 0x34, 0x93, 0xd6, 0x19, 0x99, 0x94, 0x0e, 0x0d, 0xb5, 0xdf, 0xb9,
	0x29, 0xa4, 0xca, 0xe9, 0x7a, 0x42, 0x35, 0x37, 0xbc, 0xb0, 0xb1, 0x36, 0xe8, 0x90, 0xc0, 0x26,
	0xf3, 0x65, 0x12, 0x3d, 0xce, 0x11, 0xf3, 0x95, 0xa0, 0x3e, 0x93, 0x94, 0x0b, 0x9a, 0x95, 0x86,
	0x3b, 0x89, 0xaa, 0x66, 0xa3, 0xa3, 0x1b, 0x16, 0x4c, 0xa5, 0xce, 0x1a, 0x8b, 0x0f, 0x37, 0xb4,
	0x90, 0x16, 0x57, 0xdc, 0x89, 0xb9, 0x11, 0x8b, 0xda, 0xe9, 0xd9, 0xef, 0x3e, 0xc0, 0x79, 0xd5,
	0x89, 0x70, 0xc2, 0x58, 0x72, 0x04, 0x60, 0xd3, 0xa5, 0xc8, 0xca, 0x95, 0x54, 0x79, 0xd8, 0x19,
	0x75, 0xc6, 0xc3, 0xe9, 0xd3, 0x78, 0xd7, 0x5c, 0xbc, 0x63, 0xe3, 0x8b, 0x2d, 0xc8, 0x5a, 0x22,
	0x42, 0xa1, 0x57, 0x08, 0xc7, 0xc3, 0xae, 0x17, 0x3f, 0xba, 0x42, 0xfc, 0x51, 0x38, 0xce, 0x3c,
	0x48, 0x5e, 0x42, 0xf0, 0x15, 0x93, 0x30, 0xf0, 0x7c, 0x74, 0x05, 0x3f, 0xc3, 0x84, 0x55, 0x58,
	0xf4, 0xb3, 0x0b, 0xb0, 0xab, 0x4c, 0x22, 0xd8, 0xd3, 0x46, 0xa2, 0x91, 0xee, 0xd2, 0x5f, 0xf7,
	0x0e, 0xdb, 0xc6, 0xe4, 0x13, 0x40, 0x26, 0x0b, 0xa1, 0xac, 0x44, 0x65, 0xc3, 0xee, 0x28, 0x18,
	0x0f, 0xa7, 0x93, 0x6b, 0x9b, 0x89, 0x4f, 0xb6, 0x9a, 0x53, 0xe5, 0xcc, 0x25, 0x6b, 0x99, 0x10,
	0x0a, 0xf7, 0xad, 0xe2, 0xda, 0x2e, 0xd1, 0xcd, 0x5b, 0xde, 0xc1, 0x28, 0x18, 0x0f, 0x18, 0xd9,
	0xa4, 0x76, 0x0e, 0xe4, 0x10, 0x40, 0xe2, 0xdc, 0xc9, 0x42, 0x60, 0xe9, 0xc2, 0x9e, 0xef, 0xf1,
	0x61, 0x5c, 0x6f, 0x48, 0xbc, 0xd9, 0x90, 0xf8, 0xa4, 0xd9, 0x10, 0x36, 0x90, 0xf8, 0xb9, 0x66,
	0xa3, 0xf7, 0x70, 0xf7, 0xaf, 0x9b, 0x90, 0x7b, 0x10, 0x7c, 0x13, 0x75, 0x9f, 0x03, 0x56, 0xfd,
	0x25, 0xfb, 0x70, 0x6b, 0xcd, 0x57, 0xa5, 0xf0, 0xd3, 0x1e, 0xb0, 0x3a, 0x78, 0xd7, 0x3d, 0xec,
	0x44, 0x2f, 0xa0, 0x57, 0xcd, 0x98, 0x3c, 0x81, 0xa1, 0xe2, 0x85, 0x98, 0x6b, 0x23, 0x16, 0xf2,
	0x47, 0xa3, 0x85, 0xea, 0xe8, 0xdc, 0x9f, 0x44, 0xbf, 0xba, 0x10, 0xcc, 0x30, 0x21, 0x6f, 0xa1,
	0x2f, 0x95, 0x2e, 0x9d, 0xbd, 0xe6, 0xd9, 0x67, 0x98, 0xc4, 0x67, 0x1e, 0x64, 0x8d, 0x80, 0x84,
	0x70, 0x3b, 0xc5, 0xa2, 0xe0, 0x2a, 0xf3, 0x53, 0x1e, 0xb0, 0x4d, 0x48, 0x5e, 0x43, 0x20, 0xd4,
	0xda, 0xcf, 0x67, 0x38, 0x7d, 0xfe, 0x0f, 0xc7, 0x53, 0xb5, 0xae, 0xa7, 0x5d, 0xf1, 0x51, 0x0e,
	0xfd, 0xba, 0x04, 0x39, 0x80, 0xbd, 0x66, 0x69, 0xb3, 0xb0, 0xe3, 0x5d, 0x1e, 0xb4, 0x5d, 0xce,
	0x9a, 0x1c, 0x13, 0x0b, 0xb6, 0x05, 0xc9, 0x18, 0x7a, 0xd5, 0xc7, 0xd2, 0xac, 0xe0, 0x7e, 0x5b,
	0x70, 0x2c, 0x75, 0x76, 0xa1, 0x45, 0xca, 0x3c, 0x11, 0xbd, 0x81, 0xbd, 0x4d, 0xe5, 0xff, 0x99,
	0x6e, 0xd2, 0xf7, 0x4f, 0x77, 0xf0, 0x27, 0x00, 0x00, 0xff, 0xff, 0xe2, 0x62, 0x64, 0x76, 0x3d,
	0x04, 0x00, 0x00,
}
