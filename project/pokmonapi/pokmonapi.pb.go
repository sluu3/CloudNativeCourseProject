// Proto file for movie info service. Note this is gRPC proto syntax (not Go)

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: pokmonapi/pokmonapi.proto

package pokmonapi

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// The request message containing ONE monster name
type MonsterName struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Monster string `protobuf:"bytes,1,opt,name=monster,proto3" json:"monster,omitempty"`
}

func (x *MonsterName) Reset() {
	*x = MonsterName{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pokmonapi_pokmonapi_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MonsterName) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MonsterName) ProtoMessage() {}

func (x *MonsterName) ProtoReflect() protoreflect.Message {
	mi := &file_pokmonapi_pokmonapi_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MonsterName.ProtoReflect.Descriptor instead.
func (*MonsterName) Descriptor() ([]byte, []int) {
	return file_pokmonapi_pokmonapi_proto_rawDescGZIP(), []int{0}
}

func (x *MonsterName) GetMonster() string {
	if x != nil {
		return x.Monster
	}
	return ""
}

// The reply message containing multiple monster names
type MonsterNames struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Monsters []string `protobuf:"bytes,1,rep,name=monsters,proto3" json:"monsters,omitempty"`
}

func (x *MonsterNames) Reset() {
	*x = MonsterNames{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pokmonapi_pokmonapi_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MonsterNames) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MonsterNames) ProtoMessage() {}

func (x *MonsterNames) ProtoReflect() protoreflect.Message {
	mi := &file_pokmonapi_pokmonapi_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MonsterNames.ProtoReflect.Descriptor instead.
func (*MonsterNames) Descriptor() ([]byte, []int) {
	return file_pokmonapi_pokmonapi_proto_rawDescGZIP(), []int{1}
}

func (x *MonsterNames) GetMonsters() []string {
	if x != nil {
		return x.Monsters
	}
	return nil
}

// The request message containing username and monster name
type UserAndName struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Monster string `protobuf:"bytes,2,opt,name=monster,proto3" json:"monster,omitempty"`
}

func (x *UserAndName) Reset() {
	*x = UserAndName{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pokmonapi_pokmonapi_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserAndName) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserAndName) ProtoMessage() {}

func (x *UserAndName) ProtoReflect() protoreflect.Message {
	mi := &file_pokmonapi_pokmonapi_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserAndName.ProtoReflect.Descriptor instead.
func (*UserAndName) Descriptor() ([]byte, []int) {
	return file_pokmonapi_pokmonapi_proto_rawDescGZIP(), []int{2}
}

func (x *UserAndName) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UserAndName) GetMonster() string {
	if x != nil {
		return x.Monster
	}
	return ""
}

// The request message containing monster action
type MonsterAction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Action string `protobuf:"bytes,2,opt,name=action,proto3" json:"action,omitempty"`
	GameID string `protobuf:"bytes,3,opt,name=gameID,proto3" json:"gameID,omitempty"`
}

func (x *MonsterAction) Reset() {
	*x = MonsterAction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pokmonapi_pokmonapi_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MonsterAction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MonsterAction) ProtoMessage() {}

func (x *MonsterAction) ProtoReflect() protoreflect.Message {
	mi := &file_pokmonapi_pokmonapi_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MonsterAction.ProtoReflect.Descriptor instead.
func (*MonsterAction) Descriptor() ([]byte, []int) {
	return file_pokmonapi_pokmonapi_proto_rawDescGZIP(), []int{3}
}

func (x *MonsterAction) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *MonsterAction) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

func (x *MonsterAction) GetGameID() string {
	if x != nil {
		return x.GameID
	}
	return ""
}

// The message for setting UserName
type UserName struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *UserName) Reset() {
	*x = UserName{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pokmonapi_pokmonapi_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserName) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserName) ProtoMessage() {}

func (x *UserName) ProtoReflect() protoreflect.Message {
	mi := &file_pokmonapi_pokmonapi_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserName.ProtoReflect.Descriptor instead.
func (*UserName) Descriptor() ([]byte, []int) {
	return file_pokmonapi_pokmonapi_proto_rawDescGZIP(), []int{4}
}

func (x *UserName) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type HealthRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	GameID string `protobuf:"bytes,2,opt,name=gameID,proto3" json:"gameID,omitempty"`
}

