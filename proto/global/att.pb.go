// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: att.proto

package global

import (
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

type AttType int32

const (
	AttType_Att_Begin          AttType = 0  // 32位属性开始
	AttType_Att_SecondBegin    AttType = 0  // 二级属性开始
	AttType_Att_Atk            AttType = 0  // 0 攻击力最终值
	AttType_Att_Armor          AttType = 1  // 1 护甲最终值
	AttType_Att_SelfDmgInc     AttType = 2  // 2 自身伤害加成
	AttType_Att_EnemyWoundInc  AttType = 3  // 3 敌方受伤加成
	AttType_Att_SelfDmgDec     AttType = 4  // 4 自身伤害消减
	AttType_Att_EnemyWoundDec  AttType = 5  // 5 敌方受伤消减
	AttType_Att_Crit           AttType = 6  // 6 暴击值
	AttType_Att_CritInc        AttType = 7  // 7 暴击倍数加层
	AttType_Att_Tenacity       AttType = 8  // 8 韧性
	AttType_Att_Heal           AttType = 9  // 9 治疗强度最终值
	AttType_Att_RealDmg        AttType = 10 // 10 真实伤害
	AttType_Att_MoveSpeed      AttType = 11 // 11 移动速度最终值
	AttType_Att_AtbSpeed       AttType = 12 // 12 时间槽速度最终值
	AttType_Att_EffectHit      AttType = 13 // 13 技能效果命中
	AttType_Att_EffectResist   AttType = 14 // 14 技能效果抵抗
	AttType_Att_MaxHP          AttType = 15 // 15 生命值上限最终值
	AttType_Att_CurHP          AttType = 16 // 16 当前生命值
	AttType_Att_MaxMP          AttType = 17 // 17 蓝量上限
	AttType_Att_CurMP          AttType = 18 // 18 当前蓝量
	AttType_Att_GenMP          AttType = 19 // 19 mp恢复值
	AttType_Att_MaxRage        AttType = 20 // 20 怒气值上限
	AttType_Att_GenRagePercent AttType = 21 // 21 回怒百分比
	AttType_Att_InitRage       AttType = 22 // 22 初始怒气
	AttType_Att_Hit            AttType = 23 // 23 命中
	AttType_Att_Dodge          AttType = 24 // 24 闪避
	AttType_Att_MoveScope      AttType = 25 // 25 移动范围
	AttType_Att_DmgTypeBegin   AttType = 26 // 26 各系伤害加成begin
	AttType_Att_DmgPhysics     AttType = 26 // 26 物理系伤害加成
	AttType_Att_DmgEarth       AttType = 27 // 27 地系伤害加成
	AttType_Att_DmgWater       AttType = 28 // 28 水系伤害加成
	AttType_Att_DmgFire        AttType = 29 // 29 火系伤害加成
	AttType_Att_DmgWind        AttType = 30 // 30 风系伤害加成
	AttType_Att_DmgThunder     AttType = 31 // 31 雷系伤害加成
	AttType_Att_DmgTime        AttType = 32 // 32 时系伤害加成
	AttType_Att_DmgSpace       AttType = 33 // 33 空系伤害加成
	AttType_Att_DmgSteel       AttType = 34 // 34 钢系伤害加成
	AttType_Att_DmgDeath       AttType = 35 // 35 灭系伤害加成
	AttType_Att_DmgTypeEnd     AttType = 36 // 36 各系伤害加成end
	AttType_Att_ResTypeBegin   AttType = 36 // 36 各系伤害抗性
	AttType_Att_ResPhysics     AttType = 36 // 36 物理系伤害抗性
	AttType_Att_ResEarth       AttType = 37 // 37 地系伤害抗性
	AttType_Att_ResWater       AttType = 38 // 38 水系伤害抗性
	AttType_Att_ResFire        AttType = 39 // 39 火系伤害抗性
	AttType_Att_ResWind        AttType = 40 // 40 风系伤害抗性
	AttType_Att_ResThunder     AttType = 41 // 41 雷系伤害抗性
	AttType_Att_ResTime        AttType = 42 // 42 时系伤害抗性
	AttType_Att_ResSpace       AttType = 43 // 43 空系伤害抗性
	AttType_Att_ResSteel       AttType = 44 // 44 钢系伤害抗性
	AttType_Att_ResDeath       AttType = 45 // 45 灭系伤害抗性
	AttType_Att_ResTypeEnd     AttType = 46 // 46 各系伤害抗性
	AttType_Att_SecondEnd      AttType = 46 // 二级属性结束
	AttType_Att_End            AttType = 46 // 32位属性结束
)

// Enum value maps for AttType.
var (
	AttType_name = map[int32]string{
		0: "Att_Begin",
		// Duplicate value: 0: "Att_SecondBegin",
		// Duplicate value: 0: "Att_Atk",
		1:  "Att_Armor",
		2:  "Att_SelfDmgInc",
		3:  "Att_EnemyWoundInc",
		4:  "Att_SelfDmgDec",
		5:  "Att_EnemyWoundDec",
		6:  "Att_Crit",
		7:  "Att_CritInc",
		8:  "Att_Tenacity",
		9:  "Att_Heal",
		10: "Att_RealDmg",
		11: "Att_MoveSpeed",
		12: "Att_AtbSpeed",
		13: "Att_EffectHit",
		14: "Att_EffectResist",
		15: "Att_MaxHP",
		16: "Att_CurHP",
		17: "Att_MaxMP",
		18: "Att_CurMP",
		19: "Att_GenMP",
		20: "Att_MaxRage",
		21: "Att_GenRagePercent",
		22: "Att_InitRage",
		23: "Att_Hit",
		24: "Att_Dodge",
		25: "Att_MoveScope",
		26: "Att_DmgTypeBegin",
		// Duplicate value: 26: "Att_DmgPhysics",
		27: "Att_DmgEarth",
		28: "Att_DmgWater",
		29: "Att_DmgFire",
		30: "Att_DmgWind",
		31: "Att_DmgThunder",
		32: "Att_DmgTime",
		33: "Att_DmgSpace",
		34: "Att_DmgSteel",
		35: "Att_DmgDeath",
		36: "Att_DmgTypeEnd",
		// Duplicate value: 36: "Att_ResTypeBegin",
		// Duplicate value: 36: "Att_ResPhysics",
		37: "Att_ResEarth",
		38: "Att_ResWater",
		39: "Att_ResFire",
		40: "Att_ResWind",
		41: "Att_ResThunder",
		42: "Att_ResTime",
		43: "Att_ResSpace",
		44: "Att_ResSteel",
		45: "Att_ResDeath",
		46: "Att_ResTypeEnd",
		// Duplicate value: 46: "Att_SecondEnd",
		// Duplicate value: 46: "Att_End",
	}
	AttType_value = map[string]int32{
		"Att_Begin":          0,
		"Att_SecondBegin":    0,
		"Att_Atk":            0,
		"Att_Armor":          1,
		"Att_SelfDmgInc":     2,
		"Att_EnemyWoundInc":  3,
		"Att_SelfDmgDec":     4,
		"Att_EnemyWoundDec":  5,
		"Att_Crit":           6,
		"Att_CritInc":        7,
		"Att_Tenacity":       8,
		"Att_Heal":           9,
		"Att_RealDmg":        10,
		"Att_MoveSpeed":      11,
		"Att_AtbSpeed":       12,
		"Att_EffectHit":      13,
		"Att_EffectResist":   14,
		"Att_MaxHP":          15,
		"Att_CurHP":          16,
		"Att_MaxMP":          17,
		"Att_CurMP":          18,
		"Att_GenMP":          19,
		"Att_MaxRage":        20,
		"Att_GenRagePercent": 21,
		"Att_InitRage":       22,
		"Att_Hit":            23,
		"Att_Dodge":          24,
		"Att_MoveScope":      25,
		"Att_DmgTypeBegin":   26,
		"Att_DmgPhysics":     26,
		"Att_DmgEarth":       27,
		"Att_DmgWater":       28,
		"Att_DmgFire":        29,
		"Att_DmgWind":        30,
		"Att_DmgThunder":     31,
		"Att_DmgTime":        32,
		"Att_DmgSpace":       33,
		"Att_DmgSteel":       34,
		"Att_DmgDeath":       35,
		"Att_DmgTypeEnd":     36,
		"Att_ResTypeBegin":   36,
		"Att_ResPhysics":     36,
		"Att_ResEarth":       37,
		"Att_ResWater":       38,
		"Att_ResFire":        39,
		"Att_ResWind":        40,
		"Att_ResThunder":     41,
		"Att_ResTime":        42,
		"Att_ResSpace":       43,
		"Att_ResSteel":       44,
		"Att_ResDeath":       45,
		"Att_ResTypeEnd":     46,
		"Att_SecondEnd":      46,
		"Att_End":            46,
	}
)

func (x AttType) Enum() *AttType {
	p := new(AttType)
	*p = x
	return p
}

func (x AttType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AttType) Descriptor() protoreflect.EnumDescriptor {
	return file_att_proto_enumTypes[0].Descriptor()
}

func (AttType) Type() protoreflect.EnumType {
	return &file_att_proto_enumTypes[0]
}

func (x AttType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AttType.Descriptor instead.
func (AttType) EnumDescriptor() ([]byte, []int) {
	return file_att_proto_rawDescGZIP(), []int{0}
}

// 普通属性
type Att struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AttType  AttType `protobuf:"varint,1,opt,name=AttType,proto3,enum=proto.AttType" json:"AttType,omitempty"`
	AttValue float32 `protobuf:"fixed32,2,opt,name=AttValue,proto3" json:"AttValue,omitempty"`
}

func (x *Att) Reset() {
	*x = Att{}
	if protoimpl.UnsafeEnabled {
		mi := &file_att_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Att) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Att) ProtoMessage() {}

func (x *Att) ProtoReflect() protoreflect.Message {
	mi := &file_att_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Att.ProtoReflect.Descriptor instead.
func (*Att) Descriptor() ([]byte, []int) {
	return file_att_proto_rawDescGZIP(), []int{0}
}

func (x *Att) GetAttType() AttType {
	if x != nil {
		return x.AttType
	}
	return AttType_Att_Begin
}

func (x *Att) GetAttValue() float32 {
	if x != nil {
		return x.AttValue
	}
	return 0
}

var File_att_proto protoreflect.FileDescriptor

var file_att_proto_rawDesc = []byte{
	0x0a, 0x09, 0x61, 0x74, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x4b, 0x0a, 0x03, 0x41, 0x74, 0x74, 0x12, 0x28, 0x0a, 0x07, 0x41, 0x74, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x41, 0x74, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x07, 0x41, 0x74, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x41, 0x74, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x41, 0x74, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x2a,
	0xd4, 0x07, 0x0a, 0x07, 0x41, 0x74, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0d, 0x0a, 0x09, 0x41,
	0x74, 0x74, 0x5f, 0x42, 0x65, 0x67, 0x69, 0x6e, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x41, 0x74,
	0x74, 0x5f, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x42, 0x65, 0x67, 0x69, 0x6e, 0x10, 0x00, 0x12,
	0x0b, 0x0a, 0x07, 0x41, 0x74, 0x74, 0x5f, 0x41, 0x74, 0x6b, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09,
	0x41, 0x74, 0x74, 0x5f, 0x41, 0x72, 0x6d, 0x6f, 0x72, 0x10, 0x01, 0x12, 0x12, 0x0a, 0x0e, 0x41,
	0x74, 0x74, 0x5f, 0x53, 0x65, 0x6c, 0x66, 0x44, 0x6d, 0x67, 0x49, 0x6e, 0x63, 0x10, 0x02, 0x12,
	0x15, 0x0a, 0x11, 0x41, 0x74, 0x74, 0x5f, 0x45, 0x6e, 0x65, 0x6d, 0x79, 0x57, 0x6f, 0x75, 0x6e,
	0x64, 0x49, 0x6e, 0x63, 0x10, 0x03, 0x12, 0x12, 0x0a, 0x0e, 0x41, 0x74, 0x74, 0x5f, 0x53, 0x65,
	0x6c, 0x66, 0x44, 0x6d, 0x67, 0x44, 0x65, 0x63, 0x10, 0x04, 0x12, 0x15, 0x0a, 0x11, 0x41, 0x74,
	0x74, 0x5f, 0x45, 0x6e, 0x65, 0x6d, 0x79, 0x57, 0x6f, 0x75, 0x6e, 0x64, 0x44, 0x65, 0x63, 0x10,
	0x05, 0x12, 0x0c, 0x0a, 0x08, 0x41, 0x74, 0x74, 0x5f, 0x43, 0x72, 0x69, 0x74, 0x10, 0x06, 0x12,
	0x0f, 0x0a, 0x0b, 0x41, 0x74, 0x74, 0x5f, 0x43, 0x72, 0x69, 0x74, 0x49, 0x6e, 0x63, 0x10, 0x07,
	0x12, 0x10, 0x0a, 0x0c, 0x41, 0x74, 0x74, 0x5f, 0x54, 0x65, 0x6e, 0x61, 0x63, 0x69, 0x74, 0x79,
	0x10, 0x08, 0x12, 0x0c, 0x0a, 0x08, 0x41, 0x74, 0x74, 0x5f, 0x48, 0x65, 0x61, 0x6c, 0x10, 0x09,
	0x12, 0x0f, 0x0a, 0x0b, 0x41, 0x74, 0x74, 0x5f, 0x52, 0x65, 0x61, 0x6c, 0x44, 0x6d, 0x67, 0x10,
	0x0a, 0x12, 0x11, 0x0a, 0x0d, 0x41, 0x74, 0x74, 0x5f, 0x4d, 0x6f, 0x76, 0x65, 0x53, 0x70, 0x65,
	0x65, 0x64, 0x10, 0x0b, 0x12, 0x10, 0x0a, 0x0c, 0x41, 0x74, 0x74, 0x5f, 0x41, 0x74, 0x62, 0x53,
	0x70, 0x65, 0x65, 0x64, 0x10, 0x0c, 0x12, 0x11, 0x0a, 0x0d, 0x41, 0x74, 0x74, 0x5f, 0x45, 0x66,
	0x66, 0x65, 0x63, 0x74, 0x48, 0x69, 0x74, 0x10, 0x0d, 0x12, 0x14, 0x0a, 0x10, 0x41, 0x74, 0x74,
	0x5f, 0x45, 0x66, 0x66, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x69, 0x73, 0x74, 0x10, 0x0e, 0x12,
	0x0d, 0x0a, 0x09, 0x41, 0x74, 0x74, 0x5f, 0x4d, 0x61, 0x78, 0x48, 0x50, 0x10, 0x0f, 0x12, 0x0d,
	0x0a, 0x09, 0x41, 0x74, 0x74, 0x5f, 0x43, 0x75, 0x72, 0x48, 0x50, 0x10, 0x10, 0x12, 0x0d, 0x0a,
	0x09, 0x41, 0x74, 0x74, 0x5f, 0x4d, 0x61, 0x78, 0x4d, 0x50, 0x10, 0x11, 0x12, 0x0d, 0x0a, 0x09,
	0x41, 0x74, 0x74, 0x5f, 0x43, 0x75, 0x72, 0x4d, 0x50, 0x10, 0x12, 0x12, 0x0d, 0x0a, 0x09, 0x41,
	0x74, 0x74, 0x5f, 0x47, 0x65, 0x6e, 0x4d, 0x50, 0x10, 0x13, 0x12, 0x0f, 0x0a, 0x0b, 0x41, 0x74,
	0x74, 0x5f, 0x4d, 0x61, 0x78, 0x52, 0x61, 0x67, 0x65, 0x10, 0x14, 0x12, 0x16, 0x0a, 0x12, 0x41,
	0x74, 0x74, 0x5f, 0x47, 0x65, 0x6e, 0x52, 0x61, 0x67, 0x65, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e,
	0x74, 0x10, 0x15, 0x12, 0x10, 0x0a, 0x0c, 0x41, 0x74, 0x74, 0x5f, 0x49, 0x6e, 0x69, 0x74, 0x52,
	0x61, 0x67, 0x65, 0x10, 0x16, 0x12, 0x0b, 0x0a, 0x07, 0x41, 0x74, 0x74, 0x5f, 0x48, 0x69, 0x74,
	0x10, 0x17, 0x12, 0x0d, 0x0a, 0x09, 0x41, 0x74, 0x74, 0x5f, 0x44, 0x6f, 0x64, 0x67, 0x65, 0x10,
	0x18, 0x12, 0x11, 0x0a, 0x0d, 0x41, 0x74, 0x74, 0x5f, 0x4d, 0x6f, 0x76, 0x65, 0x53, 0x63, 0x6f,
	0x70, 0x65, 0x10, 0x19, 0x12, 0x14, 0x0a, 0x10, 0x41, 0x74, 0x74, 0x5f, 0x44, 0x6d, 0x67, 0x54,
	0x79, 0x70, 0x65, 0x42, 0x65, 0x67, 0x69, 0x6e, 0x10, 0x1a, 0x12, 0x12, 0x0a, 0x0e, 0x41, 0x74,
	0x74, 0x5f, 0x44, 0x6d, 0x67, 0x50, 0x68, 0x79, 0x73, 0x69, 0x63, 0x73, 0x10, 0x1a, 0x12, 0x10,
	0x0a, 0x0c, 0x41, 0x74, 0x74, 0x5f, 0x44, 0x6d, 0x67, 0x45, 0x61, 0x72, 0x74, 0x68, 0x10, 0x1b,
	0x12, 0x10, 0x0a, 0x0c, 0x41, 0x74, 0x74, 0x5f, 0x44, 0x6d, 0x67, 0x57, 0x61, 0x74, 0x65, 0x72,
	0x10, 0x1c, 0x12, 0x0f, 0x0a, 0x0b, 0x41, 0x74, 0x74, 0x5f, 0x44, 0x6d, 0x67, 0x46, 0x69, 0x72,
	0x65, 0x10, 0x1d, 0x12, 0x0f, 0x0a, 0x0b, 0x41, 0x74, 0x74, 0x5f, 0x44, 0x6d, 0x67, 0x57, 0x69,
	0x6e, 0x64, 0x10, 0x1e, 0x12, 0x12, 0x0a, 0x0e, 0x41, 0x74, 0x74, 0x5f, 0x44, 0x6d, 0x67, 0x54,
	0x68, 0x75, 0x6e, 0x64, 0x65, 0x72, 0x10, 0x1f, 0x12, 0x0f, 0x0a, 0x0b, 0x41, 0x74, 0x74, 0x5f,
	0x44, 0x6d, 0x67, 0x54, 0x69, 0x6d, 0x65, 0x10, 0x20, 0x12, 0x10, 0x0a, 0x0c, 0x41, 0x74, 0x74,
	0x5f, 0x44, 0x6d, 0x67, 0x53, 0x70, 0x61, 0x63, 0x65, 0x10, 0x21, 0x12, 0x10, 0x0a, 0x0c, 0x41,
	0x74, 0x74, 0x5f, 0x44, 0x6d, 0x67, 0x53, 0x74, 0x65, 0x65, 0x6c, 0x10, 0x22, 0x12, 0x10, 0x0a,
	0x0c, 0x41, 0x74, 0x74, 0x5f, 0x44, 0x6d, 0x67, 0x44, 0x65, 0x61, 0x74, 0x68, 0x10, 0x23, 0x12,
	0x12, 0x0a, 0x0e, 0x41, 0x74, 0x74, 0x5f, 0x44, 0x6d, 0x67, 0x54, 0x79, 0x70, 0x65, 0x45, 0x6e,
	0x64, 0x10, 0x24, 0x12, 0x14, 0x0a, 0x10, 0x41, 0x74, 0x74, 0x5f, 0x52, 0x65, 0x73, 0x54, 0x79,
	0x70, 0x65, 0x42, 0x65, 0x67, 0x69, 0x6e, 0x10, 0x24, 0x12, 0x12, 0x0a, 0x0e, 0x41, 0x74, 0x74,
	0x5f, 0x52, 0x65, 0x73, 0x50, 0x68, 0x79, 0x73, 0x69, 0x63, 0x73, 0x10, 0x24, 0x12, 0x10, 0x0a,
	0x0c, 0x41, 0x74, 0x74, 0x5f, 0x52, 0x65, 0x73, 0x45, 0x61, 0x72, 0x74, 0x68, 0x10, 0x25, 0x12,
	0x10, 0x0a, 0x0c, 0x41, 0x74, 0x74, 0x5f, 0x52, 0x65, 0x73, 0x57, 0x61, 0x74, 0x65, 0x72, 0x10,
	0x26, 0x12, 0x0f, 0x0a, 0x0b, 0x41, 0x74, 0x74, 0x5f, 0x52, 0x65, 0x73, 0x46, 0x69, 0x72, 0x65,
	0x10, 0x27, 0x12, 0x0f, 0x0a, 0x0b, 0x41, 0x74, 0x74, 0x5f, 0x52, 0x65, 0x73, 0x57, 0x69, 0x6e,
	0x64, 0x10, 0x28, 0x12, 0x12, 0x0a, 0x0e, 0x41, 0x74, 0x74, 0x5f, 0x52, 0x65, 0x73, 0x54, 0x68,
	0x75, 0x6e, 0x64, 0x65, 0x72, 0x10, 0x29, 0x12, 0x0f, 0x0a, 0x0b, 0x41, 0x74, 0x74, 0x5f, 0x52,
	0x65, 0x73, 0x54, 0x69, 0x6d, 0x65, 0x10, 0x2a, 0x12, 0x10, 0x0a, 0x0c, 0x41, 0x74, 0x74, 0x5f,
	0x52, 0x65, 0x73, 0x53, 0x70, 0x61, 0x63, 0x65, 0x10, 0x2b, 0x12, 0x10, 0x0a, 0x0c, 0x41, 0x74,
	0x74, 0x5f, 0x52, 0x65, 0x73, 0x53, 0x74, 0x65, 0x65, 0x6c, 0x10, 0x2c, 0x12, 0x10, 0x0a, 0x0c,
	0x41, 0x74, 0x74, 0x5f, 0x52, 0x65, 0x73, 0x44, 0x65, 0x61, 0x74, 0x68, 0x10, 0x2d, 0x12, 0x12,
	0x0a, 0x0e, 0x41, 0x74, 0x74, 0x5f, 0x52, 0x65, 0x73, 0x54, 0x79, 0x70, 0x65, 0x45, 0x6e, 0x64,
	0x10, 0x2e, 0x12, 0x11, 0x0a, 0x0d, 0x41, 0x74, 0x74, 0x5f, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64,
	0x45, 0x6e, 0x64, 0x10, 0x2e, 0x12, 0x0b, 0x0a, 0x07, 0x41, 0x74, 0x74, 0x5f, 0x45, 0x6e, 0x64,
	0x10, 0x2e, 0x1a, 0x02, 0x10, 0x01, 0x42, 0x32, 0x5a, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x61, 0x73, 0x74, 0x2d, 0x65, 0x64, 0x65, 0x6e, 0x2f, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x62,
	0x61, 0x6c, 0xaa, 0x02, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_att_proto_rawDescOnce sync.Once
	file_att_proto_rawDescData = file_att_proto_rawDesc
)

func file_att_proto_rawDescGZIP() []byte {
	file_att_proto_rawDescOnce.Do(func() {
		file_att_proto_rawDescData = protoimpl.X.CompressGZIP(file_att_proto_rawDescData)
	})
	return file_att_proto_rawDescData
}

var file_att_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_att_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_att_proto_goTypes = []interface{}{
	(AttType)(0), // 0: proto.AttType
	(*Att)(nil),  // 1: proto.Att
}
var file_att_proto_depIdxs = []int32{
	0, // 0: proto.Att.AttType:type_name -> proto.AttType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_att_proto_init() }
func file_att_proto_init() {
	if File_att_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_att_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Att); i {
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
			RawDescriptor: file_att_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_att_proto_goTypes,
		DependencyIndexes: file_att_proto_depIdxs,
		EnumInfos:         file_att_proto_enumTypes,
		MessageInfos:      file_att_proto_msgTypes,
	}.Build()
	File_att_proto = out.File
	file_att_proto_rawDesc = nil
	file_att_proto_goTypes = nil
	file_att_proto_depIdxs = nil
}
