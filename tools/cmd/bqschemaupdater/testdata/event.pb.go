// Code generated by protoc-gen-go. DO NOT EDIT.
// source: go.chromium.org/luci/tools/cmd/bqschemaupdater/testdata/event.proto

package testdata

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	empty "github.com/golang/protobuf/ptypes/empty"
	_struct "github.com/golang/protobuf/ptypes/struct"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

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

func (Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_aee82e0b1f3dea43, []int{0}
}

type Property struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	ValueJson            string   `protobuf:"bytes,2,opt,name=value_json,json=valueJson,proto3" json:"value_json,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Property) Reset()         { *m = Property{} }
func (m *Property) String() string { return proto.CompactTextString(m) }
func (*Property) ProtoMessage()    {}
func (*Property) Descriptor() ([]byte, []int) {
	return fileDescriptor_aee82e0b1f3dea43, []int{0}
}

func (m *Property) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Property.Unmarshal(m, b)
}
func (m *Property) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Property.Marshal(b, m, deterministic)
}
func (m *Property) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Property.Merge(m, src)
}
func (m *Property) XXX_Size() int {
	return xxx_messageInfo_Property.Size(m)
}
func (m *Property) XXX_DiscardUnknown() {
	xxx_messageInfo_Property.DiscardUnknown(m)
}

var xxx_messageInfo_Property proto.InternalMessageInfo

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
	Properties           []*Property `protobuf:"bytes,4,rep,name=properties,proto3" json:"properties,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Input) Reset()         { *m = Input{} }
func (m *Input) String() string { return proto.CompactTextString(m) }
func (*Input) ProtoMessage()    {}
func (*Input) Descriptor() ([]byte, []int) {
	return fileDescriptor_aee82e0b1f3dea43, []int{1}
}

func (m *Input) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Input.Unmarshal(m, b)
}
func (m *Input) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Input.Marshal(b, m, deterministic)
}
func (m *Input) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Input.Merge(m, src)
}
func (m *Input) XXX_Size() int {
	return xxx_messageInfo_Input.Size(m)
}
func (m *Input) XXX_DiscardUnknown() {
	xxx_messageInfo_Input.DiscardUnknown(m)
}

var xxx_messageInfo_Input proto.InternalMessageInfo

func (m *Input) GetProperties() []*Property {
	if m != nil {
		return m.Properties
	}
	return nil
}

