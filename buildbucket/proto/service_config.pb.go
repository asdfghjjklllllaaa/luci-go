// Code generated by protoc-gen-go. DO NOT EDIT.
// source: go.chromium.org/luci/buildbucket/proto/service_config.proto

package buildbucketpb

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

// Schema of settings.cfg file, a service config.
type SettingsCfg struct {
	// Swarmbucket settings.
	Swarming *SwarmingSettings `protobuf:"bytes,1,opt,name=swarming,proto3" json:"swarming,omitempty"`
	Logdog   *LogDogSettings   `protobuf:"bytes,2,opt,name=logdog,proto3" json:"logdog,omitempty"`
	// List of Gerrit hosts to force git authentication for.
	//
	// By default public hosts are accessed anonymously, and the anonymous access
	// has very low quota. Context needs to know all such hostnames in advance to
	// be able to force authenticated access to them.
	KnownPublicGerritHosts []string `protobuf:"bytes,3,rep,name=known_public_gerrit_hosts,json=knownPublicGerritHosts,proto3" json:"known_public_gerrit_hosts,omitempty"`
	XXX_NoUnkeyedLiteral   struct{} `json:"-"`
	XXX_unrecognized       []byte   `json:"-"`
	XXX_sizecache          int32    `json:"-"`
}

func (m *SettingsCfg) Reset()         { *m = SettingsCfg{} }
func (m *SettingsCfg) String() string { return proto.CompactTextString(m) }
func (*SettingsCfg) ProtoMessage()    {}
func (*SettingsCfg) Descriptor() ([]byte, []int) {
	return fileDescriptor_4e08a93b642b6ff2, []int{0}
}

func (m *SettingsCfg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SettingsCfg.Unmarshal(m, b)
}
func (m *SettingsCfg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SettingsCfg.Marshal(b, m, deterministic)
}
func (m *SettingsCfg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SettingsCfg.Merge(m, src)
}
func (m *SettingsCfg) XXX_Size() int {
	return xxx_messageInfo_SettingsCfg.Size(m)
}
func (m *SettingsCfg) XXX_DiscardUnknown() {
	xxx_messageInfo_SettingsCfg.DiscardUnknown(m)
}

var xxx_messageInfo_SettingsCfg proto.InternalMessageInfo

func (m *SettingsCfg) GetSwarming() *SwarmingSettings {
	if m != nil {
		return m.Swarming
	}
	return nil
}

func (m *SettingsCfg) GetLogdog() *LogDogSettings {
	if m != nil {
		return m.Logdog
	}
	return nil
}

func (m *SettingsCfg) GetKnownPublicGerritHosts() []string {
	if m != nil {
		return m.KnownPublicGerritHosts
	}
	return nil
}

