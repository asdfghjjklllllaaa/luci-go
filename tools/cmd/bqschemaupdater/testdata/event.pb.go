// Code generated by protoc-gen-go. DO NOT EDIT.
// source: go.chromium.org/luci/tools/cmd/bqschemaupdater/testdata/event.proto

/*
Package testdata is a generated protocol buffer package.

It is generated from these files:
	go.chromium.org/luci/tools/cmd/bqschemaupdater/testdata/event.proto

It has these top-level messages:
	Property
	Input
	Output
	EmptyContainer
	BuildEvent
*/
package testdata

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/struct"
import google_protobuf1 "github.com/golang/protobuf/ptypes/timestamp"
import google_protobuf2 "github.com/golang/protobuf/ptypes/empty"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Status int32

const (
	Status_SUCCESS Status = 0
	Status_FAILURE Status = 1
	Status_ERROR   Status = 2
)

var Status_name = map[int32]string{
	0: "SUCCESS",
	1: "FAILURE",
	2: "ERROR",
}
var Status_value = map[string]int32{
	"SUCCESS": 0,
	"FAILURE": 1,
	"ERROR":   2,
}

func (x Status) String() string {
	return proto.EnumName(Status_name, int32(x))
}
func (Status) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Property struct {
	Name      string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	ValueJson string `protobuf:"bytes,2,opt,name=value_json,json=valueJson" json:"value_json,omitempty"`
}

func (m *Property) Reset()                    { *m = Property{} }
func (m *Property) String() string            { return proto.CompactTextString(m) }
func (*Property) ProtoMessage()               {}
func (*Property) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Property) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Property) GetValueJson() string {
	if m != nil {
		return m.ValueJson
	}
	return ""
}

type Input struct {
	Properties []*Property `protobuf:"bytes,4,rep,name=properties" json:"properties,omitempty"`
}

func (m *Input) Reset()                    { *m = Input{} }
func (m *Input) String() string            { return proto.CompactTextString(m) }
func (*Input) ProtoMessage()               {}
func (*Input) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Input) GetProperties() []*Property {
	if m != nil {
		return m.Properties
	}
	return nil
}

type Output struct {
	Properties []*Property `protobuf:"bytes,4,rep,name=properties" json:"properties,omitempty"`
}

func (m *Output) Reset()                    { *m = Output{} }
func (m *Output) String() string            { return proto.CompactTextString(m) }
func (*Output) ProtoMessage()               {}
func (*Output) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Output) GetProperties() []*Property {
	if m != nil {
		return m.Properties
	}
	return nil
}

// This entire message will be ignored.
type EmptyContainer struct {
	Empty *google_protobuf2.Empty `protobuf:"bytes,1,opt,name=empty" json:"empty,omitempty"`
}

func (m *EmptyContainer) Reset()                    { *m = EmptyContainer{} }
func (m *EmptyContainer) String() string            { return proto.CompactTextString(m) }
func (*EmptyContainer) ProtoMessage()               {}
func (*EmptyContainer) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *EmptyContainer) GetEmpty() *google_protobuf2.Empty {
	if m != nil {
		return m.Empty
	}
	return nil
}

// Build events.
//
// Line after blank line.
type BuildEvent struct {
	// Universal build id.
	BuildId string `protobuf:"bytes,1,opt,name=build_id,json=buildId" json:"build_id,omitempty"`
	// Builder name.
	Builder        string                      `protobuf:"bytes,2,opt,name=builder" json:"builder,omitempty"`
	Status         Status                      `protobuf:"varint,3,opt,name=status,enum=testdata.Status" json:"status,omitempty"`
	Input          *Input                      `protobuf:"bytes,4,opt,name=input" json:"input,omitempty"`
	Output         *Output                     `protobuf:"bytes,5,opt,name=output" json:"output,omitempty"`
	Timestamp      *google_protobuf1.Timestamp `protobuf:"bytes,6,opt,name=timestamp" json:"timestamp,omitempty"`
	Struct         *google_protobuf.Struct     `protobuf:"bytes,7,opt,name=struct" json:"struct,omitempty"`
	Empty          *google_protobuf2.Empty     `protobuf:"bytes,8,opt,name=empty" json:"empty,omitempty"`
	EmptyContainer *EmptyContainer             `protobuf:"bytes,9,opt,name=empty_container,json=emptyContainer" json:"empty_container,omitempty"`
}

