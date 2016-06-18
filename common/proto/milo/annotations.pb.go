// Code generated by protoc-gen-go.
// source: annotations.proto
// DO NOT EDIT!

/*
Package milo is a generated protocol buffer package.

It is generated from these files:
	annotations.proto

It has these top-level messages:
	FailureDetails
	Step
	Component
	Command
	Progress
	LogdogStream
	IsolateObject
	DMLink
*/
package milo

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/luci/luci-go/common/proto/google"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Status is the expressed root step of this step or substep.
type Status int32

const (
	// The step is still running.
	Status_RUNNING Status = 0
	// The step has finished successfully.
	Status_SUCCESS Status = 1
	// The step has finished unsuccessfully.
	Status_FAILURE Status = 2
	// The step has finished unexpectedly.
	Status_EXCEPTION Status = 3
)

var Status_name = map[int32]string{
	0: "RUNNING",
	1: "SUCCESS",
	2: "FAILURE",
	3: "EXCEPTION",
}
var Status_value = map[string]int32{
	"RUNNING":   0,
	"SUCCESS":   1,
	"FAILURE":   2,
	"EXCEPTION": 3,
}

func (x Status) String() string {
	return proto.EnumName(Status_name, int32(x))
}
func (Status) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// Type is the type of failure.
type FailureDetails_Type int32

const (
	// The failure is a general failure.
	FailureDetails_GENERAL FailureDetails_Type = 0
	// The failure is related to a failed infrastructure component, not an error
	// with the Step itself.
	FailureDetails_INFRA FailureDetails_Type = 1
	// The failure is due to a failed Dungeon Master dependency. This should be
	// used if a Step's external depdendency fails and the Step cannot recover
	// or proceed without it.
	FailureDetails_DM_DEPENDENCY_FAILED FailureDetails_Type = 2
)

var FailureDetails_Type_name = map[int32]string{
	0: "GENERAL",
	1: "INFRA",
	2: "DM_DEPENDENCY_FAILED",
}
var FailureDetails_Type_value = map[string]int32{
	"GENERAL":              0,
	"INFRA":                1,
	"DM_DEPENDENCY_FAILED": 2,
}

func (x FailureDetails_Type) String() string {
	return proto.EnumName(FailureDetails_Type_name, int32(x))
}
func (FailureDetails_Type) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

// FailureType provides more details on the nature of the Status.
type FailureDetails struct {
	Type FailureDetails_Type `protobuf:"varint,1,opt,name=type,enum=milo.FailureDetails_Type" json:"type,omitempty"`
	// An optional string describing the failure.
	Text string `protobuf:"bytes,2,opt,name=text" json:"text,omitempty"`
	// If the failure type is DEPENDENCY_FAILED, the failed dependencies should be
	// listed here.
	FailedDmDependency []*DMLink `protobuf:"bytes,3,rep,name=failed_dm_dependency,json=failedDmDependency" json:"failed_dm_dependency,omitempty"`
}

func (m *FailureDetails) Reset()                    { *m = FailureDetails{} }
func (m *FailureDetails) String() string            { return proto.CompactTextString(m) }
func (*FailureDetails) ProtoMessage()               {}
func (*FailureDetails) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *FailureDetails) GetFailedDmDependency() []*DMLink {
	if m != nil {
		return m.FailedDmDependency
	}
	return nil
}

