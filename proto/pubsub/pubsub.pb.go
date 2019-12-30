// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pubsub/pubsub.proto

package pubsub

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

/////////////////////////////////////////////////
// pub/sub
/////////////////////////////////////////////////
type PubStartBattle struct {
	Info                 *account.AccountInfo `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *PubStartBattle) Reset()         { *m = PubStartBattle{} }
func (m *PubStartBattle) String() string { return proto.CompactTextString(m) }
func (*PubStartBattle) ProtoMessage()    {}
func (*PubStartBattle) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce310d0bb9f289ed, []int{0}
}

func (m *PubStartBattle) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PubStartBattle.Unmarshal(m, b)
}
func (m *PubStartBattle) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PubStartBattle.Marshal(b, m, deterministic)
}
func (m *PubStartBattle) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PubStartBattle.Merge(m, src)
}
func (m *PubStartBattle) XXX_Size() int {
	return xxx_messageInfo_PubStartBattle.Size(m)
}
func (m *PubStartBattle) XXX_DiscardUnknown() {
	xxx_messageInfo_PubStartBattle.DiscardUnknown(m)
}

var xxx_messageInfo_PubStartBattle proto.InternalMessageInfo

func (m *PubStartBattle) GetInfo() *account.AccountInfo {
	if m != nil {
		return m.Info
	}
	return nil
}

type PubBattleResult struct {
	Info                 *account.AccountInfo `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
	Win                  bool                 `protobuf:"varint,2,opt,name=win,proto3" json:"win,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *PubBattleResult) Reset()         { *m = PubBattleResult{} }
func (m *PubBattleResult) String() string { return proto.CompactTextString(m) }
func (*PubBattleResult) ProtoMessage()    {}
func (*PubBattleResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce310d0bb9f289ed, []int{1}
}

func (m *PubBattleResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PubBattleResult.Unmarshal(m, b)
}
func (m *PubBattleResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PubBattleResult.Marshal(b, m, deterministic)
}
func (m *PubBattleResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PubBattleResult.Merge(m, src)
}
func (m *PubBattleResult) XXX_Size() int {
	return xxx_messageInfo_PubBattleResult.Size(m)
}
func (m *PubBattleResult) XXX_DiscardUnknown() {
	xxx_messageInfo_PubBattleResult.DiscardUnknown(m)
}

var xxx_messageInfo_PubBattleResult proto.InternalMessageInfo

func (m *PubBattleResult) GetInfo() *account.AccountInfo {
	if m != nil {
		return m.Info
	}
	return nil
}

func (m *PubBattleResult) GetWin() bool {
	if m != nil {
		return m.Win
	}
	return false
}

type PubExpirePlayer struct {
	PlayerId             int64    `protobuf:"varint,1,opt,name=player_id,json=playerId,proto3" json:"player_id,omitempty"`
	GameId               int32    `protobuf:"varint,2,opt,name=game_id,json=gameId,proto3" json:"game_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PubExpirePlayer) Reset()         { *m = PubExpirePlayer{} }
func (m *PubExpirePlayer) String() string { return proto.CompactTextString(m) }
func (*PubExpirePlayer) ProtoMessage()    {}
func (*PubExpirePlayer) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce310d0bb9f289ed, []int{2}
}

func (m *PubExpirePlayer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PubExpirePlayer.Unmarshal(m, b)
}
func (m *PubExpirePlayer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PubExpirePlayer.Marshal(b, m, deterministic)
}
func (m *PubExpirePlayer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PubExpirePlayer.Merge(m, src)
}
func (m *PubExpirePlayer) XXX_Size() int {
	return xxx_messageInfo_PubExpirePlayer.Size(m)
}
func (m *PubExpirePlayer) XXX_DiscardUnknown() {
	xxx_messageInfo_PubExpirePlayer.DiscardUnknown(m)
}

var xxx_messageInfo_PubExpirePlayer proto.InternalMessageInfo

