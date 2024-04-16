// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.26.1
// source: proto/game/game.proto

package game

import (
	common "elim5/pbs/common"
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

// 游戏步骤请求
type GameStepReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Head *common.ReqHead     `protobuf:"bytes,1,opt,name=head,proto3" json:"head,omitempty"` // 请求头
	Data *common.GameRecover `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"` // 游戏恢复数据
}

func (x *GameStepReq) Reset() {
	*x = GameStepReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_game_game_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GameStepReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GameStepReq) ProtoMessage() {}

func (x *GameStepReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_game_game_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GameStepReq.ProtoReflect.Descriptor instead.
func (*GameStepReq) Descriptor() ([]byte, []int) {
	return file_proto_game_game_proto_rawDescGZIP(), []int{0}
}

func (x *GameStepReq) GetHead() *common.ReqHead {
	if x != nil {
		return x.Head
	}
	return nil
}

func (x *GameStepReq) GetData() *common.GameRecover {
	if x != nil {
		return x.Data
	}
	return nil
}

// 游戏步骤响应
type GameStepAck struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Head *common.AckHead `protobuf:"bytes,1,opt,name=head,proto3" json:"head,omitempty"` // 请求头
}

func (x *GameStepAck) Reset() {
	*x = GameStepAck{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_game_game_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GameStepAck) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GameStepAck) ProtoMessage() {}

func (x *GameStepAck) ProtoReflect() protoreflect.Message {
	mi := &file_proto_game_game_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GameStepAck.ProtoReflect.Descriptor instead.
func (*GameStepAck) Descriptor() ([]byte, []int) {
	return file_proto_game_game_proto_rawDescGZIP(), []int{1}
}

func (x *GameStepAck) GetHead() *common.AckHead {
	if x != nil {
		return x.Head
	}
	return nil
}

// 游戏流程记录
type GameProcess struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data       *common.GameRecover `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`                                // 游戏恢复数据
	RecordId   uint64              `protobuf:"varint,2,opt,name=record_id,json=recordId,proto3" json:"record_id,omitempty"`       // 记录id
	CreateTime int64               `protobuf:"varint,3,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"` // 创建时间
}

func (x *GameProcess) Reset() {
	*x = GameProcess{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_game_game_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GameProcess) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GameProcess) ProtoMessage() {}

func (x *GameProcess) ProtoReflect() protoreflect.Message {
	mi := &file_proto_game_game_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GameProcess.ProtoReflect.Descriptor instead.
func (*GameProcess) Descriptor() ([]byte, []int) {
	return file_proto_game_game_proto_rawDescGZIP(), []int{2}
}

func (x *GameProcess) GetData() *common.GameRecover {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *GameProcess) GetRecordId() uint64 {
	if x != nil {
		return x.RecordId
	}
	return 0
}

func (x *GameProcess) GetCreateTime() int64 {
	if x != nil {
		return x.CreateTime
	}
	return 0
}

// 基础牌属性
type BaseCard struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`                          //牌id
	X       int32 `protobuf:"varint,2,opt,name=x,proto3" json:"x,omitempty"`                            //x坐标
	Y       int32 `protobuf:"varint,3,opt,name=y,proto3" json:"y,omitempty"`                            //y坐标
	Mul     int32 `protobuf:"varint,4,opt,name=mul,proto3" json:"mul,omitempty"`                        // 标签的倍数
	IsWild  bool  `protobuf:"varint,5,opt,name=is_wild,json=isWild,proto3" json:"is_wild,omitempty"`    // 是否百搭
	IsPay   bool  `protobuf:"varint,6,opt,name=is_pay,json=isPay,proto3" json:"is_pay,omitempty"`       //是否pay_table
	IsValid bool  `protobuf:"varint,7,opt,name=is_valid,json=isValid,proto3" json:"is_valid,omitempty"` //是否有效
}

func (x *BaseCard) Reset() {
	*x = BaseCard{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_game_game_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BaseCard) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BaseCard) ProtoMessage() {}

func (x *BaseCard) ProtoReflect() protoreflect.Message {
	mi := &file_proto_game_game_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BaseCard.ProtoReflect.Descriptor instead.
func (*BaseCard) Descriptor() ([]byte, []int) {
	return file_proto_game_game_proto_rawDescGZIP(), []int{3}
}

func (x *BaseCard) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *BaseCard) GetX() int32 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *BaseCard) GetY() int32 {
	if x != nil {
		return x.Y
	}
	return 0
}

func (x *BaseCard) GetMul() int32 {
	if x != nil {
		return x.Mul
	}
	return 0
}

func (x *BaseCard) GetIsWild() bool {
	if x != nil {
		return x.IsWild
	}
	return false
}

func (x *BaseCard) GetIsPay() bool {
	if x != nil {
		return x.IsPay
	}
	return false
}

func (x *BaseCard) GetIsValid() bool {
	if x != nil {
		return x.IsValid
	}
	return false
}

type BaseStepFlow struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Index int32 `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"` //索引
	Gain  int64 `protobuf:"varint,2,opt,name=gain,proto3" json:"gain,omitempty"`   //总赢取
}

