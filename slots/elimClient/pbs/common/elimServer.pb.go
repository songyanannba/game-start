// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.26.1
// source: proto/common/elimServer.proto

package common

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type SpinReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GameId    int32 `protobuf:"varint,1,opt,name=GameId,proto3" json:"GameId,omitempty"`
	SessionId int32 `protobuf:"varint,2,opt,name=SessionId,proto3" json:"SessionId,omitempty"`
	Uid       int32 `protobuf:"varint,3,opt,name=Uid,proto3" json:"Uid,omitempty"`
	FreeNum   int32 `protobuf:"varint,4,opt,name=FreeNum,proto3" json:"FreeNum,omitempty"` // 剩余免费次数
	ResNum    int32 `protobuf:"varint,5,opt,name=ResNum,proto3" json:"ResNum,omitempty"`   //  剩余Respin次数
	Raise     int64 `protobuf:"varint,6,opt,name=Raise,proto3" json:"Raise,omitempty"`
	Bet       int64 `protobuf:"varint,7,opt,name=Bet,proto3" json:"Bet,omitempty"`
}

func (x *SpinReq) Reset() {
	*x = SpinReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_common_elimServer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SpinReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SpinReq) ProtoMessage() {}

func (x *SpinReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_common_elimServer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SpinReq.ProtoReflect.Descriptor instead.
func (*SpinReq) Descriptor() ([]byte, []int) {
	return file_proto_common_elimServer_proto_rawDescGZIP(), []int{0}
}

func (x *SpinReq) GetGameId() int32 {
	if x != nil {
		return x.GameId
	}
	return 0
}

func (x *SpinReq) GetSessionId() int32 {
	if x != nil {
		return x.SessionId
	}
	return 0
}

func (x *SpinReq) GetUid() int32 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *SpinReq) GetFreeNum() int32 {
	if x != nil {
		return x.FreeNum
	}
	return 0
}

func (x *SpinReq) GetResNum() int32 {
	if x != nil {
		return x.ResNum
	}
	return 0
}

func (x *SpinReq) GetRaise() int64 {
	if x != nil {
		return x.Raise
	}
	return 0
}

func (x *SpinReq) GetBet() int64 {
	if x != nil {
		return x.Bet
	}
	return 0
}

type SpinRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg     string      `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	Code    int32       `protobuf:"varint,2,opt,name=Code,proto3" json:"Code,omitempty"`
	MsgData *NetMessage `protobuf:"bytes,3,opt,name=msgData,proto3" json:"msgData,omitempty"`
}

func (x *SpinRes) Reset() {
	*x = SpinRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_common_elimServer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SpinRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SpinRes) ProtoMessage() {}

func (x *SpinRes) ProtoReflect() protoreflect.Message {
	mi := &file_proto_common_elimServer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SpinRes.ProtoReflect.Descriptor instead.
func (*SpinRes) Descriptor() ([]byte, []int) {
	return file_proto_common_elimServer_proto_rawDescGZIP(), []int{1}
}

func (x *SpinRes) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *SpinRes) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *SpinRes) GetMsgData() *NetMessage {
	if x != nil {
		return x.MsgData
	}
	return nil
}

type SlotTestReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SID  int32 `protobuf:"varint,1,opt,name=SID,proto3" json:"SID,omitempty"`
	Bet  int32 `protobuf:"varint,2,opt,name=Bet,proto3" json:"Bet,omitempty"`
	Type int32 `protobuf:"varint,3,opt,name=Type,proto3" json:"Type,omitempty"`
}

func (x *SlotTestReq) Reset() {
	*x = SlotTestReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_common_elimServer_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SlotTestReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SlotTestReq) ProtoMessage() {}

func (x *SlotTestReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_common_elimServer_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SlotTestReq.ProtoReflect.Descriptor instead.
func (*SlotTestReq) Descriptor() ([]byte, []int) {
	return file_proto_common_elimServer_proto_rawDescGZIP(), []int{2}
}

func (x *SlotTestReq) GetSID() int32 {
	if x != nil {
		return x.SID
	}
	return 0
}

func (x *SlotTestReq) GetBet() int32 {
	if x != nil {
		return x.Bet
	}
	return 0
}

func (x *SlotTestReq) GetType() int32 {
	if x != nil {
		return x.Type
	}
	return 0
}

type SlotTestRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg  string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	Code int32  `protobuf:"varint,2,opt,name=Code,proto3" json:"Code,omitempty"`
	Data string `protobuf:"bytes,3,opt,name=Data,proto3" json:"Data,omitempty"`
}

func (x *SlotTestRes) Reset() {
	*x = SlotTestRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_common_elimServer_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SlotTestRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SlotTestRes) ProtoMessage() {}

func (x *SlotTestRes) ProtoReflect() protoreflect.Message {
	mi := &file_proto_common_elimServer_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SlotTestRes.ProtoReflect.Descriptor instead.
func (*SlotTestRes) Descriptor() ([]byte, []int) {
	return file_proto_common_elimServer_proto_rawDescGZIP(), []int{3}
}

func (x *SlotTestRes) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *SlotTestRes) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *SlotTestRes) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

type NetMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServiceId string `protobuf:"bytes,1,opt,name=ServiceId,proto3" json:"ServiceId,omitempty"`
	UId       string `protobuf:"bytes,2,opt,name=UId,proto3" json:"UId,omitempty"`
	Content   []byte `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	Type      int32  `protobuf:"varint,4,opt,name=type,proto3" json:"type,omitempty"`
	SlotID    int32  `protobuf:"varint,5,opt,name=SlotID,proto3" json:"SlotID,omitempty"`
}

