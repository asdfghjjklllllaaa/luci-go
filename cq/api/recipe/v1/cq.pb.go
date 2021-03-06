// Code generated by protoc-gen-go. DO NOT EDIT.
// source: go.chromium.org/luci/cq/api/recipe/v1/cq.proto

package recipe

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

// Input provides CQ metadata for CQ-triggered tryjob.
type Input struct {
	// If active is false, CQ isn't active for the current build.
	Active bool `protobuf:"varint,1,opt,name=active,proto3" json:"active,omitempty"`
	// If false, CQ would try to submit CL(s) if all other checks pass.
	// If true, CQ won't try to submit.
	DryRun bool `protobuf:"varint,2,opt,name=dry_run,json=dryRun,proto3" json:"dry_run,omitempty"`
	// If true, CQ will not take this build into account while deciding whether CL
	// is good or not. See also `experiment_percentage` of CQ's config file.
	Experimental bool `protobuf:"varint,3,opt,name=experimental,proto3" json:"experimental,omitempty"`
	// If true, CQ triggered this build directly, otherwise typically indicates
	// a child build triggered by a CQ triggered one (possibly indirectly).
	//
	// Can be spoofed. *DO NOT USE FOR SECURITY CHECKS.*
	//
	// One possible use is to distinguish which builds must be cancelled manually,
	// and which (top_level=True) CQ would cancel itself.
	TopLevel             bool     `protobuf:"varint,4,opt,name=top_level,json=topLevel,proto3" json:"top_level,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Input) Reset()         { *m = Input{} }
func (m *Input) String() string { return proto.CompactTextString(m) }
func (*Input) ProtoMessage()    {}
func (*Input) Descriptor() ([]byte, []int) {
	return fileDescriptor_5310ea7d0dc1a356, []int{0}
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

func (m *Input) GetActive() bool {
	if m != nil {
		return m.Active
	}
	return false
}

func (m *Input) GetDryRun() bool {
	if m != nil {
		return m.DryRun
	}
	return false
}

func (m *Input) GetExperimental() bool {
	if m != nil {
		return m.Experimental
	}
	return false
}

func (m *Input) GetTopLevel() bool {
	if m != nil {
		return m.TopLevel
	}
	return false
}

func init() {
	proto.RegisterType((*Input)(nil), "cq.recipe.Input")
}

func init() {
	proto.RegisterFile("go.chromium.org/luci/cq/api/recipe/v1/cq.proto", fileDescriptor_5310ea7d0dc1a356)
}

var fileDescriptor_5310ea7d0dc1a356 = []byte{
	// 187 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x8e, 0x3f, 0xab, 0xc2, 0x30,
	0x14, 0x47, 0xe9, 0xfb, 0xd3, 0xd7, 0x86, 0x37, 0x65, 0xd0, 0x80, 0x8b, 0x74, 0x72, 0x90, 0x04,
	0x71, 0x74, 0x73, 0x13, 0x9c, 0x3a, 0xba, 0x94, 0x9a, 0x5e, 0x6a, 0x20, 0x4d, 0x6e, 0x2f, 0x49,
	0xb1, 0xdf, 0x5e, 0x6c, 0x5d, 0xdc, 0xdc, 0xee, 0xef, 0x9c, 0x3b, 0x1c, 0x26, 0x5b, 0x2f, 0xf5,
	0x8d, 0x7c, 0x67, 0x62, 0x27, 0x3d, 0xb5, 0xca, 0x46, 0x6d, 0x94, 0xee, 0x55, 0x8d, 0x46, 0x11,
	0x68, 0x83, 0xa0, 0x86, 0x9d, 0xd2, 0xbd, 0x44, 0xf2, 0xc1, 0xf3, 0x5c, 0xf7, 0x72, 0xc6, 0xc5,
	0xc8, 0x7e, 0x4f, 0x0e, 0x63, 0xe0, 0x0b, 0x96, 0xd6, 0x3a, 0x98, 0x01, 0x44, 0xb2, 0x4e, 0x36,
	0x59, 0xf9, 0x5a, 0x7c, 0xc9, 0xfe, 0x1a, 0x1a, 0x2b, 0x8a, 0x4e, 0x7c, 0xcd, 0xa2, 0xa1, 0xb1,
	0x8c, 0x8e, 0x17, 0xec, 0x1f, 0xee, 0x08, 0x64, 0x3a, 0x70, 0xa1, 0xb6, 0xe2, 0x7b, 0xb2, 0x6f,
	0x8c, 0xaf, 0x58, 0x1e, 0x3c, 0x56, 0x16, 0x06, 0xb0, 0xe2, 0x67, 0x7a, 0xc8, 0x82, 0xc7, 0xf3,
	0x73, 0x1f, 0xe5, 0x65, 0xfb, 0x51, 0xf7, 0x61, 0xbe, 0xae, 0xe9, 0x14, 0xbf, 0x7f, 0x04, 0x00,
	0x00, 0xff, 0xff, 0xad, 0xf5, 0xb4, 0x35, 0xee, 0x00, 0x00, 0x00,
}
