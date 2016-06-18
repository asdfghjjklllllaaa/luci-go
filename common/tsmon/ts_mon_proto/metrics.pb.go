// Code generated by protoc-gen-go.
// source: metrics.proto
// DO NOT EDIT!

package ts_mon_proto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type MetricsField_FieldType int32

const (
	MetricsField_STRING MetricsField_FieldType = 1
	MetricsField_INT    MetricsField_FieldType = 2
	MetricsField_BOOL   MetricsField_FieldType = 3
)

var MetricsField_FieldType_name = map[int32]string{
	1: "STRING",
	2: "INT",
	3: "BOOL",
}
var MetricsField_FieldType_value = map[string]int32{
	"STRING": 1,
	"INT":    2,
	"BOOL":   3,
}

func (x MetricsField_FieldType) Enum() *MetricsField_FieldType {
	p := new(MetricsField_FieldType)
	*p = x
	return p
}
func (x MetricsField_FieldType) String() string {
	return proto.EnumName(MetricsField_FieldType_name, int32(x))
}
func (x *MetricsField_FieldType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(MetricsField_FieldType_value, data, "MetricsField_FieldType")
	if err != nil {
		return err
	}
	*x = MetricsField_FieldType(value)
	return nil
}
func (MetricsField_FieldType) EnumDescriptor() ([]byte, []int) { return fileDescriptor2, []int{1, 0} }

type PrecomputedDistribution_SpecType int32

const (
	PrecomputedDistribution_CANONICAL_POWERS_OF_2        PrecomputedDistribution_SpecType = 1
	PrecomputedDistribution_CANONICAL_POWERS_OF_10_P_0_2 PrecomputedDistribution_SpecType = 2
	PrecomputedDistribution_CANONICAL_POWERS_OF_10       PrecomputedDistribution_SpecType = 3
	PrecomputedDistribution_CUSTOM_PARAMETERIZED         PrecomputedDistribution_SpecType = 20
	PrecomputedDistribution_CUSTOM_BOUNDED               PrecomputedDistribution_SpecType = 21
)

var PrecomputedDistribution_SpecType_name = map[int32]string{
	1:  "CANONICAL_POWERS_OF_2",
	2:  "CANONICAL_POWERS_OF_10_P_0_2",
	3:  "CANONICAL_POWERS_OF_10",
	20: "CUSTOM_PARAMETERIZED",
	21: "CUSTOM_BOUNDED",
}
var PrecomputedDistribution_SpecType_value = map[string]int32{
	"CANONICAL_POWERS_OF_2":        1,
	"CANONICAL_POWERS_OF_10_P_0_2": 2,
	"CANONICAL_POWERS_OF_10":       3,
	"CUSTOM_PARAMETERIZED":         20,
	"CUSTOM_BOUNDED":               21,
}

func (x PrecomputedDistribution_SpecType) Enum() *PrecomputedDistribution_SpecType {
	p := new(PrecomputedDistribution_SpecType)
	*p = x
	return p
}
func (x PrecomputedDistribution_SpecType) String() string {
	return proto.EnumName(PrecomputedDistribution_SpecType_name, int32(x))
}
func (x *PrecomputedDistribution_SpecType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(PrecomputedDistribution_SpecType_value, data, "PrecomputedDistribution_SpecType")
	if err != nil {
		return err
	}
	*x = PrecomputedDistribution_SpecType(value)
	return nil
}
func (PrecomputedDistribution_SpecType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor2, []int{2, 0}
}

type MetricsData_Units int32

const (
	MetricsData_UNKNOWN_UNITS MetricsData_Units = 0
	MetricsData_SECONDS       MetricsData_Units = 1
	MetricsData_MILLISECONDS  MetricsData_Units = 2
	MetricsData_MICROSECONDS  MetricsData_Units = 3
	MetricsData_NANOSECONDS   MetricsData_Units = 4
	MetricsData_BITS          MetricsData_Units = 21
	MetricsData_BYTES         MetricsData_Units = 22
	// * 1000 bytes (not 1024).
	MetricsData_KILOBYTES MetricsData_Units = 31
	// * 1e6 (1,000,000) bytes.
	MetricsData_MEGABYTES MetricsData_Units = 32
	// * 1e9 (1,000,000,000) bytes.
	MetricsData_GIGABYTES MetricsData_Units = 33
	// * 1024 bytes.
	MetricsData_KIBIBYTES MetricsData_Units = 41
	// * 1024^2 (1,048,576) bytes.
	MetricsData_MEBIBYTES MetricsData_Units = 42
	// * 1024^3 (1,073,741,824) bytes.
	MetricsData_GIBIBYTES MetricsData_Units = 43
	// * Extended Units
	MetricsData_AMPS            MetricsData_Units = 60
	MetricsData_MILLIAMPS       MetricsData_Units = 61
	MetricsData_DEGREES_CELSIUS MetricsData_Units = 62
)

