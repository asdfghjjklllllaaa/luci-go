// Code generated by protoc-gen-go. DO NOT EDIT.
// source: go.chromium.org/luci/machine-db/api/crimson/v1/physical_hosts.proto

package crimson

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// A physical host in the database.
type PhysicalHost struct {
	// The name of this host on the network. With VLAN ID, uniquely identifies this host.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// The VLAN this host belongs to. With hostname, uniquely identifies this host.
	Vlan int64 `protobuf:"varint,2,opt,name=vlan" json:"vlan,omitempty"`
	// The machine backing this host.
	Machine string `protobuf:"bytes,3,opt,name=machine" json:"machine,omitempty"`
	// The operating system backing this host.
	Os string `protobuf:"bytes,4,opt,name=os" json:"os,omitempty"`
	// The number of VMs which can be deployed on this host.
	VmSlots int32 `protobuf:"varint,5,opt,name=vm_slots,json=vmSlots" json:"vm_slots,omitempty"`
	// A description of this host.
	Description string `protobuf:"bytes,6,opt,name=description" json:"description,omitempty"`
	// The deployment ticket associated with this host.
	DeploymentTicket string `protobuf:"bytes,7,opt,name=deployment_ticket,json=deploymentTicket" json:"deployment_ticket,omitempty"`
	// The IPv4 address associated with this host.
	Ipv4 string `protobuf:"bytes,8,opt,name=ipv4" json:"ipv4,omitempty"`
}

func (m *PhysicalHost) Reset()                    { *m = PhysicalHost{} }
func (m *PhysicalHost) String() string            { return proto.CompactTextString(m) }
func (*PhysicalHost) ProtoMessage()               {}
func (*PhysicalHost) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{0} }

func (m *PhysicalHost) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *PhysicalHost) GetVlan() int64 {
	if m != nil {
		return m.Vlan
	}
	return 0
}

func (m *PhysicalHost) GetMachine() string {
	if m != nil {
		return m.Machine
	}
	return ""
}

func (m *PhysicalHost) GetOs() string {
	if m != nil {
		return m.Os
	}
	return ""
}

func (m *PhysicalHost) GetVmSlots() int32 {
	if m != nil {
		return m.VmSlots
	}
	return 0
}

func (m *PhysicalHost) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *PhysicalHost) GetDeploymentTicket() string {
	if m != nil {
		return m.DeploymentTicket
	}
	return ""
}

func (m *PhysicalHost) GetIpv4() string {
	if m != nil {
		return m.Ipv4
	}
	return ""
}

// CreatePhysicalHostRequest is a request to create a new physical host in the database.
type CreatePhysicalHostRequest struct {
	// The host to create in the database.
	Host *PhysicalHost `protobuf:"bytes,1,opt,name=host" json:"host,omitempty"`
}

func (m *CreatePhysicalHostRequest) Reset()                    { *m = CreatePhysicalHostRequest{} }
func (m *CreatePhysicalHostRequest) String() string            { return proto.CompactTextString(m) }
func (*CreatePhysicalHostRequest) ProtoMessage()               {}
func (*CreatePhysicalHostRequest) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{1} }

func (m *CreatePhysicalHostRequest) GetHost() *PhysicalHost {
	if m != nil {
		return m.Host
	}
	return nil
}

// ListPhysicalHostsRequest is a request to list physical hosts in the database.
type ListPhysicalHostsRequest struct {
	// The names of hosts to get.
	Names []string `protobuf:"bytes,1,rep,name=names" json:"names,omitempty"`
	// The VLANs to filter retrieved hosts on.
	Vlans []int64 `protobuf:"varint,2,rep,packed,name=vlans" json:"vlans,omitempty"`
}

func (m *ListPhysicalHostsRequest) Reset()                    { *m = ListPhysicalHostsRequest{} }
func (m *ListPhysicalHostsRequest) String() string            { return proto.CompactTextString(m) }
func (*ListPhysicalHostsRequest) ProtoMessage()               {}
func (*ListPhysicalHostsRequest) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{2} }

func (m *ListPhysicalHostsRequest) GetNames() []string {
	if m != nil {
		return m.Names
	}
	return nil
}

func (m *ListPhysicalHostsRequest) GetVlans() []int64 {
	if m != nil {
		return m.Vlans
	}
	return nil
}

// ListPhysicalHostsResponse is a response to a request to list physical hosts.
type ListPhysicalHostsResponse struct {
	// The hosts matching this request.
	Hosts []*PhysicalHost `protobuf:"bytes,1,rep,name=hosts" json:"hosts,omitempty"`
}

func (m *ListPhysicalHostsResponse) Reset()                    { *m = ListPhysicalHostsResponse{} }
func (m *ListPhysicalHostsResponse) String() string            { return proto.CompactTextString(m) }
func (*ListPhysicalHostsResponse) ProtoMessage()               {}
func (*ListPhysicalHostsResponse) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{3} }

func (m *ListPhysicalHostsResponse) GetHosts() []*PhysicalHost {
	if m != nil {
		return m.Hosts
	}
	return nil
}

