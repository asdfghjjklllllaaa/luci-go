// Code generated by protoc-gen-go. DO NOT EDIT.
// source: go.chromium.org/luci/machine-db/api/crimson/v1/hosts.proto

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

// A request to delete a physical or virtual host from the database.
type DeleteHostRequest struct {
	// The name of the host to delete.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The VLAN the host belongs to.
	Vlan                 int64    `protobuf:"varint,2,opt,name=vlan,proto3" json:"vlan,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteHostRequest) Reset()         { *m = DeleteHostRequest{} }
func (m *DeleteHostRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteHostRequest) ProtoMessage()    {}
func (*DeleteHostRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_hosts_77bd50d2491322df, []int{0}
}
func (m *DeleteHostRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteHostRequest.Unmarshal(m, b)
}
func (m *DeleteHostRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteHostRequest.Marshal(b, m, deterministic)
}
func (dst *DeleteHostRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteHostRequest.Merge(dst, src)
}
func (m *DeleteHostRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteHostRequest.Size(m)
}
func (m *DeleteHostRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteHostRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteHostRequest proto.InternalMessageInfo

func (m *DeleteHostRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DeleteHostRequest) GetVlan() int64 {
	if m != nil {
		return m.Vlan
	}
	return 0
}

func init() {
	proto.RegisterType((*DeleteHostRequest)(nil), "crimson.DeleteHostRequest")
}

func init() {
	proto.RegisterFile("go.chromium.org/luci/machine-db/api/crimson/v1/hosts.proto", fileDescriptor_hosts_77bd50d2491322df)
}

var fileDescriptor_hosts_77bd50d2491322df = []byte{
	// 144 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x2c, 0xcc, 0xb1, 0x0e, 0xc2, 0x20,
	0x10, 0x80, 0xe1, 0x54, 0x8d, 0x46, 0x36, 0x99, 0x3a, 0x36, 0x4e, 0x5d, 0xe4, 0x62, 0xdc, 0x74,
	0x75, 0x70, 0xe6, 0x0d, 0x28, 0x5e, 0x0a, 0x09, 0x70, 0x15, 0x8e, 0x3e, 0xbf, 0x91, 0xb8, 0xfd,
	0xf9, 0x86, 0x5f, 0xdc, 0x67, 0x52, 0xd6, 0x65, 0x8a, 0xbe, 0x46, 0x45, 0x79, 0x86, 0x50, 0xad,
	0x87, 0x68, 0xac, 0xf3, 0x09, 0x2f, 0xef, 0x09, 0xcc, 0xe2, 0xc1, 0x66, 0x1f, 0x0b, 0x25, 0x58,
	0xaf, 0xe0, 0xa8, 0x70, 0x51, 0x4b, 0x26, 0x26, 0x79, 0xf8, 0xfb, 0xf9, 0x21, 0x4e, 0x4f, 0x0c,
	0xc8, 0xf8, 0xa2, 0xc2, 0x1a, 0x3f, 0x15, 0x0b, 0x4b, 0x29, 0x76, 0xc9, 0x44, 0xec, 0xbb, 0xa1,
	0x1b, 0x8f, 0xba, 0xf5, 0xcf, 0xd6, 0x60, 0x52, 0xbf, 0x19, 0xba, 0x71, 0xab, 0x5b, 0x4f, 0xfb,
	0x36, 0xbb, 0x7d, 0x03, 0x00, 0x00, 0xff, 0xff, 0xfc, 0x6c, 0x71, 0x17, 0x8a, 0x00, 0x00, 0x00,
}