var MetricsData_Units_name = map[int32]string{
	0:  "UNKNOWN_UNITS",
	1:  "SECONDS",
	2:  "MILLISECONDS",
	3:  "MICROSECONDS",
	4:  "NANOSECONDS",
	21: "BITS",
	22: "BYTES",
	31: "KILOBYTES",
	32: "MEGABYTES",
	33: "GIGABYTES",
	41: "KIBIBYTES",
	42: "MEBIBYTES",
	43: "GIBIBYTES",
	60: "AMPS",
	61: "MILLIAMPS",
	62: "DEGREES_CELSIUS",
}
var MetricsData_Units_value = map[string]int32{
	"UNKNOWN_UNITS":   0,
	"SECONDS":         1,
	"MILLISECONDS":    2,
	"MICROSECONDS":    3,
	"NANOSECONDS":     4,
	"BITS":            21,
	"BYTES":           22,
	"KILOBYTES":       31,
	"MEGABYTES":       32,
	"GIGABYTES":       33,
	"KIBIBYTES":       41,
	"MEBIBYTES":       42,
	"GIBIBYTES":       43,
	"AMPS":            60,
	"MILLIAMPS":       61,
	"DEGREES_CELSIUS": 62,
}

func (x MetricsData_Units) Enum() *MetricsData_Units {
	p := new(MetricsData_Units)
	*p = x
	return p
}
func (x MetricsData_Units) String() string {
	return proto.EnumName(MetricsData_Units_name, int32(x))
}
func (x *MetricsData_Units) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(MetricsData_Units_value, data, "MetricsData_Units")
	if err != nil {
		return err
	}
	*x = MetricsData_Units(value)
	return nil
}
func (MetricsData_Units) EnumDescriptor() ([]byte, []int) { return fileDescriptor2, []int{3, 0} }

type MetricsCollection struct {
	Data              []*MetricsData `protobuf:"bytes,1,rep,name=data" json:"data,omitempty"`
	StartTimestampUs  *uint64        `protobuf:"varint,2,opt,name=start_timestamp_us,json=startTimestampUs" json:"start_timestamp_us,omitempty"`
	CollectionPointId *string        `protobuf:"bytes,3,opt,name=collection_point_id,json=collectionPointId" json:"collection_point_id,omitempty"`
	XXX_unrecognized  []byte         `json:"-"`
}

func (m *MetricsCollection) Reset()                    { *m = MetricsCollection{} }
func (m *MetricsCollection) String() string            { return proto.CompactTextString(m) }
func (*MetricsCollection) ProtoMessage()               {}
func (*MetricsCollection) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *MetricsCollection) GetData() []*MetricsData {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *MetricsCollection) GetStartTimestampUs() uint64 {
	if m != nil && m.StartTimestampUs != nil {
		return *m.StartTimestampUs
	}
	return 0
}

func (m *MetricsCollection) GetCollectionPointId() string {
	if m != nil && m.CollectionPointId != nil {
		return *m.CollectionPointId
	}
	return ""
}

