// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto3_proto/proto3.proto

package proto3_proto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import test_proto "github.com/golang/protobuf/proto/test_proto"
import any "github.com/golang/protobuf/ptypes/any"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Message_Humour int32

const (
	Message_UNKNOWN     Message_Humour = 0
	Message_PUNS        Message_Humour = 1
	Message_SLAPSTICK   Message_Humour = 2
	Message_BILL_BAILEY Message_Humour = 3
)

var Message_Humour_name = map[int32]string{
	0: "UNKNOWN",
	1: "PUNS",
	2: "SLAPSTICK",
	3: "BILL_BAILEY",
}
var Message_Humour_value = map[string]int32{
	"UNKNOWN":     0,
	"PUNS":        1,
	"SLAPSTICK":   2,
	"BILL_BAILEY": 3,
}

func (x Message_Humour) String() string {
	return proto.EnumName(Message_Humour_name, int32(x))
}
func (Message_Humour) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_proto3_e706e4ff19a5dbea, []int{0, 0}
}

type Message struct {
	Name                 string                             `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Hilarity             Message_Humour                     `protobuf:"varint,2,opt,name=hilarity,enum=proto3_proto.Message_Humour" json:"hilarity,omitempty"`
	HeightInCm           uint32                             `protobuf:"varint,3,opt,name=height_in_cm,json=heightInCm" json:"height_in_cm,omitempty"`
	Data                 []byte                             `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
	ResultCount          int64                              `protobuf:"varint,7,opt,name=result_count,json=resultCount" json:"result_count,omitempty"`
	TrueScotsman         bool                               `protobuf:"varint,8,opt,name=true_scotsman,json=trueScotsman" json:"true_scotsman,omitempty"`
	Score                float32                            `protobuf:"fixed32,9,opt,name=score" json:"score,omitempty"`
	Key                  []uint64                           `protobuf:"varint,5,rep,packed,name=key" json:"key,omitempty"`
	ShortKey             []int32                            `protobuf:"varint,19,rep,packed,name=short_key,json=shortKey" json:"short_key,omitempty"`
	Nested               *Nested                            `protobuf:"bytes,6,opt,name=nested" json:"nested,omitempty"`
	RFunny               []Message_Humour                   `protobuf:"varint,16,rep,packed,name=r_funny,json=rFunny,enum=proto3_proto.Message_Humour" json:"r_funny,omitempty"`
	Terrain              map[string]*Nested                 `protobuf:"bytes,10,rep,name=terrain" json:"terrain,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Proto2Field          *test_proto.SubDefaults            `protobuf:"bytes,11,opt,name=proto2_field,json=proto2Field" json:"proto2_field,omitempty"`
	Proto2Value          map[string]*test_proto.SubDefaults `protobuf:"bytes,13,rep,name=proto2_value,json=proto2Value" json:"proto2_value,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Anything             *any.Any                           `protobuf:"bytes,14,opt,name=anything" json:"anything,omitempty"`
	ManyThings           []*any.Any                         `protobuf:"bytes,15,rep,name=many_things,json=manyThings" json:"many_things,omitempty"`
	Submessage           *Message                           `protobuf:"bytes,17,opt,name=submessage" json:"submessage,omitempty"`
	Children             []*Message                         `protobuf:"bytes,18,rep,name=children" json:"children,omitempty"`
	StringMap            map[string]string                  `protobuf:"bytes,20,rep,name=string_map,json=stringMap" json:"string_map,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	XXX_NoUnkeyedLiteral struct{}                           `json:"-"`
	XXX_unrecognized     []byte                             `json:"-"`
	XXX_sizecache        int32                              `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_proto3_e706e4ff19a5dbea, []int{0}
}
func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (dst *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(dst, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Message) GetHilarity() Message_Humour {
	if m != nil {
		return m.Hilarity
	}
	return Message_UNKNOWN
}

func (m *Message) GetHeightInCm() uint32 {
	if m != nil {
		return m.HeightInCm
	}
	return 0
}

func (m *Message) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Message) GetResultCount() int64 {
	if m != nil {
		return m.ResultCount
	}
	return 0
}

func (m *Message) GetTrueScotsman() bool {
	if m != nil {
		return m.TrueScotsman
	}
	return false
}

func (m *Message) GetScore() float32 {
	if m != nil {
		return m.Score
	}
	return 0
}

func (m *Message) GetKey() []uint64 {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *Message) GetShortKey() []int32 {
	if m != nil {
		return m.ShortKey
	}
	return nil
}

func (m *Message) GetNested() *Nested {
	if m != nil {
		return m.Nested
	}
	return nil
}

