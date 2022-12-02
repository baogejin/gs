// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: msg.proto

package myproto

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type MsgId int32

const (
	MsgId_Msg_None          MsgId = 0
	MsgId_Msg_RegisterREQ   MsgId = 1
	MsgId_Msg_RegisterACK   MsgId = 2
	MsgId_Msg_LoginREQ      MsgId = 3
	MsgId_Msg_LoginACK      MsgId = 4
	MsgId_Msg_CreateRoleREQ MsgId = 5
	MsgId_Msg_CreateRoleACK MsgId = 6
)

var MsgId_name = map[int32]string{
	0: "Msg_None",
	1: "Msg_RegisterREQ",
	2: "Msg_RegisterACK",
	3: "Msg_LoginREQ",
	4: "Msg_LoginACK",
	5: "Msg_CreateRoleREQ",
	6: "Msg_CreateRoleACK",
}

var MsgId_value = map[string]int32{
	"Msg_None":          0,
	"Msg_RegisterREQ":   1,
	"Msg_RegisterACK":   2,
	"Msg_LoginREQ":      3,
	"Msg_LoginACK":      4,
	"Msg_CreateRoleREQ": 5,
	"Msg_CreateRoleACK": 6,
}

func (x MsgId) String() string {
	return proto.EnumName(MsgId_name, int32(x))
}

func (MsgId) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c06e4cca6c2cc899, []int{0}
}

type ResultCode int32

const (
	ResultCode_Success         ResultCode = 0
	ResultCode_MsgErr          ResultCode = 1
	ResultCode_AccountExist    ResultCode = 101
	ResultCode_AccountNotExist ResultCode = 102
	ResultCode_NickNameExist   ResultCode = 103
	ResultCode_PasswordErr     ResultCode = 104
	ResultCode_AlreadyLogin    ResultCode = 105
	ResultCode_AccountEmpty    ResultCode = 106
	ResultCode_PasswordEmpty   ResultCode = 107
	ResultCode_AccountErr      ResultCode = 108
)

var ResultCode_name = map[int32]string{
	0:   "Success",
	1:   "MsgErr",
	101: "AccountExist",
	102: "AccountNotExist",
	103: "NickNameExist",
	104: "PasswordErr",
	105: "AlreadyLogin",
	106: "AccountEmpty",
	107: "PasswordEmpty",
	108: "AccountErr",
}

var ResultCode_value = map[string]int32{
	"Success":         0,
	"MsgErr":          1,
	"AccountExist":    101,
	"AccountNotExist": 102,
	"NickNameExist":   103,
	"PasswordErr":     104,
	"AlreadyLogin":    105,
	"AccountEmpty":    106,
	"PasswordEmpty":   107,
	"AccountErr":      108,
}

func (x ResultCode) String() string {
	return proto.EnumName(ResultCode_name, int32(x))
}

func (ResultCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c06e4cca6c2cc899, []int{1}
}

type RegisterREQ struct {
	Account  string `protobuf:"bytes,1,opt,name=Account,proto3" json:"Account,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=Password,proto3" json:"Password,omitempty"`
}

func (m *RegisterREQ) Reset()         { *m = RegisterREQ{} }
func (m *RegisterREQ) String() string { return proto.CompactTextString(m) }
func (*RegisterREQ) ProtoMessage()    {}
func (*RegisterREQ) Descriptor() ([]byte, []int) {
	return fileDescriptor_c06e4cca6c2cc899, []int{0}
}
func (m *RegisterREQ) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RegisterREQ) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RegisterREQ.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RegisterREQ) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterREQ.Merge(m, src)
}
func (m *RegisterREQ) XXX_Size() int {
	return m.Size()
}
func (m *RegisterREQ) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterREQ.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterREQ proto.InternalMessageInfo

func (m *RegisterREQ) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