func (x *BaseStepFlow) Reset() {
	*x = BaseStepFlow{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_game_game_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BaseStepFlow) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BaseStepFlow) ProtoMessage() {}

func (x *BaseStepFlow) ProtoReflect() protoreflect.Message {
	mi := &file_proto_game_game_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BaseStepFlow.ProtoReflect.Descriptor instead.
func (*BaseStepFlow) Descriptor() ([]byte, []int) {
	return file_proto_game_game_proto_rawDescGZIP(), []int{4}
}

func (x *BaseStepFlow) GetIndex() int32 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *BaseStepFlow) GetGain() int64 {
	if x != nil {
		return x.Gain
	}
	return 0
}

type BaseSpinStep struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type      int32 `protobuf:"varint,1,opt,name=type,proto3" json:"type,omitempty"`                            //游戏类型 1 normal_spin 2 free_spin 3: re_spin 4: re_spin_link
	Id        int32 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`                                //id
	Pid       int32 `protobuf:"varint,3,opt,name=pid,proto3" json:"pid,omitempty"`                              //pid
	Which     int32 `protobuf:"varint,4,opt,name=which,proto3" json:"which,omitempty"`                          // 配置选择
	JackpotId int32 `protobuf:"varint,5,opt,name=jackpot_id,json=jackpotId,proto3" json:"jackpot_id,omitempty"` //奖池id
	Gain      int64 `protobuf:"varint,6,opt,name=gain,proto3" json:"gain,omitempty"`                            //赢取
}

func (x *BaseSpinStep) Reset() {
	*x = BaseSpinStep{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_game_game_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BaseSpinStep) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BaseSpinStep) ProtoMessage() {}

func (x *BaseSpinStep) ProtoReflect() protoreflect.Message {
	mi := &file_proto_game_game_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BaseSpinStep.ProtoReflect.Descriptor instead.
func (*BaseSpinStep) Descriptor() ([]byte, []int) {
	return file_proto_game_game_proto_rawDescGZIP(), []int{5}
}

func (x *BaseSpinStep) GetType() int32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *BaseSpinStep) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *BaseSpinStep) GetPid() int32 {
	if x != nil {
		return x.Pid
	}
	return 0
}

func (x *BaseSpinStep) GetWhich() int32 {
	if x != nil {
		return x.Which
	}
	return 0
}

func (x *BaseSpinStep) GetJackpotId() int32 {
	if x != nil {
		return x.JackpotId
	}
	return 0
}

func (x *BaseSpinStep) GetGain() int64 {
	if x != nil {
		return x.Gain
	}
	return 0
}

type BaseSpinAck struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Opt          *common.SpinOpt `protobuf:"bytes,1,opt,name=opt,proto3" json:"opt,omitempty"`                                        // 请求信息
	TxnId        int32           `protobuf:"varint,2,opt,name=txn_id,json=txnId,proto3" json:"txn_id,omitempty"`                      //交易id
	TotalGain    int64           `protobuf:"varint,3,opt,name=total_gain,json=totalGain,proto3" json:"total_gain,omitempty"`          //总赢取
	TotalBet     int64           `protobuf:"varint,4,opt,name=total_bet,json=totalBet,proto3" json:"total_bet,omitempty"`             //总下注
	BeforeAmount int64           `protobuf:"varint,5,opt,name=before_amount,json=beforeAmount,proto3" json:"before_amount,omitempty"` //before余额
	AfterAmount  int64           `protobuf:"varint,6,opt,name=after_amount,json=afterAmount,proto3" json:"after_amount,omitempty"`    //after余额
}

func (x *BaseSpinAck) Reset() {
	*x = BaseSpinAck{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_game_game_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BaseSpinAck) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BaseSpinAck) ProtoMessage() {}

func (x *BaseSpinAck) ProtoReflect() protoreflect.Message {
	mi := &file_proto_game_game_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BaseSpinAck.ProtoReflect.Descriptor instead.
func (*BaseSpinAck) Descriptor() ([]byte, []int) {
	return file_proto_game_game_proto_rawDescGZIP(), []int{6}
}

func (x *BaseSpinAck) GetOpt() *common.SpinOpt {
	if x != nil {
		return x.Opt
	}
	return nil
}

func (x *BaseSpinAck) GetTxnId() int32 {
	if x != nil {
		return x.TxnId
	}
	return 0
}

func (x *BaseSpinAck) GetTotalGain() int64 {
	if x != nil {
		return x.TotalGain
	}
	return 0
}

func (x *BaseSpinAck) GetTotalBet() int64 {
	if x != nil {
		return x.TotalBet
	}
	return 0
}

func (x *BaseSpinAck) GetBeforeAmount() int64 {
	if x != nil {
		return x.BeforeAmount
	}
	return 0
}

func (x *BaseSpinAck) GetAfterAmount() int64 {
	if x != nil {
		return x.AfterAmount
	}
	return 0
}

// Spin 选项请求
type OptionsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Head    *common.ReqHead `protobuf:"bytes,1,opt,name=head,proto3" json:"head,omitempty"`                      // 请求头
	SlotId  int32           `protobuf:"varint,2,opt,name=slot_id,json=slotId,proto3" json:"slot_id,omitempty"`   //slot_id
	Type    int32           `protobuf:"varint,3,opt,name=type,proto3" json:"type,omitempty"`                     //类型 1:gamble 2:into free_spin
	OptInfo string          `protobuf:"bytes,4,opt,name=opt_info,json=optInfo,proto3" json:"opt_info,omitempty"` // 选项信息
}

func (x *OptionsReq) Reset() {
	*x = OptionsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_game_game_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OptionsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OptionsReq) ProtoMessage() {}

func (x *OptionsReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_game_game_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OptionsReq.ProtoReflect.Descriptor instead.
func (*OptionsReq) Descriptor() ([]byte, []int) {
	return file_proto_game_game_proto_rawDescGZIP(), []int{7}
}

func (x *OptionsReq) GetHead() *common.ReqHead {
	if x != nil {
		return x.Head
	}
	return nil
}

func (x *OptionsReq) GetSlotId() int32 {
	if x != nil {
		return x.SlotId
	}
	return 0
}

func (x *OptionsReq) GetType() int32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *OptionsReq) GetOptInfo() string {
	if x != nil {
		return x.OptInfo
	}
	return ""
}

// Spin 选项响应
type OptionsAck struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Head       *common.AckHead `protobuf:"bytes,1,opt,name=head,proto3" json:"head,omitempty"`                               // 请求头
	SlotId     int32           `protobuf:"varint,2,opt,name=slot_id,json=slotId,proto3" json:"slot_id,omitempty"`            //slot_id
	Type       int32           `protobuf:"varint,3,opt,name=type,proto3" json:"type,omitempty"`                              //类型 1:gamble 2:into free_spin
	ProtoName  string          `protobuf:"bytes,4,opt,name=proto_name,json=protoName,proto3" json:"proto_name,omitempty"`    // proto名称
	ProtoBytes []byte          `protobuf:"bytes,5,opt,name=proto_bytes,json=protoBytes,proto3" json:"proto_bytes,omitempty"` // proto字节流
}

func (x *OptionsAck) Reset() {
	*x = OptionsAck{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_game_game_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OptionsAck) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OptionsAck) ProtoMessage() {}

func (x *OptionsAck) ProtoReflect() protoreflect.Message {
	mi := &file_proto_game_game_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OptionsAck.ProtoReflect.Descriptor instead.
func (*OptionsAck) Descriptor() ([]byte, []int) {
	return file_proto_game_game_proto_rawDescGZIP(), []int{8}
}

func (x *OptionsAck) GetHead() *common.AckHead {
	if x != nil {
		return x.Head
	}
	return nil
}

func (x *OptionsAck) GetSlotId() int32 {
	if x != nil {
		return x.SlotId
	}
	return 0
}

func (x *OptionsAck) GetType() int32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *OptionsAck) GetProtoName() string {
	if x != nil {
		return x.ProtoName
	}
	return ""
}

func (x *OptionsAck) GetProtoBytes() []byte {
	if x != nil {
		return x.ProtoBytes
	}
	return nil
}

var File_proto_game_game_proto protoreflect.FileDescriptor

var file_proto_game_game_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x61, 0x6d, 0x65, 0x2f, 0x67, 0x61, 0x6d,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x67, 0x61, 0x6d, 0x65, 0x1a, 0x19, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5f, 0x0a, 0x0d, 0x67, 0x61, 0x6d, 0x65,
	0x5f, 0x73, 0x74, 0x65, 0x70, 0x5f, 0x72, 0x65, 0x71, 0x12, 0x24, 0x0a, 0x04, 0x68, 0x65, 0x61,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x72, 0x65, 0x71, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x52, 0x04, 0x68, 0x65, 0x61, 0x64, 0x12,
	0x28, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x72, 0x65, 0x63, 0x6f,
	0x76, 0x65, 0x72, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x35, 0x0a, 0x0d, 0x67, 0x61, 0x6d,
	0x65, 0x5f, 0x73, 0x74, 0x65, 0x70, 0x5f, 0x61, 0x63, 0x6b, 0x12, 0x24, 0x0a, 0x04, 0x68, 0x65,
	0x61, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x61, 0x63, 0x6b, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x52, 0x04, 0x68, 0x65, 0x61, 0x64,
	0x22, 0x76, 0x0a, 0x0c, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73,
	0x12, 0x28, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x72, 0x65, 0x63,
	0x6f, 0x76, 0x65, 0x72, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x72,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x94, 0x01, 0x0a, 0x09, 0x62, 0x61, 0x73,
	0x65, 0x5f, 0x63, 0x61, 0x72, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x0c, 0x0a, 0x01, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x01, 0x78, 0x12, 0x0c, 0x0a, 0x01, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x01, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x75, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x03, 0x6d, 0x75, 0x6c, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x73, 0x5f, 0x77, 0x69, 0x6c, 0x64, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x69, 0x73, 0x57, 0x69, 0x6c, 0x64, 0x12, 0x15, 0x0a,
	0x06, 0x69, 0x73, 0x5f, 0x70, 0x61, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x69,
	0x73, 0x50, 0x61, 0x79, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x73, 0x5f, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x69, 0x73, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x22,
	0x3a, 0x0a, 0x0e, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x73, 0x74, 0x65, 0x70, 0x5f, 0x66, 0x6c, 0x6f,
	0x77, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x12, 0x0a, 0x04, 0x67, 0x61, 0x69, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x67, 0x61, 0x69, 0x6e, 0x22, 0x8f, 0x01, 0x0a, 0x0e,
	0x62, 0x61, 0x73, 0x65, 0x5f, 0x73, 0x70, 0x69, 0x6e, 0x5f, 0x73, 0x74, 0x65, 0x70, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x03, 0x70, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x77, 0x68, 0x69, 0x63, 0x68, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x77, 0x68, 0x69, 0x63, 0x68, 0x12, 0x1d, 0x0a, 0x0a, 0x6a, 0x61,
	0x63, 0x6b, 0x70, 0x6f, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09,
	0x6a, 0x61, 0x63, 0x6b, 0x70, 0x6f, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x67, 0x61, 0x69,
	0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x67, 0x61, 0x69, 0x6e, 0x22, 0xce, 0x01,
	0x0a, 0x0d, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x73, 0x70, 0x69, 0x6e, 0x5f, 0x61, 0x63, 0x6b, 0x12,
	0x22, 0x0a, 0x03, 0x6f, 0x70, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x73, 0x70, 0x69, 0x6e, 0x5f, 0x6f, 0x70, 0x74, 0x52, 0x03,
	0x6f, 0x70, 0x74, 0x12, 0x15, 0x0a, 0x06, 0x74, 0x78, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x78, 0x6e, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x5f, 0x67, 0x61, 0x69, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x47, 0x61, 0x69, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x5f, 0x62, 0x65, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x42, 0x65, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x62, 0x65, 0x66, 0x6f, 0x72, 0x65,
	0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x62,
	0x65, 0x66, 0x6f, 0x72, 0x65, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x61,
	0x66, 0x74, 0x65, 0x72, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0b, 0x61, 0x66, 0x74, 0x65, 0x72, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x7b,
	0x0a, 0x0b, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x5f, 0x72, 0x65, 0x71, 0x12, 0x24, 0x0a,
	0x04, 0x68, 0x65, 0x61, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x72, 0x65, 0x71, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x52, 0x04, 0x68,
	0x65, 0x61, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x6c, 0x6f, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x6c, 0x6f, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x12, 0x19, 0x0a, 0x08, 0x6f, 0x70, 0x74, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6f, 0x70, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0xa0, 0x01, 0x0a, 0x0b,
	0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x5f, 0x61, 0x63, 0x6b, 0x12, 0x24, 0x0a, 0x04, 0x68,
	0x65, 0x61, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x61, 0x63, 0x6b, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x52, 0x04, 0x68, 0x65, 0x61,
	0x64, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x6c, 0x6f, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x73, 0x6c, 0x6f, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1d,
	0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a,
	0x0b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x42, 0x79, 0x74, 0x65, 0x73, 0x42, 0x10,
	0x5a, 0x0e, 0x65, 0x6c, 0x69, 0x6d, 0x35, 0x2f, 0x70, 0x62, 0x73, 0x2f, 0x67, 0x61, 0x6d, 0x65,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_game_game_proto_rawDescOnce sync.Once
	file_proto_game_game_proto_rawDescData = file_proto_game_game_proto_rawDesc
)

func file_proto_game_game_proto_rawDescGZIP() []byte {
	file_proto_game_game_proto_rawDescOnce.Do(func() {
		file_proto_game_game_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_game_game_proto_rawDescData)
	})
	return file_proto_game_game_proto_rawDescData
}

var file_proto_game_game_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_proto_game_game_proto_goTypes = []interface{}{
	(*GameStepReq)(nil),        // 0: game.game_step_req
	(*GameStepAck)(nil),        // 1: game.game_step_ack
	(*GameProcess)(nil),        // 2: game.game_process
	(*BaseCard)(nil),           // 3: game.base_card
	(*BaseStepFlow)(nil),       // 4: game.base_step_flow
	(*BaseSpinStep)(nil),       // 5: game.base_spin_step
	(*BaseSpinAck)(nil),        // 6: game.base_spin_ack
	(*OptionsReq)(nil),         // 7: game.options_req
	(*OptionsAck)(nil),         // 8: game.options_ack
	(*common.ReqHead)(nil),     // 9: common.req_head
	(*common.GameRecover)(nil), // 10: common.game_recover
	(*common.AckHead)(nil),     // 11: common.ack_head
	(*common.SpinOpt)(nil),     // 12: common.spin_opt
}
var file_proto_game_game_proto_depIdxs = []int32{
	9,  // 0: game.game_step_req.head:type_name -> common.req_head
	10, // 1: game.game_step_req.data:type_name -> common.game_recover
	11, // 2: game.game_step_ack.head:type_name -> common.ack_head
	10, // 3: game.game_process.data:type_name -> common.game_recover
	12, // 4: game.base_spin_ack.opt:type_name -> common.spin_opt
	9,  // 5: game.options_req.head:type_name -> common.req_head
	11, // 6: game.options_ack.head:type_name -> common.ack_head
	7,  // [7:7] is the sub-list for method output_type
	7,  // [7:7] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_proto_game_game_proto_init() }
func file_proto_game_game_proto_init() {
	if File_proto_game_game_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_game_game_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GameStepReq); i {
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
		file_proto_game_game_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GameStepAck); i {
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
		file_proto_game_game_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GameProcess); i {
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
		file_proto_game_game_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BaseCard); i {
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
		file_proto_game_game_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BaseStepFlow); i {
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
		file_proto_game_game_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BaseSpinStep); i {
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
		file_proto_game_game_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BaseSpinAck); i {
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
		file_proto_game_game_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OptionsReq); i {
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
		file_proto_game_game_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OptionsAck); i {
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
			RawDescriptor: file_proto_game_game_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_game_game_proto_goTypes,
		DependencyIndexes: file_proto_game_game_proto_depIdxs,
		MessageInfos:      file_proto_game_game_proto_msgTypes,
	}.Build()
	File_proto_game_game_proto = out.File
	file_proto_game_game_proto_rawDesc = nil
	file_proto_game_game_proto_goTypes = nil
	file_proto_game_game_proto_depIdxs = nil
}
