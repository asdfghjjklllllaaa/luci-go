// Code generated by protoc-gen-go. DO NOT EDIT.
// source: go.chromium.org/luci/machine-db/api/crimson/v1/machines.proto

package crimson

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf1 "google.golang.org/genproto/protobuf/field_mask"
import common "go.chromium.org/luci/machine-db/api/common/v1"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// A machine in the database.
type Machine struct {
	// The name of this machine. Uniquely identifies this machine.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// The type of platform this machine is.
	Platform string `protobuf:"bytes,2,opt,name=platform" json:"platform,omitempty"`
	// The rack this machine belongs to.
	Rack string `protobuf:"bytes,3,opt,name=rack" json:"rack,omitempty"`
	// A description of this machine.
	Description string `protobuf:"bytes,4,opt,name=description" json:"description,omitempty"`
	// The asset tag associated with this machine.
	AssetTag string `protobuf:"bytes,5,opt,name=asset_tag,json=assetTag" json:"asset_tag,omitempty"`
	// The service tag associated with this machine.
	ServiceTag string `protobuf:"bytes,6,opt,name=service_tag,json=serviceTag" json:"service_tag,omitempty"`
	// The deployment ticket associated with this machine.
	DeploymentTicket string `protobuf:"bytes,7,opt,name=deployment_ticket,json=deploymentTicket" json:"deployment_ticket,omitempty"`
	// The state of this machine.
	State common.State `protobuf:"varint,8,opt,name=state,enum=common.State" json:"state,omitempty"`
	// The datacenter this machine belongs to.
	// When creating a machine, omit this field. It will be inferred from the rack.
	Datacenter string `protobuf:"bytes,9,opt,name=datacenter" json:"datacenter,omitempty"`
}

func (m *Machine) Reset()                    { *m = Machine{} }
func (m *Machine) String() string            { return proto.CompactTextString(m) }
func (*Machine) ProtoMessage()               {}
func (*Machine) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

func (m *Machine) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Machine) GetPlatform() string {
	if m != nil {
		return m.Platform
	}
	return ""
}

func (m *Machine) GetRack() string {
	if m != nil {
		return m.Rack
	}
	return ""
}

func (m *Machine) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Machine) GetAssetTag() string {
	if m != nil {
		return m.AssetTag
	}
	return ""
}

func (m *Machine) GetServiceTag() string {
	if m != nil {
		return m.ServiceTag
	}
	return ""
}

func (m *Machine) GetDeploymentTicket() string {
	if m != nil {
		return m.DeploymentTicket
	}
	return ""
}

func (m *Machine) GetState() common.State {
	if m != nil {
		return m.State
	}
	return common.State_STATE_UNSPECIFIED
}

func (m *Machine) GetDatacenter() string {
	if m != nil {
		return m.Datacenter
	}
	return ""
}

// A request to create a new machine in the database.
type CreateMachineRequest struct {
	// The machine to create in the database.
	Machine *Machine `protobuf:"bytes,1,opt,name=machine" json:"machine,omitempty"`
}

func (m *CreateMachineRequest) Reset()                    { *m = CreateMachineRequest{} }
func (m *CreateMachineRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateMachineRequest) ProtoMessage()               {}
func (*CreateMachineRequest) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{1} }

func (m *CreateMachineRequest) GetMachine() *Machine {
	if m != nil {
		return m.Machine
	}
	return nil
}

// A request to delete a machine from the database.
type DeleteMachineRequest struct {
	// The name of the machine to delete.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *DeleteMachineRequest) Reset()                    { *m = DeleteMachineRequest{} }
func (m *DeleteMachineRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteMachineRequest) ProtoMessage()               {}
func (*DeleteMachineRequest) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{2} }

func (m *DeleteMachineRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// A request to list machines in the database.
type ListMachinesRequest struct {
	// The names of machines to get.
	Names []string `protobuf:"bytes,1,rep,name=names" json:"names,omitempty"`
	// The platforms to filter returned machines on.
	Platforms []string `protobuf:"bytes,2,rep,name=platforms" json:"platforms,omitempty"`
	// The racks to filter returned machines on.
	Racks []string `protobuf:"bytes,3,rep,name=racks" json:"racks,omitempty"`
	// The states to filter returned machines on.
	States []common.State `protobuf:"varint,4,rep,packed,name=states,enum=common.State" json:"states,omitempty"`
	// The datacenters to filter returned machines on.
	Datacenters []string `protobuf:"bytes,5,rep,name=datacenters" json:"datacenters,omitempty"`
}

func (m *ListMachinesRequest) Reset()                    { *m = ListMachinesRequest{} }
func (m *ListMachinesRequest) String() string            { return proto.CompactTextString(m) }
func (*ListMachinesRequest) ProtoMessage()               {}
func (*ListMachinesRequest) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{3} }

