// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: server/gate/gate.proto

package gate

import (
	global "bitbucket.org/funplus/server/proto/global"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GateStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GateId int32 `protobuf:"varint,1,opt,name=gate_id,json=gateId,proto3" json:"gate_id,omitempty"`
	Health int32 `protobuf:"varint,2,opt,name=health,proto3" json:"health,omitempty"`
}

func (x *GateStatus) Reset() {
	*x = GateStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_gate_gate_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GateStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GateStatus) ProtoMessage() {}

func (x *GateStatus) ProtoReflect() protoreflect.Message {
	mi := &file_server_gate_gate_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GateStatus.ProtoReflect.Descriptor instead.
func (*GateStatus) Descriptor() ([]byte, []int) {
	return file_server_gate_gate_proto_rawDescGZIP(), []int{0}
}

func (x *GateStatus) GetGateId() int32 {
	if x != nil {
		return x.GateId
	}
	return 0
}

func (x *GateStatus) GetHealth() int32 {
	if x != nil {
		return x.Health
	}
	return 0
}

type UserInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId      int64  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	AccountId   int64  `protobuf:"varint,2,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	GameId      int32  `protobuf:"varint,3,opt,name=game_id,json=gameId,proto3" json:"game_id,omitempty"`
	PlayerId    int64  `protobuf:"varint,4,opt,name=player_id,json=playerId,proto3" json:"player_id,omitempty"`
	PlayerName  string `protobuf:"bytes,5,opt,name=player_name,json=playerName,proto3" json:"player_name,omitempty"`
	PlayerLevel int32  `protobuf:"varint,6,opt,name=player_level,json=playerLevel,proto3" json:"player_level,omitempty"`
}

func (x *UserInfo) Reset() {
	*x = UserInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_gate_gate_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInfo) ProtoMessage() {}

func (x *UserInfo) ProtoReflect() protoreflect.Message {
	mi := &file_server_gate_gate_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInfo.ProtoReflect.Descriptor instead.
func (*UserInfo) Descriptor() ([]byte, []int) {
	return file_server_gate_gate_proto_rawDescGZIP(), []int{1}
}

func (x *UserInfo) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *UserInfo) GetAccountId() int64 {
	if x != nil {
		return x.AccountId
	}
	return 0
}

func (x *UserInfo) GetGameId() int32 {
	if x != nil {
		return x.GameId
	}
	return 0
}

func (x *UserInfo) GetPlayerId() int64 {
	if x != nil {
		return x.PlayerId
	}
	return 0
}

func (x *UserInfo) GetPlayerName() string {
	if x != nil {
		return x.PlayerName
	}
	return ""
}

func (x *UserInfo) GetPlayerLevel() int32 {
	if x != nil {
		return x.PlayerLevel
	}
	return 0
}

type GateEmptyMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GateEmptyMessage) Reset() {
	*x = GateEmptyMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_gate_gate_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GateEmptyMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GateEmptyMessage) ProtoMessage() {}

func (x *GateEmptyMessage) ProtoReflect() protoreflect.Message {
	mi := &file_server_gate_gate_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GateEmptyMessage.ProtoReflect.Descriptor instead.
func (*GateEmptyMessage) Descriptor() ([]byte, []int) {
	return file_server_gate_gate_proto_rawDescGZIP(), []int{2}
}

type GetGateStatusReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status *GateStatus `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *GetGateStatusReply) Reset() {
	*x = GetGateStatusReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_gate_gate_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGateStatusReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGateStatusReply) ProtoMessage() {}

func (x *GetGateStatusReply) ProtoReflect() protoreflect.Message {
	mi := &file_server_gate_gate_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGateStatusReply.ProtoReflect.Descriptor instead.
func (*GetGateStatusReply) Descriptor() ([]byte, []int) {
	return file_server_gate_gate_proto_rawDescGZIP(), []int{3}
}

func (x *GetGateStatusReply) GetStatus() *GateStatus {
	if x != nil {
		return x.Status
	}
	return nil
}

type SyncPlayerInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64              `protobuf:"varint,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
	Info   *global.PlayerInfo `protobuf:"bytes,2,opt,name=info,proto3" json:"info,omitempty"`
}

func (x *SyncPlayerInfoRequest) Reset() {
	*x = SyncPlayerInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_gate_gate_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SyncPlayerInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyncPlayerInfoRequest) ProtoMessage() {}

func (x *SyncPlayerInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_server_gate_gate_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SyncPlayerInfoRequest.ProtoReflect.Descriptor instead.
func (*SyncPlayerInfoRequest) Descriptor() ([]byte, []int) {
	return file_server_gate_gate_proto_rawDescGZIP(), []int{4}
}

func (x *SyncPlayerInfoRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *SyncPlayerInfoRequest) GetInfo() *global.PlayerInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

type SyncPlayerInfoReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *UserInfo `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
}

func (x *SyncPlayerInfoReply) Reset() {
	*x = SyncPlayerInfoReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_gate_gate_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SyncPlayerInfoReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyncPlayerInfoReply) ProtoMessage() {}