func (x *NetMessage) Reset() {
	*x = NetMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_common_elimServer_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NetMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetMessage) ProtoMessage() {}

func (x *NetMessage) ProtoReflect() protoreflect.Message {
	mi := &file_proto_common_elimServer_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetMessage.ProtoReflect.Descriptor instead.
func (*NetMessage) Descriptor() ([]byte, []int) {
	return file_proto_common_elimServer_proto_rawDescGZIP(), []int{4}
}

func (x *NetMessage) GetServiceId() string {
	if x != nil {
		return x.ServiceId
	}
	return ""
}

func (x *NetMessage) GetUId() string {
	if x != nil {
		return x.UId
	}
	return ""
}

func (x *NetMessage) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

func (x *NetMessage) GetType() int32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *NetMessage) GetSlotID() int32 {
	if x != nil {
		return x.SlotID
	}
	return 0
}

var File_proto_common_elimServer_proto protoreflect.FileDescriptor

var file_proto_common_elimServer_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x65,
	0x6c, 0x69, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x06, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x22, 0xab, 0x01, 0x0a, 0x07, 0x53, 0x70, 0x69, 0x6e,
	0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x47, 0x61, 0x6d, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x47, 0x61, 0x6d, 0x65, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x53,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09,
	0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x55, 0x69, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x55, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x46,
	0x72, 0x65, 0x65, 0x4e, 0x75, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x46, 0x72,
	0x65, 0x65, 0x4e, 0x75, 0x6d, 0x12, 0x16, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x4e, 0x75, 0x6d, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x52, 0x65, 0x73, 0x4e, 0x75, 0x6d, 0x12, 0x14, 0x0a,
	0x05, 0x52, 0x61, 0x69, 0x73, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x52, 0x61,
	0x69, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x42, 0x65, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x03, 0x42, 0x65, 0x74, 0x22, 0x5d, 0x0a, 0x07, 0x53, 0x70, 0x69, 0x6e, 0x52, 0x65, 0x73,
	0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d,
	0x73, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x2c, 0x0a, 0x07, 0x6d, 0x73, 0x67, 0x44, 0x61, 0x74,
	0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x6e, 0x65, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x07, 0x6d, 0x73, 0x67,
	0x44, 0x61, 0x74, 0x61, 0x22, 0x45, 0x0a, 0x0b, 0x53, 0x6c, 0x6f, 0x74, 0x54, 0x65, 0x73, 0x74,
	0x52, 0x65, 0x71, 0x12, 0x10, 0x0a, 0x03, 0x53, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x03, 0x53, 0x49, 0x44, 0x12, 0x10, 0x0a, 0x03, 0x42, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x03, 0x42, 0x65, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x54, 0x79, 0x70, 0x65, 0x22, 0x47, 0x0a, 0x0b, 0x53,
	0x6c, 0x6f, 0x74, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73,
	0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x12, 0x0a, 0x04,
	0x43, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x43, 0x6f, 0x64, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x44, 0x61, 0x74, 0x61, 0x22, 0x82, 0x01, 0x0a, 0x0a, 0x6e, 0x65, 0x74, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49,
	0x64, 0x12, 0x10, 0x0a, 0x03, 0x55, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x55, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x6c, 0x6f, 0x74, 0x49, 0x44, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x06, 0x53, 0x6c, 0x6f, 0x74, 0x49, 0x44, 0x32, 0x69, 0x0a, 0x0b, 0x45, 0x6c, 0x69,
	0x6d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2c, 0x0a, 0x08, 0x53, 0x6c, 0x6f, 0x74,
	0x54, 0x65, 0x73, 0x74, 0x12, 0x0f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x53, 0x70,
	0x69, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x0f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x53,
	0x70, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x12, 0x2c, 0x0a, 0x08, 0x53, 0x6c, 0x6f, 0x74, 0x53, 0x70,
	0x69, 0x6e, 0x12, 0x0f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x53, 0x70, 0x69, 0x6e,
	0x52, 0x65, 0x71, 0x1a, 0x0f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x53, 0x70, 0x69,
	0x6e, 0x52, 0x65, 0x73, 0x42, 0x17, 0x5a, 0x15, 0x65, 0x6c, 0x69, 0x6d, 0x43, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x2f, 0x70, 0x62, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_common_elimServer_proto_rawDescOnce sync.Once
	file_proto_common_elimServer_proto_rawDescData = file_proto_common_elimServer_proto_rawDesc
)