type MetricsField struct {
	Name             *string                 `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Type             *MetricsField_FieldType `protobuf:"varint,3,opt,name=type,enum=ts_mon.proto.MetricsField_FieldType,def=1" json:"type,omitempty"`
	StringValue      *string                 `protobuf:"bytes,4,opt,name=string_value,json=stringValue" json:"string_value,omitempty"`
	IntValue         *int64                  `protobuf:"varint,5,opt,name=int_value,json=intValue" json:"int_value,omitempty"`
	BoolValue        *bool                   `protobuf:"varint,6,opt,name=bool_value,json=boolValue" json:"bool_value,omitempty"`
	XXX_unrecognized []byte                  `json:"-"`
}

func (m *MetricsField) Reset()                    { *m = MetricsField{} }
func (m *MetricsField) String() string            { return proto.CompactTextString(m) }
func (*MetricsField) ProtoMessage()               {}
func (*MetricsField) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

const Default_MetricsField_Type MetricsField_FieldType = MetricsField_STRING

func (m *MetricsField) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *MetricsField) GetType() MetricsField_FieldType {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return Default_MetricsField_Type
}

func (m *MetricsField) GetStringValue() string {
	if m != nil && m.StringValue != nil {
		return *m.StringValue
	}
	return ""
}

func (m *MetricsField) GetIntValue() int64 {
	if m != nil && m.IntValue != nil {
		return *m.IntValue
	}
	return 0
}

func (m *MetricsField) GetBoolValue() bool {
	if m != nil && m.BoolValue != nil {
		return *m.BoolValue
	}
	return false
}

type PrecomputedDistribution struct {
	SpecType              *PrecomputedDistribution_SpecType `protobuf:"varint,1,opt,name=spec_type,json=specType,enum=ts_mon.proto.PrecomputedDistribution_SpecType" json:"spec_type,omitempty"`
	Width                 *float64                          `protobuf:"fixed64,2,opt,name=width,def=10" json:"width,omitempty"`
	GrowthFactor          *float64                          `protobuf:"fixed64,3,opt,name=growth_factor,json=growthFactor,def=0" json:"growth_factor,omitempty"`
	NumBuckets            *int32                            `protobuf:"varint,4,opt,name=num_buckets,json=numBuckets,def=10" json:"num_buckets,omitempty"`
	LowerBounds           []float64                         `protobuf:"fixed64,5,rep,name=lower_bounds,json=lowerBounds" json:"lower_bounds,omitempty"`
	IsCumulative          *bool                             `protobuf:"varint,6,opt,name=is_cumulative,json=isCumulative,def=0" json:"is_cumulative,omitempty"`
	Bucket                []int64                           `protobuf:"zigzag64,7,rep,name=bucket" json:"bucket,omitempty"`
	Underflow             *int64                            `protobuf:"zigzag64,8,opt,name=underflow" json:"underflow,omitempty"`
	Overflow              *int64                            `protobuf:"zigzag64,9,opt,name=overflow" json:"overflow,omitempty"`
	Mean                  *float64                          `protobuf:"fixed64,10,opt,name=mean" json:"mean,omitempty"`
	SumOfSquaredDeviation *float64                          `protobuf:"fixed64,11,opt,name=sum_of_squared_deviation,json=sumOfSquaredDeviation" json:"sum_of_squared_deviation,omitempty"`
	XXX_unrecognized      []byte                            `json:"-"`
}

func (m *PrecomputedDistribution) Reset()                    { *m = PrecomputedDistribution{} }
func (m *PrecomputedDistribution) String() string            { return proto.CompactTextString(m) }
func (*PrecomputedDistribution) ProtoMessage()               {}
func (*PrecomputedDistribution) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{2} }

const Default_PrecomputedDistribution_Width float64 = 10
const Default_PrecomputedDistribution_GrowthFactor float64 = 0
const Default_PrecomputedDistribution_NumBuckets int32 = 10
const Default_PrecomputedDistribution_IsCumulative bool = false

func (m *PrecomputedDistribution) GetSpecType() PrecomputedDistribution_SpecType {
	if m != nil && m.SpecType != nil {
		return *m.SpecType
	}
	return PrecomputedDistribution_CANONICAL_POWERS_OF_2
}

func (m *PrecomputedDistribution) GetWidth() float64 {
	if m != nil && m.Width != nil {
		return *m.Width
	}
	return Default_PrecomputedDistribution_Width
}

func (m *PrecomputedDistribution) GetGrowthFactor() float64 {
	if m != nil && m.GrowthFactor != nil {
		return *m.GrowthFactor
	}
	return Default_PrecomputedDistribution_GrowthFactor
}

func (m *PrecomputedDistribution) GetNumBuckets() int32 {
	if m != nil && m.NumBuckets != nil {
		return *m.NumBuckets
	}
	return Default_PrecomputedDistribution_NumBuckets
}

func (m *PrecomputedDistribution) GetLowerBounds() []float64 {
	if m != nil {
		return m.LowerBounds
	}
	return nil
}

func (m *PrecomputedDistribution) GetIsCumulative() bool {
	if m != nil && m.IsCumulative != nil {
		return *m.IsCumulative
	}
	return Default_PrecomputedDistribution_IsCumulative
}

func (m *PrecomputedDistribution) GetBucket() []int64 {
	if m != nil {
		return m.Bucket
	}
	return nil
}

func (m *PrecomputedDistribution) GetUnderflow() int64 {
	if m != nil && m.Underflow != nil {
		return *m.Underflow
	}
	return 0
}

func (m *PrecomputedDistribution) GetOverflow() int64 {
	if m != nil && m.Overflow != nil {
		return *m.Overflow
	}
	return 0
}

func (m *PrecomputedDistribution) GetMean() float64 {
	if m != nil && m.Mean != nil {
		return *m.Mean
	}
	return 0
}

func (m *PrecomputedDistribution) GetSumOfSquaredDeviation() float64 {
	if m != nil && m.SumOfSquaredDeviation != nil {
		return *m.SumOfSquaredDeviation
	}
	return 0
}

type MetricsData struct {
	Name                     *string                  `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	MetricNamePrefix         *string                  `protobuf:"bytes,2,opt,name=metric_name_prefix,json=metricNamePrefix" json:"metric_name_prefix,omitempty"`
	NetworkDevice            *NetworkDevice           `protobuf:"bytes,11,opt,name=network_device,json=networkDevice" json:"network_device,omitempty"`
	Task                     *Task                    `protobuf:"bytes,12,opt,name=task" json:"task,omitempty"`
	Fields                   []*MetricsField          `protobuf:"bytes,20,rep,name=fields" json:"fields,omitempty"`
	Counter                  *int64                   `protobuf:"varint,30,opt,name=counter" json:"counter,omitempty"`
	Gauge                    *int64                   `protobuf:"varint,32,opt,name=gauge" json:"gauge,omitempty"`
	NoncumulativeDoubleValue *float64                 `protobuf:"fixed64,34,opt,name=noncumulative_double_value,json=noncumulativeDoubleValue" json:"noncumulative_double_value,omitempty"`
	Distribution             *PrecomputedDistribution `protobuf:"bytes,35,opt,name=distribution" json:"distribution,omitempty"`
	StringValue              *string                  `protobuf:"bytes,36,opt,name=string_value,json=stringValue" json:"string_value,omitempty"`
	BooleanValue             *bool                    `protobuf:"varint,37,opt,name=boolean_value,json=booleanValue" json:"boolean_value,omitempty"`
	CumulativeDoubleValue    *float64                 `protobuf:"fixed64,38,opt,name=cumulative_double_value,json=cumulativeDoubleValue" json:"cumulative_double_value,omitempty"`
	StartTimestampUs         *uint64                  `protobuf:"varint,40,opt,name=start_timestamp_us,json=startTimestampUs" json:"start_timestamp_us,omitempty"`
	Units                    *MetricsData_Units       `protobuf:"varint,41,opt,name=units,enum=ts_mon.proto.MetricsData_Units" json:"units,omitempty"`
	Description              *string                  `protobuf:"bytes,43,opt,name=description" json:"description,omitempty"`
	XXX_unrecognized         []byte                   `json:"-"`
}