func (m *PubExpirePlayer) GetPlayerId() int64 {
	if m != nil {
		return m.PlayerId
	}
	return 0
}

func (m *PubExpirePlayer) GetGameId() int32 {
	if m != nil {
		return m.GameId
	}
	return 0
}

func init() {
	proto.RegisterType((*PubStartBattle)(nil), "yokai_pubsub.PubStartBattle")
	proto.RegisterType((*PubBattleResult)(nil), "yokai_pubsub.PubBattleResult")
	proto.RegisterType((*PubExpirePlayer)(nil), "yokai_pubsub.PubExpirePlayer")
}

func init() { proto.RegisterFile("pubsub/pubsub.proto", fileDescriptor_ce310d0bb9f289ed) }

var fileDescriptor_ce310d0bb9f289ed = []byte{
	// 237 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x90, 0xc1, 0x4a, 0x03, 0x31,
	0x10, 0x86, 0x59, 0xab, 0xb5, 0x46, 0x51, 0x89, 0x88, 0xa5, 0x5e, 0xca, 0x9e, 0x7a, 0x90, 0x2c,
	0xe8, 0x0b, 0x68, 0x41, 0x64, 0x6f, 0x4b, 0x7a, 0xf3, 0xb2, 0x24, 0xbb, 0x69, 0x0d, 0x6e, 0x77,
	0x42, 0x76, 0xa2, 0xf6, 0xed, 0x25, 0x99, 0xf5, 0x01, 0x3c, 0x7d, 0x33, 0x93, 0x99, 0x8f, 0xf0,
	0xb3, 0x1b, 0x17, 0xf4, 0x10, 0x74, 0x41, 0x10, 0xce, 0x03, 0x02, 0xbf, 0x38, 0xc0, 0xa7, 0xb2,
	0x35, 0xcd, 0x16, 0xb7, 0xaa, 0x69, 0x20, 0xf4, 0x58, 0x8c, 0xa4, 0xa5, 0xfc, 0x99, 0x5d, 0x56,
	0x41, 0x6f, 0x50, 0x79, 0x5c, 0x2b, 0xc4, 0xce, 0x70, 0xc1, 0x8e, 0x6d, 0xbf, 0x85, 0x79, 0xb6,
	0xcc, 0x56, 0xe7, 0x8f, 0x0b, 0x41, 0x96, 0xbf, 0xab, 0x17, 0x62, 0xd9, 0x6f, 0x41, 0xa6, 0xbd,
	0x7c, 0xc3, 0xae, 0xaa, 0xa0, 0xe9, 0x58, 0x9a, 0x21, 0x74, 0xf8, 0x5f, 0x05, 0xbf, 0x66, 0x93,
	0x6f, 0xdb, 0xcf, 0x8f, 0x96, 0xd9, 0x6a, 0x26, 0x63, 0x99, 0xbf, 0x25, 0xe9, 0xeb, 0x8f, 0xb3,
	0xde, 0x54, 0x9d, 0x3a, 0x18, 0xcf, 0xef, 0xd9, 0x99, 0x4b, 0x55, 0x6d, 0xdb, 0x64, 0x9e, 0xc8,
	0x19, 0x0d, 0xca, 0x96, 0xdf, 0xb1, 0xd3, 0x9d, 0xda, 0x9b, 0xf8, 0x14, 0x2d, 0x27, 0x72, 0x1a,
	0xdb, 0xb2, 0x5d, 0x8b, 0xf7, 0x87, 0x9d, 0xc5, 0x8f, 0xa0, 0x45, 0x03, 0xfb, 0x22, 0x7d, 0xc4,
	0x02, 0xb1, 0x1e, 0x8c, 0xff, 0x32, 0xbe, 0x48, 0x41, 0x8c, 0xd1, 0xe9, 0x69, 0xea, 0x9e, 0x7e,
	0x03, 0x00, 0x00, 0xff, 0xff, 0xe5, 0xb3, 0x56, 0x5b, 0x52, 0x01, 0x00, 0x00,
}
