// Code generated by protoc-gen-go. DO NOT EDIT.
// source: go.chromium.org/luci/machine-db/api/crimson/v1/dracs.proto

package crimson

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// A DRAC in the database.
type DRAC struct {
	// The name of this DRAC on the network. With VLAN, uniquely identifies this DRAC.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// The machine this DRAC belongs to. Uniquely identifies this DRAC.
	Machine string `protobuf:"bytes,2,opt,name=machine" json:"machine,omitempty"`
	// The IPv4 address associated with this DRAC.
	Ipv4 string `protobuf:"bytes,3,opt,name=ipv4" json:"ipv4,omitempty"`
	// The VLAN this DRAC belongs to.
	// When creating a DRAC, omit this field. It will be inferred from the IPv4 address.
	Vlan int64 `protobuf:"varint,4,opt,name=vlan" json:"vlan,omitempty"`
	// The MAC address associated with this DRAC.
	MacAddress string `protobuf:"bytes,5,opt,name=mac_address,json=macAddress" json:"mac_address,omitempty"`
	// The switch this DRAC is connected to.
	Switch string `protobuf:"bytes,6,opt,name=switch" json:"switch,omitempty"`
	// The switchport this DRAC is connected to.
	Switchport int32 `protobuf:"varint,7,opt,name=switchport" json:"switchport,omitempty"`
}

func (m *DRAC) Reset()                    { *m = DRAC{} }
func (m *DRAC) String() string            { return proto.CompactTextString(m) }
func (*DRAC) ProtoMessage()               {}
func (*DRAC) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *DRAC) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DRAC) GetMachine() string {
	if m != nil {
		return m.Machine
	}
	return ""
}

func (m *DRAC) GetIpv4() string {
	if m != nil {
		return m.Ipv4
	}
	return ""
}

func (m *DRAC) GetVlan() int64 {
	if m != nil {
		return m.Vlan
	}
	return 0
}

func (m *DRAC) GetMacAddress() string {
	if m != nil {
		return m.MacAddress
	}
	return ""
}

func (m *DRAC) GetSwitch() string {
	if m != nil {
		return m.Switch
	}
	return ""
}

func (m *DRAC) GetSwitchport() int32 {
	if m != nil {
		return m.Switchport
	}
	return 0
}

// A request to create a new DRAC in the database.
type CreateDRACRequest struct {
	// The DRAC to create in the database.
	Drac *DRAC `protobuf:"bytes,1,opt,name=drac" json:"drac,omitempty"`
}

func (m *CreateDRACRequest) Reset()                    { *m = CreateDRACRequest{} }
func (m *CreateDRACRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateDRACRequest) ProtoMessage()               {}
func (*CreateDRACRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

func (m *CreateDRACRequest) GetDrac() *DRAC {
	if m != nil {
		return m.Drac
	}
	return nil
}

// A request to list DRACs in the database.
type ListDRACsRequest struct {
	// The names of DRACs to get.
	Names []string `protobuf:"bytes,1,rep,name=names" json:"names,omitempty"`
	// The machines to filter retrieved DRACs on.
	Machines []string `protobuf:"bytes,2,rep,name=machines" json:"machines,omitempty"`
	// The IPv4 addresses to filter returned DRACs on.
	Ipv4S []string `protobuf:"bytes,3,rep,name=ipv4s" json:"ipv4s,omitempty"`
	// The VLANs to filter returned DRACs on.
	Vlans []int64 `protobuf:"varint,4,rep,packed,name=vlans" json:"vlans,omitempty"`
	// The MAC addresses to filter retrieved DRACs on.
	MacAddresses []string `protobuf:"bytes,5,rep,name=mac_addresses,json=macAddresses" json:"mac_addresses,omitempty"`
	// The switches to filter retrieved DRACs on.
	Switches []string `protobuf:"bytes,6,rep,name=switches" json:"switches,omitempty"`
}

func (m *ListDRACsRequest) Reset()                    { *m = ListDRACsRequest{} }
func (m *ListDRACsRequest) String() string            { return proto.CompactTextString(m) }
func (*ListDRACsRequest) ProtoMessage()               {}
func (*ListDRACsRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{2} }

func (m *ListDRACsRequest) GetNames() []string {
	if m != nil {
		return m.Names
	}
	return nil
}

func (m *ListDRACsRequest) GetMachines() []string {
	if m != nil {
		return m.Machines
	}
	return nil
}

func (m *ListDRACsRequest) GetIpv4S() []string {
	if m != nil {
		return m.Ipv4S
	}
	return nil
}

func (m *ListDRACsRequest) GetVlans() []int64 {
	if m != nil {
		return m.Vlans
	}
	return nil
}

func (m *ListDRACsRequest) GetMacAddresses() []string {
	if m != nil {
		return m.MacAddresses
	}
	return nil
}

func (m *ListDRACsRequest) GetSwitches() []string {
	if m != nil {
		return m.Switches
	}
	return nil
}

// A response containing a list of DRACs in the database.
type ListDRACsResponse struct {
	// The DRACs matching this request.
	Dracs []*DRAC `protobuf:"bytes,1,rep,name=dracs" json:"dracs,omitempty"`
}

func (m *ListDRACsResponse) Reset()                    { *m = ListDRACsResponse{} }
func (m *ListDRACsResponse) String() string            { return proto.CompactTextString(m) }
func (*ListDRACsResponse) ProtoMessage()               {}
func (*ListDRACsResponse) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{3} }

