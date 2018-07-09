// Code generated by protoc-gen-go. DO NOT EDIT.
// source: go.chromium.org/luci/buildbucket/proto/common.proto

package buildbucketpb // import "go.chromium.org/luci/buildbucket/proto"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Status of a build or a step.
type Status int32

const (
	// Unspecified state. Meaning depends on the context.
	Status_STATUS_UNSPECIFIED Status = 0
	// Build was scheduled, but did not start or end yet.
	Status_SCHEDULED Status = 1
	// Build/step has started.
	Status_STARTED Status = 2
	// A union of all terminal statuses.
	// Can be used in BuildPredicate.status.
	// A concrete build cannot have this status.
	// Can be used as a bitmask to check that a build ended.
	Status_ENDED_MASK Status = 4
	// A build/step ended successfully.
	Status_SUCCESS Status = 12
	// A build/step ended unsuccessfully due to its Build.Input,
	// e.g. tests failed, and NOT due to a build infrastructure failure.
	Status_FAILURE Status = 20
	// A build/step ended unsuccessfully due to a failure independent of the input,
	// e.g. swarming failed, or the recipe was unable to read the patch from gerrit..
	Status_INFRA_FAILURE Status = 36
	// A build was cancelled explicitly, e.g. via an RPC.
	Status_CANCELED Status = 68
)

var Status_name = map[int32]string{
	0:  "STATUS_UNSPECIFIED",
	1:  "SCHEDULED",
	2:  "STARTED",
	4:  "ENDED_MASK",
	12: "SUCCESS",
	20: "FAILURE",
	36: "INFRA_FAILURE",
	68: "CANCELED",
}
var Status_value = map[string]int32{
	"STATUS_UNSPECIFIED": 0,
	"SCHEDULED":          1,
	"STARTED":            2,
	"ENDED_MASK":         4,
	"SUCCESS":            12,
	"FAILURE":            20,
	"INFRA_FAILURE":      36,
	"CANCELED":           68,
}

func (x Status) String() string {
	return proto.EnumName(Status_name, int32(x))
}
func (Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_common_10866d3ce07b8e7b, []int{0}
}

// A Gerrit patchset.
type GerritChange struct {
	// Gerrit hostname, e.g. "chromium-review.googlesource.com".
	Host string `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	// Gerrit project, e.g. "chromium/src".
	Project string `protobuf:"bytes,2,opt,name=project,proto3" json:"project,omitempty"`
	// Change number, e.g. 12345.
	Change int64 `protobuf:"varint,3,opt,name=change,proto3" json:"change,omitempty"`
	// Patch set number, e.g. 1.
	Patchset             int64    `protobuf:"varint,4,opt,name=patchset,proto3" json:"patchset,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GerritChange) Reset()         { *m = GerritChange{} }
func (m *GerritChange) String() string { return proto.CompactTextString(m) }
func (*GerritChange) ProtoMessage()    {}
func (*GerritChange) Descriptor() ([]byte, []int) {
	return fileDescriptor_common_10866d3ce07b8e7b, []int{0}
}
func (m *GerritChange) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GerritChange.Unmarshal(m, b)
}
func (m *GerritChange) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GerritChange.Marshal(b, m, deterministic)
}
func (dst *GerritChange) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GerritChange.Merge(dst, src)
}
func (m *GerritChange) XXX_Size() int {
	return xxx_messageInfo_GerritChange.Size(m)
}
func (m *GerritChange) XXX_DiscardUnknown() {
	xxx_messageInfo_GerritChange.DiscardUnknown(m)
}

var xxx_messageInfo_GerritChange proto.InternalMessageInfo

func (m *GerritChange) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *GerritChange) GetProject() string {
	if m != nil {
		return m.Project
	}
	return ""
}

func (m *GerritChange) GetChange() int64 {
	if m != nil {
		return m.Change
	}
	return 0
}

func (m *GerritChange) GetPatchset() int64 {
	if m != nil {
		return m.Patchset
	}
	return 0
}

// A landed Git commit hosted on Gitiles.
type GitilesCommit struct {
	// Gitiles hostname, e.g. "chromium.googlesource.com".
	Host string `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	// Repository name on the host, e.g. "chromium/src".
	Project string `protobuf:"bytes,2,opt,name=project,proto3" json:"project,omitempty"`
	// Commit HEX SHA1.
	Id string `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`
	// Commit ref, e.g. "refs/heads/master".
	// NOT a branch name: if specified, must start with "refs/".
	Ref                  string   `protobuf:"bytes,4,opt,name=ref,proto3" json:"ref,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GitilesCommit) Reset()         { *m = GitilesCommit{} }
func (m *GitilesCommit) String() string { return proto.CompactTextString(m) }
func (*GitilesCommit) ProtoMessage()    {}
func (*GitilesCommit) Descriptor() ([]byte, []int) {
	return fileDescriptor_common_10866d3ce07b8e7b, []int{1}
}
func (m *GitilesCommit) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GitilesCommit.Unmarshal(m, b)
}
func (m *GitilesCommit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GitilesCommit.Marshal(b, m, deterministic)
}
func (dst *GitilesCommit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GitilesCommit.Merge(dst, src)
}
func (m *GitilesCommit) XXX_Size() int {
	return xxx_messageInfo_GitilesCommit.Size(m)
}
func (m *GitilesCommit) XXX_DiscardUnknown() {
	xxx_messageInfo_GitilesCommit.DiscardUnknown(m)
}