func (m *MetricsData) Reset()                    { *m = MetricsData{} }
func (m *MetricsData) String() string            { return proto.CompactTextString(m) }
func (*MetricsData) ProtoMessage()               {}
func (*MetricsData) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{3} }

func (m *MetricsData) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *MetricsData) GetMetricNamePrefix() string {
	if m != nil && m.MetricNamePrefix != nil {
		return *m.MetricNamePrefix
	}
	return ""
}

func (m *MetricsData) GetNetworkDevice() *NetworkDevice {
	if m != nil {
		return m.NetworkDevice
	}
	return nil
}

func (m *MetricsData) GetTask() *Task {
	if m != nil {
		return m.Task
	}
	return nil
}

func (m *MetricsData) GetFields() []*MetricsField {
	if m != nil {
		return m.Fields
	}
	return nil
}

func (m *MetricsData) GetCounter() int64 {
	if m != nil && m.Counter != nil {
		return *m.Counter
	}
	return 0
}

func (m *MetricsData) GetGauge() int64 {
	if m != nil && m.Gauge != nil {
		return *m.Gauge
	}
	return 0
}

func (m *MetricsData) GetNoncumulativeDoubleValue() float64 {
	if m != nil && m.NoncumulativeDoubleValue != nil {
		return *m.NoncumulativeDoubleValue
	}
	return 0
}