func (x *SyncPlayerInfoReply) ProtoReflect() protoreflect.Message {
	mi := &file_server_gate_gate_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SyncPlayerInfoReply.ProtoReflect.Descriptor instead.
func (*SyncPlayerInfoReply) Descriptor() ([]byte, []int) {
	return file_server_gate_gate_proto_rawDescGZIP(), []int{5}
}

func (x *SyncPlayerInfoReply) GetInfo() *UserInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

var File_server_gate_gate_proto protoreflect.FileDescriptor

var file_server_gate_gate_proto_rawDesc = []byte{
	0x0a, 0x16, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x2f, 0x67, 0x61,
	0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x67, 0x61, 0x74, 0x65, 0x1a, 0x13,
	0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2f, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x3d, 0x0a, 0x0a, 0x47, 0x61, 0x74, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x17, 0x0a, 0x07, 0x67, 0x61, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x67, 0x61, 0x74, 0x65, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65,
	0x61, 0x6c, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x68, 0x65, 0x61, 0x6c,
	0x74, 0x68, 0x22, 0xbc, 0x01, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x67, 0x61, 0x6d, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x67, 0x61, 0x6d, 0x65, 0x49, 0x64,
	0x12, 0x1b, 0x0a, 0x09, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a,
	0x0b, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x21,
	0x0a, 0x0c, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4c, 0x65, 0x76, 0x65,
	0x6c, 0x22, 0x12, 0x0a, 0x10, 0x47, 0x61, 0x74, 0x65, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x3e, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x47, 0x61, 0x74, 0x65,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x28, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x67, 0x61,
	0x74, 0x65, 0x2e, 0x47, 0x61, 0x74, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x57, 0x0a, 0x15, 0x53, 0x79, 0x6e, 0x63, 0x50, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2e, 0x50, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x22, 0x39,
	0x0a, 0x13, 0x53, 0x79, 0x6e, 0x63, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x22, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x32, 0x9e, 0x01, 0x0a, 0x0b, 0x67, 0x61,
	0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x43, 0x0a, 0x0d, 0x47, 0x65, 0x74,
	0x47, 0x61, 0x74, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x61, 0x74,
	0x65, 0x2e, 0x47, 0x61, 0x74, 0x65, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x1a, 0x18, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x47, 0x61, 0x74,
	0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x4a,
	0x0a, 0x0e, 0x53, 0x79, 0x6e, 0x63, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x1b, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x2e, 0x53, 0x79, 0x6e, 0x63, 0x50, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e,
	0x67, 0x61, 0x74, 0x65, 0x2e, 0x53, 0x79, 0x6e, 0x63, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x30, 0x5a, 0x2e, 0x62, 0x69,
	0x74, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x66, 0x75, 0x6e, 0x70,
	0x6c, 0x75, 0x73, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_server_gate_gate_proto_rawDescOnce sync.Once
	file_server_gate_gate_proto_rawDescData = file_server_gate_gate_proto_rawDesc
)

func file_server_gate_gate_proto_rawDescGZIP() []byte {
	file_server_gate_gate_proto_rawDescOnce.Do(func() {
		file_server_gate_gate_proto_rawDescData = protoimpl.X.CompressGZIP(file_server_gate_gate_proto_rawDescData)
	})
	return file_server_gate_gate_proto_rawDescData
}

var file_server_gate_gate_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_server_gate_gate_proto_goTypes = []interface{}{
	(*GateStatus)(nil),            // 0: gate.GateStatus
	(*UserInfo)(nil),              // 1: gate.UserInfo
	(*GateEmptyMessage)(nil),      // 2: gate.GateEmptyMessage
	(*GetGateStatusReply)(nil),    // 3: gate.GetGateStatusReply
	(*SyncPlayerInfoRequest)(nil), // 4: gate.SyncPlayerInfoRequest
	(*SyncPlayerInfoReply)(nil),   // 5: gate.SyncPlayerInfoReply
	(*global.PlayerInfo)(nil),     // 6: global.PlayerInfo
}
var file_server_gate_gate_proto_depIdxs = []int32{
	0, // 0: gate.GetGateStatusReply.status:type_name -> gate.GateStatus
	6, // 1: gate.SyncPlayerInfoRequest.info:type_name -> global.PlayerInfo
	1, // 2: gate.SyncPlayerInfoReply.info:type_name -> gate.UserInfo
	2, // 3: gate.gateService.GetGateStatus:input_type -> gate.GateEmptyMessage
	4, // 4: gate.gateService.SyncPlayerInfo:input_type -> gate.SyncPlayerInfoRequest
	3, // 5: gate.gateService.GetGateStatus:output_type -> gate.GetGateStatusReply
	5, // 6: gate.gateService.SyncPlayerInfo:output_type -> gate.SyncPlayerInfoReply
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_server_gate_gate_proto_init() }
func file_server_gate_gate_proto_init() {
	if File_server_gate_gate_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_server_gate_gate_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GateStatus); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_server_gate_gate_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_server_gate_gate_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GateEmptyMessage); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_server_gate_gate_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGateStatusReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_server_gate_gate_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SyncPlayerInfoRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_server_gate_gate_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SyncPlayerInfoReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_server_gate_gate_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_server_gate_gate_proto_goTypes,
		DependencyIndexes: file_server_gate_gate_proto_depIdxs,
		MessageInfos:      file_server_gate_gate_proto_msgTypes,
	}.Build()
	File_server_gate_gate_proto = out.File
	file_server_gate_gate_proto_rawDesc = nil
	file_server_gate_gate_proto_goTypes = nil
	file_server_gate_gate_proto_depIdxs = nil
}