func file_proto_common_elimServer_proto_rawDescGZIP() []byte {
	file_proto_common_elimServer_proto_rawDescOnce.Do(func() {
		file_proto_common_elimServer_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_common_elimServer_proto_rawDescData)
	})
	return file_proto_common_elimServer_proto_rawDescData
}

var file_proto_common_elimServer_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_common_elimServer_proto_goTypes = []interface{}{
	(*SpinReq)(nil),     // 0: common.SpinReq
	(*SpinRes)(nil),     // 1: common.SpinRes
	(*SlotTestReq)(nil), // 2: common.SlotTestReq
	(*SlotTestRes)(nil), // 3: common.SlotTestRes
	(*NetMessage)(nil),  // 4: common.netMessage
}
var file_proto_common_elimServer_proto_depIdxs = []int32{
	4, // 0: common.SpinRes.msgData:type_name -> common.netMessage
	0, // 1: common.ElimService.SlotTest:input_type -> common.SpinReq
	0, // 2: common.ElimService.SlotSpin:input_type -> common.SpinReq
	1, // 3: common.ElimService.SlotTest:output_type -> common.SpinRes
	1, // 4: common.ElimService.SlotSpin:output_type -> common.SpinRes
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_common_elimServer_proto_init() }
func file_proto_common_elimServer_proto_init() {
	if File_proto_common_elimServer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_common_elimServer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SpinReq); i {
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
		file_proto_common_elimServer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SpinRes); i {
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
		file_proto_common_elimServer_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SlotTestReq); i {
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
		file_proto_common_elimServer_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SlotTestRes); i {
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
		file_proto_common_elimServer_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NetMessage); i {
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
			RawDescriptor: file_proto_common_elimServer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_common_elimServer_proto_goTypes,
		DependencyIndexes: file_proto_common_elimServer_proto_depIdxs,
		MessageInfos:      file_proto_common_elimServer_proto_msgTypes,
	}.Build()
	File_proto_common_elimServer_proto = out.File
	file_proto_common_elimServer_proto_rawDesc = nil
	file_proto_common_elimServer_proto_goTypes = nil
	file_proto_common_elimServer_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ElimServiceClient is the client API for ElimService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ElimServiceClient interface {
	SlotTest(ctx context.Context, in *SpinReq, opts ...grpc.CallOption) (*SpinRes, error)
	SlotSpin(ctx context.Context, in *SpinReq, opts ...grpc.CallOption) (*SpinRes, error)
}

type elimServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewElimServiceClient(cc grpc.ClientConnInterface) ElimServiceClient {
	return &elimServiceClient{cc}
}

func (c *elimServiceClient) SlotTest(ctx context.Context, in *SpinReq, opts ...grpc.CallOption) (*SpinRes, error) {
	out := new(SpinRes)
	err := c.cc.Invoke(ctx, "/common.ElimService/SlotTest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *elimServiceClient) SlotSpin(ctx context.Context, in *SpinReq, opts ...grpc.CallOption) (*SpinRes, error) {
	out := new(SpinRes)
	err := c.cc.Invoke(ctx, "/common.ElimService/SlotSpin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ElimServiceServer is the server API for ElimService service.
type ElimServiceServer interface {
	SlotTest(context.Context, *SpinReq) (*SpinRes, error)
	SlotSpin(context.Context, *SpinReq) (*SpinRes, error)
}

// UnimplementedElimServiceServer can be embedded to have forward compatible implementations.
type UnimplementedElimServiceServer struct {
}

func (*UnimplementedElimServiceServer) SlotTest(context.Context, *SpinReq) (*SpinRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SlotTest not implemented")
}
func (*UnimplementedElimServiceServer) SlotSpin(context.Context, *SpinReq) (*SpinRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SlotSpin not implemented")
}

func RegisterElimServiceServer(s *grpc.Server, srv ElimServiceServer) {
	s.RegisterService(&_ElimService_serviceDesc, srv)
}

func _ElimService_SlotTest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SpinReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ElimServiceServer).SlotTest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/common.ElimService/SlotTest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ElimServiceServer).SlotTest(ctx, req.(*SpinReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ElimService_SlotSpin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SpinReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ElimServiceServer).SlotSpin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/common.ElimService/SlotSpin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ElimServiceServer).SlotSpin(ctx, req.(*SpinReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _ElimService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "common.ElimService",
	HandlerType: (*ElimServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SlotTest",
			Handler:    _ElimService_SlotTest_Handler,
		},
		{
			MethodName: "SlotSpin",
			Handler:    _ElimService_SlotSpin_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/common/elimServer.proto",
}