func (m *Message) GetRFunny() []Message_Humour {
	if m != nil {
		return m.RFunny
	}
	return nil
}

func (m *Message) GetTerrain() map[string]*Nested {
	if m != nil {
		return m.Terrain
	}
	return nil
}

func (m *Message) GetProto2Field() *test_proto.SubDefaults {
	if m != nil {
		return m.Proto2Field
	}
	return nil
}

func (m *Message) GetProto2Value() map[string]*test_proto.SubDefaults {
	if m != nil {
		return m.Proto2Value
	}
	return nil
}

func (m *Message) GetAnything() *any.Any {
	if m != nil {
		return m.Anything
	}
	return nil
}

func (m *Message) GetManyThings() []*any.Any {
	if m != nil {
		return m.ManyThings
	}
	return nil
}

func (m *Message) GetSubmessage() *Message {
	if m != nil {
		return m.Submessage
	}
	return nil
}

func (m *Message) GetChildren() []*Message {
	if m != nil {
		return m.Children
	}
	return nil
}

func (m *Message) GetStringMap() map[string]string {
	if m != nil {
		return m.StringMap
	}
	return nil
}

type Nested struct {
	Bunny                string   `protobuf:"bytes,1,opt,name=bunny" json:"bunny,omitempty"`
	Cute                 bool     `protobuf:"varint,2,opt,name=cute" json:"cute,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Nested) Reset()         { *m = Nested{} }
func (m *Nested) String() string { return proto.CompactTextString(m) }
func (*Nested) ProtoMessage()    {}
func (*Nested) Descriptor() ([]byte, []int) {
	return fileDescriptor_proto3_e706e4ff19a5dbea, []int{1}
}
func (m *Nested) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Nested.Unmarshal(m, b)
}
func (m *Nested) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Nested.Marshal(b, m, deterministic)
}
func (dst *Nested) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Nested.Merge(dst, src)
}
func (m *Nested) XXX_Size() int {
	return xxx_messageInfo_Nested.Size(m)
}
func (m *Nested) XXX_DiscardUnknown() {
	xxx_messageInfo_Nested.DiscardUnknown(m)
}

var xxx_messageInfo_Nested proto.InternalMessageInfo

func (m *Nested) GetBunny() string {
	if m != nil {
		return m.Bunny
	}
	return ""
}

func (m *Nested) GetCute() bool {
	if m != nil {
		return m.Cute
	}
	return false
}

type MessageWithMap struct {
	ByteMapping          map[bool][]byte `protobuf:"bytes,1,rep,name=byte_mapping,json=byteMapping" json:"byte_mapping,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *MessageWithMap) Reset()         { *m = MessageWithMap{} }
func (m *MessageWithMap) String() string { return proto.CompactTextString(m) }
func (*MessageWithMap) ProtoMessage()    {}
func (*MessageWithMap) Descriptor() ([]byte, []int) {
	return fileDescriptor_proto3_e706e4ff19a5dbea, []int{2}
}
func (m *MessageWithMap) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MessageWithMap.Unmarshal(m, b)
}
func (m *MessageWithMap) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MessageWithMap.Marshal(b, m, deterministic)
}
func (dst *MessageWithMap) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessageWithMap.Merge(dst, src)
}
func (m *MessageWithMap) XXX_Size() int {
	return xxx_messageInfo_MessageWithMap.Size(m)
}
func (m *MessageWithMap) XXX_DiscardUnknown() {
	xxx_messageInfo_MessageWithMap.DiscardUnknown(m)
}

var xxx_messageInfo_MessageWithMap proto.InternalMessageInfo

func (m *MessageWithMap) GetByteMapping() map[bool][]byte {
	if m != nil {
		return m.ByteMapping
	}
	return nil
}

