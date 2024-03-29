// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.3
// source: game_comm.proto

package pb

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

type DoType int32

const (
	DoType_COMMON DoType = 0
	DoType_DI_YI  DoType = 1
	DoType_DI_ER  DoType = 2
	DoType_DI_SAN DoType = 3
)

// Enum value maps for DoType.
var (
	DoType_name = map[int32]string{
		0: "COMMON",
		1: "DI_YI",
		2: "DI_ER",
		3: "DI_SAN",
	}
	DoType_value = map[string]int32{
		"COMMON": 0,
		"DI_YI":  1,
		"DI_ER":  2,
		"DI_SAN": 3,
	}
)

func (x DoType) Enum() *DoType {
	p := new(DoType)
	*p = x
	return p
}

func (x DoType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DoType) Descriptor() protoreflect.EnumDescriptor {
	return file_game_comm_proto_enumTypes[0].Descriptor()
}

func (DoType) Type() protoreflect.EnumType {
	return &file_game_comm_proto_enumTypes[0]
}

func (x DoType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DoType.Descriptor instead.
func (DoType) EnumDescriptor() ([]byte, []int) {
	return file_game_comm_proto_rawDescGZIP(), []int{0}
}

type NetMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServiceId string `protobuf:"bytes,1,opt,name=ServiceId,proto3" json:"ServiceId,omitempty"`
	UId       string `protobuf:"bytes,2,opt,name=UId,proto3" json:"UId,omitempty"`
	Content   []byte `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	Type      int32  `protobuf:"varint,4,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *NetMessage) Reset() {
	*x = NetMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_comm_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NetMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetMessage) ProtoMessage() {}

func (x *NetMessage) ProtoReflect() protoreflect.Message {
	mi := &file_game_comm_proto_msgTypes[0]
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
	return file_game_comm_proto_rawDescGZIP(), []int{0}
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

type GameMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GameMessage) Reset() {
	*x = GameMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_comm_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GameMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GameMessage) ProtoMessage() {}

func (x *GameMessage) ProtoReflect() protoreflect.Message {
	mi := &file_game_comm_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GameMessage.ProtoReflect.Descriptor instead.
func (*GameMessage) Descriptor() ([]byte, []int) {
	return file_game_comm_proto_rawDescGZIP(), []int{1}
}

var File_game_comm_proto protoreflect.FileDescriptor

var file_game_comm_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x6a, 0x0a, 0x0a, 0x6e, 0x65, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x1c, 0x0a, 0x09, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x10, 0x0a,
	0x03, 0x55, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x55, 0x49, 0x64, 0x12,
	0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x0d, 0x0a,
	0x0b, 0x67, 0x61, 0x6d, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2a, 0x36, 0x0a, 0x06,
	0x44, 0x6f, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0a, 0x0a, 0x06, 0x43, 0x4f, 0x4d, 0x4d, 0x4f, 0x4e,
	0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x44, 0x49, 0x5f, 0x59, 0x49, 0x10, 0x01, 0x12, 0x09, 0x0a,
	0x05, 0x44, 0x49, 0x5f, 0x45, 0x52, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x44, 0x49, 0x5f, 0x53,
	0x41, 0x4e, 0x10, 0x03, 0x42, 0x05, 0x5a, 0x03, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_game_comm_proto_rawDescOnce sync.Once
	file_game_comm_proto_rawDescData = file_game_comm_proto_rawDesc
)

func file_game_comm_proto_rawDescGZIP() []byte {
	file_game_comm_proto_rawDescOnce.Do(func() {
		file_game_comm_proto_rawDescData = protoimpl.X.CompressGZIP(file_game_comm_proto_rawDescData)
	})
	return file_game_comm_proto_rawDescData
}

var file_game_comm_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_game_comm_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_game_comm_proto_goTypes = []interface{}{
	(DoType)(0),         // 0: DoType
	(*NetMessage)(nil),  // 1: netMessage
	(*GameMessage)(nil), // 2: gameMessage
}
var file_game_comm_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_game_comm_proto_init() }
func file_game_comm_proto_init() {
	if File_game_comm_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_game_comm_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_game_comm_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GameMessage); i {
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
			RawDescriptor: file_game_comm_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_game_comm_proto_goTypes,
		DependencyIndexes: file_game_comm_proto_depIdxs,
		EnumInfos:         file_game_comm_proto_enumTypes,
		MessageInfos:      file_game_comm_proto_msgTypes,
	}.Build()
	File_game_comm_proto = out.File
	file_game_comm_proto_rawDesc = nil
	file_game_comm_proto_goTypes = nil
	file_game_comm_proto_depIdxs = nil
}