func (m *ListMachinesRequest) GetNames() []string {
	if m != nil {
		return m.Names
	}
	return nil
}

func (m *ListMachinesRequest) GetPlatforms() []string {
	if m != nil {
		return m.Platforms
	}
	return nil
}

func (m *ListMachinesRequest) GetRacks() []string {
	if m != nil {
		return m.Racks
	}
	return nil
}

func (m *ListMachinesRequest) GetStates() []common.State {
	if m != nil {
		return m.States
	}
	return nil
}

func (m *ListMachinesRequest) GetDatacenters() []string {
	if m != nil {
		return m.Datacenters
	}
	return nil
}

// A response containing a list of machines in the database.
type ListMachinesResponse struct {
	// The machines matching this request.
	Machines []*Machine `protobuf:"bytes,1,rep,name=machines" json:"machines,omitempty"`
}

func (m *ListMachinesResponse) Reset()                    { *m = ListMachinesResponse{} }
func (m *ListMachinesResponse) String() string            { return proto.CompactTextString(m) }
func (*ListMachinesResponse) ProtoMessage()               {}
func (*ListMachinesResponse) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{4} }

func (m *ListMachinesResponse) GetMachines() []*Machine {
	if m != nil {
		return m.Machines
	}
	return nil
}

// A request to update a machine in the database.
type UpdateMachineRequest struct {
	// The machine to update in the database.
	Machine *Machine `protobuf:"bytes,1,opt,name=machine" json:"machine,omitempty"`
	// The fields to update in the machine.
	UpdateMask *google_protobuf1.FieldMask `protobuf:"bytes,2,opt,name=update_mask,json=updateMask" json:"update_mask,omitempty"`
}

func (m *UpdateMachineRequest) Reset()                    { *m = UpdateMachineRequest{} }
func (m *UpdateMachineRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateMachineRequest) ProtoMessage()               {}
func (*UpdateMachineRequest) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{5} }

func (m *UpdateMachineRequest) GetMachine() *Machine {
	if m != nil {
		return m.Machine
	}
	return nil
}

func (m *UpdateMachineRequest) GetUpdateMask() *google_protobuf1.FieldMask {
	if m != nil {
		return m.UpdateMask
	}
	return nil
}

func init() {
	proto.RegisterType((*Machine)(nil), "crimson.Machine")
	proto.RegisterType((*CreateMachineRequest)(nil), "crimson.CreateMachineRequest")
	proto.RegisterType((*DeleteMachineRequest)(nil), "crimson.DeleteMachineRequest")
	proto.RegisterType((*ListMachinesRequest)(nil), "crimson.ListMachinesRequest")
	proto.RegisterType((*ListMachinesResponse)(nil), "crimson.ListMachinesResponse")
	proto.RegisterType((*UpdateMachineRequest)(nil), "crimson.UpdateMachineRequest")
}

func init() {
	proto.RegisterFile("go.chromium.org/luci/machine-db/api/crimson/v1/machines.proto", fileDescriptor4)
}