func (m *BuildEvent) Reset()                    { *m = BuildEvent{} }
func (m *BuildEvent) String() string            { return proto.CompactTextString(m) }
func (*BuildEvent) ProtoMessage()               {}
func (*BuildEvent) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *BuildEvent) GetBuildId() string {
	if m != nil {
		return m.BuildId
	}
	return ""
}

func (m *BuildEvent) GetBuilder() string {
	if m != nil {
		return m.Builder
	}
	return ""
}

func (m *BuildEvent) GetStatus() Status {
	if m != nil {
		return m.Status
	}
	return Status_SUCCESS
}

func (m *BuildEvent) GetInput() *Input {
	if m != nil {
		return m.Input
	}
	return nil
}

func (m *BuildEvent) GetOutput() *Output {
	if m != nil {
		return m.Output
	}
	return nil
}

func (m *BuildEvent) GetTimestamp() *google_protobuf1.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *BuildEvent) GetStruct() *google_protobuf.Struct {
	if m != nil {
		return m.Struct
	}
	return nil
}

func (m *BuildEvent) GetEmpty() *google_protobuf2.Empty {
	if m != nil {
		return m.Empty
	}
	return nil
}

func (m *BuildEvent) GetEmptyContainer() *EmptyContainer {
	if m != nil {
		return m.EmptyContainer
	}
	return nil
}

func init() {
	proto.RegisterType((*Property)(nil), "testdata.Property")
	proto.RegisterType((*Input)(nil), "testdata.Input")
	proto.RegisterType((*Output)(nil), "testdata.Output")
	proto.RegisterType((*EmptyContainer)(nil), "testdata.EmptyContainer")
	proto.RegisterType((*BuildEvent)(nil), "testdata.BuildEvent")
	proto.RegisterEnum("testdata.Status", Status_name, Status_value)
}