func (m *RegisterREQ) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type RegisterACK struct {
	Ret ResultCode `protobuf:"varint,1,opt,name=Ret,proto3,enum=myproto.ResultCode" json:"Ret,omitempty"`
}

func (m *RegisterACK) Reset()         { *m = RegisterACK{} }
func (m *RegisterACK) String() string { return proto.CompactTextString(m) }
func (*RegisterACK) ProtoMessage()    {}
func (*RegisterACK) Descriptor() ([]byte, []int) {
	return fileDescriptor_c06e4cca6c2cc899, []int{1}
}
func (m *RegisterACK) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RegisterACK) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RegisterACK.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RegisterACK) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterACK.Merge(m, src)
}
func (m *RegisterACK) XXX_Size() int {
	return m.Size()
}
func (m *RegisterACK) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterACK.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterACK proto.InternalMessageInfo

func (m *RegisterACK) GetRet() ResultCode {
	if m != nil {
		return m.Ret
	}
	return ResultCode_Success
}

type LoginREQ struct {
	Account  string `protobuf:"bytes,1,opt,name=Account,proto3" json:"Account,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=Password,proto3" json:"Password,omitempty"`
}

func (m *LoginREQ) Reset()         { *m = LoginREQ{} }
func (m *LoginREQ) String() string { return proto.CompactTextString(m) }
func (*LoginREQ) ProtoMessage()    {}
func (*LoginREQ) Descriptor() ([]byte, []int) {
	return fileDescriptor_c06e4cca6c2cc899, []int{2}
}
func (m *LoginREQ) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LoginREQ) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LoginREQ.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LoginREQ) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginREQ.Merge(m, src)
}
func (m *LoginREQ) XXX_Size() int {
	return m.Size()
}
func (m *LoginREQ) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginREQ.DiscardUnknown(m)
}

var xxx_messageInfo_LoginREQ proto.InternalMessageInfo

func (m *LoginREQ) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

func (m *LoginREQ) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type LoginACK struct {
	Ret     ResultCode `protobuf:"varint,1,opt,name=Ret,proto3,enum=myproto.ResultCode" json:"Ret,omitempty"`
	Uid     uint64     `protobuf:"varint,2,opt,name=Uid,proto3" json:"Uid,omitempty"`
	HasRole bool       `protobuf:"varint,3,opt,name=HasRole,proto3" json:"HasRole,omitempty"`
}

func (m *LoginACK) Reset()         { *m = LoginACK{} }
func (m *LoginACK) String() string { return proto.CompactTextString(m) }
func (*LoginACK) ProtoMessage()    {}
func (*LoginACK) Descriptor() ([]byte, []int) {
	return fileDescriptor_c06e4cca6c2cc899, []int{3}
}
func (m *LoginACK) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LoginACK) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LoginACK.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LoginACK) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginACK.Merge(m, src)
}
func (m *LoginACK) XXX_Size() int {
	return m.Size()
}
func (m *LoginACK) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginACK.DiscardUnknown(m)
}

var xxx_messageInfo_LoginACK proto.InternalMessageInfo

func (m *LoginACK) GetRet() ResultCode {
	if m != nil {
		return m.Ret
	}
	return ResultCode_Success
}

func (m *LoginACK) GetUid() uint64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *LoginACK) GetHasRole() bool {
	if m != nil {
		return m.HasRole
	}
	return false
}

type PlayerInfo struct {
	Uid      uint64 `protobuf:"varint,1,opt,name=Uid,proto3" json:"Uid,omitempty"`
	NickName string `protobuf:"bytes,2,opt,name=NickName,proto3" json:"NickName,omitempty"`
}

func (m *PlayerInfo) Reset()         { *m = PlayerInfo{} }
func (m *PlayerInfo) String() string { return proto.CompactTextString(m) }
func (*PlayerInfo) ProtoMessage()    {}
func (*PlayerInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_c06e4cca6c2cc899, []int{4}
}
func (m *PlayerInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PlayerInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PlayerInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PlayerInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlayerInfo.Merge(m, src)
}
func (m *PlayerInfo) XXX_Size() int {
	return m.Size()
}
func (m *PlayerInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_PlayerInfo.DiscardUnknown(m)
}