var fileDescriptor4 = []byte{
	// 478 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x53, 0x4f, 0x8f, 0xd3, 0x3e,
	0x10, 0x55, 0xff, 0xb7, 0x53, 0xfd, 0x7e, 0x5a, 0x4c, 0x0f, 0x56, 0x41, 0x10, 0x05, 0x21, 0x55,
	0x0b, 0x38, 0xa2, 0xdc, 0x40, 0x5c, 0x60, 0xc5, 0x89, 0xbd, 0x94, 0xe5, 0x5c, 0xb9, 0xc9, 0x34,
	0x6b, 0x35, 0x8e, 0x83, 0xed, 0xac, 0xc4, 0x89, 0x8f, 0xc2, 0x95, 0x8f, 0x89, 0x32, 0x4e, 0xda,
	0x45, 0xbb, 0x07, 0x24, 0x6e, 0x9e, 0xf7, 0x9e, 0x67, 0xec, 0x37, 0x33, 0xf0, 0x3e, 0x37, 0x22,
	0xbd, 0xb6, 0x46, 0xab, 0x5a, 0x0b, 0x63, 0xf3, 0xa4, 0xa8, 0x53, 0x95, 0x68, 0x99, 0x5e, 0xab,
	0x12, 0x5f, 0x65, 0xbb, 0x44, 0x56, 0x2a, 0x49, 0xad, 0xd2, 0xce, 0x94, 0xc9, 0xcd, 0xeb, 0x8e,
	0x71, 0xa2, 0xb2, 0xc6, 0x1b, 0x36, 0x69, 0xa9, 0x65, 0x94, 0x1b, 0x93, 0x17, 0x98, 0x10, 0xbc,
	0xab, 0xf7, 0xc9, 0x5e, 0x61, 0x91, 0x6d, 0xb5, 0x74, 0x87, 0x20, 0x5d, 0xbe, 0xfd, 0xab, 0x4a,
	0x46, 0xeb, 0x50, 0xc8, 0x79, 0xe9, 0xbb, 0x32, 0xf1, 0xcf, 0x3e, 0x4c, 0x2e, 0x83, 0x92, 0x31,
	0x18, 0x96, 0x52, 0x23, 0xef, 0x45, 0xbd, 0xd5, 0x6c, 0x43, 0x67, 0xb6, 0x84, 0x69, 0x55, 0x48,
	0xbf, 0x37, 0x56, 0xf3, 0x3e, 0xe1, 0xc7, 0xb8, 0xd1, 0x5b, 0x99, 0x1e, 0xf8, 0x20, 0xe8, 0x9b,
	0x33, 0x8b, 0x60, 0x9e, 0xa1, 0x4b, 0xad, 0xaa, 0xbc, 0x32, 0x25, 0x1f, 0x12, 0x75, 0x1b, 0x62,
	0x8f, 0x60, 0x26, 0x9d, 0x43, 0xbf, 0xf5, 0x32, 0xe7, 0xa3, 0x90, 0x92, 0x80, 0x2b, 0x99, 0xb3,
	0xa7, 0x30, 0x77, 0x68, 0x6f, 0x54, 0x8a, 0x44, 0x8f, 0x89, 0x86, 0x16, 0x6a, 0x04, 0x2f, 0xe0,
	0x41, 0x86, 0x55, 0x61, 0xbe, 0x6b, 0x2c, 0xfd, 0xd6, 0xab, 0xf4, 0x80, 0x9e, 0x4f, 0x48, 0x76,
	0x76, 0x22, 0xae, 0x08, 0x67, 0xcf, 0x60, 0x44, 0x9f, 0xe5, 0xd3, 0xa8, 0xb7, 0xfa, 0x7f, 0xfd,
	0x9f, 0x08, 0x26, 0x88, 0x2f, 0x0d, 0xb8, 0x09, 0x1c, 0x7b, 0x02, 0x90, 0x49, 0x2f, 0x53, 0x2c,
	0x3d, 0x5a, 0x3e, 0x0b, 0x15, 0x4f, 0x48, 0xfc, 0x01, 0x16, 0x1f, 0x2d, 0x4a, 0x8f, 0xad, 0x4d,
	0x1b, 0xfc, 0x56, 0xa3, 0xf3, 0xec, 0x1c, 0x26, 0xad, 0xc5, 0x64, 0xd8, 0x7c, 0x7d, 0x26, 0xda,
	0x96, 0x89, 0x4e, 0xd9, 0x09, 0xe2, 0x73, 0x58, 0x5c, 0x60, 0x81, 0x77, 0x72, 0xdc, 0xe3, 0x78,
	0xfc, 0xab, 0x07, 0x0f, 0x3f, 0x2b, 0xe7, 0x5b, 0xa9, 0xeb, 0xb4, 0x0b, 0x18, 0x35, 0xbc, 0xe3,
	0xbd, 0x68, 0xb0, 0x9a, 0x6d, 0x42, 0xc0, 0x1e, 0xc3, 0xac, 0xeb, 0x87, 0xe3, 0x7d, 0x62, 0x4e,
	0x40, 0x73, 0xa7, 0xe9, 0x8a, 0xe3, 0x83, 0x70, 0x87, 0x02, 0xf6, 0x1c, 0xc6, 0x61, 0x06, 0xf8,
	0x30, 0x1a, 0xdc, 0xf5, 0xa5, 0x25, 0xa9, 0x95, 0x47, 0x1b, 0x1c, 0x1f, 0x51, 0x8a, 0xdb, 0x50,
	0x7c, 0x01, 0x8b, 0x3f, 0x5f, 0xea, 0x2a, 0x53, 0x3a, 0x64, 0x2f, 0x61, 0xda, 0x4d, 0x33, 0xbd,
	0xf6, 0x3e, 0x6f, 0x8e, 0x8a, 0xf8, 0x07, 0x2c, 0xbe, 0x56, 0xd9, 0x3f, 0x19, 0xcc, 0xde, 0xc1,
	0xbc, 0xa6, 0x1c, 0xb4, 0x17, 0x34, 0xa9, 0xf3, 0xf5, 0x52, 0x84, 0xd5, 0x11, 0xdd, 0xea, 0x88,
	0x4f, 0xcd, 0xea, 0x5c, 0x4a, 0x77, 0xd8, 0x40, 0xdd, 0x96, 0x74, 0x87, 0xdd, 0x98, 0xf8, 0x37,
	0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0xa4, 0xa1, 0x69, 0x45, 0xb2, 0x03, 0x00, 0x00,
}