// Generic step or substep state.
type Step struct {
	// The command-line invocation of the step, expressed as an argument vector.
	Command *Command `protobuf:"bytes,1,opt,name=command" json:"command,omitempty"`
	// Optional information detailing the failure. This may be populated if the
	// Step's top-level command Status is set to FAILURE.
	FailureDetails *FailureDetails `protobuf:"bytes,2,opt,name=failure_details,json=failureDetails" json:"failure_details,omitempty"`
	// Base Component information describing the high-level Step.
	StepComponent *Component `protobuf:"bytes,3,opt,name=step_component,json=stepComponent" json:"step_component,omitempty"`
	// Sub-components that are part of the Step.
	Components []*Component `protobuf:"bytes,4,rep,name=components" json:"components,omitempty"`
	// Substeps will be constructed as extensions on the parent LogDog stream.
	//
	// For example, if the parent's logging base path is:
	// luci/dm/QUEST/ATTEMPT/EXECUTION/+/
	//
	// Its substep #0 will have logging base path:
	// luci/dm/QUEST/ATTEMPT/EXECUTION/+/steps/0
	//
	// ... which can have known log stream names appended to it for the full
	// log stream path. The following appendages are part of the standard
	// Milo protocol expectations:
	// - .../stdout: A text stream containing the Step's STDOUT.
	// - .../stderr: A text stream containing the Step's STDERR.
	// - .../annotation: A text stream containing the Step's annotation stream
	//                   protobuf (Step message protobuf).
	//
	// For example:
	// - luci/dm/QUEST/ATTEMPT/EXECUTION/+/steps/0/stdout
	// - luci/dm/QUEST/ATTEMPT/EXECUTION/+/steps/0/annotations
	SubstepLogdogNameBase []string `protobuf:"bytes,5,rep,name=substep_logdog_name_base,json=substepLogdogNameBase" json:"substep_logdog_name_base,omitempty"`
}

func (m *Step) Reset()                    { *m = Step{} }
func (m *Step) String() string            { return proto.CompactTextString(m) }
func (*Step) ProtoMessage()               {}
func (*Step) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Step) GetCommand() *Command {
	if m != nil {
		return m.Command
	}
	return nil
}

func (m *Step) GetFailureDetails() *FailureDetails {
	if m != nil {
		return m.FailureDetails
	}
	return nil
}

func (m *Step) GetStepComponent() *Component {
	if m != nil {
		return m.StepComponent
	}
	return nil
}

func (m *Step) GetComponents() []*Component {
	if m != nil {
		return m.Components
	}
	return nil
}

// A Component represents a renderable state.
type Component struct {
	// The display name of the Component.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// Type classifies the information associated with the Note.
	Status Status `protobuf:"varint,2,opt,name=status,enum=milo.Status" json:"status,omitempty"`
	// When the step started, expressed as an RFC3339 string using Z (UTC)
	// timezone.
	Started *google_protobuf.Timestamp `protobuf:"bytes,3,opt,name=started" json:"started,omitempty"`
	// When the step ended, expressed as an RFC3339 string using Z (UTC) timezone.
	Ended *google_protobuf.Timestamp `protobuf:"bytes,4,opt,name=ended" json:"ended,omitempty"`
	// Arbitrary lines of component text. Each string here is a consecutive line,
	// and should not contain newlines.
	Text []string `protobuf:"bytes,5,rep,name=text" json:"text,omitempty"`
	// The Component's progress.
	Progress *Progress `protobuf:"bytes,6,opt,name=progress" json:"progress,omitempty"`
	// The primary link for this Component. This is the link that interaction
	// with the Component will use.
	Link *Component_Link `protobuf:"bytes,7,opt,name=link" json:"link,omitempty"`
	// Additional links related to the Component. These will be rendered alongside
	// the component.
	OtherLinks []*Component_Link     `protobuf:"bytes,8,rep,name=other_links,json=otherLinks" json:"other_links,omitempty"`
	Property   []*Component_Property `protobuf:"bytes,9,rep,name=property" json:"property,omitempty"`
}

func (m *Component) Reset()                    { *m = Component{} }
func (m *Component) String() string            { return proto.CompactTextString(m) }
func (*Component) ProtoMessage()               {}
func (*Component) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Component) GetStarted() *google_protobuf.Timestamp {
	if m != nil {
		return m.Started
	}
	return nil
}

func (m *Component) GetEnded() *google_protobuf.Timestamp {
	if m != nil {
		return m.Ended
	}
	return nil
}

func (m *Component) GetProgress() *Progress {
	if m != nil {
		return m.Progress
	}
	return nil
}

func (m *Component) GetLink() *Component_Link {
	if m != nil {
		return m.Link
	}
	return nil
}

