// Code generated by protoc-gen-go. DO NOT EDIT.
// source: go.chromium.org/luci/milo/api/config/settings.proto

package config

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Settings represents the format for the global (service) config for Milo.
type Settings struct {
	Buildbot    *Settings_Buildbot    `protobuf:"bytes,1,opt,name=buildbot" json:"buildbot,omitempty"`
	Buildbucket *Settings_Buildbucket `protobuf:"bytes,2,opt,name=buildbucket" json:"buildbucket,omitempty"`
	Swarming    *Settings_Swarming    `protobuf:"bytes,3,opt,name=swarming" json:"swarming,omitempty"`
}

func (m *Settings) Reset()                    { *m = Settings{} }
func (m *Settings) String() string            { return proto.CompactTextString(m) }
func (*Settings) ProtoMessage()               {}
func (*Settings) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *Settings) GetBuildbot() *Settings_Buildbot {
	if m != nil {
		return m.Buildbot
	}
	return nil
}

func (m *Settings) GetBuildbucket() *Settings_Buildbucket {
	if m != nil {
		return m.Buildbucket
	}
	return nil
}

func (m *Settings) GetSwarming() *Settings_Swarming {
	if m != nil {
		return m.Swarming
	}
	return nil
}

type Settings_Buildbot struct {
	// internal_reader is the infra-auth group that is allowed to read internal
	// buildbot data.
	InternalReader string `protobuf:"bytes,1,opt,name=internal_reader,json=internalReader" json:"internal_reader,omitempty"`
	// public_subscription is the name of the pubsub topic where public builds come in
	// from
	PublicSubscription string `protobuf:"bytes,2,opt,name=public_subscription,json=publicSubscription" json:"public_subscription,omitempty"`
	// internal_subscription is the name of the pubsub topic where internal builds
	// come in from
	InternalSubscription string `protobuf:"bytes,3,opt,name=internal_subscription,json=internalSubscription" json:"internal_subscription,omitempty"`
}

func (m *Settings_Buildbot) Reset()                    { *m = Settings_Buildbot{} }
func (m *Settings_Buildbot) String() string            { return proto.CompactTextString(m) }
func (*Settings_Buildbot) ProtoMessage()               {}
func (*Settings_Buildbot) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0, 0} }

func (m *Settings_Buildbot) GetInternalReader() string {
	if m != nil {
		return m.InternalReader
	}
	return ""
}

func (m *Settings_Buildbot) GetPublicSubscription() string {
	if m != nil {
		return m.PublicSubscription
	}
	return ""
}

func (m *Settings_Buildbot) GetInternalSubscription() string {
	if m != nil {
		return m.InternalSubscription
	}
	return ""
}

type Settings_Buildbucket struct {
	// name is the user friendly name of the Buildbucket instance we're pointing to.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// host is the hostname of the buildbucket instance we're pointing to (sans schema).
	Host string `protobuf:"bytes,2,opt,name=host" json:"host,omitempty"`
	// project is the name of the Google Cloud project that the pubsub topic
	// belongs to.
	Project string `protobuf:"bytes,3,opt,name=project" json:"project,omitempty"`
}

func (m *Settings_Buildbucket) Reset()                    { *m = Settings_Buildbucket{} }
func (m *Settings_Buildbucket) String() string            { return proto.CompactTextString(m) }
func (*Settings_Buildbucket) ProtoMessage()               {}
func (*Settings_Buildbucket) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0, 1} }

func (m *Settings_Buildbucket) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Settings_Buildbucket) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *Settings_Buildbucket) GetProject() string {
	if m != nil {
		return m.Project
	}
	return ""
}

type Settings_Swarming struct {
	// default_host is the hostname of the swarming host Milo defaults to, if
	// none is specified.  Default host is implicitly an allowed host.
	DefaultHost string `protobuf:"bytes,1,opt,name=default_host,json=defaultHost" json:"default_host,omitempty"`
	// allowed_hosts is a whitelist of hostnames of swarming instances
	// that Milo is allowed to talk to.  This is specified here for security
	// reasons, because Milo will hand out its oauth2 token to a swarming host.
	AllowedHosts []string `protobuf:"bytes,2,rep,name=allowed_hosts,json=allowedHosts" json:"allowed_hosts,omitempty"`
}

func (m *Settings_Swarming) Reset()                    { *m = Settings_Swarming{} }
func (m *Settings_Swarming) String() string            { return proto.CompactTextString(m) }
func (*Settings_Swarming) ProtoMessage()               {}
func (*Settings_Swarming) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0, 2} }

func (m *Settings_Swarming) GetDefaultHost() string {
	if m != nil {
		return m.DefaultHost
	}
	return ""
}

func (m *Settings_Swarming) GetAllowedHosts() []string {
	if m != nil {
		return m.AllowedHosts
	}
	return nil
}