// Swarmbucket settings.
type SwarmingSettings struct {
	// Swarmbucket build URLs will point to this Milo instance.
	MiloHostname string `protobuf:"bytes,2,opt,name=milo_hostname,json=miloHostname,proto3" json:"milo_hostname,omitempty"`
	// These caches are available to all builders implicitly.
	// A builder may override a cache specified here.
	GlobalCaches []*Builder_CacheEntry `protobuf:"bytes,4,rep,name=global_caches,json=globalCaches,proto3" json:"global_caches,omitempty"`
	// Packages available to the user executable in $PATH.
	// Installed in "{TASK_RUN_DIR}/cipd_bin_packages".
	// "{TASK_RUN_DIR}/cipd_bin_packages" and
	// "{TASK_RUN_DIR}/cipd_bin_packages/bin" are prepended to $PATH.
	UserPackages []*SwarmingSettings_Package `protobuf:"bytes,5,rep,name=user_packages,json=userPackages,proto3" json:"user_packages,omitempty"`
	// Package of LUCI runner,
	// https://chromium.googlesource.com/infra/luci/luci-go/+/HEAD/buildbucket/cmd/luci_runner
	// used to run LUCI executables.
	LuciRunnerPackage *SwarmingSettings_Package `protobuf:"bytes,6,opt,name=luci_runner_package,json=luciRunnerPackage,proto3" json:"luci_runner_package,omitempty"`
	// CIPD package of kitchen binary. DEPRECATED. TODO(nodir): remove.
	KitchenPackage       *SwarmingSettings_Package `protobuf:"bytes,7,opt,name=kitchen_package,json=kitchenPackage,proto3" json:"kitchen_package,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *SwarmingSettings) Reset()         { *m = SwarmingSettings{} }
func (m *SwarmingSettings) String() string { return proto.CompactTextString(m) }
func (*SwarmingSettings) ProtoMessage()    {}
func (*SwarmingSettings) Descriptor() ([]byte, []int) {
	return fileDescriptor_4e08a93b642b6ff2, []int{1}
}

func (m *SwarmingSettings) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SwarmingSettings.Unmarshal(m, b)
}
func (m *SwarmingSettings) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SwarmingSettings.Marshal(b, m, deterministic)
}
func (m *SwarmingSettings) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SwarmingSettings.Merge(m, src)
}
func (m *SwarmingSettings) XXX_Size() int {
	return xxx_messageInfo_SwarmingSettings.Size(m)
}
func (m *SwarmingSettings) XXX_DiscardUnknown() {
	xxx_messageInfo_SwarmingSettings.DiscardUnknown(m)
}

var xxx_messageInfo_SwarmingSettings proto.InternalMessageInfo

func (m *SwarmingSettings) GetMiloHostname() string {
	if m != nil {
		return m.MiloHostname
	}
	return ""
}

func (m *SwarmingSettings) GetGlobalCaches() []*Builder_CacheEntry {
	if m != nil {
		return m.GlobalCaches
	}
	return nil
}

func (m *SwarmingSettings) GetUserPackages() []*SwarmingSettings_Package {
	if m != nil {
		return m.UserPackages
	}
	return nil
}

func (m *SwarmingSettings) GetLuciRunnerPackage() *SwarmingSettings_Package {
	if m != nil {
		return m.LuciRunnerPackage
	}
	return nil
}

func (m *SwarmingSettings) GetKitchenPackage() *SwarmingSettings_Package {
	if m != nil {
		return m.KitchenPackage
	}
	return nil
}

// CIPD package. Does not specify installation path.
type SwarmingSettings_Package struct {
	// CIPD package name, e.g. "infra/python/cpython/${platform}"
	PackageName string `protobuf:"bytes,1,opt,name=package_name,json=packageName,proto3" json:"package_name,omitempty"`
	// CIPD instance version, e.g. "version:2.7.15.chromium14".
	// Used for non-canary builds.
	Version string `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	// CIPD instance version for canary builds.
	// Defaults to version.
	VersionCanary string `protobuf:"bytes,3,opt,name=version_canary,json=versionCanary,proto3" json:"version_canary,omitempty"`
	// Include in builders matching the predicate.
	Builders *BuilderPredicate `protobuf:"bytes,4,opt,name=builders,proto3" json:"builders,omitempty"`
	// Subdirectory to install the package into, relative to the installation
	// root. Useful if installing two packages at the same root would conflict.
	Subdir               string   `protobuf:"bytes,5,opt,name=subdir,proto3" json:"subdir,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SwarmingSettings_Package) Reset()         { *m = SwarmingSettings_Package{} }
func (m *SwarmingSettings_Package) String() string { return proto.CompactTextString(m) }
func (*SwarmingSettings_Package) ProtoMessage()    {}
func (*SwarmingSettings_Package) Descriptor() ([]byte, []int) {
	return fileDescriptor_4e08a93b642b6ff2, []int{1, 0}
}

func (m *SwarmingSettings_Package) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SwarmingSettings_Package.Unmarshal(m, b)
}
func (m *SwarmingSettings_Package) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SwarmingSettings_Package.Marshal(b, m, deterministic)
}
func (m *SwarmingSettings_Package) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SwarmingSettings_Package.Merge(m, src)
}
func (m *SwarmingSettings_Package) XXX_Size() int {
	return xxx_messageInfo_SwarmingSettings_Package.Size(m)
}
func (m *SwarmingSettings_Package) XXX_DiscardUnknown() {
	xxx_messageInfo_SwarmingSettings_Package.DiscardUnknown(m)
}

var xxx_messageInfo_SwarmingSettings_Package proto.InternalMessageInfo

func (m *SwarmingSettings_Package) GetPackageName() string {
	if m != nil {
		return m.PackageName
	}
	return ""
}

func (m *SwarmingSettings_Package) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *SwarmingSettings_Package) GetVersionCanary() string {
	if m != nil {
		return m.VersionCanary
	}
	return ""
}

func (m *SwarmingSettings_Package) GetBuilders() *BuilderPredicate {
	if m != nil {
		return m.Builders
	}
	return nil
}

func (m *SwarmingSettings_Package) GetSubdir() string {
	if m != nil {
		return m.Subdir
	}
	return ""
}

type LogDogSettings struct {
	// Hostname of the LogDog instance to use, e.g. "logs.chromium.org".
	Hostname             string   `protobuf:"bytes,1,opt,name=hostname,proto3" json:"hostname,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogDogSettings) Reset()         { *m = LogDogSettings{} }
func (m *LogDogSettings) String() string { return proto.CompactTextString(m) }
func (*LogDogSettings) ProtoMessage()    {}
func (*LogDogSettings) Descriptor() ([]byte, []int) {
	return fileDescriptor_4e08a93b642b6ff2, []int{2}
}

func (m *LogDogSettings) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogDogSettings.Unmarshal(m, b)
}
func (m *LogDogSettings) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogDogSettings.Marshal(b, m, deterministic)
}
func (m *LogDogSettings) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogDogSettings.Merge(m, src)
}
func (m *LogDogSettings) XXX_Size() int {
	return xxx_messageInfo_LogDogSettings.Size(m)
}
func (m *LogDogSettings) XXX_DiscardUnknown() {
	xxx_messageInfo_LogDogSettings.DiscardUnknown(m)
}

