// Code generated by protoc-gen-go.
// source: msg.proto
// DO NOT EDIT!

/*
Package ProtobufTest is a generated protocol buffer package.

It is generated from these files:
	msg.proto

It has these top-level messages:
	TestMessage
*/
package ProtobufTest

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type TestMessage_ItemType int32

const (
	TestMessage_TypeX TestMessage_ItemType = 0
	TestMessage_TypeY TestMessage_ItemType = 1
	TestMessage_TypeZ TestMessage_ItemType = 2
)

var TestMessage_ItemType_name = map[int32]string{
	0: "TypeX",
	1: "TypeY",
	2: "TypeZ",
}
var TestMessage_ItemType_value = map[string]int32{
	"TypeX": 0,
	"TypeY": 1,
	"TypeZ": 2,
}

func (x TestMessage_ItemType) Enum() *TestMessage_ItemType {
	p := new(TestMessage_ItemType)
	*p = x
	return p
}
func (x TestMessage_ItemType) String() string {
	return proto.EnumName(TestMessage_ItemType_name, int32(x))
}
func (x *TestMessage_ItemType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(TestMessage_ItemType_value, data, "TestMessage_ItemType")
	if err != nil {
		return err
	}
	*x = TestMessage_ItemType(value)
	return nil
}
func (TestMessage_ItemType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

type TestMessage struct {
	ClientName       *string                `protobuf:"bytes,1,req,name=clientName" json:"clientName,omitempty"`
	ClientId         *int32                 `protobuf:"varint,2,req,name=clientId" json:"clientId,omitempty"`
	Description      *string                `protobuf:"bytes,3,opt,name=description,def=NONE" json:"description,omitempty"`
	Messageitems     []*TestMessage_MsgItem `protobuf:"bytes,4,rep,name=messageitems" json:"messageitems,omitempty"`
	XXX_unrecognized []byte                 `json:"-"`
}

func (m *TestMessage) Reset()                    { *m = TestMessage{} }
func (m *TestMessage) String() string            { return proto.CompactTextString(m) }
func (*TestMessage) ProtoMessage()               {}
func (*TestMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

const Default_TestMessage_Description string = "NONE"

func (m *TestMessage) GetClientName() string {
	if m != nil && m.ClientName != nil {
		return *m.ClientName
	}
	return ""
}

func (m *TestMessage) GetClientId() int32 {
	if m != nil && m.ClientId != nil {
		return *m.ClientId
	}
	return 0
}

func (m *TestMessage) GetDescription() string {
	if m != nil && m.Description != nil {
		return *m.Description
	}
	return Default_TestMessage_Description
}

func (m *TestMessage) GetMessageitems() []*TestMessage_MsgItem {
	if m != nil {
		return m.Messageitems
	}
	return nil
}

type TestMessage_MsgItem struct {
	Id               *int32                `protobuf:"varint,1,req,name=id" json:"id,omitempty"`
	ItemName         *string               `protobuf:"bytes,2,opt,name=itemName" json:"itemName,omitempty"`
	ItemValue        *int32                `protobuf:"varint,3,opt,name=itemValue" json:"itemValue,omitempty"`
	ItemType         *TestMessage_ItemType `protobuf:"varint,4,opt,name=itemType,enum=ProtobufTest.TestMessage_ItemType" json:"itemType,omitempty"`
	XXX_unrecognized []byte                `json:"-"`
}

func (m *TestMessage_MsgItem) Reset()                    { *m = TestMessage_MsgItem{} }
func (m *TestMessage_MsgItem) String() string            { return proto.CompactTextString(m) }
func (*TestMessage_MsgItem) ProtoMessage()               {}
func (*TestMessage_MsgItem) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

func (m *TestMessage_MsgItem) GetId() int32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *TestMessage_MsgItem) GetItemName() string {
	if m != nil && m.ItemName != nil {
		return *m.ItemName
	}
	return ""
}

func (m *TestMessage_MsgItem) GetItemValue() int32 {
	if m != nil && m.ItemValue != nil {
		return *m.ItemValue
	}
	return 0
}

func (m *TestMessage_MsgItem) GetItemType() TestMessage_ItemType {
	if m != nil && m.ItemType != nil {
		return *m.ItemType
	}
	return TestMessage_TypeX
}

func init() {
	proto.RegisterType((*TestMessage)(nil), "ProtobufTest.TestMessage")
	proto.RegisterType((*TestMessage_MsgItem)(nil), "ProtobufTest.TestMessage.MsgItem")
	proto.RegisterEnum("ProtobufTest.TestMessage_ItemType", TestMessage_ItemType_name, TestMessage_ItemType_value)
}

func init() { proto.RegisterFile("msg.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 236 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x74, 0x8e, 0x31, 0x4f, 0xc3, 0x30,
	0x10, 0x85, 0x89, 0xdb, 0x88, 0xfa, 0x52, 0xa1, 0x70, 0x93, 0x61, 0x0a, 0x99, 0x2c, 0x21, 0x79,
	0xa8, 0x90, 0x90, 0xd8, 0x19, 0x32, 0x34, 0x30, 0x54, 0x08, 0xd8, 0x42, 0x73, 0x44, 0x96, 0xea,
	0xc6, 0xaa, 0xdd, 0x81, 0x3f, 0xc4, 0xef, 0xc4, 0x4e, 0x1a, 0xa9, 0x0b, 0xd3, 0x7d, 0x7a, 0x7a,
	0xba, 0xef, 0x01, 0x37, 0xae, 0x53, 0xf6, 0xd0, 0xfb, 0x1e, 0x97, 0xaf, 0xf1, 0x7c, 0x1d, 0xbf,
	0x37, 0xe4, 0x7c, 0xf9, 0xcb, 0x20, 0x8b, 0xb0, 0x26, 0xe7, 0x9a, 0x8e, 0x10, 0x01, 0xb6, 0x3b,
	0x4d, 0x7b, 0x5f, 0x37, 0x86, 0x44, 0x52, 0x30, 0xc9, 0x31, 0x87, 0xc5, 0x98, 0x55, 0xad, 0x60,
	0x21, 0x49, 0xf1, 0x06, 0xb2, 0x96, 0xdc, 0xf6, 0xa0, 0xad, 0xd7, 0xfd, 0x5e, 0xcc, 0x8a, 0x44,
	0xf2, 0xa7, 0x79, 0xfd, 0x52, 0x3f, 0xe3, 0x23, 0x2c, 0xcd, 0xf8, 0x4b, 0x7b, 0x32, 0x4e, 0xcc,
	0x8b, 0x99, 0xcc, 0x56, 0x77, 0xea, 0xdc, 0xaa, 0xce, 0x8c, 0x6a, 0xed, 0xba, 0x2a, 0x34, 0x6f,
	0x2d, 0x5c, 0x9e, 0x30, 0x6c, 0x60, 0xba, 0x1d, 0xe4, 0x69, 0x94, 0xc7, 0x47, 0xc3, 0x1c, 0x16,
	0x3d, 0x78, 0x0d, 0x3c, 0x26, 0x6f, 0xcd, 0xee, 0x48, 0x83, 0x3a, 0xc5, 0x87, 0xb1, 0xb4, 0xf9,
	0xb1, 0x14, 0x84, 0x89, 0xbc, 0x5a, 0x95, 0xff, 0x0b, 0xab, 0x53, 0xb3, 0xbc, 0x87, 0xc5, 0xc4,
	0xc8, 0x21, 0x8d, 0xf7, 0x3d, 0xbf, 0x98, 0xf0, 0x23, 0x4f, 0x26, 0xfc, 0xcc, 0xd9, 0x5f, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x1f, 0x4b, 0xb8, 0xdf, 0x42, 0x01, 0x00, 0x00,
}
