// Code generated by protoc-gen-go. DO NOT EDIT.
// source: go.chromium.org/luci/common/bq/testdata/testmessage.proto

package testdata

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	empty "github.com/golang/protobuf/ptypes/empty"
	_struct "github.com/golang/protobuf/ptypes/struct"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type TestMessage_FOO int32

const (
	TestMessage_X TestMessage_FOO = 0
	TestMessage_Y TestMessage_FOO = 1
	TestMessage_Z TestMessage_FOO = 2
)

var TestMessage_FOO_name = map[int32]string{
	0: "X",
	1: "Y",
	2: "Z",
}

var TestMessage_FOO_value = map[string]int32{
	"X": 0,
	"Y": 1,
	"Z": 2,
}

func (x TestMessage_FOO) String() string {
	return proto.EnumName(TestMessage_FOO_name, int32(x))
}

func (TestMessage_FOO) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_bc9d446aa0b4e493, []int{0, 0}
}

type TestMessage struct {
	Name           string               `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Timestamp      *timestamp.Timestamp `protobuf:"bytes,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Nested         *NestedTestMessage   `protobuf:"bytes,3,opt,name=nested,proto3" json:"nested,omitempty"`
	RepeatedNested []*NestedTestMessage `protobuf:"bytes,4,rep,name=repeated_nested,json=repeatedNested,proto3" json:"repeated_nested,omitempty"`
	Struct         *_struct.Struct      `protobuf:"bytes,5,opt,name=struct,proto3" json:"struct,omitempty"`
	Foo            TestMessage_FOO      `protobuf:"varint,6,opt,name=foo,proto3,enum=testdata.TestMessage_FOO" json:"foo,omitempty"`
	FooRepeated    []TestMessage_FOO    `protobuf:"varint,7,rep,packed,name=foo_repeated,json=fooRepeated,proto3,enum=testdata.TestMessage_FOO" json:"foo_repeated,omitempty"`
	Empty          *empty.Empty         `protobuf:"bytes,8,opt,name=empty,proto3" json:"empty,omitempty"`
	Empties        []*empty.Empty       `protobuf:"bytes,9,rep,name=empties,proto3" json:"empties,omitempty"`
	Duration       *duration.Duration   `protobuf:"bytes,10,opt,name=duration,proto3" json:"duration,omitempty"`
	// Types that are valid to be assigned to OneOf:
	//	*TestMessage_First
	//	*TestMessage_Second
	OneOf                isTestMessage_OneOf `protobuf_oneof:"one_of"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *TestMessage) Reset()         { *m = TestMessage{} }
func (m *TestMessage) String() string { return proto.CompactTextString(m) }
func (*TestMessage) ProtoMessage()    {}
func (*TestMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_bc9d446aa0b4e493, []int{0}
}

func (m *TestMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestMessage.Unmarshal(m, b)
}
func (m *TestMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestMessage.Marshal(b, m, deterministic)
}
func (m *TestMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestMessage.Merge(m, src)
}
func (m *TestMessage) XXX_Size() int {
	return xxx_messageInfo_TestMessage.Size(m)
}
func (m *TestMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_TestMessage.DiscardUnknown(m)
}

var xxx_messageInfo_TestMessage proto.InternalMessageInfo

func (m *TestMessage) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TestMessage) GetTimestamp() *timestamp.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *TestMessage) GetNested() *NestedTestMessage {
	if m != nil {
		return m.Nested
	}
	return nil
}

func (m *TestMessage) GetRepeatedNested() []*NestedTestMessage {
	if m != nil {
		return m.RepeatedNested
	}
	return nil
}

func (m *TestMessage) GetStruct() *_struct.Struct {
	if m != nil {
		return m.Struct
	}
	return nil
}

func (m *TestMessage) GetFoo() TestMessage_FOO {
	if m != nil {
		return m.Foo
	}
	return TestMessage_X
}

func (m *TestMessage) GetFooRepeated() []TestMessage_FOO {
	if m != nil {
		return m.FooRepeated
	}
	return nil
}

func (m *TestMessage) GetEmpty() *empty.Empty {
	if m != nil {
		return m.Empty
	}
	return nil
}

func (m *TestMessage) GetEmpties() []*empty.Empty {
	if m != nil {
		return m.Empties
	}
	return nil
}

func (m *TestMessage) GetDuration() *duration.Duration {
	if m != nil {
		return m.Duration
	}
	return nil
}

type isTestMessage_OneOf interface {
	isTestMessage_OneOf()
}

type TestMessage_First struct {
	First *NestedTestMessage `protobuf:"bytes,11,opt,name=first,proto3,oneof"`
}

type TestMessage_Second struct {
	Second *NestedTestMessage `protobuf:"bytes,12,opt,name=second,proto3,oneof"`
}

func (*TestMessage_First) isTestMessage_OneOf() {}

func (*TestMessage_Second) isTestMessage_OneOf() {}

func (m *TestMessage) GetOneOf() isTestMessage_OneOf {
	if m != nil {
		return m.OneOf
	}
	return nil
}

func (m *TestMessage) GetFirst() *NestedTestMessage {
	if x, ok := m.GetOneOf().(*TestMessage_First); ok {
		return x.First
	}
	return nil
}

func (m *TestMessage) GetSecond() *NestedTestMessage {
	if x, ok := m.GetOneOf().(*TestMessage_Second); ok {
		return x.Second
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*TestMessage) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*TestMessage_First)(nil),
		(*TestMessage_Second)(nil),
	}
}

type NestedTestMessage struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NestedTestMessage) Reset()         { *m = NestedTestMessage{} }
func (m *NestedTestMessage) String() string { return proto.CompactTextString(m) }
func (*NestedTestMessage) ProtoMessage()    {}
func (*NestedTestMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_bc9d446aa0b4e493, []int{1}
}

func (m *NestedTestMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NestedTestMessage.Unmarshal(m, b)
}
func (m *NestedTestMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NestedTestMessage.Marshal(b, m, deterministic)
}
func (m *NestedTestMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NestedTestMessage.Merge(m, src)
}
func (m *NestedTestMessage) XXX_Size() int {
	return xxx_messageInfo_NestedTestMessage.Size(m)
}
func (m *NestedTestMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_NestedTestMessage.DiscardUnknown(m)
}

var xxx_messageInfo_NestedTestMessage proto.InternalMessageInfo

func (m *NestedTestMessage) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterEnum("testdata.TestMessage_FOO", TestMessage_FOO_name, TestMessage_FOO_value)
	proto.RegisterType((*TestMessage)(nil), "testdata.TestMessage")
	proto.RegisterType((*NestedTestMessage)(nil), "testdata.NestedTestMessage")
}

func init() {
	proto.RegisterFile("go.chromium.org/luci/common/bq/testdata/testmessage.proto", fileDescriptor_bc9d446aa0b4e493)
}

var fileDescriptor_bc9d446aa0b4e493 = []byte{
	// 430 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x41, 0x6f, 0xd3, 0x40,
	0x10, 0x85, 0xeb, 0xba, 0x76, 0x93, 0x49, 0x55, 0xc2, 0x1e, 0x60, 0x9b, 0x22, 0xb0, 0x72, 0x21,
	0x12, 0x68, 0x8d, 0x1a, 0x55, 0x02, 0x89, 0x13, 0x2a, 0x15, 0x17, 0x88, 0xb4, 0xf4, 0x00, 0x5c,
	0x22, 0xc7, 0x1e, 0x1b, 0x4b, 0x5d, 0x4f, 0xf0, 0xae, 0x0f, 0xfc, 0x16, 0xfe, 0x2c, 0xf2, 0xda,
	0xdb, 0x56, 0x31, 0x50, 0x4e, 0x1e, 0xed, 0xfb, 0xde, 0xbc, 0x19, 0xef, 0xc2, 0x9b, 0x82, 0x44,
	0xfa, 0xbd, 0x26, 0x55, 0x36, 0x4a, 0x50, 0x5d, 0xc4, 0xd7, 0x4d, 0x5a, 0xc6, 0x29, 0x29, 0x45,
	0x55, 0xbc, 0xf9, 0x11, 0x1b, 0xd4, 0x26, 0x4b, 0x4c, 0x62, 0x0b, 0x85, 0x5a, 0x27, 0x05, 0x8a,
	0x6d, 0x4d, 0x86, 0xd8, 0xc8, 0x69, 0xb3, 0xa7, 0x05, 0x51, 0x71, 0x8d, 0xb1, 0x3d, 0xdf, 0x34,
	0x79, 0x9c, 0x35, 0x75, 0x62, 0x4a, 0xaa, 0x3a, 0x72, 0x76, 0xba, 0xab, 0xa3, 0xda, 0x9a, 0x9f,
	0xbd, 0xf8, 0x6c, 0x57, 0x34, 0xa5, 0x42, 0x6d, 0x12, 0xb5, 0xed, 0x81, 0x27, 0xbb, 0x80, 0x36,
	0x75, 0x93, 0x9a, 0x4e, 0x9d, 0xff, 0x0a, 0x60, 0x72, 0x85, 0xda, 0x7c, 0xec, 0x66, 0x63, 0x0c,
	0x0e, 0xaa, 0x44, 0x21, 0xf7, 0x22, 0x6f, 0x31, 0x96, 0xb6, 0x66, 0xaf, 0x61, 0x7c, 0xd3, 0x94,
	0xef, 0x47, 0xde, 0x62, 0x72, 0x36, 0x13, 0x5d, 0x57, 0xe1, 0xba, 0x8a, 0x2b, 0x47, 0xc8, 0x5b,
	0x98, 0x2d, 0x21, 0xac, 0x50, 0x1b, 0xcc, 0xb8, 0x6f, 0x6d, 0xa7, 0xc2, 0x2d, 0x2d, 0x3e, 0xd9,
	0xf3, 0x3b, 0xd1, 0xb2, 0x47, 0xd9, 0x05, 0x3c, 0xa8, 0x71, 0x8b, 0x89, 0xc1, 0x6c, 0xdd, 0xbb,
	0x0f, 0x22, 0xff, 0x3e, 0xf7, 0xb1, 0xf3, 0x74, 0x12, 0x8b, 0x21, 0xec, 0x16, 0xe5, 0x81, 0x8d,
	0x7e, 0x3c, 0x98, 0xf8, 0xb3, 0x95, 0x65, 0x8f, 0xb1, 0x17, 0xe0, 0xe7, 0x44, 0x3c, 0x8c, 0xbc,
	0xc5, 0xf1, 0xd9, 0xc9, 0x6d, 0xd4, 0x9d, 0x10, 0x71, 0xb9, 0x5a, 0xc9, 0x96, 0x62, 0x6f, 0xe1,
	0x28, 0x27, 0x5a, 0xbb, 0x4c, 0x7e, 0x18, 0xf9, 0xff, 0x76, 0x4d, 0x72, 0x22, 0xd9, 0xd3, 0xec,
	0x25, 0x04, 0xf6, 0x0a, 0xf9, 0xc8, 0x8e, 0xf6, 0x68, 0x30, 0xda, 0xfb, 0x56, 0x95, 0x1d, 0xc4,
	0x5e, 0xc1, 0x61, 0x5b, 0x94, 0xa8, 0xf9, 0xd8, 0xfe, 0x87, 0xbf, 0xf1, 0x0e, 0x63, 0xe7, 0x30,
	0x72, 0x4f, 0x88, 0x83, 0x8d, 0x38, 0x19, 0x58, 0x2e, 0x7a, 0x40, 0xde, 0xa0, 0x6c, 0x09, 0x41,
	0x5e, 0xd6, 0xda, 0xf0, 0xc9, 0xbd, 0x97, 0xf5, 0x61, 0x4f, 0x76, 0x2c, 0x3b, 0x87, 0x50, 0x63,
	0x4a, 0x55, 0xc6, 0x8f, 0xfe, 0xc7, 0xd5, 0xc3, 0xf3, 0x19, 0xf8, 0x97, 0xab, 0x15, 0x0b, 0xc0,
	0xfb, 0x32, 0xdd, 0x6b, 0x3f, 0x5f, 0xa7, 0x5e, 0xfb, 0xf9, 0x36, 0xdd, 0x7f, 0x37, 0x82, 0x90,
	0x2a, 0x5c, 0x53, 0x3e, 0x7f, 0x0e, 0x0f, 0x07, 0x4d, 0xfe, 0xf4, 0x44, 0x37, 0xa1, 0xdd, 0x6b,
	0xf9, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x83, 0x77, 0x81, 0x83, 0x90, 0x03, 0x00, 0x00,
}
