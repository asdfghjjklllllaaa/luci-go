// Code generated by protoc-gen-go.
// source: service_account.proto
// DO NOT EDIT!

package tokenserver

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/luci/luci-go/common/proto/google"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// ServiceAccount describes a Cloud IAM Service Account.
//
// See https://cloud.google.com/iam/reference/rest/v1/projects.serviceAccounts#ServiceAccount
type ServiceAccount struct {
	ProjectId      string                     `protobuf:"bytes,1,opt,name=project_id,json=projectId" json:"project_id,omitempty"`
	UniqueId       string                     `protobuf:"bytes,2,opt,name=unique_id,json=uniqueId" json:"unique_id,omitempty"`
	Email          string                     `protobuf:"bytes,3,opt,name=email" json:"email,omitempty"`
	DisplayName    string                     `protobuf:"bytes,4,opt,name=display_name,json=displayName" json:"display_name,omitempty"`
	Oauth2ClientId string                     `protobuf:"bytes,5,opt,name=oauth2_client_id,json=oauth2ClientId" json:"oauth2_client_id,omitempty"`
	Fqdn           string                     `protobuf:"bytes,6,opt,name=fqdn" json:"fqdn,omitempty"`
	Registered     *google_protobuf.Timestamp `protobuf:"bytes,7,opt,name=registered" json:"registered,omitempty"`
}

func (m *ServiceAccount) Reset()                    { *m = ServiceAccount{} }
func (m *ServiceAccount) String() string            { return proto.CompactTextString(m) }
func (*ServiceAccount) ProtoMessage()               {}
func (*ServiceAccount) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *ServiceAccount) GetRegistered() *google_protobuf.Timestamp {
	if m != nil {
		return m.Registered
	}
	return nil
}

func init() {
	proto.RegisterType((*ServiceAccount)(nil), "tokenserver.ServiceAccount")
}

func init() { proto.RegisterFile("service_account.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 247 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x4c, 0xcf, 0xcb, 0x4b, 0x03, 0x31,
	0x10, 0x06, 0x70, 0x56, 0xdb, 0xea, 0xce, 0x4a, 0x91, 0xa0, 0x10, 0x2a, 0xe2, 0xe3, 0xd4, 0xd3,
	0x16, 0xea, 0xcd, 0x9b, 0x78, 0xf2, 0xe2, 0xa1, 0x7a, 0x5f, 0xd2, 0x64, 0xba, 0x46, 0x77, 0x93,
	0x6d, 0x1e, 0x82, 0x7f, 0xbd, 0xa6, 0x93, 0x16, 0xbc, 0x25, 0xdf, 0x6f, 0x98, 0xe1, 0x83, 0x4b,
	0x8f, 0xee, 0x5b, 0x4b, 0x6c, 0x84, 0x94, 0x36, 0x9a, 0x50, 0x0f, 0xce, 0x06, 0xcb, 0xaa, 0x60,
	0xbf, 0xd0, 0xec, 0x0c, 0xdd, 0xec, 0xa6, 0xb5, 0xb6, 0xed, 0x70, 0x41, 0xb4, 0x8e, 0x9b, 0x45,
	0xd0, 0x3d, 0xfa, 0x20, 0xfa, 0x21, 0x4f, 0xdf, 0xff, 0x16, 0x30, 0x7d, 0xcb, 0x7b, 0x9e, 0xf2,
	0x1a, 0x76, 0x0d, 0x90, 0xec, 0x13, 0x65, 0x68, 0xb4, 0xe2, 0xc5, 0x6d, 0x31, 0x2f, 0x57, 0xe5,
	0x3e, 0x79, 0x51, 0xec, 0x0a, 0xca, 0x68, 0xf4, 0x36, 0xe2, 0x4e, 0x8f, 0x48, 0x4f, 0x73, 0x90,
	0xf0, 0x02, 0xc6, 0xd8, 0x0b, 0xdd, 0xf1, 0x63, 0x82, 0xfc, 0x61, 0x77, 0x70, 0xa6, 0xb4, 0x1f,
	0x3a, 0xf1, 0xd3, 0x18, 0xd1, 0x23, 0x1f, 0x11, 0x56, 0xfb, 0xec, 0x35, 0x45, 0x6c, 0x0e, 0xe7,
	0x56, 0xc4, 0xf0, 0xb1, 0x6c, 0x64, 0xa7, 0xd1, 0xd0, 0xe9, 0x31, 0x8d, 0x4d, 0x73, 0xfe, 0x4c,
	0x71, 0x3a, 0xc1, 0x60, 0xb4, 0xd9, 0x2a, 0xc3, 0x27, 0xa4, 0xf4, 0x66, 0x8f, 0x00, 0x0e, 0x5b,
	0xed, 0x03, 0x3a, 0x54, 0xfc, 0x24, 0x49, 0xb5, 0x9c, 0xd5, 0xb9, 0x7b, 0x7d, 0xe8, 0x5e, 0xbf,
	0x1f, 0xba, 0xaf, 0xfe, 0x4d, 0xaf, 0x27, 0xe4, 0x0f, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xe9,
	0x06, 0x3e, 0x17, 0x4f, 0x01, 0x00, 0x00,
}