var xxx_messageInfo_GitilesCommit proto.InternalMessageInfo

func (m *GitilesCommit) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *GitilesCommit) GetProject() string {
	if m != nil {
		return m.Project
	}
	return ""
}

func (m *GitilesCommit) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *GitilesCommit) GetRef() string {
	if m != nil {
		return m.Ref
	}
	return ""
}

// A key-value pair of strings.
type StringPair struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StringPair) Reset()         { *m = StringPair{} }
func (m *StringPair) String() string { return proto.CompactTextString(m) }
func (*StringPair) ProtoMessage()    {}
func (*StringPair) Descriptor() ([]byte, []int) {
	return fileDescriptor_common_10866d3ce07b8e7b, []int{2}
}
func (m *StringPair) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StringPair.Unmarshal(m, b)
}
func (m *StringPair) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StringPair.Marshal(b, m, deterministic)
}
func (dst *StringPair) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StringPair.Merge(dst, src)
}
func (m *StringPair) XXX_Size() int {
	return xxx_messageInfo_StringPair.Size(m)
}
func (m *StringPair) XXX_DiscardUnknown() {
	xxx_messageInfo_StringPair.DiscardUnknown(m)
}

var xxx_messageInfo_StringPair proto.InternalMessageInfo

func (m *StringPair) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *StringPair) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

// Half-open time range.
type TimeRange struct {
	// Inclusive lower boundary. Optional.
	StartTime *timestamp.Timestamp `protobuf:"bytes,1,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	// Exclusive upper boundary. Optional.
	EndTime              *timestamp.Timestamp `protobuf:"bytes,2,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *TimeRange) Reset()         { *m = TimeRange{} }
func (m *TimeRange) String() string { return proto.CompactTextString(m) }
func (*TimeRange) ProtoMessage()    {}
func (*TimeRange) Descriptor() ([]byte, []int) {
	return fileDescriptor_common_10866d3ce07b8e7b, []int{3}
}
func (m *TimeRange) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TimeRange.Unmarshal(m, b)
}
func (m *TimeRange) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TimeRange.Marshal(b, m, deterministic)
}
func (dst *TimeRange) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TimeRange.Merge(dst, src)
}
func (m *TimeRange) XXX_Size() int {
	return xxx_messageInfo_TimeRange.Size(m)
}
func (m *TimeRange) XXX_DiscardUnknown() {
	xxx_messageInfo_TimeRange.DiscardUnknown(m)
}

var xxx_messageInfo_TimeRange proto.InternalMessageInfo

func (m *TimeRange) GetStartTime() *timestamp.Timestamp {
	if m != nil {
		return m.StartTime
	}
	return nil
}

func (m *TimeRange) GetEndTime() *timestamp.Timestamp {
	if m != nil {
		return m.EndTime
	}
	return nil
}

func init() {
	proto.RegisterType((*GerritChange)(nil), "buildbucket.v2.GerritChange")
	proto.RegisterType((*GitilesCommit)(nil), "buildbucket.v2.GitilesCommit")
	proto.RegisterType((*StringPair)(nil), "buildbucket.v2.StringPair")
	proto.RegisterType((*TimeRange)(nil), "buildbucket.v2.TimeRange")
	proto.RegisterEnum("buildbucket.v2.Status", Status_name, Status_value)
}

func init() {
	proto.RegisterFile("go.chromium.org/luci/buildbucket/proto/common.proto", fileDescriptor_common_10866d3ce07b8e7b)
}

