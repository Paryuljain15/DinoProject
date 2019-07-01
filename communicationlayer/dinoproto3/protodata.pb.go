// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protodata.proto

package dinoproto3

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

//in proto 3 we dont need the required r optional
type Animal struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	AnimalType           string   `protobuf:"bytes,2,opt,name=animal_type,json=animalType,proto3" json:"animal_type,omitempty"`
	Nickname             string   `protobuf:"bytes,3,opt,name=nickname,proto3" json:"nickname,omitempty"`
	Zone                 int32    `protobuf:"varint,4,opt,name=zone,proto3" json:"zone,omitempty"`
	Age                  int32    `protobuf:"varint,5,opt,name=age,proto3" json:"age,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Animal) Reset()         { *m = Animal{} }
func (m *Animal) String() string { return proto.CompactTextString(m) }
func (*Animal) ProtoMessage()    {}
func (*Animal) Descriptor() ([]byte, []int) {
	return fileDescriptor_cb955849bc7471cb, []int{0}
}

func (m *Animal) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Animal.Unmarshal(m, b)
}
func (m *Animal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Animal.Marshal(b, m, deterministic)
}
func (m *Animal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Animal.Merge(m, src)
}
func (m *Animal) XXX_Size() int {
	return xxx_messageInfo_Animal.Size(m)
}
func (m *Animal) XXX_DiscardUnknown() {
	xxx_messageInfo_Animal.DiscardUnknown(m)
}

var xxx_messageInfo_Animal proto.InternalMessageInfo

func (m *Animal) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Animal) GetAnimalType() string {
	if m != nil {
		return m.AnimalType
	}
	return ""
}

func (m *Animal) GetNickname() string {
	if m != nil {
		return m.Nickname
	}
	return ""
}

func (m *Animal) GetZone() int32 {
	if m != nil {
		return m.Zone
	}
	return 0
}

func (m *Animal) GetAge() int32 {
	if m != nil {
		return m.Age
	}
	return 0
}

func init() {
	proto.RegisterType((*Animal)(nil), "dinoproto3.animal")
}

func init() { proto.RegisterFile("protodata.proto", fileDescriptor_cb955849bc7471cb) }

var fileDescriptor_cb955849bc7471cb = []byte{
	// 146 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2f, 0x28, 0xca, 0x2f,
	0xc9, 0x4f, 0x49, 0x2c, 0x49, 0xd4, 0x03, 0xb3, 0x84, 0xb8, 0x52, 0x32, 0xf3, 0xf2, 0xc1, 0x4c,
	0x63, 0xa5, 0x6a, 0x2e, 0xb6, 0xc4, 0xbc, 0xcc, 0xdc, 0xc4, 0x1c, 0x21, 0x3e, 0x2e, 0xa6, 0xcc,
	0x14, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xd6, 0x20, 0xa6, 0xcc, 0x14, 0x21, 0x79, 0x2e, 0x6e, 0x88,
	0x4c, 0x7c, 0x49, 0x65, 0x41, 0xaa, 0x04, 0x93, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x17, 0x44, 0x28,
	0xa4, 0xb2, 0x20, 0x55, 0x48, 0x8a, 0x8b, 0x23, 0x2f, 0x33, 0x39, 0x3b, 0x2f, 0x31, 0x37, 0x55,
	0x82, 0x19, 0x2c, 0x0b, 0xe7, 0x0b, 0x09, 0x71, 0xb1, 0x54, 0xe5, 0xe7, 0xa5, 0x4a, 0xb0, 0x80,
	0x8d, 0x03, 0xb3, 0x85, 0x04, 0xb8, 0x98, 0x13, 0xd3, 0x53, 0x25, 0x58, 0xc1, 0x42, 0x20, 0x66,
	0x12, 0x1b, 0xc4, 0x11, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x2e, 0x1e, 0xd9, 0x14, 0xa2, 0x00,
	0x00, 0x00,
}
