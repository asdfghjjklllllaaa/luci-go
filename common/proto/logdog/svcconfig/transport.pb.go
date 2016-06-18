// Code generated by protoc-gen-go.
// source: transport.proto
// DO NOT EDIT!

package svcconfig

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Transport is the transport configuration.
type Transport struct {
	// Type is the transport configuration that is being used.
	//
	// Types that are valid to be assigned to Type:
	//	*Transport_Pubsub
	Type isTransport_Type `protobuf_oneof:"Type"`
}

func (m *Transport) Reset()                    { *m = Transport{} }
func (m *Transport) String() string            { return proto.CompactTextString(m) }
func (*Transport) ProtoMessage()               {}
func (*Transport) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

type isTransport_Type interface {
	isTransport_Type()
}

type Transport_Pubsub struct {
	Pubsub *Transport_PubSub `protobuf:"bytes,1,opt,name=pubsub,oneof"`
}

func (*Transport_Pubsub) isTransport_Type() {}

func (m *Transport) GetType() isTransport_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *Transport) GetPubsub() *Transport_PubSub {
	if x, ok := m.GetType().(*Transport_Pubsub); ok {
		return x.Pubsub
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Transport) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Transport_OneofMarshaler, _Transport_OneofUnmarshaler, _Transport_OneofSizer, []interface{}{
		(*Transport_Pubsub)(nil),
	}
}

func _Transport_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Transport)
	// Type
	switch x := m.Type.(type) {
	case *Transport_Pubsub:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Pubsub); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Transport.Type has unexpected type %T", x)
	}
	return nil
}

func _Transport_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Transport)
	switch tag {
	case 1: // Type.pubsub
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Transport_PubSub)
		err := b.DecodeMessage(msg)
		m.Type = &Transport_Pubsub{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Transport_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Transport)
	// Type
	switch x := m.Type.(type) {
	case *Transport_Pubsub:
		s := proto.Size(x.Pubsub)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// PubSub is a transport configuration for Google Cloud Pub/Sub.
type Transport_PubSub struct {
	// The name of the authentication group for administrators.
	Project string `protobuf:"bytes,1,opt,name=project" json:"project,omitempty"`
	// The name of the authentication group for administrators.
	Topic string `protobuf:"bytes,2,opt,name=topic" json:"topic,omitempty"`
	// The name of the authentication group for administrators.
	Subscription string `protobuf:"bytes,3,opt,name=subscription" json:"subscription,omitempty"`
}

func (m *Transport_PubSub) Reset()                    { *m = Transport_PubSub{} }
func (m *Transport_PubSub) String() string            { return proto.CompactTextString(m) }
func (*Transport_PubSub) ProtoMessage()               {}
func (*Transport_PubSub) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{0, 0} }

func init() {
	proto.RegisterType((*Transport)(nil), "svcconfig.Transport")
	proto.RegisterType((*Transport_PubSub)(nil), "svcconfig.Transport.PubSub")
}

func init() { proto.RegisterFile("transport.proto", fileDescriptor4) }

var fileDescriptor4 = []byte{
	// 167 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x2f, 0x29, 0x4a, 0xcc,
	0x2b, 0x2e, 0xc8, 0x2f, 0x2a, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x2c, 0x2e, 0x4b,
	0x4e, 0xce, 0xcf, 0x4b, 0xcb, 0x4c, 0x57, 0x5a, 0xc1, 0xc8, 0xc5, 0x19, 0x02, 0x93, 0x16, 0x32,
	0xe5, 0x62, 0x2b, 0x28, 0x4d, 0x2a, 0x2e, 0x4d, 0x92, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x36, 0x92,
	0xd6, 0x83, 0xab, 0xd4, 0x83, 0xab, 0xd2, 0x0b, 0x28, 0x4d, 0x0a, 0x2e, 0x4d, 0xf2, 0x60, 0x08,
	0x82, 0x2a, 0x96, 0x8a, 0xe1, 0x62, 0x83, 0x88, 0x09, 0x49, 0x70, 0xb1, 0x03, 0xad, 0xc8, 0x4a,
	0x4d, 0x2e, 0x01, 0x9b, 0xc0, 0x19, 0x04, 0xe3, 0x0a, 0x89, 0x70, 0xb1, 0x96, 0xe4, 0x17, 0x64,
	0x26, 0x4b, 0x30, 0x81, 0xc5, 0x21, 0x1c, 0x21, 0x25, 0x2e, 0x1e, 0xa0, 0x01, 0xc5, 0xc9, 0x45,
	0x99, 0x05, 0x25, 0x99, 0xf9, 0x79, 0x12, 0xcc, 0x60, 0x49, 0x14, 0x31, 0x27, 0x36, 0x2e, 0x96,
	0x90, 0xca, 0x82, 0xd4, 0x24, 0x36, 0xb0, 0xe3, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x39,
	0x0e, 0x08, 0x8a, 0xcf, 0x00, 0x00, 0x00,
}
