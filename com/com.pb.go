// Code generated by protoc-gen-go.
// source: com.proto
// DO NOT EDIT!

/*
Package com is a generated protocol buffer package.

It is generated from these files:
	com.proto

It has these top-level messages:
	UserInfo
*/
package com

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

type UserInfo struct {
	Id               *string `protobuf:"bytes,1,req,name=id" json:"id,omitempty"`
	Nickname         *string `protobuf:"bytes,2,req,name=nickname" json:"nickname,omitempty"`
	Coin             *int32  `protobuf:"varint,3,opt,name=coin" json:"coin,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *UserInfo) Reset()                    { *m = UserInfo{} }
func (m *UserInfo) String() string            { return proto.CompactTextString(m) }
func (*UserInfo) ProtoMessage()               {}
func (*UserInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *UserInfo) GetId() string {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return ""
}

func (m *UserInfo) GetNickname() string {
	if m != nil && m.Nickname != nil {
		return *m.Nickname
	}
	return ""
}

func (m *UserInfo) GetCoin() int32 {
	if m != nil && m.Coin != nil {
		return *m.Coin
	}
	return 0
}

func init() {
	proto.RegisterType((*UserInfo)(nil), "com.UserInfo")
}

var fileDescriptor0 = []byte{
	// 100 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x4c, 0xce, 0xcf, 0xd5,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x06, 0x32, 0x95, 0xbc, 0xb8, 0x38, 0x42, 0x8b, 0x53,
	0x8b, 0x3c, 0xf3, 0xd2, 0xf2, 0x85, 0xf8, 0xb8, 0x98, 0x32, 0x53, 0x24, 0x18, 0x15, 0x98, 0x34,
	0x38, 0x83, 0x80, 0x2c, 0x21, 0x29, 0x2e, 0x8e, 0xbc, 0xcc, 0xe4, 0xec, 0xbc, 0xc4, 0xdc, 0x54,
	0x09, 0x26, 0xb0, 0x28, 0x9c, 0x2f, 0x24, 0xc4, 0xc5, 0x92, 0x9c, 0x9f, 0x99, 0x27, 0xc1, 0xac,
	0xc0, 0xa8, 0xc1, 0x1a, 0x04, 0x66, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0xaf, 0x94, 0x1d, 0x49,
	0x5c, 0x00, 0x00, 0x00,
}