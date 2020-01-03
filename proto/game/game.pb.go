// Code generated by protoc-gen-go. DO NOT EDIT.
// source: game/game.proto

package game

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	account "github.com/yokaiio/yokai_server/proto/account"
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

type GetAccountByIDRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAccountByIDRequest) Reset()         { *m = GetAccountByIDRequest{} }
func (m *GetAccountByIDRequest) String() string { return proto.CompactTextString(m) }
func (*GetAccountByIDRequest) ProtoMessage()    {}
func (*GetAccountByIDRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a9278d664c0c01e, []int{0}
}

func (m *GetAccountByIDRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAccountByIDRequest.Unmarshal(m, b)
}
func (m *GetAccountByIDRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAccountByIDRequest.Marshal(b, m, deterministic)
}
func (m *GetAccountByIDRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAccountByIDRequest.Merge(m, src)
}
func (m *GetAccountByIDRequest) XXX_Size() int {
	return xxx_messageInfo_GetAccountByIDRequest.Size(m)
}
func (m *GetAccountByIDRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAccountByIDRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetAccountByIDRequest proto.InternalMessageInfo

func (m *GetAccountByIDRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type GetAccountByIDReply struct {
	Info                 *account.AccountInfo `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *GetAccountByIDReply) Reset()         { *m = GetAccountByIDReply{} }
func (m *GetAccountByIDReply) String() string { return proto.CompactTextString(m) }
func (*GetAccountByIDReply) ProtoMessage()    {}
func (*GetAccountByIDReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a9278d664c0c01e, []int{1}
}

func (m *GetAccountByIDReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAccountByIDReply.Unmarshal(m, b)
}
func (m *GetAccountByIDReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAccountByIDReply.Marshal(b, m, deterministic)
}
func (m *GetAccountByIDReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAccountByIDReply.Merge(m, src)
}
func (m *GetAccountByIDReply) XXX_Size() int {
	return xxx_messageInfo_GetAccountByIDReply.Size(m)
}
func (m *GetAccountByIDReply) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAccountByIDReply.DiscardUnknown(m)
}

var xxx_messageInfo_GetAccountByIDReply proto.InternalMessageInfo

func (m *GetAccountByIDReply) GetInfo() *account.AccountInfo {
	if m != nil {
		return m.Info
	}
	return nil
}

type GetRemoteLitePlayerRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRemoteLitePlayerRequest) Reset()         { *m = GetRemoteLitePlayerRequest{} }
func (m *GetRemoteLitePlayerRequest) String() string { return proto.CompactTextString(m) }
func (*GetRemoteLitePlayerRequest) ProtoMessage()    {}
func (*GetRemoteLitePlayerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a9278d664c0c01e, []int{2}
}

func (m *GetRemoteLitePlayerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRemoteLitePlayerRequest.Unmarshal(m, b)
}
func (m *GetRemoteLitePlayerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRemoteLitePlayerRequest.Marshal(b, m, deterministic)
}
func (m *GetRemoteLitePlayerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRemoteLitePlayerRequest.Merge(m, src)
}
func (m *GetRemoteLitePlayerRequest) XXX_Size() int {
	return xxx_messageInfo_GetRemoteLitePlayerRequest.Size(m)
}
func (m *GetRemoteLitePlayerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRemoteLitePlayerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRemoteLitePlayerRequest proto.InternalMessageInfo

func (m *GetRemoteLitePlayerRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type GetRemoteLitePlayerReply struct {
	Info                 *LitePlayer `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *GetRemoteLitePlayerReply) Reset()         { *m = GetRemoteLitePlayerReply{} }
func (m *GetRemoteLitePlayerReply) String() string { return proto.CompactTextString(m) }
func (*GetRemoteLitePlayerReply) ProtoMessage()    {}
func (*GetRemoteLitePlayerReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a9278d664c0c01e, []int{3}
}

func (m *GetRemoteLitePlayerReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRemoteLitePlayerReply.Unmarshal(m, b)
}
func (m *GetRemoteLitePlayerReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRemoteLitePlayerReply.Marshal(b, m, deterministic)
}
func (m *GetRemoteLitePlayerReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRemoteLitePlayerReply.Merge(m, src)
}
func (m *GetRemoteLitePlayerReply) XXX_Size() int {
	return xxx_messageInfo_GetRemoteLitePlayerReply.Size(m)
}
func (m *GetRemoteLitePlayerReply) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRemoteLitePlayerReply.DiscardUnknown(m)
}

var xxx_messageInfo_GetRemoteLitePlayerReply proto.InternalMessageInfo

func (m *GetRemoteLitePlayerReply) GetInfo() *LitePlayer {
	if m != nil {
		return m.Info
	}
	return nil
}

func init() {
	proto.RegisterType((*GetAccountByIDRequest)(nil), "yokai_game.GetAccountByIDRequest")
	proto.RegisterType((*GetAccountByIDReply)(nil), "yokai_game.GetAccountByIDReply")
	proto.RegisterType((*GetRemoteLitePlayerRequest)(nil), "yokai_game.GetRemoteLitePlayerRequest")
	proto.RegisterType((*GetRemoteLitePlayerReply)(nil), "yokai_game.GetRemoteLitePlayerReply")
}

func init() { proto.RegisterFile("game/game.proto", fileDescriptor_2a9278d664c0c01e) }

var fileDescriptor_2a9278d664c0c01e = []byte{
	// 285 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xc1, 0x4a, 0x03, 0x31,
	0x10, 0x86, 0x6d, 0x15, 0x0f, 0x53, 0xa8, 0x18, 0xa9, 0x94, 0x5c, 0xd4, 0x45, 0x54, 0x4a, 0xc9,
	0x42, 0x7d, 0x02, 0x8b, 0xba, 0x14, 0x3c, 0xc8, 0x0a, 0x1e, 0xbc, 0x94, 0x74, 0x3b, 0xad, 0xc1,
	0xdd, 0xcd, 0x9a, 0x66, 0x0b, 0x79, 0x47, 0x1f, 0x4a, 0x76, 0x52, 0xb1, 0x5d, 0x5a, 0xbd, 0x64,
	0xc2, 0xf0, 0xfd, 0xfc, 0xf3, 0x4f, 0x02, 0x47, 0x73, 0x99, 0x61, 0x58, 0x1d, 0xa2, 0x30, 0xda,
	0x6a, 0x06, 0x4e, 0x7f, 0x48, 0x35, 0xae, 0x3a, 0xbc, 0x23, 0x93, 0x44, 0x97, 0xb9, 0x0d, 0x57,
	0xd5, 0x23, 0xfc, 0x98, 0x34, 0x45, 0x2a, 0x1d, 0x1a, 0xdf, 0x0a, 0xae, 0xa1, 0x13, 0xa1, 0xbd,
	0xf3, 0xd8, 0xd0, 0x8d, 0xee, 0x63, 0xfc, 0x2c, 0x71, 0x61, 0x59, 0x1b, 0x9a, 0x6a, 0xda, 0x6d,
	0x9c, 0x37, 0x6e, 0xf6, 0xe3, 0xa6, 0x9a, 0x06, 0x0f, 0x70, 0x52, 0x07, 0x8b, 0xd4, 0x31, 0x01,
	0x07, 0x2a, 0x9f, 0x69, 0x02, 0x5b, 0x03, 0x2e, 0xfc, 0x10, 0x3f, 0xb6, 0x2b, 0x7c, 0x94, 0xcf,
	0x74, 0x4c, 0x5c, 0xd0, 0x07, 0x1e, 0xa1, 0x8d, 0x31, 0xd3, 0x16, 0x9f, 0x94, 0xc5, 0x67, 0x1a,
	0x66, 0x97, 0xe9, 0x23, 0x74, 0xb7, 0xd2, 0x95, 0x73, 0x6f, 0xc3, 0xf9, 0x54, 0xfc, 0xc6, 0x17,
	0x6b, 0x28, 0x31, 0x83, 0xaf, 0x06, 0xb4, 0x22, 0x99, 0xe1, 0x0b, 0x9a, 0xa5, 0x4a, 0x90, 0xbd,
	0x42, 0x7b, 0x33, 0x0c, 0xbb, 0x58, 0xd7, 0x6f, 0xdd, 0x08, 0x3f, 0xfb, 0x0b, 0x29, 0x52, 0x17,
	0xec, 0x31, 0xa4, 0x25, 0xd5, 0xe7, 0x65, 0x57, 0x35, 0xe5, 0x8e, 0xf8, 0xfc, 0xf2, 0x5f, 0x8e,
	0x6c, 0x86, 0xfd, 0xb7, 0xde, 0x5c, 0xd9, 0xf7, 0x72, 0x22, 0x12, 0x9d, 0x85, 0xa4, 0x51, 0xda,
	0xd7, 0xf1, 0x02, 0xcd, 0x12, 0x4d, 0x48, 0xaf, 0x4b, 0xdf, 0x63, 0x72, 0x48, 0xf7, 0xdb, 0xef,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xa4, 0x42, 0xbc, 0x71, 0x32, 0x02, 0x00, 0x00,
}
