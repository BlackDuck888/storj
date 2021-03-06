// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: overlay.proto

package pb

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Restriction_Operator int32

const (
	Restriction_LT  Restriction_Operator = 0
	Restriction_EQ  Restriction_Operator = 1
	Restriction_GT  Restriction_Operator = 2
	Restriction_LTE Restriction_Operator = 3
	Restriction_GTE Restriction_Operator = 4
)

var Restriction_Operator_name = map[int32]string{
	0: "LT",
	1: "EQ",
	2: "GT",
	3: "LTE",
	4: "GTE",
}

var Restriction_Operator_value = map[string]int32{
	"LT":  0,
	"EQ":  1,
	"GT":  2,
	"LTE": 3,
	"GTE": 4,
}

func (x Restriction_Operator) String() string {
	return proto.EnumName(Restriction_Operator_name, int32(x))
}

func (Restriction_Operator) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_61fc82527fbe24ad, []int{1, 0}
}

type Restriction_Operand int32

const (
	Restriction_FREE_BANDWIDTH Restriction_Operand = 0
	Restriction_FREE_DISK      Restriction_Operand = 1
)

var Restriction_Operand_name = map[int32]string{
	0: "FREE_BANDWIDTH",
	1: "FREE_DISK",
}

var Restriction_Operand_value = map[string]int32{
	"FREE_BANDWIDTH": 0,
	"FREE_DISK":      1,
}

func (x Restriction_Operand) String() string {
	return proto.EnumName(Restriction_Operand_name, int32(x))
}

func (Restriction_Operand) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_61fc82527fbe24ad, []int{1, 1}
}