func (m *Component) GetOtherLinks() []*Component_Link {
	if m != nil {
		return m.OtherLinks
	}
	return nil
}

func (m *Component) GetProperty() []*Component_Property {
	if m != nil {
		return m.Property
	}
	return nil
}

// A Link is an optional label followed by a typed link to an external
// resource.
type Component_Link struct {
	// An optional display label for the link.
	Label string `protobuf:"bytes,1,opt,name=label" json:"label,omitempty"`
	// Types that are valid to be assigned to Value:
	//	*Component_Link_Url
	//	*Component_Link_LogdogStream
	//	*Component_Link_IsolateObject
	//	*Component_Link_DmLink
	Value isComponent_Link_Value `protobuf_oneof:"value"`
}

func (m *Component_Link) Reset()                    { *m = Component_Link{} }
func (m *Component_Link) String() string            { return proto.CompactTextString(m) }
func (*Component_Link) ProtoMessage()               {}
func (*Component_Link) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2, 0} }

type isComponent_Link_Value interface {
	isComponent_Link_Value()
}

type Component_Link_Url struct {
	Url string `protobuf:"bytes,2,opt,name=url,oneof"`
}
type Component_Link_LogdogStream struct {
	LogdogStream *LogdogStream `protobuf:"bytes,3,opt,name=logdog_stream,json=logdogStream,oneof"`
}
type Component_Link_IsolateObject struct {
	IsolateObject *IsolateObject `protobuf:"bytes,4,opt,name=isolate_object,json=isolateObject,oneof"`
}
type Component_Link_DmLink struct {
	DmLink *DMLink `protobuf:"bytes,5,opt,name=dm_link,json=dmLink,oneof"`
}

func (*Component_Link_Url) isComponent_Link_Value()           {}
func (*Component_Link_LogdogStream) isComponent_Link_Value()  {}
func (*Component_Link_IsolateObject) isComponent_Link_Value() {}
func (*Component_Link_DmLink) isComponent_Link_Value()        {}

func (m *Component_Link) GetValue() isComponent_Link_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *Component_Link) GetUrl() string {
	if x, ok := m.GetValue().(*Component_Link_Url); ok {
		return x.Url
	}
	return ""
}

func (m *Component_Link) GetLogdogStream() *LogdogStream {
	if x, ok := m.GetValue().(*Component_Link_LogdogStream); ok {
		return x.LogdogStream
	}
	return nil
}

func (m *Component_Link) GetIsolateObject() *IsolateObject {
	if x, ok := m.GetValue().(*Component_Link_IsolateObject); ok {
		return x.IsolateObject
	}
	return nil
}

func (m *Component_Link) GetDmLink() *DMLink {
	if x, ok := m.GetValue().(*Component_Link_DmLink); ok {
		return x.DmLink
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Component_Link) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Component_Link_OneofMarshaler, _Component_Link_OneofUnmarshaler, _Component_Link_OneofSizer, []interface{}{
		(*Component_Link_Url)(nil),
		(*Component_Link_LogdogStream)(nil),
		(*Component_Link_IsolateObject)(nil),
		(*Component_Link_DmLink)(nil),
	}
}

func _Component_Link_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Component_Link)
	// value
	switch x := m.Value.(type) {
	case *Component_Link_Url:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.Url)
	case *Component_Link_LogdogStream:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.LogdogStream); err != nil {
			return err
		}
	case *Component_Link_IsolateObject:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.IsolateObject); err != nil {
			return err
		}
	case *Component_Link_DmLink:
		b.EncodeVarint(5<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.DmLink); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Component_Link.Value has unexpected type %T", x)
	}
	return nil
}

