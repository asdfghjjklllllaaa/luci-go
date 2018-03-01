// Code generated by protoc-gen-go. DO NOT EDIT.
// source: go.chromium.org/luci/machine-db/api/crimson/v1/switches.proto

package crimson

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import common "go.chromium.org/luci/machine-db/api/common/v1"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// A switch in the database.
type Switch struct {
	// The name of this switch. Uniquely identifies this switch.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// A description of this switch.
	Description string `protobuf:"bytes,2,opt,name=description" json:"description,omitempty"`
	// The number of ports this switch has.
	Ports int32 `protobuf:"varint,3,opt,name=ports" json:"ports,omitempty"`
	// The datacenter this switch belongs to.
	Datacenter string `protobuf:"bytes,4,opt,name=datacenter" json:"datacenter,omitempty"`
	// The rack this switch belongs to.
	Rack string `protobuf:"bytes,5,opt,name=rack" json:"rack,omitempty"`
	// The state of this switch.
	State common.State `protobuf:"varint,6,opt,name=state,enum=common.State" json:"state,omitempty"`
}

func (m *Switch) Reset()                    { *m = Switch{} }
func (m *Switch) String() string            { return proto.CompactTextString(m) }
func (*Switch) ProtoMessage()               {}
func (*Switch) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{0} }

func (m *Switch) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Switch) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Switch) GetPorts() int32 {
	if m != nil {
		return m.Ports
	}
	return 0
}

func (m *Switch) GetDatacenter() string {
	if m != nil {
		return m.Datacenter
	}
	return ""
}

func (m *Switch) GetRack() string {
	if m != nil {
		return m.Rack
	}
	return ""
}

func (m *Switch) GetState() common.State {
	if m != nil {
		return m.State
	}
	return common.State_STATE_UNSPECIFIED
}

// A request to list switches in the database.
type ListSwitchesRequest struct {
	// The names of switches to retrieve.
	Names []string `protobuf:"bytes,1,rep,name=names" json:"names,omitempty"`
	// The datacenters to filter retrieved switches on.
	Datacenters []string `protobuf:"bytes,2,rep,name=datacenters" json:"datacenters,omitempty"`
	// The racks to filter retrieved switches on.
	Racks []string `protobuf:"bytes,3,rep,name=racks" json:"racks,omitempty"`
}

func (m *ListSwitchesRequest) Reset()                    { *m = ListSwitchesRequest{} }
func (m *ListSwitchesRequest) String() string            { return proto.CompactTextString(m) }
func (*ListSwitchesRequest) ProtoMessage()               {}
func (*ListSwitchesRequest) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{1} }

func (m *ListSwitchesRequest) GetNames() []string {
	if m != nil {
		return m.Names
	}
	return nil
}

func (m *ListSwitchesRequest) GetDatacenters() []string {
	if m != nil {
		return m.Datacenters
	}
	return nil
}

func (m *ListSwitchesRequest) GetRacks() []string {
	if m != nil {
		return m.Racks
	}
	return nil
}

// A response containing a list of switches in the database.
type ListSwitchesResponse struct {
	// The switches matching the request.
	Switches []*Switch `protobuf:"bytes,1,rep,name=switches" json:"switches,omitempty"`
}

func (m *ListSwitchesResponse) Reset()                    { *m = ListSwitchesResponse{} }
func (m *ListSwitchesResponse) String() string            { return proto.CompactTextString(m) }
func (*ListSwitchesResponse) ProtoMessage()               {}
func (*ListSwitchesResponse) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{2} }

func (m *ListSwitchesResponse) GetSwitches() []*Switch {
	if m != nil {
		return m.Switches
	}
	return nil
}

func init() {
	proto.RegisterType((*Switch)(nil), "crimson.Switch")
	proto.RegisterType((*ListSwitchesRequest)(nil), "crimson.ListSwitchesRequest")
	proto.RegisterType((*ListSwitchesResponse)(nil), "crimson.ListSwitchesResponse")
}

func init() {
	proto.RegisterFile("go.chromium.org/luci/machine-db/api/crimson/v1/switches.proto", fileDescriptor11)
}

