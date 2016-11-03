// Code generated by protoc-gen-go.
// source: github.com/luci/luci-go/dm/api/distributor/swarming/v1/result.proto
// DO NOT EDIT!

package swarmingV1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is the swarming-specific result for Executions run via swarming.
type Result struct {
	ExitCode int64 `protobuf:"varint,1,opt,name=exit_code,json=exitCode" json:"exit_code,omitempty"`
	// The isolated hash of the output directory
	IsolatedOutdir *IsolatedRef `protobuf:"bytes,2,opt,name=isolated_outdir,json=isolatedOutdir" json:"isolated_outdir,omitempty"`
	// The pinned cipd packages that this task actually used.
	CipdPins *CipdSpec `protobuf:"bytes,3,opt,name=cipd_pins,json=cipdPins" json:"cipd_pins,omitempty"`
	// The captured snapshot dimensions that the bot actually had.
	SnapshotDimensions map[string]string `protobuf:"bytes,4,rep,name=snapshot_dimensions,json=snapshotDimensions" json:"snapshot_dimensions,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *Result) Reset()                    { *m = Result{} }
func (m *Result) String() string            { return proto.CompactTextString(m) }
func (*Result) ProtoMessage()               {}
func (*Result) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

func (m *Result) GetIsolatedOutdir() *IsolatedRef {
	if m != nil {
		return m.IsolatedOutdir
	}
	return nil
}

func (m *Result) GetCipdPins() *CipdSpec {
	if m != nil {
		return m.CipdPins
	}
	return nil
}

func (m *Result) GetSnapshotDimensions() map[string]string {
	if m != nil {
		return m.SnapshotDimensions
	}
	return nil
}

func init() {
	proto.RegisterType((*Result)(nil), "swarmingV1.Result")
}

func init() {
	proto.RegisterFile("github.com/luci/luci-go/dm/api/distributor/swarming/v1/result.proto", fileDescriptor4)
}

var fileDescriptor4 = []byte{
	// 315 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x90, 0xcf, 0x4b, 0xfb, 0x30,
	0x18, 0xc6, 0xe9, 0xfa, 0xfd, 0x8e, 0x35, 0x03, 0x95, 0x38, 0x58, 0x99, 0x97, 0xe1, 0x69, 0x08,
	0x36, 0x6c, 0x5e, 0xc4, 0x93, 0x32, 0x07, 0x7a, 0x52, 0x32, 0xf0, 0xe2, 0xa1, 0x74, 0x4d, 0xb6,
	0xbd, 0xd8, 0x26, 0x21, 0x3f, 0xa6, 0xfb, 0xe7, 0xfc, 0xdb, 0xa4, 0x69, 0xa7, 0x03, 0xf1, 0xb2,
	0x4b, 0x48, 0xf2, 0x7e, 0x9e, 0x87, 0xf7, 0x79, 0xd0, 0x74, 0x05, 0x76, 0xed, 0x16, 0x49, 0x2e,
	0x4b, 0x52, 0xb8, 0x1c, 0xfc, 0x71, 0xb9, 0x92, 0x84, 0x95, 0x24, 0x53, 0x40, 0x18, 0x18, 0xab,
	0x61, 0xe1, 0xac, 0xd4, 0xc4, 0xbc, 0x67, 0xba, 0x04, 0xb1, 0x22, 0x9b, 0x31, 0xd1, 0xdc, 0xb8,
	0xc2, 0x26, 0x4a, 0x4b, 0x2b, 0x31, 0xda, 0x4d, 0x5e, 0xc6, 0x83, 0xbb, 0x03, 0x0d, 0x73, 0x50,
	0xac, 0xb6, 0x1b, 0x3c, 0x1c, 0x68, 0x01, 0x46, 0x16, 0x99, 0xe5, 0xa9, 0xe6, 0xcb, 0xda, 0xe9,
	0xfc, 0xb3, 0x85, 0xda, 0xd4, 0x6f, 0x8a, 0xcf, 0x50, 0xc4, 0x3f, 0xc0, 0xa6, 0xb9, 0x64, 0x3c,
	0x0e, 0x86, 0xc1, 0x28, 0xa4, 0x9d, 0xea, 0x63, 0x2a, 0x19, 0xc7, 0xb7, 0xe8, 0xb8, 0x11, 0xb3,
	0x54, 0x3a, 0xcb, 0x40, 0xc7, 0xad, 0x61, 0x30, 0xea, 0x4e, 0xfa, 0xc9, 0x4f, 0xb4, 0xe4, 0xb1,
	0x41, 0x28, 0x5f, 0xd2, 0xa3, 0x1d, 0xff, 0xe4, 0x71, 0x3c, 0x46, 0x51, 0x95, 0x20, 0x55, 0x20,
	0x4c, 0x1c, 0x7a, 0x6d, 0x6f, 0x5f, 0x3b, 0x05, 0xc5, 0xe6, 0x8a, 0xe7, 0xb4, 0x53, 0x61, 0xcf,
	0x20, 0x0c, 0x7e, 0x45, 0xa7, 0x46, 0x64, 0xca, 0xac, 0xa5, 0x4d, 0x19, 0x94, 0x5c, 0x18, 0x90,
	0xc2, 0xc4, 0xff, 0x86, 0xe1, 0xa8, 0x3b, 0xb9, 0xd8, 0x17, 0xd7, 0x11, 0x92, 0x79, 0x43, 0xdf,
	0x7f, 0xc3, 0x33, 0x61, 0xf5, 0x96, 0x62, 0xf3, 0x6b, 0x30, 0x98, 0xa1, 0xfe, 0x1f, 0x38, 0x3e,
	0x41, 0xe1, 0x1b, 0xdf, 0xfa, 0x0e, 0x22, 0x5a, 0x5d, 0x71, 0x0f, 0xfd, 0xdf, 0x64, 0x85, 0xe3,
	0x3e, 0x74, 0x44, 0xeb, 0xc7, 0x4d, 0xeb, 0x3a, 0x58, 0xb4, 0x7d, 0x8f, 0x57, 0x5f, 0x01, 0x00,
	0x00, 0xff, 0xff, 0x16, 0xba, 0xd0, 0x69, 0x27, 0x02, 0x00, 0x00,
}