var xxx_messageInfo_PlayerInfo proto.InternalMessageInfo

func (m *PlayerInfo) GetUid() uint64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *PlayerInfo) GetNickName() string {
	if m != nil {
		return m.NickName
	}
	return ""
}

type CreateRoleREQ struct {
	NickName string `protobuf:"bytes,1,opt,name=NickName,proto3" json:"NickName,omitempty"`
}

func (m *CreateRoleREQ) Reset()         { *m = CreateRoleREQ{} }
func (m *CreateRoleREQ) String() string { return proto.CompactTextString(m) }
func (*CreateRoleREQ) ProtoMessage()    {}
func (*CreateRoleREQ) Descriptor() ([]byte, []int) {
	return fileDescriptor_c06e4cca6c2cc899, []int{5}
}
func (m *CreateRoleREQ) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CreateRoleREQ) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CreateRoleREQ.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CreateRoleREQ) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRoleREQ.Merge(m, src)
}
func (m *CreateRoleREQ) XXX_Size() int {
	return m.Size()
}
func (m *CreateRoleREQ) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRoleREQ.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRoleREQ proto.InternalMessageInfo

func (m *CreateRoleREQ) GetNickName() string {
	if m != nil {
		return m.NickName
	}
	return ""
}

type CreateRoleACK struct {
	Ret ResultCode `protobuf:"varint,1,opt,name=Ret,proto3,enum=myproto.ResultCode" json:"Ret,omitempty"`
	Uid uint64     `protobuf:"varint,2,opt,name=Uid,proto3" json:"Uid,omitempty"`
}

func (m *CreateRoleACK) Reset()         { *m = CreateRoleACK{} }
func (m *CreateRoleACK) String() string { return proto.CompactTextString(m) }
func (*CreateRoleACK) ProtoMessage()    {}
func (*CreateRoleACK) Descriptor() ([]byte, []int) {
	return fileDescriptor_c06e4cca6c2cc899, []int{6}
}
func (m *CreateRoleACK) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CreateRoleACK) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CreateRoleACK.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CreateRoleACK) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRoleACK.Merge(m, src)
}
func (m *CreateRoleACK) XXX_Size() int {
	return m.Size()
}
func (m *CreateRoleACK) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRoleACK.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRoleACK proto.InternalMessageInfo

func (m *CreateRoleACK) GetRet() ResultCode {
	if m != nil {
		return m.Ret
	}
	return ResultCode_Success
}

func (m *CreateRoleACK) GetUid() uint64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func init() {
	proto.RegisterEnum("myproto.MsgId", MsgId_name, MsgId_value)
	proto.RegisterEnum("myproto.ResultCode", ResultCode_name, ResultCode_value)
	proto.RegisterType((*RegisterREQ)(nil), "myproto.RegisterREQ")
	proto.RegisterType((*RegisterACK)(nil), "myproto.RegisterACK")
	proto.RegisterType((*LoginREQ)(nil), "myproto.LoginREQ")
	proto.RegisterType((*LoginACK)(nil), "myproto.LoginACK")
	proto.RegisterType((*PlayerInfo)(nil), "myproto.PlayerInfo")
	proto.RegisterType((*CreateRoleREQ)(nil), "myproto.CreateRoleREQ")
	proto.RegisterType((*CreateRoleACK)(nil), "myproto.CreateRoleACK")
}

func init() { proto.RegisterFile("msg.proto", fileDescriptor_c06e4cca6c2cc899) }