func (m *MetricsData) GetDistribution() *PrecomputedDistribution {
	if m != nil {
		return m.Distribution
	}
	return nil
}

func (m *MetricsData) GetStringValue() string {
	if m != nil && m.StringValue != nil {
		return *m.StringValue
	}
	return ""
}

func (m *MetricsData) GetBooleanValue() bool {
	if m != nil && m.BooleanValue != nil {
		return *m.BooleanValue
	}
	return false
}

func (m *MetricsData) GetCumulativeDoubleValue() float64 {
	if m != nil && m.CumulativeDoubleValue != nil {
		return *m.CumulativeDoubleValue
	}
	return 0
}

func (m *MetricsData) GetStartTimestampUs() uint64 {
	if m != nil && m.StartTimestampUs != nil {
		return *m.StartTimestampUs
	}
	return 0
}

func (m *MetricsData) GetUnits() MetricsData_Units {
	if m != nil && m.Units != nil {
		return *m.Units
	}
	return MetricsData_UNKNOWN_UNITS
}

func (m *MetricsData) GetDescription() string {
	if m != nil && m.Description != nil {
		return *m.Description
	}
	return ""
}

func init() {
	proto.RegisterType((*MetricsCollection)(nil), "ts_mon.proto.MetricsCollection")
	proto.RegisterType((*MetricsField)(nil), "ts_mon.proto.MetricsField")
	proto.RegisterType((*PrecomputedDistribution)(nil), "ts_mon.proto.PrecomputedDistribution")
	proto.RegisterType((*MetricsData)(nil), "ts_mon.proto.MetricsData")
	proto.RegisterEnum("ts_mon.proto.MetricsField_FieldType", MetricsField_FieldType_name, MetricsField_FieldType_value)
	proto.RegisterEnum("ts_mon.proto.PrecomputedDistribution_SpecType", PrecomputedDistribution_SpecType_name, PrecomputedDistribution_SpecType_value)
	proto.RegisterEnum("ts_mon.proto.MetricsData_Units", MetricsData_Units_name, MetricsData_Units_value)
}