type InfoResponse struct {
	Type                 NodeType      `protobuf:"varint,2,opt,name=type,proto3,enum=node.NodeType" json:"type,omitempty"`
	Operator             *NodeOperator `protobuf:"bytes,3,opt,name=operator,proto3" json:"operator,omitempty"`
	Capacity             *NodeCapacity `protobuf:"bytes,4,opt,name=capacity,proto3" json:"capacity,omitempty"`
	Version              *NodeVersion  `protobuf:"bytes,5,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *InfoResponse) Reset()         { *m = InfoResponse{} }
func (m *InfoResponse) String() string { return proto.CompactTextString(m) }
func (*InfoResponse) ProtoMessage()    {}
func (*InfoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_61fc82527fbe24ad, []int{0}
}
func (m *InfoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InfoResponse.Unmarshal(m, b)
}
func (m *InfoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InfoResponse.Marshal(b, m, deterministic)
}
func (m *InfoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InfoResponse.Merge(m, src)
}
func (m *InfoResponse) XXX_Size() int {
	return xxx_messageInfo_InfoResponse.Size(m)
}
func (m *InfoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_InfoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_InfoResponse proto.InternalMessageInfo

func (m *InfoResponse) GetType() NodeType {
	if m != nil {
		return m.Type
	}
	return NodeType_INVALID
}

func (m *InfoResponse) GetOperator() *NodeOperator {
	if m != nil {
		return m.Operator
	}
	return nil
}

func (m *InfoResponse) GetCapacity() *NodeCapacity {
	if m != nil {
		return m.Capacity
	}
	return nil
}

func (m *InfoResponse) GetVersion() *NodeVersion {
	if m != nil {
		return m.Version
	}
	return nil
}

type Restriction struct {
	Operator             Restriction_Operator `protobuf:"varint,1,opt,name=operator,proto3,enum=overlay.Restriction_Operator" json:"operator,omitempty"`
	Operand              Restriction_Operand  `protobuf:"varint,2,opt,name=operand,proto3,enum=overlay.Restriction_Operand" json:"operand,omitempty"`
	Value                int64                `protobuf:"varint,3,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Restriction) Reset()         { *m = Restriction{} }
func (m *Restriction) String() string { return proto.CompactTextString(m) }
func (*Restriction) ProtoMessage()    {}
func (*Restriction) Descriptor() ([]byte, []int) {
	return fileDescriptor_61fc82527fbe24ad, []int{1}
}
func (m *Restriction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Restriction.Unmarshal(m, b)
}
func (m *Restriction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Restriction.Marshal(b, m, deterministic)
}
func (m *Restriction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Restriction.Merge(m, src)
}
func (m *Restriction) XXX_Size() int {
	return xxx_messageInfo_Restriction.Size(m)
}
func (m *Restriction) XXX_DiscardUnknown() {
	xxx_messageInfo_Restriction.DiscardUnknown(m)
}

var xxx_messageInfo_Restriction proto.InternalMessageInfo

func (m *Restriction) GetOperator() Restriction_Operator {
	if m != nil {
		return m.Operator
	}
	return Restriction_LT
}

func (m *Restriction) GetOperand() Restriction_Operand {
	if m != nil {
		return m.Operand
	}
	return Restriction_FREE_BANDWIDTH
}

func (m *Restriction) GetValue() int64 {
	if m != nil {
		return m.Value
	}
	return 0
}

func init() {
	proto.RegisterEnum("overlay.Restriction_Operator", Restriction_Operator_name, Restriction_Operator_value)
	proto.RegisterEnum("overlay.Restriction_Operand", Restriction_Operand_name, Restriction_Operand_value)
	proto.RegisterType((*InfoResponse)(nil), "overlay.InfoResponse")
	proto.RegisterType((*Restriction)(nil), "overlay.Restriction")
}

func init() { proto.RegisterFile("overlay.proto", fileDescriptor_61fc82527fbe24ad) }

var fileDescriptor_61fc82527fbe24ad = []byte{
	// 325 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0x41, 0x4b, 0xeb, 0x40,
	0x10, 0xc7, 0xbb, 0x49, 0xda, 0xf4, 0x4d, 0x5f, 0xcb, 0xbe, 0xe5, 0x1d, 0x42, 0x51, 0x28, 0x39,
	0x15, 0x94, 0x1c, 0xaa, 0x08, 0x1e, 0xad, 0x8d, 0xb5, 0x58, 0x2a, 0xae, 0x41, 0xc1, 0x8b, 0xa4,
	0xcd, 0x5a, 0x02, 0x65, 0x67, 0x49, 0x62, 0x21, 0x9f, 0xce, 0xef, 0xe5, 0x49, 0xb2, 0x49, 0xda,
	0x22, 0x7a, 0x9a, 0xff, 0xcc, 0xff, 0xb7, 0xbb, 0x33, 0xb3, 0xd0, 0xc5, 0xad, 0x48, 0x36, 0x61,
	0xee, 0xa9, 0x04, 0x33, 0x64, 0x76, 0x95, 0xf6, 0x61, 0x8d, 0x6b, 0x2c, 0x8b, 0x7d, 0x90, 0x18,
	0x89, 0x52, 0xbb, 0x1f, 0x04, 0xfe, 0xce, 0xe4, 0x1b, 0x72, 0x91, 0x2a, 0x94, 0xa9, 0x60, 0x2e,
	0x58, 0x59, 0xae, 0x84, 0x63, 0x0c, 0xc8, 0xb0, 0x37, 0xea, 0x79, 0x9a, 0x5d, 0x60, 0x24, 0x82,
	0x5c, 0x09, 0xae, 0x3d, 0xe6, 0x41, 0x1b, 0x95, 0x48, 0xc2, 0x0c, 0x13, 0xc7, 0x1c, 0x90, 0x61,
	0x67, 0xc4, 0xf6, 0xdc, 0x7d, 0xe5, 0xf0, 0x1d, 0x53, 0xf0, 0xab, 0x50, 0x85, 0xab, 0x38, 0xcb,
	0x1d, 0xeb, 0x3b, 0x7f, 0x5d, 0x39, 0x7c, 0xc7, 0xb0, 0x13, 0xb0, 0xb7, 0x22, 0x49, 0x63, 0x94,
	0x4e, 0x53, 0xe3, 0xff, 0xf6, 0xf8, 0x53, 0x69, 0xf0, 0x9a, 0x70, 0x3f, 0x09, 0x74, 0xb8, 0x48,
	0xb3, 0x24, 0x5e, 0x65, 0x31, 0x4a, 0x76, 0x79, 0xd0, 0x1c, 0xd1, 0x43, 0x1c, 0x7b, 0xf5, 0x52,
	0x0e, 0x38, 0xef, 0x87, 0x3e, 0x2f, 0xc0, 0xd6, 0x5a, 0x46, 0xd5, 0xf8, 0x47, 0xbf, 0x9f, 0x94,
	0x11, 0xaf, 0x61, 0xf6, 0x1f, 0x9a, 0xdb, 0x70, 0xf3, 0x2e, 0xf4, 0x32, 0x4c, 0x5e, 0x26, 0xee,
	0x39, 0xb4, 0xeb, 0x37, 0x58, 0x0b, 0x8c, 0x79, 0x40, 0x1b, 0x45, 0xf4, 0x1f, 0x28, 0x29, 0xe2,
	0x34, 0xa0, 0x06, 0xb3, 0xc1, 0x9c, 0x07, 0x3e, 0x35, 0x0b, 0x31, 0x0d, 0x7c, 0x6a, 0xb9, 0xa7,
	0x60, 0x57, 0xf7, 0x33, 0x06, 0xbd, 0x1b, 0xee, 0xfb, 0xaf, 0xe3, 0xab, 0xc5, 0xe4, 0x79, 0x36,
	0x09, 0x6e, 0x69, 0x83, 0x75, 0xe1, 0x8f, 0xae, 0x4d, 0x66, 0x8f, 0x77, 0x94, 0x8c, 0xad, 0x17,
	0x43, 0x2d, 0x97, 0x2d, 0xfd, 0x97, 0x67, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x39, 0x65, 0xff,
	0x9d, 0xfd, 0x01, 0x00, 0x00,
}