var fileDescriptor_common_10866d3ce07b8e7b = []byte{
	// 435 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0x4f, 0x6f, 0x9b, 0x40,
	0x10, 0xc5, 0x0b, 0x76, 0x6d, 0x33, 0xb1, 0x2d, 0xba, 0x8a, 0x22, 0xe4, 0x4b, 0x23, 0xab, 0x87,
	0xa8, 0x07, 0x90, 0x92, 0xb4, 0x52, 0xd5, 0x13, 0x85, 0x75, 0x6a, 0x35, 0xb5, 0x22, 0x16, 0x2e,
	0xbd, 0x20, 0xfe, 0x6c, 0xf0, 0x36, 0xc0, 0xa2, 0x65, 0x89, 0xd4, 0x43, 0xcf, 0xfd, 0xda, 0x15,
	0x8b, 0x5d, 0xf9, 0xd6, 0xf6, 0x36, 0xbf, 0xf7, 0xf6, 0xcd, 0xa0, 0x27, 0xe0, 0xa6, 0xe0, 0x76,
	0xb6, 0x17, 0xbc, 0x62, 0x5d, 0x65, 0x73, 0x51, 0x38, 0x65, 0x97, 0x31, 0x27, 0xed, 0x58, 0x99,
	0xa7, 0x5d, 0xf6, 0x44, 0xa5, 0xd3, 0x08, 0x2e, 0xb9, 0x93, 0xf1, 0xaa, 0xe2, 0xb5, 0xad, 0x00,
	0x2d, 0x4f, 0x7c, 0xfb, 0xf9, 0x7a, 0xf5, 0xba, 0xe0, 0xbc, 0x28, 0xe9, 0xf0, 0x34, 0xed, 0x1e,
	0x1d, 0xc9, 0x2a, 0xda, 0xca, 0xa4, 0x6a, 0x86, 0xc0, 0xba, 0x81, 0xf9, 0x1d, 0x15, 0x82, 0x49,
	0x6f, 0x9f, 0xd4, 0x05, 0x45, 0x08, 0xc6, 0x7b, 0xde, 0x4a, 0x4b, 0xbb, 0xd4, 0xae, 0x8c, 0x40,
	0xcd, 0xc8, 0x82, 0x69, 0x23, 0xf8, 0x77, 0x9a, 0x49, 0x4b, 0x57, 0xf2, 0x11, 0xd1, 0x05, 0x4c,
	0x32, 0x95, 0xb3, 0x46, 0x97, 0xda, 0xd5, 0x28, 0x38, 0x10, 0x5a, 0xc1, 0xac, 0x49, 0x64, 0xb6,
	0x6f, 0xa9, 0xb4, 0xc6, 0xca, 0xf9, 0xc3, 0xeb, 0x18, 0x16, 0x77, 0x4c, 0xb2, 0x92, 0xb6, 0x1e,
	0xaf, 0x2a, 0x26, 0xff, 0xf3, 0xe4, 0x12, 0x74, 0x96, 0xab, 0x73, 0x46, 0xa0, 0xb3, 0x1c, 0x99,
	0x30, 0x12, 0xf4, 0x51, 0x5d, 0x31, 0x82, 0x7e, 0x5c, 0xdf, 0x02, 0x10, 0x29, 0x58, 0x5d, 0x3c,
	0x24, 0x4c, 0xf4, 0xfe, 0x13, 0xfd, 0x71, 0x58, 0xde, 0x8f, 0xe8, 0x1c, 0x5e, 0x3e, 0x27, 0x65,
	0x47, 0x0f, 0x9b, 0x07, 0x58, 0xff, 0x04, 0x23, 0x64, 0x15, 0x0d, 0xd4, 0xf7, 0x7f, 0x00, 0x68,
	0x65, 0x22, 0x64, 0xdc, 0xd7, 0xa5, 0xb2, 0x67, 0xd7, 0x2b, 0x7b, 0xe8, 0xd2, 0x3e, 0x76, 0x69,
	0x87, 0xc7, 0x2e, 0x03, 0x43, 0xbd, 0xee, 0x19, 0xbd, 0x83, 0x19, 0xad, 0xf3, 0x21, 0xa8, 0xff,
	0x35, 0x38, 0xa5, 0x75, 0xde, 0xd3, 0xdb, 0x5f, 0x1a, 0x4c, 0x88, 0x4c, 0x64, 0xd7, 0xa2, 0x0b,
	0x40, 0x24, 0x74, 0xc3, 0x88, 0xc4, 0xd1, 0x8e, 0x3c, 0x60, 0x6f, 0xbb, 0xd9, 0x62, 0xdf, 0x7c,
	0x81, 0x16, 0x60, 0x10, 0xef, 0x33, 0xf6, 0xa3, 0x7b, 0xec, 0x9b, 0x1a, 0x3a, 0x83, 0x29, 0x09,
	0xdd, 0x20, 0xc4, 0xbe, 0xa9, 0xa3, 0x25, 0x00, 0xde, 0xf9, 0xd8, 0x8f, 0xbf, 0xba, 0xe4, 0x8b,
	0x39, 0x56, 0x66, 0xe4, 0x79, 0x98, 0x10, 0x73, 0xde, 0xc3, 0xc6, 0xdd, 0xde, 0x47, 0x01, 0x36,
	0xcf, 0xd1, 0x2b, 0x58, 0x6c, 0x77, 0x9b, 0xc0, 0x8d, 0x8f, 0xd2, 0x1b, 0x34, 0x87, 0x99, 0xe7,
	0xee, 0x3c, 0xdc, 0xef, 0xf5, 0x3f, 0xbd, 0xff, 0x76, 0xfb, 0x6f, 0x7f, 0xde, 0xc7, 0x13, 0xa5,
	0x49, 0xd3, 0x89, 0x12, 0x6f, 0x7e, 0x07, 0x00, 0x00, 0xff, 0xff, 0x63, 0x1a, 0xc8, 0x9b, 0xb8,
	0x02, 0x00, 0x00,
}