var fileDescriptor_c06e4cca6c2cc899 = []byte{
	// 439 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x52, 0x5d, 0x8b, 0xd3, 0x40,
	0x14, 0xcd, 0x6c, 0x76, 0xdb, 0xf4, 0x76, 0x3f, 0x66, 0x67, 0x11, 0xc2, 0x3e, 0x84, 0x25, 0x22,
	0x2c, 0x2b, 0x44, 0x50, 0x9f, 0x7c, 0xb2, 0x86, 0xc2, 0x2e, 0x6b, 0xc3, 0x3a, 0xe2, 0x8b, 0x20,
	0x12, 0x93, 0x69, 0x8c, 0x4d, 0x32, 0x65, 0x26, 0x45, 0xf3, 0x2f, 0xf4, 0xd7, 0xf8, 0x17, 0x7c,
	0xec, 0xa3, 0x8f, 0xd2, 0xfe, 0x11, 0x99, 0x34, 0xd3, 0xa6, 0xf5, 0x45, 0xf4, 0x29, 0x73, 0xce,
	0x9c, 0x7b, 0xce, 0xbd, 0xb9, 0x03, 0xbd, 0x5c, 0x26, 0xde, 0x54, 0xf0, 0x92, 0x93, 0x6e, 0x5e,
	0xd5, 0x07, 0xd7, 0x87, 0x3e, 0x65, 0x49, 0x2a, 0x4b, 0x26, 0xe8, 0xf0, 0x15, 0xb1, 0xa1, 0x3b,
	0x88, 0x22, 0x3e, 0x2b, 0x4a, 0x1b, 0x5d, 0xa0, 0xcb, 0x1e, 0xd5, 0x90, 0x9c, 0x83, 0x75, 0x17,
	0x4a, 0xf9, 0x99, 0x8b, 0xd8, 0xde, 0xab, 0xaf, 0xd6, 0xd8, 0x7d, 0xba, 0x31, 0x19, 0xf8, 0xb7,
	0xe4, 0x01, 0x98, 0x94, 0xad, 0x0c, 0x8e, 0x1f, 0x9f, 0x79, 0x4d, 0x94, 0x47, 0x99, 0x9c, 0x65,
	0xa5, 0xcf, 0x63, 0x46, 0xd5, 0xbd, 0xfb, 0x1c, 0xac, 0x97, 0x3c, 0x49, 0x8b, 0x7f, 0xcf, 0x7d,
	0xd7, 0x38, 0xfc, 0x7d, 0x28, 0xc1, 0x60, 0xbe, 0x49, 0x57, 0x4e, 0xfb, 0x54, 0x1d, 0x55, 0xf4,
	0x75, 0x28, 0x29, 0xcf, 0x98, 0x6d, 0x5e, 0xa0, 0x4b, 0x8b, 0x6a, 0xe8, 0x3e, 0x03, 0xb8, 0xcb,
	0xc2, 0x8a, 0x89, 0x9b, 0x62, 0xcc, 0x75, 0x25, 0xda, 0x54, 0x9e, 0x83, 0x15, 0xa4, 0xd1, 0x24,
	0x08, 0x73, 0xa6, 0x5b, 0xd3, 0xd8, 0x7d, 0x08, 0x47, 0xbe, 0x60, 0x61, 0xc9, 0x94, 0x93, 0x9a,
	0xb0, 0x2d, 0x46, 0x3b, 0xe2, 0xeb, 0xb6, 0xf8, 0x7f, 0x86, 0xb9, 0xfa, 0x86, 0xe0, 0x60, 0x24,
	0x93, 0x9b, 0x98, 0x1c, 0x82, 0x35, 0x92, 0xc9, 0xfb, 0x80, 0x17, 0x0c, 0x1b, 0xe4, 0x0c, 0x4e,
	0x14, 0x6a, 0xad, 0x1a, 0xa3, 0x5d, 0x72, 0xe0, 0xdf, 0xe2, 0x3d, 0x82, 0xe1, 0x50, 0x91, 0x7a,
	0x33, 0xd8, 0xdc, 0x62, 0x94, 0x66, 0x9f, 0xdc, 0x83, 0x53, 0xc5, 0x6c, 0x0d, 0x88, 0x0f, 0xfe,
	0xa4, 0x95, 0xba, 0x73, 0xf5, 0x1d, 0x01, 0x6c, 0x3a, 0x27, 0x7d, 0xe8, 0xbe, 0x9e, 0x45, 0x11,
	0x93, 0x12, 0x1b, 0x04, 0xa0, 0x33, 0x92, 0xc9, 0x50, 0x08, 0x8c, 0x54, 0x4e, 0xb3, 0xf4, 0xe1,
	0x97, 0x54, 0x96, 0x98, 0xa9, 0x06, 0x1b, 0x26, 0xe0, 0x0d, 0x39, 0x26, 0xa7, 0x70, 0xa4, 0x7f,
	0xdc, 0x8a, 0x4a, 0xc8, 0x09, 0xf4, 0xf5, 0x9b, 0x50, 0x56, 0x1f, 0x6b, 0xab, 0x4c, 0xb0, 0x30,
	0xae, 0xea, 0xae, 0x71, 0xda, 0x36, 0xcf, 0xa7, 0x65, 0x85, 0x3f, 0x29, 0x9f, 0x75, 0x51, 0x4d,
	0x4d, 0xc8, 0x31, 0x80, 0x16, 0x09, 0x81, 0xb3, 0x17, 0xf7, 0x7f, 0x2c, 0x1c, 0x34, 0x5f, 0x38,
	0xe8, 0xd7, 0xc2, 0x41, 0x5f, 0x97, 0x8e, 0x31, 0x5f, 0x3a, 0xc6, 0xcf, 0xa5, 0x63, 0xbc, 0xed,
	0x79, 0x8f, 0x9a, 0xa5, 0x7c, 0xe8, 0xd4, 0x9f, 0x27, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0x28,
	0x72, 0x64, 0xa1, 0x5e, 0x03, 0x00, 0x00,
}