func _Component_Link_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Component_Link)
	switch tag {
	case 2: // value.url
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Value = &Component_Link_Url{x}
		return true, err
	case 3: // value.logdog_stream
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(LogdogStream)
		err := b.DecodeMessage(msg)
		m.Value = &Component_Link_LogdogStream{msg}
		return true, err
	case 4: // value.isolate_object
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(IsolateObject)
		err := b.DecodeMessage(msg)
		m.Value = &Component_Link_IsolateObject{msg}
		return true, err
	case 5: // value.dm_link
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(DMLink)
		err := b.DecodeMessage(msg)
		m.Value = &Component_Link_DmLink{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Component_Link_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Component_Link)
	// value
	switch x := m.Value.(type) {
	case *Component_Link_Url:
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.Url)))
		n += len(x.Url)
	case *Component_Link_LogdogStream:
		s := proto.Size(x.LogdogStream)
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Component_Link_IsolateObject:
		s := proto.Size(x.IsolateObject)
		n += proto.SizeVarint(4<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Component_Link_DmLink:
		s := proto.Size(x.DmLink)
		n += proto.SizeVarint(5<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// Property is an arbitrary key/value (build) property.
type Component_Property struct {
	// name is the property name.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// value is the optional property value.
	Value string `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
}

func (m *Component_Property) Reset()                    { *m = Component_Property{} }
func (m *Component_Property) String() string            { return proto.CompactTextString(m) }
func (*Component_Property) ProtoMessage()               {}
func (*Component_Property) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2, 1} }

// Command contains information about a command-line invocation.
type Command struct {
	// The command-line invocation, expressed as an argument vector.
	CommandLine []string `protobuf:"bytes,1,rep,name=command_line,json=commandLine" json:"command_line,omitempty"`
	// The current working directory.
	Cwd     string               `protobuf:"bytes,2,opt,name=cwd" json:"cwd,omitempty"`
	Environ *Command_Environment `protobuf:"bytes,3,opt,name=environ" json:"environ,omitempty"`
}

func (m *Command) Reset()                    { *m = Command{} }
func (m *Command) String() string            { return proto.CompactTextString(m) }
func (*Command) ProtoMessage()               {}
func (*Command) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Command) GetEnviron() *Command_Environment {
	if m != nil {
		return m.Environ
	}
	return nil
}

// Environment represents the state of a process' environment.
type Command_Environment struct {
	// The entries that compose the environment.
	Entries []*Command_Environment_Entry `protobuf:"bytes,1,rep,name=entries" json:"entries,omitempty"`
}

func (m *Command_Environment) Reset()                    { *m = Command_Environment{} }
func (m *Command_Environment) String() string            { return proto.CompactTextString(m) }
func (*Command_Environment) ProtoMessage()               {}
func (*Command_Environment) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3, 0} }

func (m *Command_Environment) GetEntries() []*Command_Environment_Entry {
	if m != nil {
		return m.Entries
	}
	return nil
}

// Entry is a single key/value environment entry.
type Command_Environment_Entry struct {
	// Name is the name of the environment variable.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// Value is the value of the environment variable. This may be empty.
	Value string `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
}

func (m *Command_Environment_Entry) Reset()                    { *m = Command_Environment_Entry{} }
func (m *Command_Environment_Entry) String() string            { return proto.CompactTextString(m) }
func (*Command_Environment_Entry) ProtoMessage()               {}
func (*Command_Environment_Entry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3, 0, 0} }

// Progress expresses a Component's overall progress. It does this using
// arbitrary "progress units", wich are discrete units of work measured by the
// Component that are either completed or not completed.
//
// A simple construction for "percentage complete" is to set `total` to 100 and
// `completed` to the percentage value.
type Progress struct {
	// The total number of progress units. If missing or zero, no progress is
	// expressed.
	Total int32 `protobuf:"varint,1,opt,name=total" json:"total,omitempty"`
	// The number of completed progress units. This must always be less than or
	// equal to `total`. If omitted, it is implied to be zero.
	Completed int32 `protobuf:"varint,2,opt,name=completed" json:"completed,omitempty"`
}

func (m *Progress) Reset()                    { *m = Progress{} }
func (m *Progress) String() string            { return proto.CompactTextString(m) }
func (*Progress) ProtoMessage()               {}
func (*Progress) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

// LogdogStream is a LogDog stream link.
type LogdogStream struct {
	// The stream's server. If omitted, the server is the same server that this
	// annotation stream is homed on.
	Server string `protobuf:"bytes,1,opt,name=server" json:"server,omitempty"`
	// The log Prefix. If empty, the prefix is the same prefix as this annotation
	// stream.
	Prefix string `protobuf:"bytes,2,opt,name=prefix" json:"prefix,omitempty"`
	// The log name.
	Name string `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
}

func (m *LogdogStream) Reset()                    { *m = LogdogStream{} }
func (m *LogdogStream) String() string            { return proto.CompactTextString(m) }
func (*LogdogStream) ProtoMessage()               {}
func (*LogdogStream) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

// IsolateObject is an Isolate service object specification.
type IsolateObject struct {
	// The Isolate server. If empty, this is the default Isolate server specified
	// by the project's LUCI config.
	Server string `protobuf:"bytes,1,opt,name=server" json:"server,omitempty"`
	// The isolate object hash.
	Hash string `protobuf:"bytes,2,opt,name=hash" json:"hash,omitempty"`
}

func (m *IsolateObject) Reset()                    { *m = IsolateObject{} }
func (m *IsolateObject) String() string            { return proto.CompactTextString(m) }
func (*IsolateObject) ProtoMessage()               {}
func (*IsolateObject) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

// Dependency is a Dungeon Master execution specification.
type DMLink struct {
	// The Dungeon Master server. If empty, this is the default Isolate server
	// specified by the project's LUCI config.
	Server string `protobuf:"bytes,1,opt,name=server" json:"server,omitempty"`
	// The quest name.
	Quest string `protobuf:"bytes,2,opt,name=quest" json:"quest,omitempty"`
	// The attempt number.
	Attempt int64 `protobuf:"varint,3,opt,name=attempt" json:"attempt,omitempty"`
	// The execution number.
	Execution int64 `protobuf:"varint,4,opt,name=execution" json:"execution,omitempty"`
}

func (m *DMLink) Reset()                    { *m = DMLink{} }
func (m *DMLink) String() string            { return proto.CompactTextString(m) }
func (*DMLink) ProtoMessage()               {}
func (*DMLink) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func init() {
	proto.RegisterType((*FailureDetails)(nil), "milo.FailureDetails")
	proto.RegisterType((*Step)(nil), "milo.Step")
	proto.RegisterType((*Component)(nil), "milo.Component")
	proto.RegisterType((*Component_Link)(nil), "milo.Component.Link")
	proto.RegisterType((*Component_Property)(nil), "milo.Component.Property")
	proto.RegisterType((*Command)(nil), "milo.Command")
	proto.RegisterType((*Command_Environment)(nil), "milo.Command.Environment")
	proto.RegisterType((*Command_Environment_Entry)(nil), "milo.Command.Environment.Entry")
	proto.RegisterType((*Progress)(nil), "milo.Progress")
	proto.RegisterType((*LogdogStream)(nil), "milo.LogdogStream")
	proto.RegisterType((*IsolateObject)(nil), "milo.IsolateObject")
	proto.RegisterType((*DMLink)(nil), "milo.DMLink")
	proto.RegisterEnum("milo.Status", Status_name, Status_value)
	proto.RegisterEnum("milo.FailureDetails_Type", FailureDetails_Type_name, FailureDetails_Type_value)
}

func init() { proto.RegisterFile("annotations.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 907 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x55, 0xeb, 0x6e, 0xe3, 0x44,
	0x14, 0xde, 0x5c, 0x9c, 0x6c, 0x4e, 0x9a, 0x6c, 0x18, 0x02, 0xf2, 0x46, 0x48, 0xbb, 0x44, 0x48,
	0x54, 0x2b, 0xe1, 0x42, 0x96, 0xab, 0x80, 0x4a, 0xdd, 0xc4, 0x5d, 0x22, 0x75, 0xdd, 0x6a, 0xd2,
	0x4a, 0xf0, 0xcb, 0x72, 0xe2, 0x49, 0x6a, 0xf0, 0x0d, 0x7b, 0x5c, 0x5a, 0xf1, 0x12, 0xfc, 0xe4,
	0xa5, 0x78, 0x02, 0x5e, 0x82, 0x47, 0xe0, 0xcc, 0xc5, 0x4e, 0x02, 0xad, 0x10, 0xff, 0xce, 0xe5,
	0xfb, 0xe6, 0xcc, 0xb9, 0xcd, 0xc0, 0x5b, 0x5e, 0x1c, 0x27, 0xdc, 0xe3, 0x41, 0x12, 0xe7, 0x56,
	0x9a, 0x25, 0x3c, 0x21, 0xcd, 0x28, 0x08, 0x93, 0xd1, 0xb3, 0x4d, 0x92, 0x6c, 0x42, 0x76, 0x24,
	0x6d, 0xcb, 0x62, 0x7d, 0xc4, 0x83, 0x88, 0xe5, 0xdc, 0x8b, 0x52, 0x05, 0x1b, 0xff, 0x51, 0x83,
	0xfe, 0xa9, 0x17, 0x84, 0x45, 0xc6, 0x66, 0x8c, 0xa3, 0x90, 0x93, 0x8f, 0xa0, 0xc9, 0xef, 0x52,
	0x66, 0xd6, 0x9e, 0xd7, 0x0e, 0xfb, 0x93, 0xa7, 0x96, 0x38, 0xc8, 0xda, 0xc7, 0x58, 0x97, 0x08,
	0xa0, 0x12, 0x46, 0x08, 0xc2, 0xd9, 0x2d, 0x37, 0xeb, 0x08, 0xef, 0x50, 0x29, 0x93, 0x63, 0x18,
	0xae, 0x11, 0xc7, 0x7c, 0xd7, 0x8f, 0x5c, 0x9f, 0xa5, 0x2c, 0xf6, 0x59, 0xbc, 0xba, 0x33, 0x1b,
	0xcf, 0x1b, 0x87, 0xdd, 0xc9, 0x81, 0x3a, 0x72, 0xf6, 0xe6, 0x2c, 0x88, 0x7f, 0xa2, 0x44, 0x21,
	0x67, 0xd1, 0xac, 0xc2, 0x8d, 0xbf, 0x84, 0xa6, 0x88, 0x40, 0xba, 0xd0, 0x7e, 0x6d, 0x3b, 0x36,
	0x3d, 0x39, 0x1b, 0x3c, 0x22, 0x1d, 0x30, 0xe6, 0xce, 0x29, 0x3d, 0x19, 0xd4, 0x88, 0x09, 0xc3,
	0xd9, 0x1b, 0x77, 0x66, 0x5f, 0xd8, 0xce, 0xcc, 0x76, 0xa6, 0x3f, 0xb8, 0xa7, 0x27, 0xf3, 0x33,
	0x7b, 0x36, 0xa8, 0x8f, 0x7f, 0xab, 0x43, 0x73, 0xc1, 0x59, 0x4a, 0x3e, 0x84, 0xf6, 0x2a, 0x89,
	0x22, 0x2f, 0xf6, 0x65, 0x22, 0xdd, 0x49, 0x4f, 0x45, 0x9d, 0x2a, 0x23, 0x2d, 0xbd, 0xe4, 0x5b,
	0x78, 0xb2, 0x56, 0xc9, 0xe1, 0x4d, 0x65, 0x76, 0x32, 0x95, 0xee, 0x64, 0x78, 0x5f, 0xe6, 0xb4,
	0xbf, 0xde, 0xaf, 0xd6, 0xe7, 0xd0, 0xcf, 0x31, 0x9e, 0x8b, 0xc7, 0xa5, 0x49, 0xcc, 0x62, 0x8e,
	0x49, 0x0a, 0xf6, 0x93, 0x2a, 0x9c, 0x32, 0xd3, 0x9e, 0x80, 0x55, 0x2a, 0x39, 0x02, 0xa8, 0x28,
	0xb9, 0xd9, 0x94, 0x85, 0xf9, 0x17, 0x67, 0x07, 0x42, 0xbe, 0x00, 0x33, 0x2f, 0x96, 0x32, 0x56,
	0x98, 0x6c, 0xfc, 0x64, 0xe3, 0xc6, 0x5e, 0xc4, 0xdc, 0xa5, 0x97, 0x33, 0xd3, 0x40, 0x7a, 0x87,
	0xbe, 0xa3, 0xfd, 0x67, 0xd2, 0xed, 0xa0, 0xf7, 0x15, 0x3a, 0xc7, 0xbf, 0x1b, 0xd0, 0xd9, 0xc6,
	0xc5, 0x76, 0x09, 0x9e, 0x2c, 0x0a, 0xb6, 0x4b, 0xc8, 0xe4, 0x03, 0x68, 0xe1, 0x4c, 0xf0, 0x42,
	0x65, 0xde, 0x2f, 0x1b, 0xb4, 0x90, 0x36, 0xaa, 0x7d, 0xe4, 0x53, 0x68, 0xa3, 0x94, 0x71, 0xe6,
	0xeb, 0x14, 0x47, 0x96, 0x9a, 0x2e, 0xab, 0x9c, 0x2e, 0xeb, 0xb2, 0x9c, 0x2e, 0x5a, 0x42, 0xc9,
	0xc7, 0x60, 0x88, 0xae, 0xfa, 0x98, 0xe2, 0x7f, 0x71, 0x14, 0xb0, 0x1a, 0x28, 0x95, 0x94, 0x1a,
	0xa8, 0x17, 0xf0, 0x18, 0x09, 0x9b, 0x8c, 0xe5, 0xb9, 0xd9, 0x92, 0x07, 0xf5, 0xd5, 0x1d, 0x2f,
	0xb4, 0x95, 0x56, 0x7e, 0x72, 0x08, 0xcd, 0x10, 0x07, 0xcb, 0x6c, 0xef, 0x76, 0xb1, 0x2a, 0x80,
	0x25, 0x87, 0x4e, 0x22, 0xc8, 0x67, 0xd0, 0x4d, 0xf8, 0x35, 0xcb, 0x5c, 0xa1, 0xe5, 0xe6, 0x63,
	0xd9, 0x84, 0xfb, 0x09, 0x20, 0x81, 0x42, 0x14, 0x85, 0x10, 0xc1, 0x52, 0x96, 0xf1, 0x3b, 0xb3,
	0x23, 0x39, 0xe6, 0x3f, 0x39, 0x17, 0xda, 0x4f, 0x2b, 0xe4, 0xe8, 0xcf, 0x1a, 0x34, 0x05, 0x9f,
	0x0c, 0xc1, 0x08, 0xbd, 0x25, 0x0b, 0x75, 0x0b, 0x94, 0x82, 0x59, 0x37, 0x8a, 0x2c, 0x54, 0x5b,
	0xf4, 0xdd, 0x23, 0x2a, 0x14, 0xf2, 0x15, 0xf4, 0x74, 0xab, 0x73, 0x9e, 0x31, 0x2f, 0xd2, 0x75,
	0x27, 0x2a, 0x9a, 0x6a, 0xf3, 0x42, 0x7a, 0x90, 0x71, 0x10, 0xee, 0xe8, 0xe4, 0x1b, 0xe8, 0x07,
	0x79, 0x12, 0x7a, 0x9c, 0xb9, 0xc9, 0xf2, 0x47, 0xb6, 0xe2, 0xba, 0xfe, 0x6f, 0x2b, 0xee, 0x5c,
	0xf9, 0xce, 0xa5, 0x0b, 0xc9, 0xbd, 0x60, 0xd7, 0x20, 0x96, 0x07, 0x17, 0x57, 0x56, 0xd1, 0x90,
	0xb4, 0xbd, 0x95, 0x45, 0x7c, 0xcb, 0x8f, 0x84, 0xf4, 0xaa, 0x0d, 0xc6, 0x8d, 0x17, 0x16, 0x6c,
	0x84, 0x35, 0x29, 0x73, 0xbe, 0x77, 0xc4, 0x86, 0x1a, 0xa8, 0x9f, 0x09, 0xa5, 0x8c, 0xff, 0xaa,
	0x41, 0x5b, 0x2f, 0x24, 0x79, 0x1f, 0x0e, 0xf4, 0x4a, 0x8a, 0xc0, 0x82, 0x2d, 0xda, 0xdf, 0xd5,
	0x36, 0x8c, 0xc6, 0xc8, 0x00, 0x1a, 0xab, 0x5f, 0x7c, 0x7d, 0x84, 0x10, 0xc9, 0x4b, 0x68, 0xb3,
	0xf8, 0x26, 0xc8, 0x92, 0x58, 0xd7, 0xe6, 0xe9, 0xde, 0x96, 0x5b, 0xb6, 0x72, 0x46, 0x62, 0x99,
	0x4a, 0xe4, 0xe8, 0x57, 0xe8, 0xee, 0xd8, 0xb1, 0xca, 0xe8, 0xe1, 0x59, 0xc0, 0x72, 0x19, 0xb3,
	0x3b, 0x79, 0xf6, 0xe0, 0x19, 0x28, 0xf3, 0xec, 0x8e, 0x96, 0xf8, 0xd1, 0x27, 0x60, 0x48, 0xcb,
	0xff, 0x48, 0xf9, 0x58, 0x16, 0x4a, 0x4d, 0x2a, 0x22, 0x38, 0x3e, 0xdb, 0x6a, 0x12, 0x0c, 0xaa,
	0x14, 0xf2, 0x1e, 0x74, 0xc4, 0xda, 0x87, 0x4c, 0x6c, 0x5a, 0x5d, 0x7a, 0xb6, 0x86, 0x31, 0x85,
	0x83, 0xdd, 0xc6, 0x93, 0x77, 0x71, 0x77, 0x59, 0x76, 0xc3, 0x32, 0x1d, 0x5b, 0x6b, 0xc2, 0x9e,
	0x66, 0x6c, 0x1d, 0xdc, 0xea, 0xf0, 0x5a, 0xab, 0x6e, 0xda, 0xd8, 0xde, 0x74, 0xfc, 0x35, 0xf4,
	0xf6, 0x06, 0xe2, 0xc1, 0x43, 0x91, 0x7c, 0xed, 0xe5, 0xd7, 0xe5, 0x5b, 0x2f, 0xe4, 0x71, 0x0c,
	0x2d, 0x35, 0x16, 0x0f, 0xb2, 0x30, 0xcd, 0x9f, 0x0b, 0xdc, 0xf1, 0xb2, 0x10, 0x52, 0xc1, 0x37,
	0xbc, 0xed, 0x71, 0xce, 0xa2, 0x54, 0xbd, 0x98, 0x0d, 0x5a, 0xaa, 0xa2, 0x00, 0xec, 0x96, 0xad,
	0x0a, 0xf1, 0x9d, 0xc9, 0xb1, 0x6d, 0xd0, 0xad, 0xe1, 0xc5, 0x31, 0xb4, 0xd4, 0xc3, 0x24, 0x7e,
	0x07, 0x7a, 0xe5, 0x38, 0x73, 0xe7, 0x35, 0xfe, 0x0e, 0xa8, 0x2c, 0xae, 0xa6, 0x53, 0x7b, 0xb1,
	0xc0, 0xff, 0x01, 0x15, 0xf1, 0x23, 0x5c, 0x51, 0x7b, 0x50, 0x27, 0x3d, 0xe8, 0xd8, 0xdf, 0x4f,
	0xed, 0x8b, 0xcb, 0xf9, 0xb9, 0x33, 0x68, 0x2c, 0x5b, 0xf2, 0xe5, 0x79, 0xf9, 0x77, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xd3, 0xc2, 0x04, 0x70, 0x34, 0x07, 0x00, 0x00,
}
