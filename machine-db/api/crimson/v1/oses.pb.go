// Code generated by protoc-gen-go. DO NOT EDIT.
// source: go.chromium.org/luci/machine-db/api/crimson/v1/oses.proto

package crimson

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// OSesRequest is a request to retrieve operating systems.
type OSesRequest struct {
	// The names of operating systems to retrieve.
	Names []string `protobuf:"bytes,1,rep,name=names" json:"names,omitempty"`
}

func (m *OSesRequest) Reset()                    { *m = OSesRequest{} }
func (m *OSesRequest) String() string            { return proto.CompactTextString(m) }
func (*OSesRequest) ProtoMessage()               {}
func (*OSesRequest) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

func (m *OSesRequest) GetNames() []string {
	if m != nil {
		return m.Names
	}
	return nil
}

// OS describes an operating system.
type OS struct {
	// The name of this operating system. Uniquely identifies this operating system.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// A description of this operating system.
	Description string `protobuf:"bytes,2,opt,name=description" json:"description,omitempty"`
}

func (m *OS) Reset()                    { *m = OS{} }
func (m *OS) String() string            { return proto.CompactTextString(m) }
func (*OS) ProtoMessage()               {}
func (*OS) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{1} }

func (m *OS) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *OS) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

// OSesResponse is a response to a request to retrieve operating systems.
type OSesResponse struct {
	// The operating systems matching the request.
	Oses []*OS `protobuf:"bytes,1,rep,name=oses" json:"oses,omitempty"`
}

func (m *OSesResponse) Reset()                    { *m = OSesResponse{} }
func (m *OSesResponse) String() string            { return proto.CompactTextString(m) }
func (*OSesResponse) ProtoMessage()               {}
func (*OSesResponse) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{2} }

func (m *OSesResponse) GetOses() []*OS {
	if m != nil {
		return m.Oses
	}
	return nil
}

func init() {
	proto.RegisterType((*OSesRequest)(nil), "crimson.OSesRequest")
	proto.RegisterType((*OS)(nil), "crimson.OS")
	proto.RegisterType((*OSesResponse)(nil), "crimson.OSesResponse")
}

func init() {
	proto.RegisterFile("go.chromium.org/luci/machine-db/api/crimson/v1/oses.proto", fileDescriptor4)
}

var fileDescriptor4 = []byte{
	// 189 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x8e, 0xbf, 0x4e, 0x85, 0x30,
	0x14, 0x87, 0xc3, 0xf5, 0xaa, 0xe1, 0xd4, 0xa9, 0x71, 0x60, 0x93, 0xe0, 0xc2, 0x62, 0x1b, 0x75,
	0xd2, 0x97, 0x20, 0x29, 0x4f, 0x00, 0xe5, 0x04, 0x9a, 0xd8, 0x9e, 0xda, 0x03, 0x3e, 0xbf, 0xa1,
	0x30, 0xdc, 0xed, 0xf7, 0x67, 0xf8, 0x3e, 0xf8, 0x9a, 0x49, 0xd9, 0x25, 0x91, 0x77, 0x9b, 0x57,
	0x94, 0x66, 0xfd, 0xb3, 0x59, 0xa7, 0xfd, 0x60, 0x17, 0x17, 0xf0, 0x6d, 0x1a, 0xf5, 0x10, 0x9d,
	0xb6, 0xc9, 0x79, 0xa6, 0xa0, 0xff, 0xde, 0x35, 0x31, 0xb2, 0x8a, 0x89, 0x56, 0x92, 0x8f, 0xe7,
	0xdc, 0xbc, 0x82, 0xe8, 0x7a, 0x64, 0x83, 0xbf, 0x1b, 0xf2, 0x2a, 0x9f, 0xe1, 0x3e, 0x0c, 0x1e,
	0xb9, 0x2a, 0xea, 0xbb, 0xb6, 0x34, 0x47, 0x69, 0xbe, 0xe1, 0xd2, 0xf5, 0x52, 0xc2, 0x75, 0xaf,
	0x55, 0x51, 0x17, 0x6d, 0x69, 0x72, 0x96, 0x35, 0x88, 0x09, 0xd9, 0x26, 0x17, 0x57, 0x47, 0xa1,
	0xba, 0xe4, 0xeb, 0x76, 0x6a, 0x34, 0x3c, 0x1d, 0x00, 0x8e, 0x14, 0x18, 0xe5, 0x0b, 0x5c, 0x77,
	0x8f, 0x0c, 0x10, 0x1f, 0x42, 0x9d, 0x22, 0xaa, 0xeb, 0x4d, 0x3e, 0xc6, 0x87, 0x6c, 0xf8, 0xf9,
	0x1f, 0x00, 0x00, 0xff, 0xff, 0x78, 0x5e, 0x6f, 0x1d, 0xde, 0x00, 0x00, 0x00,
}