var fileDescriptor11 = []byte{
	// 294 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x51, 0xc1, 0x4e, 0xc4, 0x20,
	0x10, 0x0d, 0xbb, 0xdb, 0xea, 0xb2, 0x51, 0x13, 0xdc, 0x03, 0xf1, 0x60, 0x9a, 0x7a, 0x69, 0x62,
	0x84, 0xb8, 0xde, 0x4c, 0x3c, 0x79, 0xf5, 0x44, 0xbf, 0x80, 0xa5, 0x64, 0x4b, 0x14, 0xa8, 0x40,
	0xf5, 0xab, 0xfc, 0x47, 0x03, 0x74, 0x37, 0xf5, 0xe6, 0x6d, 0xe6, 0xbd, 0x99, 0xf7, 0x5e, 0x66,
	0xe0, 0xcb, 0xc1, 0x12, 0xd1, 0x3b, 0xab, 0xd5, 0xa8, 0x89, 0x75, 0x07, 0xfa, 0x31, 0x0a, 0x45,
	0x35, 0x17, 0xbd, 0x32, 0xf2, 0xa1, 0xdb, 0x53, 0x3e, 0x28, 0x2a, 0x9c, 0xd2, 0xde, 0x1a, 0xfa,
	0xf5, 0x48, 0xfd, 0xb7, 0x0a, 0xa2, 0x97, 0x9e, 0x0c, 0xce, 0x06, 0x8b, 0xce, 0x26, 0xea, 0xe6,
	0xf9, 0x5f, 0x3a, 0x56, 0xeb, 0x49, 0x26, 0xf0, 0x70, 0x14, 0xa9, 0x7f, 0x00, 0x2c, 0xdb, 0xa4,
	0x8b, 0x10, 0x5c, 0x19, 0xae, 0x25, 0x06, 0x15, 0x68, 0xd6, 0x2c, 0xd5, 0xa8, 0x82, 0x9b, 0x4e,
	0x7a, 0xe1, 0xd4, 0x10, 0x94, 0x35, 0x78, 0x91, 0xa8, 0x39, 0x84, 0xb6, 0xb0, 0x18, 0xac, 0x0b,
	0x1e, 0x2f, 0x2b, 0xd0, 0x14, 0x2c, 0x37, 0xe8, 0x16, 0xc2, 0x8e, 0x07, 0x2e, 0xa4, 0x09, 0xd2,
	0xe1, 0x55, 0x5a, 0x9b, 0x21, 0xd1, 0xcb, 0x71, 0xf1, 0x8e, 0x8b, 0xec, 0x15, 0x6b, 0x74, 0x07,
	0x8b, 0x14, 0x0d, 0x97, 0x15, 0x68, 0x2e, 0x77, 0x17, 0x24, 0x47, 0x26, 0x6d, 0x04, 0x59, 0xe6,
	0x6a, 0x01, 0xaf, 0xdf, 0x94, 0x0f, 0xed, 0x74, 0x0a, 0x26, 0x3f, 0x47, 0xe9, 0x43, 0x4c, 0x11,
	0xf3, 0x7a, 0x0c, 0xaa, 0x65, 0xb3, 0x66, 0xb9, 0x49, 0xe9, 0x4f, 0x9e, 0x1e, 0x2f, 0x12, 0x37,
	0x87, 0xe2, 0x5e, 0xf4, 0x8e, 0xe9, 0xd3, 0x5e, 0x6a, 0xea, 0x57, 0xb8, 0xfd, 0x6b, 0xe2, 0x07,
	0x6b, 0xbc, 0x44, 0xf7, 0xf0, 0xfc, 0xf8, 0x83, 0x64, 0xb4, 0xd9, 0x5d, 0x91, 0xe9, 0x09, 0x24,
	0x0f, 0xb3, 0xd3, 0xc0, 0xbe, 0x4c, 0x07, 0x7e, 0xfa, 0x0d, 0x00, 0x00, 0xff, 0xff, 0xd6, 0xf6,
	0x93, 0x8c, 0xe6, 0x01, 0x00, 0x00,
}