func init() {
	proto.RegisterFile("go.chromium.org/luci/tools/cmd/bqschemaupdater/testdata/event.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 461 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xdb, 0x8b, 0xd3, 0x40,
	0x14, 0xc6, 0xed, 0xb6, 0x49, 0x9b, 0x53, 0xe8, 0x96, 0x79, 0xd0, 0xb1, 0x2a, 0x96, 0x82, 0x50,
	0x44, 0x33, 0x50, 0x5f, 0x04, 0x2f, 0xb0, 0x96, 0x08, 0x15, 0x61, 0x65, 0xe2, 0x3e, 0x97, 0x69,
	0x32, 0x76, 0x23, 0x49, 0x26, 0xce, 0x65, 0x61, 0xff, 0x13, 0xff, 0x5c, 0xc9, 0x49, 0x62, 0xf6,
	0xf2, 0x22, 0xbe, 0xcd, 0x39, 0xdf, 0xef, 0x4c, 0xce, 0xf7, 0x4d, 0x60, 0x7b, 0x54, 0x61, 0x72,
	0xa9, 0x55, 0x91, 0xb9, 0x22, 0x54, 0xfa, 0xc8, 0x72, 0x97, 0x64, 0xcc, 0x2a, 0x95, 0x1b, 0x96,
	0x14, 0x29, 0x3b, 0xfc, 0x32, 0xc9, 0xa5, 0x2c, 0x84, 0xab, 0x52, 0x61, 0xa5, 0x66, 0x56, 0x1a,
	0x9b, 0x0a, 0x2b, 0x98, 0xbc, 0x92, 0xa5, 0x0d, 0x2b, 0xad, 0xac, 0x22, 0x93, 0xae, 0xbb, 0x78,
	0x7a, 0x54, 0xea, 0x98, 0x4b, 0x86, 0xfd, 0x83, 0xfb, 0xc1, 0x8c, 0xd5, 0x2e, 0x69, 0xb9, 0xc5,
	0xf3, 0xbb, 0xaa, 0xcd, 0x0a, 0x69, 0xac, 0x28, 0xaa, 0x16, 0x78, 0x72, 0x17, 0x90, 0x45, 0x65,
	0xaf, 0x1b, 0x71, 0xf5, 0x01, 0x26, 0xdf, 0xb4, 0xaa, 0xa4, 0xb6, 0xd7, 0x84, 0xc0, 0xa8, 0x14,
	0x85, 0xa4, 0x83, 0xe5, 0x60, 0x1d, 0x70, 0x3c, 0x93, 0x67, 0x00, 0x57, 0x22, 0x77, 0x72, 0xff,
	0xd3, 0xa8, 0x92, 0x9e, 0xa0, 0x12, 0x60, 0xe7, 0x8b, 0x51, 0xe5, 0xea, 0x1d, 0x78, 0xbb, 0xb2,
	0x72, 0x96, 0x6c, 0x00, 0xaa, 0xe6, 0x9e, 0x4c, 0x1a, 0x3a, 0x5a, 0x0e, 0xd7, 0xd3, 0x0d, 0x09,
	0x3b, 0x0b, 0x61, 0xf7, 0x0d, 0x7e, 0x83, 0x5a, 0xbd, 0x07, 0xff, 0xdc, 0xd9, 0xff, 0x9d, 0xfe,
	0x08, 0xb3, 0xa8, 0x36, 0xb2, 0x55, 0xa5, 0x15, 0x59, 0x29, 0x35, 0x79, 0x05, 0x1e, 0x5a, 0x43,
	0x03, 0xd3, 0xcd, 0xc3, 0xb0, 0x31, 0x1e, 0x76, 0xc6, 0x43, 0xe4, 0x79, 0x03, 0xad, 0x7e, 0x0f,
	0x01, 0x3e, 0xb9, 0x2c, 0x4f, 0xa3, 0x3a, 0x74, 0xf2, 0x18, 0x26, 0x87, 0xba, 0xda, 0x67, 0x69,
	0x1b, 0xc0, 0x18, 0xeb, 0x5d, 0x4a, 0x28, 0x34, 0x47, 0xa9, 0xdb, 0x00, 0xba, 0x92, 0xac, 0xc1,
	0x37, 0x56, 0x58, 0x67, 0xe8, 0x70, 0x39, 0x58, 0xcf, 0x36, 0xf3, 0x7e, 0xe7, 0x18, 0xfb, 0xbc,
	0xd5, 0xc9, 0x0b, 0xf0, 0xb2, 0x3a, 0x28, 0x3a, 0xc2, 0xdd, 0x4e, 0x7b, 0x10, 0xf3, 0xe3, 0x8d,
	0x5a, 0x5f, 0xa8, 0x30, 0x12, 0xea, 0x21, 0x77, 0xe3, 0xc2, 0x26, 0x2a, 0xde, 0xea, 0xe4, 0x2d,
	0x04, 0x7f, 0x1f, 0x9a, 0xfa, 0x08, 0x2f, 0xee, 0x19, 0xfe, 0xde, 0x11, 0xbc, 0x87, 0x09, 0xab,
	0x97, 0xae, 0x7f, 0x20, 0x3a, 0xc6, 0xb1, 0x47, 0xf7, 0xc6, 0x62, 0x94, 0x79, 0x8b, 0xf5, 0xb9,
	0x4e, 0xfe, 0x21, 0x57, 0x72, 0x06, 0xa7, 0x78, 0xd8, 0x27, 0xdd, 0xc3, 0xd0, 0x00, 0xe7, 0x68,
	0xef, 0xe5, 0xf6, 0xc3, 0xf1, 0x99, 0xbc, 0x55, 0xbf, 0x7c, 0x0d, 0x7e, 0x13, 0x1f, 0x99, 0xc2,
	0x38, 0xbe, 0xd8, 0x6e, 0xa3, 0x38, 0x9e, 0x3f, 0xa8, 0x8b, 0xcf, 0x67, 0xbb, 0xaf, 0x17, 0x3c,
	0x9a, 0x0f, 0x48, 0x00, 0x5e, 0xc4, 0xf9, 0x39, 0x9f, 0x9f, 0x1c, 0x7c, 0x5c, 0xe4, 0xcd, 0x9f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xfe, 0xef, 0x22, 0xe7, 0x77, 0x03, 0x00, 0x00,
}