type Output struct {
	Properties           []*Property `protobuf:"bytes,4,rep,name=properties,proto3" json:"properties,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Output) Reset()         { *m = Output{} }
func (m *Output) String() string { return proto.CompactTextString(m) }
func (*Output) ProtoMessage()    {}
func (*Output) Descriptor() ([]byte, []int) {
	return fileDescriptor_aee82e0b1f3dea43, []int{2}
}

func (m *Output) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Output.Unmarshal(m, b)
}
func (m *Output) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Output.Marshal(b, m, deterministic)
}
func (m *Output) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Output.Merge(m, src)
}
func (m *Output) XXX_Size() int {
	return xxx_messageInfo_Output.Size(m)
}
func (m *Output) XXX_DiscardUnknown() {
	xxx_messageInfo_Output.DiscardUnknown(m)
}

var xxx_messageInfo_Output proto.InternalMessageInfo

func (m *Output) GetProperties() []*Property {
	if m != nil {
		return m.Properties
	}
	return nil
}

// This entire message will be ignored.
type EmptyContainer struct {
	Empty                *empty.Empty `protobuf:"bytes,1,opt,name=empty,proto3" json:"empty,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *EmptyContainer) Reset()         { *m = EmptyContainer{} }
func (m *EmptyContainer) String() string { return proto.CompactTextString(m) }
func (*EmptyContainer) ProtoMessage()    {}
func (*EmptyContainer) Descriptor() ([]byte, []int) {
	return fileDescriptor_aee82e0b1f3dea43, []int{3}
}

func (m *EmptyContainer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EmptyContainer.Unmarshal(m, b)
}
func (m *EmptyContainer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EmptyContainer.Marshal(b, m, deterministic)
}
func (m *EmptyContainer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EmptyContainer.Merge(m, src)
}
func (m *EmptyContainer) XXX_Size() int {
	return xxx_messageInfo_EmptyContainer.Size(m)
}
func (m *EmptyContainer) XXX_DiscardUnknown() {
	xxx_messageInfo_EmptyContainer.DiscardUnknown(m)
}

var xxx_messageInfo_EmptyContainer proto.InternalMessageInfo

func (m *EmptyContainer) GetEmpty() *empty.Empty {
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
	BuildId string `protobuf:"bytes,1,opt,name=build_id,json=buildId,proto3" json:"build_id,omitempty"`
	// Builder name.
	Builder              string               `protobuf:"bytes,2,opt,name=builder,proto3" json:"builder,omitempty"`
	Status               Status               `protobuf:"varint,3,opt,name=status,proto3,enum=testdata.Status" json:"status,omitempty"`
	Input                *Input               `protobuf:"bytes,4,opt,name=input,proto3" json:"input,omitempty"`
	Output               *Output              `protobuf:"bytes,5,opt,name=output,proto3" json:"output,omitempty"`
	Timestamp            *timestamp.Timestamp `protobuf:"bytes,6,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Struct               *_struct.Struct      `protobuf:"bytes,7,opt,name=struct,proto3" json:"struct,omitempty"`
	Empty                *empty.Empty         `protobuf:"bytes,8,opt,name=empty,proto3" json:"empty,omitempty"`
	EmptyContainer       *EmptyContainer      `protobuf:"bytes,9,opt,name=empty_container,json=emptyContainer,proto3" json:"empty_container,omitempty"`
	Duration             *duration.Duration   `protobuf:"bytes,10,opt,name=duration,proto3" json:"duration,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *BuildEvent) Reset()         { *m = BuildEvent{} }
func (m *BuildEvent) String() string { return proto.CompactTextString(m) }
func (*BuildEvent) ProtoMessage()    {}
func (*BuildEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_aee82e0b1f3dea43, []int{4}
}

func (m *BuildEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BuildEvent.Unmarshal(m, b)
}
func (m *BuildEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BuildEvent.Marshal(b, m, deterministic)
}
func (m *BuildEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BuildEvent.Merge(m, src)
}
func (m *BuildEvent) XXX_Size() int {
	return xxx_messageInfo_BuildEvent.Size(m)
}
func (m *BuildEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_BuildEvent.DiscardUnknown(m)
}

var xxx_messageInfo_BuildEvent proto.InternalMessageInfo

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

func (m *BuildEvent) GetTimestamp() *timestamp.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *BuildEvent) GetStruct() *_struct.Struct {
	if m != nil {
		return m.Struct
	}
	return nil
}

func (m *BuildEvent) GetEmpty() *empty.Empty {
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

func (m *BuildEvent) GetDuration() *duration.Duration {
	if m != nil {
		return m.Duration
	}
	return nil
}

func init() {
	proto.RegisterEnum("testdata.Status", Status_name, Status_value)
	proto.RegisterType((*Property)(nil), "testdata.Property")
	proto.RegisterType((*Input)(nil), "testdata.Input")
	proto.RegisterType((*Output)(nil), "testdata.Output")
	proto.RegisterType((*EmptyContainer)(nil), "testdata.EmptyContainer")
	proto.RegisterType((*BuildEvent)(nil), "testdata.BuildEvent")
}

func init() {
	proto.RegisterFile("go.chromium.org/luci/tools/cmd/bqschemaupdater/testdata/event.proto", fileDescriptor_aee82e0b1f3dea43)
}

var fileDescriptor_aee82e0b1f3dea43 = []byte{
	// 492 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0x5b, 0x6b, 0xdb, 0x4c,
	0x10, 0x86, 0x3f, 0xc7, 0xb6, 0x2c, 0x8d, 0xc1, 0x31, 0x7b, 0xf1, 0x75, 0xe3, 0x9e, 0x8c, 0xa1,
	0x60, 0x4a, 0x2b, 0x81, 0x4b, 0xa1, 0xd0, 0x03, 0xa4, 0xae, 0x0a, 0x2e, 0x85, 0x94, 0x55, 0x73,
	0x6d, 0xd6, 0xd2, 0xd6, 0x51, 0x91, 0xb4, 0xea, 0x1e, 0x02, 0xf9, 0x95, 0xfd, 0x4b, 0x45, 0x23,
	0x29, 0x4a, 0xec, 0x9b, 0xd2, 0xbb, 0x9d, 0x79, 0x9f, 0x99, 0xdd, 0x79, 0x47, 0x82, 0xf5, 0x5e,
	0xfa, 0xf1, 0x95, 0x92, 0x79, 0x6a, 0x73, 0x5f, 0xaa, 0x7d, 0x90, 0xd9, 0x38, 0x0d, 0x8c, 0x94,
	0x99, 0x0e, 0xe2, 0x3c, 0x09, 0x76, 0xbf, 0x74, 0x7c, 0x25, 0x72, 0x6e, 0xcb, 0x84, 0x1b, 0xa1,
	0x02, 0x23, 0xb4, 0x49, 0xb8, 0xe1, 0x81, 0xb8, 0x16, 0x85, 0xf1, 0x4b, 0x25, 0x8d, 0x24, 0x6e,
	0x9b, 0x9d, 0x3d, 0xd9, 0x4b, 0xb9, 0xcf, 0x44, 0x80, 0xf9, 0x9d, 0xfd, 0x11, 0x24, 0x56, 0x71,
	0x93, 0xca, 0xa2, 0x26, 0x67, 0x0f, 0x0f, 0x75, 0x91, 0x97, 0xe6, 0xa6, 0x11, 0x1f, 0x1d, 0x8a,
	0xda, 0x28, 0x1b, 0x37, 0x97, 0xcc, 0x9e, 0x1e, 0xaa, 0x26, 0xcd, 0x85, 0x36, 0x3c, 0x2f, 0x6b,
	0x60, 0xf1, 0x1e, 0xdc, 0x6f, 0x4a, 0x96, 0x42, 0x99, 0x1b, 0x42, 0x60, 0x50, 0xf0, 0x5c, 0xd0,
	0xde, 0xbc, 0xb7, 0xf4, 0x18, 0x9e, 0xc9, 0x63, 0x80, 0x6b, 0x9e, 0x59, 0xb1, 0xfd, 0xa9, 0x65,
	0x41, 0x4f, 0x50, 0xf1, 0x30, 0xf3, 0x45, 0xcb, 0x62, 0xf1, 0x16, 0x86, 0x9b, 0xa2, 0xb4, 0x86,
	0xac, 0x00, 0xca, 0xba, 0x4f, 0x2a, 0x34, 0x1d, 0xcc, 0xfb, 0xcb, 0xf1, 0x8a, 0xf8, 0xed, 0x88,
	0x7e, 0x7b, 0x07, 0xbb, 0x43, 0x2d, 0xde, 0x81, 0x73, 0x61, 0xcd, 0xbf, 0x56, 0x7f, 0x80, 0x49,
	0x58, 0xf9, 0xb0, 0x96, 0x85, 0xe1, 0x69, 0x21, 0x14, 0x79, 0x01, 0x43, 0x74, 0x06, 0x07, 0x18,
	0xaf, 0xfe, 0xf7, 0xeb, 0xe1, 0xfd, 0x76, 0x78, 0x1f, 0x79, 0x56, 0x43, 0x8b, 0xdf, 0x7d, 0x80,
	0x8f, 0x36, 0xcd, 0x92, 0xb0, 0x5a, 0x0a, 0x39, 0x03, 0x77, 0x57, 0x45, 0xdb, 0x34, 0x69, 0x0c,
	0x18, 0x61, 0xbc, 0x49, 0x08, 0x85, 0xfa, 0x28, 0x54, 0x63, 0x40, 0x1b, 0x92, 0x25, 0x38, 0xda,
	0x70, 0x63, 0x35, 0xed, 0xcf, 0x7b, 0xcb, 0xc9, 0x6a, 0xda, 0xbd, 0x39, 0xc2, 0x3c, 0x6b, 0x74,
	0xf2, 0x0c, 0x86, 0x69, 0x65, 0x14, 0x1d, 0xe0, 0xdb, 0x4e, 0x3b, 0x10, 0xfd, 0x63, 0xb5, 0x5a,
	0x35, 0x94, 0x68, 0x09, 0x1d, 0x22, 0x77, 0xa7, 0x61, 0x6d, 0x15, 0x6b, 0x74, 0xf2, 0x06, 0xbc,
	0xdb, 0x5d, 0x52, 0x07, 0xe1, 0xd9, 0xd1, 0xc0, 0xdf, 0x5b, 0x82, 0x75, 0x30, 0x09, 0xaa, 0x47,
	0x57, 0xdf, 0x08, 0x1d, 0x61, 0xd9, 0x83, 0xa3, 0xb2, 0x08, 0x65, 0xd6, 0x60, 0x9d, 0xaf, 0xee,
	0x5f, 0xf8, 0x4a, 0xce, 0xe1, 0x14, 0x0f, 0xdb, 0xb8, 0x5d, 0x0c, 0xf5, 0xb0, 0x8e, 0x76, 0xb3,
	0xdc, 0x5f, 0x1c, 0x9b, 0x88, 0xfb, 0x8b, 0x7c, 0x0d, 0x6e, 0xfb, 0x0b, 0x50, 0xc0, 0xda, 0xb3,
	0xa3, 0x3b, 0x3f, 0x35, 0x00, 0xbb, 0x45, 0x9f, 0xbf, 0x04, 0xa7, 0x76, 0x9d, 0x8c, 0x61, 0x14,
	0x5d, 0xae, 0xd7, 0x61, 0x14, 0x4d, 0xff, 0xab, 0x82, 0xcf, 0xe7, 0x9b, 0xaf, 0x97, 0x2c, 0x9c,
	0xf6, 0x88, 0x07, 0xc3, 0x90, 0xb1, 0x0b, 0x36, 0x3d, 0xd9, 0x39, 0xd8, 0xeb, 0xd5, 0x9f, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x8f, 0x2b, 0xbf, 0x1a, 0xce, 0x03, 0x00, 0x00,
}
