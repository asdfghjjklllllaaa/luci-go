// Code generated by protoc-gen-go. DO NOT EDIT.
// source: go.chromium.org/luci/machine-db/api/crimson/v1/platforms.proto

package crimson

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// A platform in the database.
type Platform struct {
	// The name of this platform. Uniquely identifies this platform.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// A description of this platform.
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	// The manufacturer of this platform.
	Manufacturer         string   `protobuf:"bytes,3,opt,name=manufacturer,proto3" json:"manufacturer,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Platform) Reset()         { *m = Platform{} }
func (m *Platform) String() string { return proto.CompactTextString(m) }
func (*Platform) ProtoMessage()    {}
func (*Platform) Descriptor() ([]byte, []int) {
	return fileDescriptor_platforms_91eb230287ccad6a, []int{0}
}
func (m *Platform) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Platform.Unmarshal(m, b)
}
func (m *Platform) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Platform.Marshal(b, m, deterministic)
}
func (dst *Platform) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Platform.Merge(dst, src)
}
func (m *Platform) XXX_Size() int {
	return xxx_messageInfo_Platform.Size(m)
}
func (m *Platform) XXX_DiscardUnknown() {
	xxx_messageInfo_Platform.DiscardUnknown(m)
}

var xxx_messageInfo_Platform proto.InternalMessageInfo

func (m *Platform) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Platform) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Platform) GetManufacturer() string {
	if m != nil {
		return m.Manufacturer
	}
	return ""
}

// A request to list platforms in the database.
type ListPlatformsRequest struct {
	// The names of platforms to retrieve.
	Names []string `protobuf:"bytes,1,rep,name=names,proto3" json:"names,omitempty"`
	// The manufacturers to filter retrieved platforms on.
	Manufacturers        []string `protobuf:"bytes,2,rep,name=manufacturers,proto3" json:"manufacturers,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListPlatformsRequest) Reset()         { *m = ListPlatformsRequest{} }
func (m *ListPlatformsRequest) String() string { return proto.CompactTextString(m) }
func (*ListPlatformsRequest) ProtoMessage()    {}
func (*ListPlatformsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_platforms_91eb230287ccad6a, []int{1}
}
func (m *ListPlatformsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListPlatformsRequest.Unmarshal(m, b)
}
func (m *ListPlatformsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListPlatformsRequest.Marshal(b, m, deterministic)
}
func (dst *ListPlatformsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListPlatformsRequest.Merge(dst, src)
}
func (m *ListPlatformsRequest) XXX_Size() int {
	return xxx_messageInfo_ListPlatformsRequest.Size(m)
}
func (m *ListPlatformsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListPlatformsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListPlatformsRequest proto.InternalMessageInfo

func (m *ListPlatformsRequest) GetNames() []string {
	if m != nil {
		return m.Names
	}
	return nil
}

func (m *ListPlatformsRequest) GetManufacturers() []string {
	if m != nil {
		return m.Manufacturers
	}
	return nil
}

// A response containing a list of platforms in the database.
type ListPlatformsResponse struct {
	// The platforms matching the request.
	Platforms            []*Platform `protobuf:"bytes,1,rep,name=platforms,proto3" json:"platforms,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *ListPlatformsResponse) Reset()         { *m = ListPlatformsResponse{} }
func (m *ListPlatformsResponse) String() string { return proto.CompactTextString(m) }
func (*ListPlatformsResponse) ProtoMessage()    {}
func (*ListPlatformsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_platforms_91eb230287ccad6a, []int{2}
}
func (m *ListPlatformsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListPlatformsResponse.Unmarshal(m, b)
}
func (m *ListPlatformsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListPlatformsResponse.Marshal(b, m, deterministic)
}
func (dst *ListPlatformsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListPlatformsResponse.Merge(dst, src)
}
func (m *ListPlatformsResponse) XXX_Size() int {
	return xxx_messageInfo_ListPlatformsResponse.Size(m)
}
func (m *ListPlatformsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListPlatformsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListPlatformsResponse proto.InternalMessageInfo

func (m *ListPlatformsResponse) GetPlatforms() []*Platform {
	if m != nil {
		return m.Platforms
	}
	return nil
}

func init() {
	proto.RegisterType((*Platform)(nil), "crimson.Platform")
	proto.RegisterType((*ListPlatformsRequest)(nil), "crimson.ListPlatformsRequest")
	proto.RegisterType((*ListPlatformsResponse)(nil), "crimson.ListPlatformsResponse")
}

func init() {
	proto.RegisterFile("go.chromium.org/luci/machine-db/api/crimson/v1/platforms.proto", fileDescriptor_platforms_91eb230287ccad6a)
}

var fileDescriptor_platforms_91eb230287ccad6a = []byte{
	// 229 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0x4f, 0x4f, 0x03, 0x21,
	0x14, 0xc4, 0xb3, 0xad, 0xff, 0xf6, 0x55, 0x0f, 0x92, 0x9a, 0x70, 0xdc, 0x10, 0x0f, 0xbd, 0x08,
	0x51, 0xef, 0x9e, 0x3d, 0x78, 0x30, 0x7c, 0x03, 0xca, 0xd2, 0x96, 0xa4, 0xf0, 0x90, 0x07, 0x7e,
	0x7e, 0x23, 0x6e, 0xb5, 0xf5, 0x06, 0xf3, 0x1b, 0x66, 0x32, 0xc0, 0xcb, 0x16, 0xa5, 0xdd, 0x65,
	0x0c, 0xbe, 0x06, 0x89, 0x79, 0xab, 0xf6, 0xd5, 0x7a, 0x15, 0x8c, 0xdd, 0xf9, 0xe8, 0x1e, 0xc6,
	0xb5, 0x32, 0xc9, 0x2b, 0x9b, 0x7d, 0x20, 0x8c, 0xea, 0xf3, 0x51, 0xa5, 0xbd, 0x29, 0x1b, 0xcc,
	0x81, 0x64, 0xca, 0x58, 0x90, 0x5d, 0x4e, 0x4c, 0x8c, 0x70, 0xf5, 0x3e, 0x31, 0xc6, 0xe0, 0x2c,
	0x9a, 0xe0, 0x78, 0x37, 0x74, 0xab, 0x5e, 0xb7, 0x33, 0x1b, 0x60, 0x31, 0x3a, 0xb2, 0xd9, 0xa7,
	0xe2, 0x31, 0xf2, 0x59, 0x43, 0xc7, 0x12, 0x13, 0x70, 0x1d, 0x4c, 0xac, 0x1b, 0x63, 0x4b, 0xcd,
	0x2e, 0xf3, 0x79, 0xb3, 0x9c, 0x68, 0x42, 0xc3, 0xf2, 0xcd, 0x53, 0x39, 0x34, 0x91, 0x76, 0x1f,
	0xd5, 0x51, 0x61, 0x4b, 0x38, 0xff, 0x6e, 0x21, 0xde, 0x0d, 0xf3, 0x55, 0xaf, 0x7f, 0x2e, 0xec,
	0x1e, 0x6e, 0x8e, 0x5f, 0x13, 0x9f, 0x35, 0x7a, 0x2a, 0x8a, 0x57, 0xb8, 0xfb, 0x97, 0x49, 0x09,
	0x23, 0x39, 0xa6, 0xa0, 0xff, 0x9d, 0xdb, 0x82, 0x17, 0x4f, 0xb7, 0x72, 0xda, 0x2b, 0x0f, 0x76,
	0xfd, 0xe7, 0x59, 0x5f, 0xb4, 0x3f, 0x79, 0xfe, 0x0a, 0x00, 0x00, 0xff, 0xff, 0x2f, 0x5f, 0x53,
	0xf1, 0x55, 0x01, 0x00, 0x00,
}