func init() {
	proto.RegisterType((*PhysicalHost)(nil), "crimson.PhysicalHost")
	proto.RegisterType((*CreatePhysicalHostRequest)(nil), "crimson.CreatePhysicalHostRequest")
	proto.RegisterType((*ListPhysicalHostsRequest)(nil), "crimson.ListPhysicalHostsRequest")
	proto.RegisterType((*ListPhysicalHostsResponse)(nil), "crimson.ListPhysicalHostsResponse")
}

func init() {
	proto.RegisterFile("go.chromium.org/luci/machine-db/api/crimson/v1/physical_hosts.proto", fileDescriptor7)
}

var fileDescriptor7 = []byte{
	// 334 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0xcf, 0x6a, 0x83, 0x40,
	0x10, 0xc6, 0x51, 0x63, 0x4c, 0x26, 0xa5, 0xb4, 0x4b, 0x0b, 0x9b, 0x9b, 0x78, 0xb2, 0x84, 0x2a,
	0xfd, 0xf3, 0x06, 0x81, 0x90, 0x43, 0x0f, 0xc5, 0xf6, 0x1e, 0x8c, 0x59, 0xe2, 0x52, 0xd7, 0xb1,
	0xce, 0x46, 0xc8, 0xcb, 0xf6, 0x59, 0xca, 0xae, 0x86, 0xa6, 0xd0, 0xde, 0xe6, 0xfb, 0xf6, 0xe7,
	0x38, 0xf3, 0x0d, 0x2c, 0xf7, 0x98, 0x14, 0x65, 0x8b, 0x4a, 0x1e, 0x54, 0x82, 0xed, 0x3e, 0xad,
	0x0e, 0x85, 0x4c, 0x55, 0x5e, 0x94, 0xb2, 0x16, 0xf7, 0xbb, 0x6d, 0x9a, 0x37, 0x32, 0x2d, 0x5a,
	0xa9, 0x08, 0xeb, 0xb4, 0x7b, 0x48, 0x9b, 0xf2, 0x48, 0xb2, 0xc8, 0xab, 0x4d, 0x89, 0xa4, 0x29,
	0x69, 0x5a, 0xd4, 0xc8, 0x82, 0x01, 0x88, 0xbe, 0x1c, 0xb8, 0x78, 0x1d, 0x88, 0x35, 0x92, 0x66,
	0x0c, 0x46, 0x75, 0xae, 0x04, 0x77, 0x42, 0x27, 0x9e, 0x66, 0xb6, 0x36, 0x5e, 0x57, 0xe5, 0x35,
	0x77, 0x43, 0x27, 0xf6, 0x32, 0x5b, 0x33, 0x0e, 0xc1, 0xf0, 0x4f, 0xee, 0x59, 0xf4, 0x24, 0xd9,
	0x25, 0xb8, 0x48, 0x7c, 0x64, 0x4d, 0x17, 0x89, 0xcd, 0x61, 0xd2, 0xa9, 0x0d, 0x55, 0xa8, 0x89,
	0xfb, 0xa1, 0x13, 0xfb, 0x59, 0xd0, 0xa9, 0x37, 0x23, 0x59, 0x08, 0xb3, 0x9d, 0xa0, 0xa2, 0x95,
	0x8d, 0x96, 0x58, 0xf3, 0xb1, 0xfd, 0xe6, 0xdc, 0x62, 0x0b, 0xb8, 0xde, 0x89, 0xa6, 0xc2, 0xa3,
	0x12, 0xb5, 0xde, 0x68, 0x59, 0x7c, 0x08, 0xcd, 0x03, 0xcb, 0x5d, 0xfd, 0x3c, 0xbc, 0x5b, 0xdf,
	0xcc, 0x29, 0x9b, 0xee, 0x99, 0x4f, 0xfa, 0xd9, 0x4d, 0x1d, 0xad, 0x60, 0xbe, 0x6c, 0x45, 0xae,
	0xc5, 0xf9, 0x96, 0x99, 0xf8, 0x3c, 0x08, 0xd2, 0xec, 0x0e, 0x46, 0x26, 0x15, 0xbb, 0xec, 0xec,
	0xf1, 0x36, 0x19, 0x52, 0x49, 0x7e, 0xb1, 0x16, 0x89, 0x56, 0xc0, 0x5f, 0x24, 0xe9, 0xf3, 0x17,
	0x3a, 0xb5, 0xb9, 0x01, 0xdf, 0xe4, 0x44, 0xdc, 0x09, 0xbd, 0x78, 0x9a, 0xf5, 0xc2, 0xb8, 0x26,
	0x29, 0xe2, 0x6e, 0xe8, 0xc5, 0x5e, 0xd6, 0x8b, 0x68, 0x0d, 0xf3, 0x3f, 0xfa, 0x50, 0x83, 0x35,
	0x09, 0xb6, 0x00, 0xdf, 0x5e, 0xc9, 0x36, 0xfa, 0x77, 0xa0, 0x9e, 0xd9, 0x8e, 0xed, 0x29, 0x9f,
	0xbe, 0x03, 0x00, 0x00, 0xff, 0xff, 0x55, 0x78, 0x6b, 0x18, 0x11, 0x02, 0x00, 0x00,
}