type IntMap struct {
	Rtt                  map[int32]int32 `protobuf:"bytes,1,rep,name=rtt" json:"rtt,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *IntMap) Reset()         { *m = IntMap{} }
func (m *IntMap) String() string { return proto.CompactTextString(m) }
func (*IntMap) ProtoMessage()    {}
func (*IntMap) Descriptor() ([]byte, []int) {
	return fileDescriptor_proto3_e706e4ff19a5dbea, []int{3}
}
func (m *IntMap) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IntMap.Unmarshal(m, b)
}
func (m *IntMap) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IntMap.Marshal(b, m, deterministic)
}
func (dst *IntMap) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IntMap.Merge(dst, src)
}
func (m *IntMap) XXX_Size() int {
	return xxx_messageInfo_IntMap.Size(m)
}
func (m *IntMap) XXX_DiscardUnknown() {
	xxx_messageInfo_IntMap.DiscardUnknown(m)
}

var xxx_messageInfo_IntMap proto.InternalMessageInfo

func (m *IntMap) GetRtt() map[int32]int32 {
	if m != nil {
		return m.Rtt
	}
	return nil
}

type IntMaps struct {
	Maps                 []*IntMap `protobuf:"bytes,1,rep,name=maps" json:"maps,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *IntMaps) Reset()         { *m = IntMaps{} }
func (m *IntMaps) String() string { return proto.CompactTextString(m) }
func (*IntMaps) ProtoMessage()    {}
func (*IntMaps) Descriptor() ([]byte, []int) {
	return fileDescriptor_proto3_e706e4ff19a5dbea, []int{4}
}
func (m *IntMaps) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IntMaps.Unmarshal(m, b)
}
func (m *IntMaps) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IntMaps.Marshal(b, m, deterministic)
}
func (dst *IntMaps) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IntMaps.Merge(dst, src)
}
func (m *IntMaps) XXX_Size() int {
	return xxx_messageInfo_IntMaps.Size(m)
}
func (m *IntMaps) XXX_DiscardUnknown() {
	xxx_messageInfo_IntMaps.DiscardUnknown(m)
}

var xxx_messageInfo_IntMaps proto.InternalMessageInfo

func (m *IntMaps) GetMaps() []*IntMap {
	if m != nil {
		return m.Maps
	}
	return nil
}

func init() {
	proto.RegisterType((*Message)(nil), "proto3_proto.Message")
	proto.RegisterMapType((map[string]*test_proto.SubDefaults)(nil), "proto3_proto.Message.Proto2ValueEntry")
	proto.RegisterMapType((map[string]string)(nil), "proto3_proto.Message.StringMapEntry")
	proto.RegisterMapType((map[string]*Nested)(nil), "proto3_proto.Message.TerrainEntry")
	proto.RegisterType((*Nested)(nil), "proto3_proto.Nested")
	proto.RegisterType((*MessageWithMap)(nil), "proto3_proto.MessageWithMap")
	proto.RegisterMapType((map[bool][]byte)(nil), "proto3_proto.MessageWithMap.ByteMappingEntry")
	proto.RegisterType((*IntMap)(nil), "proto3_proto.IntMap")
	proto.RegisterMapType((map[int32]int32)(nil), "proto3_proto.IntMap.RttEntry")
	proto.RegisterType((*IntMaps)(nil), "proto3_proto.IntMaps")
	proto.RegisterEnum("proto3_proto.Message_Humour", Message_Humour_name, Message_Humour_value)
}

func init() { proto.RegisterFile("proto3_proto/proto3.proto", fileDescriptor_proto3_e706e4ff19a5dbea) }