func init() { proto.RegisterFile("metrics.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 1044 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x55, 0xdb, 0x6e, 0xdb, 0x46,
	0x10, 0x2d, 0x75, 0xb1, 0xa5, 0x91, 0x94, 0xd0, 0x1b, 0xdb, 0x61, 0x9c, 0xb4, 0x51, 0xe4, 0x24,
	0x70, 0x9c, 0x56, 0x70, 0x0c, 0xb4, 0x05, 0x8c, 0xb4, 0xa8, 0x2e, 0xb4, 0x41, 0x58, 0x26, 0x85,
	0xa5, 0xd4, 0xa0, 0x7d, 0x59, 0xd0, 0xe4, 0xca, 0x21, 0x2c, 0x91, 0x0a, 0x2f, 0x76, 0xfb, 0x19,
	0xfd, 0x83, 0xbe, 0xf7, 0x63, 0xfa, 0x29, 0x7d, 0xee, 0x5b, 0x97, 0xb3, 0x94, 0x2d, 0x21, 0x36,
	0xd0, 0x17, 0x81, 0x73, 0xce, 0x99, 0x9d, 0x9d, 0xd9, 0x99, 0x11, 0x34, 0x66, 0x3c, 0x89, 0x7c,
	0x37, 0x6e, 0xcf, 0xa3, 0x30, 0x09, 0x49, 0x3d, 0x89, 0xd9, 0x2c, 0x0c, 0xa4, 0xb5, 0xd3, 0x74,
	0xdc, 0x4f, 0xa9, 0x1f, 0xfb, 0x89, 0x1f, 0x06, 0x2c, 0xe0, 0xc9, 0x75, 0x18, 0x5d, 0x32, 0x8f,
	0x5f, 0xf9, 0x2e, 0xcf, 0x15, 0xdb, 0xcb, 0x8a, 0xc4, 0x89, 0x2f, 0x25, 0xde, 0xfa, 0x53, 0x81,
	0x8d, 0x33, 0x79, 0x72, 0x2f, 0x9c, 0x4e, 0xb9, 0x9b, 0x09, 0xc8, 0x37, 0x50, 0xf2, 0x9c, 0xc4,
	0xd1, 0x94, 0x66, 0x71, 0xaf, 0x76, 0xf8, 0xa4, 0xbd, 0x1c, 0xac, 0x9d, 0xcb, 0xfb, 0x42, 0x40,
	0x51, 0x46, 0xbe, 0x06, 0x12, 0x27, 0x4e, 0x94, 0xb0, 0xc4, 0x9f, 0x71, 0xf1, 0x35, 0x9b, 0xb3,
	0x34, 0xd6, 0x0a, 0x4d, 0x65, 0xaf, 0x44, 0x55, 0x64, 0x46, 0x0b, 0x62, 0x1c, 0x93, 0x36, 0x3c,
	0x72, 0x6f, 0x42, 0xb1, 0x79, 0xe8, 0x07, 0x09, 0xf3, 0x3d, 0xad, 0x28, 0xe4, 0x55, 0xba, 0x71,
	0x4b, 0x0d, 0x33, 0xc6, 0xf0, 0x5a, 0xff, 0x28, 0x50, 0xcf, 0x63, 0x1e, 0xfb, 0x7c, 0xea, 0x11,
	0x02, 0xa5, 0xc0, 0x99, 0x71, 0x71, 0xbb, 0xcc, 0x03, 0xbf, 0xc9, 0x4f, 0x50, 0x4a, 0x7e, 0x9f,
	0x73, 0x3c, 0xe5, 0xc1, 0xe1, 0xcb, 0x3b, 0x6f, 0x8c, 0xde, 0x6d, 0xfc, 0x1d, 0x09, 0xed, 0xd1,
	0x9a, 0x3d, 0xa2, 0x86, 0x79, 0x42, 0xd1, 0x93, 0xbc, 0x80, 0x7a, 0x2c, 0x64, 0xc1, 0x05, 0xbb,
	0x72, 0xa6, 0x29, 0xd7, 0x4a, 0x78, 0x7a, 0x4d, 0x62, 0x3f, 0x67, 0x10, 0x79, 0x0a, 0xd5, 0xec,
	0xb2, 0x92, 0x2f, 0x0b, 0xbe, 0x48, 0x2b, 0x02, 0x90, 0xe4, 0x97, 0x00, 0xe7, 0x61, 0x38, 0xcd,
	0xd9, 0x35, 0xc1, 0x56, 0x68, 0x35, 0x43, 0x90, 0x6e, 0xed, 0x43, 0xf5, 0x26, 0x32, 0x01, 0xc8,
	0x63, 0xab, 0x0a, 0x59, 0x87, 0xa2, 0x61, 0x8e, 0xd4, 0x02, 0xa9, 0x40, 0xa9, 0x6b, 0x59, 0x03,
	0xb5, 0xd8, 0xfa, 0xbb, 0x04, 0x8f, 0x87, 0x11, 0x77, 0xc3, 0xd9, 0x3c, 0x4d, 0xb8, 0xd7, 0xf7,
	0xb3, 0x4b, 0x9c, 0xa7, 0xf8, 0x34, 0xa7, 0x50, 0x8d, 0xe7, 0xdc, 0x65, 0x98, 0xad, 0x82, 0xd9,
	0xb6, 0x57, 0xb3, 0xbd, 0xc7, 0xb3, 0x6d, 0x0b, 0xb7, 0x2c, 0x3a, 0xad, 0xc4, 0xf9, 0x17, 0xd1,
	0xa0, 0x7c, 0xed, 0x7b, 0xc9, 0x47, 0x7c, 0x2b, 0xe5, 0xa8, 0xf0, 0xee, 0x80, 0x4a, 0x80, 0xbc,
	0x86, 0xc6, 0x45, 0x14, 0x5e, 0x27, 0x1f, 0xd9, 0xc4, 0x71, 0x93, 0x30, 0xc2, 0xc2, 0x2a, 0x47,
	0xca, 0x01, 0xad, 0x4b, 0xfc, 0x18, 0x61, 0xb2, 0x0b, 0xb5, 0x20, 0x9d, 0xb1, 0xf3, 0xd4, 0xbd,
	0xe4, 0x49, 0x8c, 0x45, 0x2b, 0xe3, 0x39, 0x20, 0xe0, 0xae, 0x44, 0xb3, 0xd2, 0x4e, 0xc3, 0x6b,
	0x1e, 0xb1, 0xf3, 0x30, 0x0d, 0xbc, 0x58, 0x94, 0xae, 0xb8, 0xa7, 0xd0, 0x1a, 0x62, 0x5d, 0x84,
	0xc8, 0x3e, 0x34, 0xfc, 0x98, 0xb9, 0xe9, 0x2c, 0x9d, 0x3a, 0x89, 0x7f, 0x95, 0x17, 0xf0, 0xa8,
	0x3c, 0x71, 0xa6, 0x31, 0xa7, 0x75, 0x3f, 0xee, 0xdd, 0x50, 0x64, 0x1b, 0xd6, 0x64, 0x3c, 0x6d,
	0x5d, 0x1c, 0x44, 0x68, 0x6e, 0x91, 0x67, 0x50, 0x15, 0x67, 0xf1, 0x68, 0x22, 0x0e, 0xd6, 0x2a,
	0xc2, 0x9f, 0xd0, 0x5b, 0x80, 0xec, 0x40, 0x25, 0xbc, 0xca, 0xc9, 0x2a, 0x92, 0x37, 0x76, 0xd6,
	0x51, 0x33, 0xee, 0x04, 0x1a, 0x64, 0x49, 0x52, 0xfc, 0x26, 0xdf, 0x83, 0x16, 0x8b, 0xcc, 0xc2,
	0x09, 0x8b, 0x3f, 0xa5, 0x4e, 0xc4, 0x3d, 0x1c, 0x28, 0x27, 0x2b, 0xa5, 0x56, 0x43, 0xdd, 0x96,
	0xe0, 0xad, 0x89, 0x2d, 0xd9, 0xfe, 0x82, 0x6c, 0xfd, 0xa1, 0x40, 0x65, 0x51, 0x6b, 0xf2, 0x04,
	0xb6, 0x7a, 0x1d, 0xd3, 0x32, 0x8d, 0x5e, 0x67, 0xc0, 0x86, 0xd6, 0x07, 0x9d, 0xda, 0xcc, 0x3a,
	0x66, 0x87, 0xe2, 0xe1, 0x9b, 0xf0, 0xec, 0x2e, 0xea, 0xdd, 0x01, 0x1b, 0xb2, 0x03, 0xa1, 0x28,
	0x88, 0x2b, 0x6f, 0xdf, 0xad, 0x50, 0x8b, 0xe2, 0xe9, 0x36, 0x7b, 0x63, 0x7b, 0x64, 0x9d, 0xb1,
	0x61, 0x87, 0x76, 0xce, 0xf4, 0x91, 0x4e, 0x8d, 0x5f, 0xf5, 0xbe, 0xba, 0x29, 0x92, 0x79, 0x90,
	0x33, 0x5d, 0x6b, 0x6c, 0xf6, 0x05, 0xb6, 0xd5, 0xfa, 0x6b, 0x1d, 0x6a, 0x4b, 0x73, 0xbb, 0x34,
	0x42, 0x85, 0x9b, 0x11, 0x12, 0x53, 0x2c, 0x77, 0x0c, 0xcb, 0x4c, 0x36, 0x8f, 0xf8, 0xc4, 0xff,
	0x0d, 0x3b, 0xa3, 0x4a, 0x55, 0xc9, 0x98, 0x82, 0x18, 0x22, 0x4e, 0xba, 0xf0, 0x60, 0x75, 0xd1,
	0x60, 0x51, 0x6a, 0x87, 0x4f, 0x57, 0x9b, 0xd1, 0x94, 0x9a, 0x3e, 0x4a, 0x68, 0x23, 0x58, 0x36,
	0x45, 0x93, 0x95, 0xb2, 0x55, 0xa4, 0xd5, 0xd1, 0x93, 0xac, 0x7a, 0x8e, 0x04, 0x43, 0x91, 0x27,
	0x87, 0xb0, 0x36, 0xc9, 0x66, 0x27, 0xd6, 0x36, 0x71, 0x21, 0xed, 0xdc, 0x3f, 0xde, 0x34, 0x57,
	0x8a, 0xfa, 0xac, 0xbb, 0xa2, 0xb5, 0x12, 0x1e, 0x69, 0x5f, 0xe1, 0xa4, 0x2e, 0x4c, 0xb2, 0x09,
	0xe5, 0x0b, 0x27, 0xbd, 0xe0, 0x5a, 0x13, 0x71, 0x69, 0x90, 0xf7, 0xb0, 0x13, 0x84, 0xc1, 0x6d,
	0x03, 0x32, 0x2f, 0x4c, 0xcf, 0xa7, 0x3c, 0x1f, 0xe7, 0x16, 0x3e, 0xb8, 0xb6, 0xa2, 0xe8, 0xa3,
	0x40, 0x0e, 0xbf, 0x01, 0x75, 0x6f, 0x69, 0xd6, 0xb4, 0x5d, 0xcc, 0xe8, 0xd5, 0xff, 0x1a, 0x4c,
	0xba, 0xe2, 0xfa, 0xd9, 0x1e, 0x7a, 0xf9, 0xf9, 0x1e, 0xda, 0x85, 0x46, 0xb6, 0x58, 0x44, 0x97,
	0xe6, 0x9a, 0x57, 0xb8, 0x6d, 0xea, 0x39, 0x28, 0x45, 0xdf, 0xc1, 0xe3, 0xfb, 0xb2, 0x79, 0x2d,
	0xdb, 0xf7, 0xee, 0x54, 0xee, 0x5e, 0xe6, 0x7b, 0xf7, 0x2c, 0xf3, 0x6f, 0xa1, 0x9c, 0x06, 0xbe,
	0x98, 0xfc, 0x37, 0xb8, 0x8a, 0x9e, 0xdf, 0xfb, 0x57, 0xd1, 0x1e, 0x67, 0x32, 0x2a, 0xd5, 0xa2,
	0xf7, 0x6b, 0x1e, 0x8f, 0xdd, 0xc8, 0x9f, 0x63, 0xb9, 0xde, 0xca, 0x1c, 0x97, 0xa0, 0xd6, 0xbf,
	0x0a, 0x94, 0xd1, 0x85, 0x6c, 0x40, 0x63, 0x6c, 0x9e, 0x9a, 0xd6, 0x07, 0x93, 0x8d, 0x4d, 0x63,
	0x64, 0xab, 0x5f, 0x90, 0x1a, 0xac, 0xdb, 0x7a, 0xcf, 0x32, 0xfb, 0xb6, 0x98, 0x23, 0x55, 0xfc,
	0x3d, 0x18, 0x83, 0x81, 0xb1, 0x40, 0x0a, 0x12, 0xe9, 0x51, 0x6b, 0x81, 0x14, 0xc9, 0x43, 0xa8,
	0x99, 0x62, 0x92, 0x16, 0x40, 0x09, 0x97, 0x6d, 0x76, 0xd6, 0x16, 0xa9, 0x42, 0xb9, 0xfb, 0xcb,
	0x48, 0xb7, 0xd5, 0x6d, 0xd2, 0x80, 0xea, 0xa9, 0x31, 0xb0, 0xa4, 0xf9, 0x3c, 0x33, 0xcf, 0xf4,
	0x93, 0x8e, 0x34, 0x9b, 0x99, 0x79, 0x62, 0x2c, 0xcc, 0x17, 0x52, 0xdc, 0x35, 0xa4, 0xf9, 0x46,
	0x8a, 0x17, 0xe6, 0xbe, 0x14, 0x2f, 0xcc, 0xb7, 0x59, 0xb8, 0xce, 0xd9, 0xd0, 0x56, 0xdf, 0xa3,
	0x2e, 0xbb, 0x2d, 0x9a, 0x3f, 0x90, 0x47, 0xf0, 0xb0, 0xaf, 0x9f, 0x50, 0x5d, 0xb7, 0x59, 0x4f,
	0x1f, 0xd8, 0xc6, 0xd8, 0x56, 0x7f, 0xfc, 0x2f, 0x00, 0x00, 0xff, 0xff, 0x96, 0xa9, 0xad, 0x82,
	0xec, 0x07, 0x00, 0x00,
}