var xxx_messageInfo_LogDogSettings proto.InternalMessageInfo

func (m *LogDogSettings) GetHostname() string {
	if m != nil {
		return m.Hostname
	}
	return ""
}

// A predicate for a builder.
type BuilderPredicate struct {
	// OR-connected list of regular expressions for a string
	// "{project}/{bucket}/{builder}".
	// Each regex is wrapped in ^ and $ automatically.
	// Examples:
	//
	//   # All builders in "chromium" project
	//   regex: "chromium/.+"
	//   # A specific builder.
	//   regex: "infra/ci/infra-continuous-trusty-64"
	//
	// Defaults to [".*"].
	Regex []string `protobuf:"bytes,1,rep,name=regex,proto3" json:"regex,omitempty"`
	// Like regex field, but negation. Negation always wins.
	RegexExclude         []string `protobuf:"bytes,2,rep,name=regex_exclude,json=regexExclude,proto3" json:"regex_exclude,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BuilderPredicate) Reset()         { *m = BuilderPredicate{} }
func (m *BuilderPredicate) String() string { return proto.CompactTextString(m) }
func (*BuilderPredicate) ProtoMessage()    {}
func (*BuilderPredicate) Descriptor() ([]byte, []int) {
	return fileDescriptor_4e08a93b642b6ff2, []int{3}
}

func (m *BuilderPredicate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BuilderPredicate.Unmarshal(m, b)
}
func (m *BuilderPredicate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BuilderPredicate.Marshal(b, m, deterministic)
}
func (m *BuilderPredicate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BuilderPredicate.Merge(m, src)
}
func (m *BuilderPredicate) XXX_Size() int {
	return xxx_messageInfo_BuilderPredicate.Size(m)
}
func (m *BuilderPredicate) XXX_DiscardUnknown() {
	xxx_messageInfo_BuilderPredicate.DiscardUnknown(m)
}

var xxx_messageInfo_BuilderPredicate proto.InternalMessageInfo

func (m *BuilderPredicate) GetRegex() []string {
	if m != nil {
		return m.Regex
	}
	return nil
}

func (m *BuilderPredicate) GetRegexExclude() []string {
	if m != nil {
		return m.RegexExclude
	}
	return nil
}

func init() {
	proto.RegisterType((*SettingsCfg)(nil), "buildbucket.SettingsCfg")
	proto.RegisterType((*SwarmingSettings)(nil), "buildbucket.SwarmingSettings")
	proto.RegisterType((*SwarmingSettings_Package)(nil), "buildbucket.SwarmingSettings.Package")
	proto.RegisterType((*LogDogSettings)(nil), "buildbucket.LogDogSettings")
	proto.RegisterType((*BuilderPredicate)(nil), "buildbucket.BuilderPredicate")
}

func init() {
	proto.RegisterFile("go.chromium.org/luci/buildbucket/proto/service_config.proto", fileDescriptor_4e08a93b642b6ff2)
}

var fileDescriptor_4e08a93b642b6ff2 = []byte{
	// 525 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0x4b, 0x6b, 0xdb, 0x40,
	0x10, 0xc7, 0x51, 0x1c, 0x3f, 0x32, 0x7e, 0xd4, 0xdd, 0x96, 0xa0, 0xba, 0x94, 0xba, 0x2e, 0x01,
	0x1f, 0x8a, 0x0c, 0x49, 0x29, 0x84, 0xdc, 0xe2, 0x84, 0x86, 0xd0, 0x06, 0xa3, 0xd0, 0x4b, 0x2f,
	0x42, 0x5a, 0x6f, 0xd6, 0x5b, 0xcb, 0xbb, 0x66, 0x57, 0xca, 0xe3, 0x9b, 0xb5, 0xb7, 0x7e, 0xb4,
	0xb2, 0xa3, 0xb5, 0x62, 0xa7, 0x50, 0xdc, 0xdb, 0xce, 0x7f, 0xfe, 0xf3, 0xf3, 0x78, 0x66, 0x04,
	0x27, 0x5c, 0x05, 0x74, 0xa6, 0xd5, 0x42, 0xe4, 0x8b, 0x40, 0x69, 0x3e, 0x4a, 0x73, 0x2a, 0x46,
	0x49, 0x2e, 0xd2, 0x69, 0x92, 0xd3, 0x39, 0xcb, 0x46, 0x4b, 0xad, 0x32, 0x35, 0x32, 0x4c, 0xdf,
	0x0a, 0xca, 0x22, 0xaa, 0xe4, 0x8d, 0xe0, 0x01, 0x8a, 0xa4, 0xb9, 0xe6, 0xeb, 0x6d, 0x4b, 0x5a,
	0x6a, 0xf5, 0x83, 0xd1, 0x6c, 0x83, 0x34, 0xf8, 0xe9, 0x41, 0xf3, 0x9a, 0x65, 0x99, 0x90, 0xdc,
	0x8c, 0x6f, 0x38, 0x39, 0x86, 0x86, 0xb9, 0x8b, 0xf5, 0x42, 0x48, 0xee, 0x7b, 0x7d, 0x6f, 0xd8,
	0x3c, 0x7c, 0x13, 0xac, 0xa1, 0x82, 0x6b, 0x97, 0x5c, 0xd5, 0x84, 0xa5, 0x9d, 0x1c, 0x41, 0x2d,
	0x55, 0x7c, 0xaa, 0xb8, 0xbf, 0x83, 0x85, 0xaf, 0x37, 0x0a, 0xbf, 0x28, 0x7e, 0xa6, 0x1e, 0xcb,
	0x9c, 0x95, 0x1c, 0xc3, 0xab, 0xb9, 0x54, 0x77, 0x32, 0x5a, 0xe6, 0x49, 0x2a, 0x68, 0xc4, 0x99,
	0xd6, 0x22, 0x8b, 0x66, 0xca, 0x64, 0xc6, 0xaf, 0xf4, 0x2b, 0xc3, 0xbd, 0x70, 0x1f, 0x0d, 0x13,
	0xcc, 0x7f, 0xc6, 0xf4, 0x85, 0xcd, 0x0e, 0x7e, 0xed, 0x42, 0xf7, 0x69, 0x3b, 0xe4, 0x3d, 0xb4,
	0x17, 0x22, 0x55, 0x08, 0x90, 0xf1, 0x82, 0x61, 0x2f, 0x7b, 0x61, 0xcb, 0x8a, 0x17, 0x4e, 0x23,
	0x67, 0xd0, 0xe6, 0xa9, 0x4a, 0xe2, 0x34, 0xa2, 0x31, 0x9d, 0x31, 0xe3, 0xef, 0xf6, 0x2b, 0xc3,
	0xe6, 0xe1, 0xdb, 0x8d, 0x86, 0x4f, 0xed, 0x9b, 0xe9, 0x60, 0x6c, 0x2d, 0xe7, 0x32, 0xd3, 0x0f,
	0x61, 0xab, 0xa8, 0x42, 0xc5, 0x90, 0x4b, 0x68, 0xe7, 0x86, 0xe9, 0x68, 0x19, 0xd3, 0x79, 0xcc,
	0x99, 0xf1, 0xab, 0x48, 0x39, 0xf8, 0xe7, 0xbc, 0x82, 0x49, 0xe1, 0x0e, 0x5b, 0xb6, 0xd6, 0x05,
	0x86, 0x7c, 0x83, 0x17, 0x76, 0x6b, 0x91, 0xce, 0xa5, 0x7c, 0x44, 0xfa, 0x35, 0x1c, 0xe4, 0x96,
	0xc4, 0xe7, 0x96, 0x10, 0x22, 0xc0, 0x49, 0xe4, 0x0a, 0x9e, 0xcd, 0x45, 0x46, 0x67, 0x4c, 0x96,
	0xc8, 0xfa, 0xff, 0x20, 0x3b, 0xae, 0xda, 0xc5, 0xbd, 0xdf, 0x1e, 0xd4, 0x57, 0xec, 0x77, 0xd0,
	0x72, 0xcc, 0x08, 0x07, 0xed, 0xe1, 0xa0, 0x9b, 0x4e, 0xbb, 0xb2, 0x73, 0xf6, 0xa1, 0x7e, 0xcb,
	0xb4, 0x11, 0x4a, 0xba, 0x35, 0xac, 0x42, 0x72, 0x00, 0x1d, 0xf7, 0x8c, 0x68, 0x2c, 0x63, 0xfd,
	0xe0, 0x57, 0xd0, 0xd0, 0x76, 0xea, 0x18, 0x45, 0x7b, 0x8d, 0x49, 0xb1, 0x06, 0xbb, 0xa3, 0xbf,
	0xaf, 0xd1, 0xed, 0x68, 0xa2, 0xd9, 0x54, 0xd0, 0x38, 0x63, 0x61, 0x69, 0x27, 0xfb, 0x50, 0x33,
	0x79, 0x32, 0x15, 0xda, 0xaf, 0x22, 0xd9, 0x45, 0x97, 0xbb, 0x0d, 0xaf, 0xbb, 0x33, 0xf8, 0x00,
	0x9d, 0xcd, 0x83, 0x24, 0x3d, 0x68, 0x94, 0x37, 0x53, 0xfc, 0x95, 0x32, 0x1e, 0x7c, 0x85, 0xee,
	0xd3, 0x5f, 0x22, 0x2f, 0xa1, 0xaa, 0x19, 0x67, 0xf7, 0xbe, 0x87, 0x47, 0x5a, 0x04, 0xf6, 0xfc,
	0xf0, 0x11, 0xb1, 0x7b, 0x9a, 0xe6, 0x53, 0x7b, 0x7e, 0x36, 0xdb, 0x42, 0xf1, 0xbc, 0xd0, 0x4e,
	0x3f, 0x7d, 0xff, 0xb8, 0xdd, 0x27, 0x7b, 0xb2, 0xa6, 0x2c, 0x93, 0xa4, 0x86, 0xe2, 0xd1, 0x9f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x5d, 0x49, 0x27, 0xe4, 0x3b, 0x04, 0x00, 0x00,
}