var fileDescriptor_proto3_e706e4ff19a5dbea = []byte{
	// 774 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x94, 0x6f, 0x8f, 0xdb, 0x44,
	0x10, 0xc6, 0x71, 0x9c, 0x3f, 0xce, 0xd8, 0x77, 0x35, 0x4b, 0x2a, 0xb6, 0x01, 0x24, 0x13, 0x10,
	0xb2, 0x10, 0xf5, 0x41, 0xaa, 0x43, 0x55, 0x55, 0x81, 0xee, 0x8e, 0x56, 0x44, 0x77, 0x17, 0xa2,
	0xcd, 0x95, 0x13, 0xaf, 0xac, 0x4d, 0x6e, 0x93, 0x58, 0xc4, 0xeb, 0xe0, 0x5d, 0x23, 0xf9, 0x0b,
	0xf0, 0x41, 0xf8, 0xa4, 0x68, 0x77, 0x9d, 0xd4, 0xa9, 0x5c, 0xfa, 0x2a, 0xbb, 0x8f, 0x7f, 0x33,
	0xcf, 0x78, 0x66, 0x1c, 0x78, 0xb2, 0xcb, 0x33, 0x99, 0x3d, 0x8b, 0xf5, 0xcf, 0x99, 0xb9, 0x44,
	0xfa, 0x07, 0x79, 0xf5, 0x47, 0xc3, 0x27, 0xeb, 0x2c, 0x5b, 0x6f, 0x99, 0x41, 0x16, 0xc5, 0xea,
	0x8c, 0xf2, 0xd2, 0x80, 0xc3, 0xc7, 0x92, 0x09, 0x59, 0x65, 0x50, 0x47, 0x23, 0x8f, 0xfe, 0xe9,
	0x43, 0xef, 0x96, 0x09, 0x41, 0xd7, 0x0c, 0x21, 0x68, 0x73, 0x9a, 0x32, 0x6c, 0x05, 0x56, 0xd8,
	0x27, 0xfa, 0x8c, 0x9e, 0x83, 0xb3, 0x49, 0xb6, 0x34, 0x4f, 0x64, 0x89, 0x5b, 0x81, 0x15, 0x9e,
	0x8e, 0x3f, 0x8f, 0xea, 0x96, 0x51, 0x15, 0x1c, 0xfd, 0x5a, 0xa4, 0x59, 0x91, 0x93, 0x03, 0x8d,
	0x02, 0xf0, 0x36, 0x2c, 0x59, 0x6f, 0x64, 0x9c, 0xf0, 0x78, 0x99, 0x62, 0x3b, 0xb0, 0xc2, 0x13,
	0x02, 0x46, 0x9b, 0xf0, 0xab, 0x54, 0xf9, 0x3d, 0x50, 0x49, 0x71, 0x3b, 0xb0, 0x42, 0x8f, 0xe8,
	0x33, 0xfa, 0x12, 0xbc, 0x9c, 0x89, 0x62, 0x2b, 0xe3, 0x65, 0x56, 0x70, 0x89, 0x7b, 0x81, 0x15,
	0xda, 0xc4, 0x35, 0xda, 0x95, 0x92, 0xd0, 0x57, 0x70, 0x22, 0xf3, 0x82, 0xc5, 0x62, 0x99, 0x49,
	0x91, 0x52, 0x8e, 0x9d, 0xc0, 0x0a, 0x1d, 0xe2, 0x29, 0x71, 0x5e, 0x69, 0x68, 0x00, 0x1d, 0xb1,
	0xcc, 0x72, 0x86, 0xfb, 0x81, 0x15, 0xb6, 0x88, 0xb9, 0x20, 0x1f, 0xec, 0x3f, 0x59, 0x89, 0x3b,
	0x81, 0x1d, 0xb6, 0x89, 0x3a, 0xa2, 0xcf, 0xa0, 0x2f, 0x36, 0x59, 0x2e, 0x63, 0xa5, 0x7f, 0x12,
	0xd8, 0x61, 0x87, 0x38, 0x5a, 0xb8, 0x66, 0x25, 0xfa, 0x0e, 0xba, 0x9c, 0x09, 0xc9, 0x1e, 0x70,
	0x37, 0xb0, 0x42, 0x77, 0x3c, 0x38, 0x7e, 0xf5, 0xa9, 0x7e, 0x46, 0x2a, 0x06, 0x9d, 0x43, 0x2f,
	0x8f, 0x57, 0x05, 0xe7, 0x25, 0xf6, 0x03, 0xfb, 0x83, 0x9d, 0xea, 0xe6, 0xaf, 0x15, 0x8b, 0x5e,
	0x42, 0x4f, 0xb2, 0x3c, 0xa7, 0x09, 0xc7, 0x10, 0xd8, 0xa1, 0x3b, 0x1e, 0x35, 0x87, 0xdd, 0x19,
	0xe8, 0x15, 0x97, 0x79, 0x49, 0xf6, 0x21, 0xe8, 0x05, 0x98, 0x0d, 0x18, 0xc7, 0xab, 0x84, 0x6d,
	0x1f, 0xb0, 0xab, 0x0b, 0xfd, 0x34, 0x7a, 0x3b, 0xed, 0x68, 0x5e, 0x2c, 0x7e, 0x61, 0x2b, 0x5a,
	0x6c, 0xa5, 0x20, 0xae, 0x81, 0x5f, 0x2b, 0x16, 0x4d, 0x0e, 0xb1, 0x7f, 0xd3, 0x6d, 0xc1, 0xf0,
	0x89, 0xb6, 0xff, 0xa6, 0xd9, 0x7e, 0xa6, 0xc9, 0xdf, 0x15, 0x68, 0x4a, 0xa8, 0x52, 0x69, 0x05,
	0x7d, 0x0f, 0x0e, 0xe5, 0xa5, 0xdc, 0x24, 0x7c, 0x8d, 0x4f, 0xab, 0x5e, 0x99, 0x5d, 0x8c, 0xf6,
	0xbb, 0x18, 0x5d, 0xf0, 0x92, 0x1c, 0x28, 0x74, 0x0e, 0x6e, 0x4a, 0x79, 0x19, 0xeb, 0x9b, 0xc0,
	0x8f, 0xb4, 0x77, 0x73, 0x10, 0x28, 0xf0, 0x4e, 0x73, 0xe8, 0x1c, 0x40, 0x14, 0x8b, 0xd4, 0x14,
	0x85, 0x3f, 0xd6, 0x56, 0x8f, 0x1b, 0x2b, 0x26, 0x35, 0x10, 0xfd, 0x00, 0xce, 0x72, 0x93, 0x6c,
	0x1f, 0x72, 0xc6, 0x31, 0xd2, 0x56, 0xef, 0x09, 0x3a, 0x60, 0xe8, 0x0a, 0x40, 0xc8, 0x3c, 0xe1,
	0xeb, 0x38, 0xa5, 0x3b, 0x3c, 0xd0, 0x41, 0x5f, 0x37, 0xf7, 0x66, 0xae, 0xb9, 0x5b, 0xba, 0x33,
	0x9d, 0xe9, 0x8b, 0xfd, 0x7d, 0x38, 0x03, 0xaf, 0x3e, 0xb7, 0xfd, 0x02, 0x9a, 0x2f, 0x4c, 0x2f,
	0xe0, 0xb7, 0xd0, 0x31, 0xdd, 0x6f, 0xfd, 0xcf, 0x8a, 0x19, 0xe4, 0x45, 0xeb, 0xb9, 0x35, 0xbc,
	0x07, 0xff, 0xdd, 0x51, 0x34, 0x64, 0x7d, 0x7a, 0x9c, 0xf5, 0xbd, 0xfb, 0x50, 0x4b, 0xfc, 0x12,
	0x4e, 0x8f, 0xdf, 0xa3, 0x21, 0xed, 0xa0, 0x9e, 0xb6, 0x5f, 0x8b, 0x1e, 0xfd, 0x0c, 0x5d, 0xb3,
	0xd7, 0xc8, 0x85, 0xde, 0x9b, 0xe9, 0xf5, 0xf4, 0xb7, 0xfb, 0xa9, 0xff, 0x11, 0x72, 0xa0, 0x3d,
	0x7b, 0x33, 0x9d, 0xfb, 0x16, 0x3a, 0x81, 0xfe, 0xfc, 0xe6, 0x62, 0x36, 0xbf, 0x9b, 0x5c, 0x5d,
	0xfb, 0x2d, 0xf4, 0x08, 0xdc, 0xcb, 0xc9, 0xcd, 0x4d, 0x7c, 0x79, 0x31, 0xb9, 0x79, 0xf5, 0x87,
	0x6f, 0x8f, 0xc6, 0xd0, 0x35, 0x2f, 0xab, 0x4c, 0x16, 0xfa, 0x2b, 0x32, 0xc6, 0xe6, 0xa2, 0xfe,
	0x2c, 0x96, 0x85, 0x34, 0xce, 0x0e, 0xd1, 0xe7, 0xd1, 0xbf, 0x16, 0x9c, 0x56, 0x33, 0xb8, 0x4f,
	0xe4, 0xe6, 0x96, 0xee, 0xd0, 0x0c, 0xbc, 0x45, 0x29, 0x99, 0x9a, 0xd9, 0x4e, 0x2d, 0xa3, 0xa5,
	0xe7, 0xf6, 0xb4, 0x71, 0x6e, 0x55, 0x4c, 0x74, 0x59, 0x4a, 0x76, 0x6b, 0xf8, 0x6a, 0xb5, 0x17,
	0x6f, 0x95, 0xe1, 0x4f, 0xe0, 0xbf, 0x0b, 0xd4, 0x3b, 0xe3, 0x34, 0x74, 0xc6, 0xab, 0x77, 0xe6,
	0x2f, 0xe8, 0x4e, 0xb8, 0x54, 0xb5, 0x9d, 0x81, 0x9d, 0x4b, 0x59, 0x95, 0xf4, 0xc5, 0x71, 0x49,
	0x06, 0x89, 0x88, 0x94, 0xa6, 0x04, 0x45, 0x0e, 0x7f, 0x04, 0x67, 0x2f, 0xd4, 0x2d, 0x3b, 0x0d,
	0x96, 0x9d, 0xba, 0xe5, 0x33, 0xe8, 0x99, 0x7c, 0x02, 0x85, 0xd0, 0x4e, 0xe9, 0x4e, 0x54, 0xa6,
	0x83, 0x26, 0x53, 0xa2, 0x89, 0x45, 0xd7, 0x3c, 0xfa, 0x2f, 0x00, 0x00, 0xff, 0xff, 0x99, 0x24,
	0x6b, 0x12, 0x6d, 0x06, 0x00, 0x00,
}