func (x *HealthRequest) Reset() {
	*x = HealthRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pokmonapi_pokmonapi_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HealthRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealthRequest) ProtoMessage() {}

func (x *HealthRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pokmonapi_pokmonapi_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HealthRequest.ProtoReflect.Descriptor instead.
func (*HealthRequest) Descriptor() ([]byte, []int) {
	return file_pokmonapi_pokmonapi_proto_rawDescGZIP(), []int{5}
}

func (x *HealthRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *HealthRequest) GetGameID() string {
	if x != nil {
		return x.GameID
	}
	return ""
}

type HealthPoints struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Health     int32  `protobuf:"varint,1,opt,name=health,proto3" json:"health,omitempty"`
	WhoseTurn  string `protobuf:"bytes,2,opt,name=whoseTurn,proto3" json:"whoseTurn,omitempty"`
	LastAttack string `protobuf:"bytes,3,opt,name=lastAttack,proto3" json:"lastAttack,omitempty"`
	Damage     int32  `protobuf:"varint,4,opt,name=damage,proto3" json:"damage,omitempty"`
}

func (x *HealthPoints) Reset() {
	*x = HealthPoints{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pokmonapi_pokmonapi_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HealthPoints) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealthPoints) ProtoMessage() {}

func (x *HealthPoints) ProtoReflect() protoreflect.Message {
	mi := &file_pokmonapi_pokmonapi_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HealthPoints.ProtoReflect.Descriptor instead.
func (*HealthPoints) Descriptor() ([]byte, []int) {
	return file_pokmonapi_pokmonapi_proto_rawDescGZIP(), []int{6}
}

func (x *HealthPoints) GetHealth() int32 {
	if x != nil {
		return x.Health
	}
	return 0
}

func (x *HealthPoints) GetWhoseTurn() string {
	if x != nil {
		return x.WhoseTurn
	}
	return ""
}

func (x *HealthPoints) GetLastAttack() string {
	if x != nil {
		return x.LastAttack
	}
	return ""
}

func (x *HealthPoints) GetDamage() int32 {
	if x != nil {
		return x.Damage
	}
	return 0
}

type GameStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OpponentName    string `protobuf:"bytes,1,opt,name=opponentName,proto3" json:"opponentName,omitempty"`
	OpponentMonster string `protobuf:"bytes,2,opt,name=opponentMonster,proto3" json:"opponentMonster,omitempty"`
	OpponentHealth  int32  `protobuf:"varint,3,opt,name=opponentHealth,proto3" json:"opponentHealth,omitempty"`
	WhoseTurn       string `protobuf:"bytes,4,opt,name=whoseTurn,proto3" json:"whoseTurn,omitempty"`
	MyHealth        int32  `protobuf:"varint,5,opt,name=myHealth,proto3" json:"myHealth,omitempty"`
	MyMonster       string `protobuf:"bytes,6,opt,name=myMonster,proto3" json:"myMonster,omitempty"`
	Code            string `protobuf:"bytes,7,opt,name=code,proto3" json:"code,omitempty"`
	GameID          string `protobuf:"bytes,8,opt,name=gameID,proto3" json:"gameID,omitempty"`
}

func (x *GameStatus) Reset() {
	*x = GameStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pokmonapi_pokmonapi_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GameStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GameStatus) ProtoMessage() {}

func (x *GameStatus) ProtoReflect() protoreflect.Message {
	mi := &file_pokmonapi_pokmonapi_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GameStatus.ProtoReflect.Descriptor instead.
func (*GameStatus) Descriptor() ([]byte, []int) {
	return file_pokmonapi_pokmonapi_proto_rawDescGZIP(), []int{7}
}

func (x *GameStatus) GetOpponentName() string {
	if x != nil {
		return x.OpponentName
	}
	return ""
}

func (x *GameStatus) GetOpponentMonster() string {
	if x != nil {
		return x.OpponentMonster
	}
	return ""
}

func (x *GameStatus) GetOpponentHealth() int32 {
	if x != nil {
		return x.OpponentHealth
	}
	return 0
}

func (x *GameStatus) GetWhoseTurn() string {
	if x != nil {
		return x.WhoseTurn
	}
	return ""
}

func (x *GameStatus) GetMyHealth() int32 {
	if x != nil {
		return x.MyHealth
	}
	return 0
}

func (x *GameStatus) GetMyMonster() string {
	if x != nil {
		return x.MyMonster
	}
	return ""
}

func (x *GameStatus) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *GameStatus) GetGameID() string {
	if x != nil {
		return x.GameID
	}
	return ""
}

type OpponentStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Health int32  `protobuf:"varint,2,opt,name=health,proto3" json:"health,omitempty"`
	Action string `protobuf:"bytes,3,opt,name=action,proto3" json:"action,omitempty"`
}

func (x *OpponentStatus) Reset() {
	*x = OpponentStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pokmonapi_pokmonapi_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OpponentStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OpponentStatus) ProtoMessage() {}

func (x *OpponentStatus) ProtoReflect() protoreflect.Message {
	mi := &file_pokmonapi_pokmonapi_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OpponentStatus.ProtoReflect.Descriptor instead.
func (*OpponentStatus) Descriptor() ([]byte, []int) {
	return file_pokmonapi_pokmonapi_proto_rawDescGZIP(), []int{8}
}

func (x *OpponentStatus) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *OpponentStatus) GetHealth() int32 {
	if x != nil {
		return x.Health
	}
	return 0
}

func (x *OpponentStatus) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

type RequestInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *RequestInfo) Reset() {
	*x = RequestInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pokmonapi_pokmonapi_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestInfo) ProtoMessage() {}

func (x *RequestInfo) ProtoReflect() protoreflect.Message {
	mi := &file_pokmonapi_pokmonapi_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestInfo.ProtoReflect.Descriptor instead.
func (*RequestInfo) Descriptor() ([]byte, []int) {
	return file_pokmonapi_pokmonapi_proto_rawDescGZIP(), []int{9}
}

func (x *RequestInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type AttackActions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Actions []string `protobuf:"bytes,1,rep,name=actions,proto3" json:"actions,omitempty"`
}

func (x *AttackActions) Reset() {
	*x = AttackActions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pokmonapi_pokmonapi_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AttackActions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AttackActions) ProtoMessage() {}

func (x *AttackActions) ProtoReflect() protoreflect.Message {
	mi := &file_pokmonapi_pokmonapi_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AttackActions.ProtoReflect.Descriptor instead.
func (*AttackActions) Descriptor() ([]byte, []int) {
	return file_pokmonapi_pokmonapi_proto_rawDescGZIP(), []int{10}
}

func (x *AttackActions) GetActions() []string {
	if x != nil {
		return x.Actions
	}
	return nil
}

type Status struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *Status) Reset() {
	*x = Status{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pokmonapi_pokmonapi_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Status) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Status) ProtoMessage() {}

func (x *Status) ProtoReflect() protoreflect.Message {
	mi := &file_pokmonapi_pokmonapi_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Status.ProtoReflect.Descriptor instead.
func (*Status) Descriptor() ([]byte, []int) {
	return file_pokmonapi_pokmonapi_proto_rawDescGZIP(), []int{11}
}

func (x *Status) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

var File_pokmonapi_pokmonapi_proto protoreflect.FileDescriptor

var file_pokmonapi_pokmonapi_proto_rawDesc = []byte{
	0x0a, 0x19, 0x70, 0x6f, 0x6b, 0x6d, 0x6f, 0x6e, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x6f, 0x6b, 0x6d,
	0x6f, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x70, 0x6f, 0x6b,
	0x6d, 0x6f, 0x6e, 0x61, 0x70, 0x69, 0x22, 0x27, 0x0a, 0x0b, 0x4d, 0x6f, 0x6e, 0x73, 0x74, 0x65,
	0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72, 0x22,
	0x2a, 0x0a, 0x0c, 0x4d, 0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x12,
	0x1a, 0x0a, 0x08, 0x6d, 0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x08, 0x6d, 0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72, 0x73, 0x22, 0x3b, 0x0a, 0x0b, 0x55,
	0x73, 0x65, 0x72, 0x41, 0x6e, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6d, 0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72, 0x22, 0x53, 0x0a, 0x0d, 0x4d, 0x6f, 0x6e, 0x73,
	0x74, 0x65, 0x72, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x61, 0x6d, 0x65, 0x49, 0x44, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x67, 0x61, 0x6d, 0x65, 0x49, 0x44, 0x22, 0x1e, 0x0a,
	0x08, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x3b, 0x0a,
	0x0d, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x61, 0x6d, 0x65, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x67, 0x61, 0x6d, 0x65, 0x49, 0x44, 0x22, 0x7c, 0x0a, 0x0c, 0x48, 0x65,
	0x61, 0x6c, 0x74, 0x68, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65,
	0x61, 0x6c, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x68, 0x65, 0x61, 0x6c,
	0x74, 0x68, 0x12, 0x1c, 0x0a, 0x09, 0x77, 0x68, 0x6f, 0x73, 0x65, 0x54, 0x75, 0x72, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x77, 0x68, 0x6f, 0x73, 0x65, 0x54, 0x75, 0x72, 0x6e,
	0x12, 0x1e, 0x0a, 0x0a, 0x6c, 0x61, 0x73, 0x74, 0x41, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6c, 0x61, 0x73, 0x74, 0x41, 0x74, 0x74, 0x61, 0x63, 0x6b,
	0x12, 0x16, 0x0a, 0x06, 0x64, 0x61, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x64, 0x61, 0x6d, 0x61, 0x67, 0x65, 0x22, 0x86, 0x02, 0x0a, 0x0a, 0x47, 0x61, 0x6d,
	0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x6f, 0x70, 0x70, 0x6f, 0x6e,
	0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6f,
	0x70, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x6f,
	0x70, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x4d, 0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x6f, 0x70, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x4d, 0x6f,
	0x6e, 0x73, 0x74, 0x65, 0x72, 0x12, 0x26, 0x0a, 0x0e, 0x6f, 0x70, 0x70, 0x6f, 0x6e, 0x65, 0x6e,
	0x74, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x6f,
	0x70, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x12, 0x1c, 0x0a,
	0x09, 0x77, 0x68, 0x6f, 0x73, 0x65, 0x54, 0x75, 0x72, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x77, 0x68, 0x6f, 0x73, 0x65, 0x54, 0x75, 0x72, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x6d,
	0x79, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x6d,
	0x79, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x12, 0x1c, 0x0a, 0x09, 0x6d, 0x79, 0x4d, 0x6f, 0x6e,
	0x73, 0x74, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6d, 0x79, 0x4d, 0x6f,
	0x6e, 0x73, 0x74, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x61, 0x6d,
	0x65, 0x49, 0x44, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x67, 0x61, 0x6d, 0x65, 0x49,
	0x44, 0x22, 0x54, 0x0a, 0x0e, 0x4f, 0x70, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x6c, 0x74,
	0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x12,
	0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x21, 0x0a, 0x0b, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x29, 0x0a, 0x0d, 0x41, 0x74,
	0x74, 0x61, 0x63, 0x6b, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x1c, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x32, 0x92, 0x05, 0x0a, 0x0a, 0x50, 0x6f, 0x6b, 0x6d, 0x6f, 0x6e, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x43, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x16, 0x2e, 0x70, 0x6f, 0x6b, 0x6d, 0x6f, 0x6e, 0x61, 0x70, 0x69,
	0x2e, 0x4d, 0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x1a, 0x17, 0x2e, 0x70,
	0x6f, 0x6b, 0x6d, 0x6f, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x4d, 0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72,
	0x4e, 0x61, 0x6d, 0x65, 0x73, 0x22, 0x00, 0x12, 0x3d, 0x0a, 0x0e, 0x53, 0x65, 0x74, 0x4d, 0x6f,
	0x6e, 0x73, 0x74, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x16, 0x2e, 0x70, 0x6f, 0x6b, 0x6d,
	0x6f, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x41, 0x6e, 0x64, 0x4e, 0x61, 0x6d,
	0x65, 0x1a, 0x11, 0x2e, 0x70, 0x6f, 0x6b, 0x6d, 0x6f, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x22, 0x00, 0x12, 0x44, 0x0a, 0x0d, 0x4d, 0x6f, 0x6e, 0x73, 0x74, 0x65,
	0x72, 0x41, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x12, 0x18, 0x2e, 0x70, 0x6f, 0x6b, 0x6d, 0x6f, 0x6e,
	0x61, 0x70, 0x69, 0x2e, 0x4d, 0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72, 0x41, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x1a, 0x17, 0x2e, 0x70, 0x6f, 0x6b, 0x6d, 0x6f, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x48, 0x65,
	0x61, 0x6c, 0x74, 0x68, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x22, 0x00, 0x12, 0x35, 0x0a, 0x09,
	0x4a, 0x6f, 0x69, 0x6e, 0x51, 0x75, 0x65, 0x75, 0x65, 0x12, 0x13, 0x2e, 0x70, 0x6f, 0x6b, 0x6d,
	0x6f, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x1a, 0x11,
	0x2e, 0x70, 0x6f, 0x6b, 0x6d, 0x6f, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x22, 0x00, 0x12, 0x35, 0x0a, 0x09, 0x4c, 0x65, 0x61, 0x76, 0x65, 0x47, 0x61, 0x6d, 0x65,
	0x12, 0x13, 0x2e, 0x70, 0x6f, 0x6b, 0x6d, 0x6f, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x4e, 0x61, 0x6d, 0x65, 0x1a, 0x11, 0x2e, 0x70, 0x6f, 0x6b, 0x6d, 0x6f, 0x6e, 0x61, 0x70,
	0x69, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x0b, 0x53, 0x65,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x13, 0x2e, 0x70, 0x6f, 0x6b, 0x6d,
	0x6f, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x1a, 0x11,
	0x2e, 0x70, 0x6f, 0x6b, 0x6d, 0x6f, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x22, 0x00, 0x12, 0x46, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68,
	0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x12, 0x18, 0x2e, 0x70, 0x6f, 0x6b, 0x6d, 0x6f, 0x6e, 0x61,
	0x70, 0x69, 0x2e, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x17, 0x2e, 0x70, 0x6f, 0x6b, 0x6d, 0x6f, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x48, 0x65, 0x61,
	0x6c, 0x74, 0x68, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x22, 0x00, 0x12, 0x3e, 0x0a, 0x0b, 0x47,
	0x65, 0x74, 0x47, 0x61, 0x6d, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x16, 0x2e, 0x70, 0x6f, 0x6b,
	0x6d, 0x6f, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x6e,
	0x66, 0x6f, 0x1a, 0x15, 0x2e, 0x70, 0x6f, 0x6b, 0x6d, 0x6f, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x47,
	0x61, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x00, 0x12, 0x46, 0x0a, 0x0f, 0x47,
	0x65, 0x74, 0x4f, 0x70, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x16,
	0x2e, 0x70, 0x6f, 0x6b, 0x6d, 0x6f, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x19, 0x2e, 0x70, 0x6f, 0x6b, 0x6d, 0x6f, 0x6e, 0x61,
	0x70, 0x69, 0x2e, 0x4f, 0x70, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x16, 0x2e, 0x70, 0x6f, 0x6b, 0x6d, 0x6f, 0x6e, 0x61, 0x70, 0x69,
	0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x18, 0x2e, 0x70,
	0x6f, 0x6b, 0x6d, 0x6f, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x41, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x41,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x00, 0x42, 0x13, 0x5a, 0x11, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x2f, 0x70, 0x6f, 0x6b, 0x6d, 0x6f, 0x6e, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pokmonapi_pokmonapi_proto_rawDescOnce sync.Once
	file_pokmonapi_pokmonapi_proto_rawDescData = file_pokmonapi_pokmonapi_proto_rawDesc
)

func file_pokmonapi_pokmonapi_proto_rawDescGZIP() []byte {
	file_pokmonapi_pokmonapi_proto_rawDescOnce.Do(func() {
		file_pokmonapi_pokmonapi_proto_rawDescData = protoimpl.X.CompressGZIP(file_pokmonapi_pokmonapi_proto_rawDescData)
	})
	return file_pokmonapi_pokmonapi_proto_rawDescData
}

var file_pokmonapi_pokmonapi_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_pokmonapi_pokmonapi_proto_goTypes = []interface{}{
	(*MonsterName)(nil),    // 0: pokmonapi.MonsterName
	(*MonsterNames)(nil),   // 1: pokmonapi.MonsterNames
	(*UserAndName)(nil),    // 2: pokmonapi.UserAndName
	(*MonsterAction)(nil),  // 3: pokmonapi.MonsterAction
	(*UserName)(nil),       // 4: pokmonapi.UserName
	(*HealthRequest)(nil),  // 5: pokmonapi.HealthRequest
	(*HealthPoints)(nil),   // 6: pokmonapi.HealthPoints
	(*GameStatus)(nil),     // 7: pokmonapi.GameStatus
	(*OpponentStatus)(nil), // 8: pokmonapi.OpponentStatus
	(*RequestInfo)(nil),    // 9: pokmonapi.RequestInfo
	(*AttackActions)(nil),  // 10: pokmonapi.AttackActions
	(*Status)(nil),         // 11: pokmonapi.Status
}
var file_pokmonapi_pokmonapi_proto_depIdxs = []int32{
	0,  // 0: pokmonapi.PokmonInfo.GetMonsterInfo:input_type -> pokmonapi.MonsterName
	2,  // 1: pokmonapi.PokmonInfo.SetMonsterInfo:input_type -> pokmonapi.UserAndName
	3,  // 2: pokmonapi.PokmonInfo.MonsterAttack:input_type -> pokmonapi.MonsterAction
	4,  // 3: pokmonapi.PokmonInfo.JoinQueue:input_type -> pokmonapi.UserName
	4,  // 4: pokmonapi.PokmonInfo.LeaveGame:input_type -> pokmonapi.UserName
	4,  // 5: pokmonapi.PokmonInfo.SetUserName:input_type -> pokmonapi.UserName
	5,  // 6: pokmonapi.PokmonInfo.GetHealthPoints:input_type -> pokmonapi.HealthRequest
	9,  // 7: pokmonapi.PokmonInfo.GetGameInfo:input_type -> pokmonapi.RequestInfo
	9,  // 8: pokmonapi.PokmonInfo.GetOpponentInfo:input_type -> pokmonapi.RequestInfo
	9,  // 9: pokmonapi.PokmonInfo.GetActionInfo:input_type -> pokmonapi.RequestInfo
	1,  // 10: pokmonapi.PokmonInfo.GetMonsterInfo:output_type -> pokmonapi.MonsterNames
	11, // 11: pokmonapi.PokmonInfo.SetMonsterInfo:output_type -> pokmonapi.Status
	6,  // 12: pokmonapi.PokmonInfo.MonsterAttack:output_type -> pokmonapi.HealthPoints
	11, // 13: pokmonapi.PokmonInfo.JoinQueue:output_type -> pokmonapi.Status
	11, // 14: pokmonapi.PokmonInfo.LeaveGame:output_type -> pokmonapi.Status
	11, // 15: pokmonapi.PokmonInfo.SetUserName:output_type -> pokmonapi.Status
	6,  // 16: pokmonapi.PokmonInfo.GetHealthPoints:output_type -> pokmonapi.HealthPoints
	7,  // 17: pokmonapi.PokmonInfo.GetGameInfo:output_type -> pokmonapi.GameStatus
	8,  // 18: pokmonapi.PokmonInfo.GetOpponentInfo:output_type -> pokmonapi.OpponentStatus
	10, // 19: pokmonapi.PokmonInfo.GetActionInfo:output_type -> pokmonapi.AttackActions
	10, // [10:20] is the sub-list for method output_type
	0,  // [0:10] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_pokmonapi_pokmonapi_proto_init() }
func file_pokmonapi_pokmonapi_proto_init() {
	if File_pokmonapi_pokmonapi_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pokmonapi_pokmonapi_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MonsterName); i {
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
		file_pokmonapi_pokmonapi_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MonsterNames); i {
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
		file_pokmonapi_pokmonapi_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserAndName); i {
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
		file_pokmonapi_pokmonapi_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MonsterAction); i {
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
		file_pokmonapi_pokmonapi_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserName); i {
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
		file_pokmonapi_pokmonapi_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HealthRequest); i {
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
		file_pokmonapi_pokmonapi_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HealthPoints); i {
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
		file_pokmonapi_pokmonapi_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GameStatus); i {
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
		file_pokmonapi_pokmonapi_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OpponentStatus); i {
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
		file_pokmonapi_pokmonapi_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestInfo); i {
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
		file_pokmonapi_pokmonapi_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AttackActions); i {
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
		file_pokmonapi_pokmonapi_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Status); i {
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
			RawDescriptor: file_pokmonapi_pokmonapi_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pokmonapi_pokmonapi_proto_goTypes,
		DependencyIndexes: file_pokmonapi_pokmonapi_proto_depIdxs,
		MessageInfos:      file_pokmonapi_pokmonapi_proto_msgTypes,
	}.Build()
	File_pokmonapi_pokmonapi_proto = out.File
	file_pokmonapi_pokmonapi_proto_rawDesc = nil
	file_pokmonapi_pokmonapi_proto_goTypes = nil
	file_pokmonapi_pokmonapi_proto_depIdxs = nil
}