func (m *RegisterREQ) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RegisterREQ) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RegisterREQ) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Password) > 0 {
		i -= len(m.Password)
		copy(dAtA[i:], m.Password)
		i = encodeVarintMsg(dAtA, i, uint64(len(m.Password)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Account) > 0 {
		i -= len(m.Account)
		copy(dAtA[i:], m.Account)
		i = encodeVarintMsg(dAtA, i, uint64(len(m.Account)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *RegisterACK) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RegisterACK) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RegisterACK) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Ret != 0 {
		i = encodeVarintMsg(dAtA, i, uint64(m.Ret))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *LoginREQ) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LoginREQ) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LoginREQ) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Password) > 0 {
		i -= len(m.Password)
		copy(dAtA[i:], m.Password)
		i = encodeVarintMsg(dAtA, i, uint64(len(m.Password)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Account) > 0 {
		i -= len(m.Account)
		copy(dAtA[i:], m.Account)
		i = encodeVarintMsg(dAtA, i, uint64(len(m.Account)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *LoginACK) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LoginACK) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LoginACK) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.HasRole {
		i--
		if m.HasRole {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x18
	}
	if m.Uid != 0 {
		i = encodeVarintMsg(dAtA, i, uint64(m.Uid))
		i--
		dAtA[i] = 0x10
	}
	if m.Ret != 0 {
		i = encodeVarintMsg(dAtA, i, uint64(m.Ret))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *PlayerInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PlayerInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PlayerInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.NickName) > 0 {
		i -= len(m.NickName)
		copy(dAtA[i:], m.NickName)
		i = encodeVarintMsg(dAtA, i, uint64(len(m.NickName)))
		i--
		dAtA[i] = 0x12
	}
	if m.Uid != 0 {
		i = encodeVarintMsg(dAtA, i, uint64(m.Uid))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *CreateRoleREQ) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CreateRoleREQ) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CreateRoleREQ) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.NickName) > 0 {
		i -= len(m.NickName)
		copy(dAtA[i:], m.NickName)
		i = encodeVarintMsg(dAtA, i, uint64(len(m.NickName)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *CreateRoleACK) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CreateRoleACK) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CreateRoleACK) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Uid != 0 {
		i = encodeVarintMsg(dAtA, i, uint64(m.Uid))
		i--
		dAtA[i] = 0x10
	}
	if m.Ret != 0 {
		i = encodeVarintMsg(dAtA, i, uint64(m.Ret))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintMsg(dAtA []byte, offset int, v uint64) int {
	offset -= sovMsg(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *RegisterREQ) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Account)
	if l > 0 {
		n += 1 + l + sovMsg(uint64(l))
	}
	l = len(m.Password)
	if l > 0 {
		n += 1 + l + sovMsg(uint64(l))
	}
	return n
}