func (m *ListDRACsResponse) GetDracs() []*DRAC {
	if m != nil {
		return m.Dracs
	}
	return nil
}

func init() {
	proto.RegisterType((*DRAC)(nil), "crimson.DRAC")
	proto.RegisterType((*CreateDRACRequest)(nil), "crimson.CreateDRACRequest")
	proto.RegisterType((*ListDRACsRequest)(nil), "crimson.ListDRACsRequest")
	proto.RegisterType((*ListDRACsResponse)(nil), "crimson.ListDRACsResponse")
}

func init() {
	proto.RegisterFile("go.chromium.org/luci/machine-db/api/crimson/v1/dracs.proto", fileDescriptor2)
}

var fileDescriptor2 = []byte{
	// 345 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x91, 0xcb, 0x4e, 0xeb, 0x30,
	0x10, 0x86, 0x95, 0xe6, 0xd2, 0xd3, 0xe9, 0xa9, 0x74, 0x6a, 0x1d, 0x21, 0x8b, 0x05, 0x84, 0x74,
	0x93, 0x0d, 0x89, 0xb8, 0x08, 0x21, 0x76, 0x55, 0x59, 0xb2, 0xf2, 0x0b, 0x20, 0xd7, 0xb1, 0x5a,
	0x4b, 0x75, 0x1c, 0x3c, 0x69, 0x79, 0x2e, 0x16, 0xbc, 0x1f, 0xb2, 0x9d, 0x42, 0x25, 0x76, 0xf3,
	0xff, 0x33, 0x63, 0xcf, 0x37, 0x03, 0x4f, 0x1b, 0x53, 0x89, 0xad, 0x35, 0x5a, 0xed, 0x75, 0x65,
	0xec, 0xa6, 0xde, 0xed, 0x85, 0xaa, 0x35, 0x17, 0x5b, 0xd5, 0xca, 0xeb, 0x66, 0x5d, 0xf3, 0x4e,
	0xd5, 0xc2, 0x2a, 0x8d, 0xa6, 0xad, 0x0f, 0x37, 0x75, 0x63, 0xb9, 0xc0, 0xaa, 0xb3, 0xa6, 0x37,
	0x64, 0x3c, 0xf8, 0xc5, 0x67, 0x04, 0xc9, 0x33, 0x5b, 0xae, 0x08, 0x81, 0xa4, 0xe5, 0x5a, 0xd2,
	0x28, 0x8f, 0xca, 0x09, 0xf3, 0x31, 0xa1, 0x30, 0x1e, 0x9e, 0xa3, 0x23, 0x6f, 0x1f, 0xa5, 0xab,
	0x56, 0xdd, 0xe1, 0x9e, 0xc6, 0xa1, 0xda, 0xc5, 0xce, 0x3b, 0xec, 0x78, 0x4b, 0x93, 0x3c, 0x2a,
	0x63, 0xe6, 0x63, 0x72, 0x09, 0x53, 0xcd, 0xc5, 0x2b, 0x6f, 0x1a, 0x2b, 0x11, 0x69, 0xea, 0xcb,
	0x41, 0x73, 0xb1, 0x0c, 0x0e, 0x39, 0x83, 0x0c, 0xdf, 0x55, 0x2f, 0xb6, 0x34, 0xf3, 0xb9, 0x41,
	0x91, 0x0b, 0x80, 0x10, 0x75, 0xc6, 0xf6, 0x74, 0x9c, 0x47, 0x65, 0xca, 0x4e, 0x9c, 0xe2, 0x01,
	0xe6, 0x2b, 0x2b, 0x79, 0x2f, 0xdd, 0xf0, 0x4c, 0xbe, 0xed, 0x25, 0xf6, 0xe4, 0x0a, 0x12, 0x07,
	0xe9, 0x19, 0xa6, 0xb7, 0xb3, 0x6a, 0x80, 0xac, 0x7c, 0x8d, 0x4f, 0x15, 0x1f, 0x11, 0xfc, 0x7b,
	0x51, 0xd8, 0x3b, 0x0b, 0x8f, 0x7d, 0xff, 0x21, 0x75, 0xbc, 0x48, 0xa3, 0x3c, 0x2e, 0x27, 0x2c,
	0x08, 0x72, 0x0e, 0x7f, 0x06, 0x5c, 0xa4, 0x23, 0x9f, 0xf8, 0xd6, 0xae, 0xc3, 0x31, 0x23, 0x8d,
	0x43, 0x87, 0x17, 0xce, 0x75, 0xd4, 0x48, 0x93, 0x3c, 0x2e, 0x63, 0x16, 0x04, 0x59, 0xc0, 0xec,
	0x64, 0x07, 0xd2, 0x6d, 0xc1, 0xf5, 0xfc, 0xfd, 0xd9, 0x42, 0xf8, 0x2c, 0xd0, 0x49, 0xa4, 0x59,
	0xf8, 0xec, 0xa8, 0x8b, 0x47, 0x98, 0x9f, 0x8c, 0x8c, 0x9d, 0x69, 0x51, 0x92, 0x05, 0xa4, 0xfe,
	0xa0, 0x7e, 0xe6, 0x5f, 0xb0, 0x21, 0xb7, 0xce, 0xfc, 0xb5, 0xef, 0xbe, 0x02, 0x00, 0x00, 0xff,
	0xff, 0x12, 0x69, 0x49, 0x35, 0x2b, 0x02, 0x00, 0x00,
}