func init() {
	proto.RegisterType((*Settings)(nil), "config.Settings")
	proto.RegisterType((*Settings_Buildbot)(nil), "config.Settings.Buildbot")
	proto.RegisterType((*Settings_Buildbucket)(nil), "config.Settings.Buildbucket")
	proto.RegisterType((*Settings_Swarming)(nil), "config.Settings.Swarming")
}

func init() {
	proto.RegisterFile("go.chromium.org/luci/milo/api/config/settings.proto", fileDescriptor1)
}

var fileDescriptor1 = []byte{
	// 327 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0xbf, 0x4e, 0xc3, 0x30,
	0x10, 0xc6, 0xd5, 0x06, 0x95, 0xf4, 0x52, 0x40, 0x32, 0x20, 0x85, 0x8a, 0xa1, 0xc0, 0x40, 0xa7,
	0x44, 0xa2, 0x62, 0x65, 0x60, 0x62, 0x43, 0x72, 0x1f, 0xa0, 0x72, 0x1c, 0x37, 0x35, 0x38, 0x76,
	0xe4, 0x3f, 0xea, 0xb3, 0xf0, 0xa8, 0x6c, 0x28, 0x8e, 0x13, 0x85, 0xa1, 0xdb, 0xe5, 0xbb, 0xdf,
	0xf7, 0xdd, 0x5d, 0x64, 0xd8, 0x54, 0x2a, 0xa3, 0x07, 0xad, 0x6a, 0xee, 0xea, 0x4c, 0xe9, 0x2a,
	0x17, 0x8e, 0xf2, 0xbc, 0xe6, 0x42, 0xe5, 0xa4, 0xe1, 0x39, 0x55, 0x72, 0xcf, 0xab, 0xdc, 0x30,
	0x6b, 0xb9, 0xac, 0x4c, 0xd6, 0x68, 0x65, 0x15, 0x9a, 0x75, 0xf2, 0xe3, 0x6f, 0x04, 0xf1, 0x36,
	0xb4, 0xd0, 0x2b, 0xc4, 0x85, 0xe3, 0xa2, 0x2c, 0x94, 0x4d, 0x27, 0xab, 0xc9, 0x3a, 0x79, 0xb9,
	0xcb, 0x3a, 0x2e, 0xeb, 0x99, 0xec, 0x3d, 0x00, 0x78, 0x40, 0xd1, 0x1b, 0x24, 0x5d, 0xed, 0xe8,
	0x37, 0xb3, 0xe9, 0xd4, 0x3b, 0xef, 0x4f, 0x38, 0x3d, 0x83, 0xc7, 0x86, 0x76, 0xac, 0x39, 0x12,
	0x5d, 0x73, 0x59, 0xa5, 0xd1, 0x89, 0xb1, 0xdb, 0x00, 0xe0, 0x01, 0x5d, 0xfe, 0x4c, 0x20, 0xee,
	0xb7, 0x41, 0xcf, 0x70, 0xc5, 0xa5, 0x65, 0x5a, 0x12, 0xb1, 0xd3, 0x8c, 0x94, 0x4c, 0xfb, 0x0b,
	0xe6, 0xf8, 0xb2, 0x97, 0xb1, 0x57, 0x51, 0x0e, 0xd7, 0x8d, 0x2b, 0x04, 0xa7, 0x3b, 0xe3, 0x0a,
	0x43, 0x35, 0x6f, 0x2c, 0x57, 0xd2, 0x2f, 0x3d, 0xc7, 0xa8, 0x6b, 0x6d, 0x47, 0x1d, 0xb4, 0x81,
	0xdb, 0x21, 0xf9, 0x9f, 0x25, 0xf2, 0x96, 0x9b, 0xbe, 0x39, 0x36, 0x2d, 0x3f, 0x21, 0x19, 0x9d,
	0x8b, 0x10, 0x9c, 0x49, 0x52, 0xb3, 0xb0, 0x92, 0xaf, 0x5b, 0xed, 0xa0, 0x8c, 0x0d, 0x93, 0x7d,
	0x8d, 0x52, 0x38, 0x6f, 0xb4, 0xfa, 0x62, 0xd4, 0x86, 0xf4, 0xfe, 0x73, 0x89, 0x21, 0xee, 0x7f,
	0x01, 0x7a, 0x80, 0x45, 0xc9, 0xf6, 0xc4, 0x09, 0xbb, 0xf3, 0x09, 0x5d, 0x6a, 0x12, 0xb4, 0x8f,
	0x36, 0xe8, 0x09, 0x2e, 0x88, 0x10, 0xea, 0xc8, 0x4a, 0x8f, 0x98, 0x74, 0xba, 0x8a, 0xd6, 0x73,
	0xbc, 0x08, 0x62, 0xcb, 0x98, 0x62, 0xe6, 0x9f, 0xc2, 0xe6, 0x2f, 0x00, 0x00, 0xff, 0xff, 0x5a,
	0x7c, 0xd0, 0x8a, 0x41, 0x02, 0x00, 0x00,
}