func (m *RegisterACK) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Ret != 0 {
		n += 1 + sovMsg(uint64(m.Ret))
	}
	return n
}

func (m *LoginREQ) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Account)
	if l > 0 {
		n += 1 + l + sovMsg(uint64(l))
	}
	l = len(m.Password)
	if l > 0 {
		n += 1 + l + sovMsg(uint64(l))
	}
	return n
}

func (m *LoginACK) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Ret != 0 {
		n += 1 + sovMsg(uint64(m.Ret))
	}
	if m.Uid != 0 {
		n += 1 + sovMsg(uint64(m.Uid))
	}
	if m.HasRole {
		n += 2
	}
	return n
}

func (m *PlayerInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Uid != 0 {
		n += 1 + sovMsg(uint64(m.Uid))
	}
	l = len(m.NickName)
	if l > 0 {
		n += 1 + l + sovMsg(uint64(l))
	}
	return n
}

func (m *CreateRoleREQ) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.NickName)
	if l > 0 {
		n += 1 + l + sovMsg(uint64(l))
	}
	return n
}

func (m *CreateRoleACK) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Ret != 0 {
		n += 1 + sovMsg(uint64(m.Ret))
	}
	if m.Uid != 0 {
		n += 1 + sovMsg(uint64(m.Uid))
	}
	return n
}

func sovMsg(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMsg(x uint64) (n int) {
	return sovMsg(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *RegisterREQ) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsg
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: RegisterREQ: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RegisterREQ: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Account", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMsg
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Account = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Password", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMsg
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Password = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMsg(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMsg
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *RegisterACK) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsg
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: RegisterACK: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RegisterACK: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ret", wireType)
			}
			m.Ret = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Ret |= ResultCode(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipMsg(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMsg
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *LoginREQ) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsg
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: LoginREQ: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LoginREQ: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Account", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMsg
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Account = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Password", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMsg
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Password = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMsg(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMsg
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *LoginACK) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsg
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: LoginACK: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LoginACK: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ret", wireType)
			}
			m.Ret = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Ret |= ResultCode(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Uid", wireType)
			}
			m.Uid = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Uid |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field HasRole", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.HasRole = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipMsg(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMsg
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *PlayerInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsg
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: PlayerInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PlayerInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Uid", wireType)
			}
			m.Uid = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Uid |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NickName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMsg
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NickName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMsg(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMsg
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *CreateRoleREQ) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsg
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CreateRoleREQ: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CreateRoleREQ: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NickName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMsg
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NickName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMsg(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMsg
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *CreateRoleACK) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsg
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CreateRoleACK: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CreateRoleACK: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ret", wireType)
			}
			m.Ret = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Ret |= ResultCode(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Uid", wireType)
			}
			m.Uid = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Uid |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipMsg(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMsg
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipMsg(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMsg
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowMsg
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowMsg
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthMsg
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMsg
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMsg
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMsg        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMsg          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMsg = fmt.Errorf("proto: unexpected end of group")
)
