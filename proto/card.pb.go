// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: card.proto

package proto

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

type CardStatusCode int32

const (
	CardStatusCode_OK             CardStatusCode = 0
	CardStatusCode_ERROR          CardStatusCode = 1
	CardStatusCode_INTERNAL_ERROR CardStatusCode = 2
	CardStatusCode_FORBIDDEN      CardStatusCode = 3
	CardStatusCode_BAD_HASH       CardStatusCode = 4
)

// Enum value maps for CardStatusCode.
var (
	CardStatusCode_name = map[int32]string{
		0: "OK",
		1: "ERROR",
		2: "INTERNAL_ERROR",
		3: "FORBIDDEN",
		4: "BAD_HASH",
	}
	CardStatusCode_value = map[string]int32{
		"OK":             0,
		"ERROR":          1,
		"INTERNAL_ERROR": 2,
		"FORBIDDEN":      3,
		"BAD_HASH":       4,
	}
)

func (x CardStatusCode) Enum() *CardStatusCode {
	p := new(CardStatusCode)
	*p = x
	return p
}

func (x CardStatusCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CardStatusCode) Descriptor() protoreflect.EnumDescriptor {
	return file_card_proto_enumTypes[0].Descriptor()
}

func (CardStatusCode) Type() protoreflect.EnumType {
	return &file_card_proto_enumTypes[0]
}

func (x CardStatusCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CardStatusCode.Descriptor instead.
func (CardStatusCode) EnumDescriptor() ([]byte, []int) {
	return file_card_proto_rawDescGZIP(), []int{0}
}

type Card struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CardId    string `protobuf:"bytes,1,opt,name=card_id,json=cardId,proto3" json:"card_id,omitempty"`
	Username  string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Content   string `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	Timestamp int64  `protobuf:"varint,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Hash      string `protobuf:"bytes,5,opt,name=hash,proto3" json:"hash,omitempty"`
}

func (x *Card) Reset() {
	*x = Card{}
	if protoimpl.UnsafeEnabled {
		mi := &file_card_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Card) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Card) ProtoMessage() {}

func (x *Card) ProtoReflect() protoreflect.Message {
	mi := &file_card_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Card.ProtoReflect.Descriptor instead.
func (*Card) Descriptor() ([]byte, []int) {
	return file_card_proto_rawDescGZIP(), []int{0}
}

func (x *Card) GetCardId() string {
	if x != nil {
		return x.CardId
	}
	return ""
}

func (x *Card) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *Card) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Card) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *Card) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

type CardRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CardId string `protobuf:"bytes,1,opt,name=card_id,json=cardId,proto3" json:"card_id,omitempty"`
}

func (x *CardRequest) Reset() {
	*x = CardRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_card_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CardRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CardRequest) ProtoMessage() {}

func (x *CardRequest) ProtoReflect() protoreflect.Message {
	mi := &file_card_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CardRequest.ProtoReflect.Descriptor instead.
func (*CardRequest) Descriptor() ([]byte, []int) {
	return file_card_proto_rawDescGZIP(), []int{1}
}

func (x *CardRequest) GetCardId() string {
	if x != nil {
		return x.CardId
	}
	return ""
}

type CardResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode    CardStatusCode `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3,enum=proto.CardStatusCode" json:"status_code,omitempty"`
	StatusMessage string         `protobuf:"bytes,2,opt,name=status_message,json=statusMessage,proto3" json:"status_message,omitempty"`
	CardId        string         `protobuf:"bytes,3,opt,name=card_id,json=cardId,proto3" json:"card_id,omitempty"`
	Card          *Card          `protobuf:"bytes,4,opt,name=card,proto3" json:"card,omitempty"`
}

func (x *CardResponse) Reset() {
	*x = CardResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_card_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CardResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CardResponse) ProtoMessage() {}

func (x *CardResponse) ProtoReflect() protoreflect.Message {
	mi := &file_card_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CardResponse.ProtoReflect.Descriptor instead.
func (*CardResponse) Descriptor() ([]byte, []int) {
	return file_card_proto_rawDescGZIP(), []int{2}
}

func (x *CardResponse) GetStatusCode() CardStatusCode {
	if x != nil {
		return x.StatusCode
	}
	return CardStatusCode_OK
}

func (x *CardResponse) GetStatusMessage() string {
	if x != nil {
		return x.StatusMessage
	}
	return ""
}

func (x *CardResponse) GetCardId() string {
	if x != nil {
		return x.CardId
	}
	return ""
}

func (x *CardResponse) GetCard() *Card {
	if x != nil {
		return x.Card
	}
	return nil
}

var File_card_proto protoreflect.FileDescriptor

var file_card_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x63, 0x61, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x87, 0x01, 0x0a, 0x04, 0x43, 0x61, 0x72, 0x64, 0x12, 0x17, 0x0a, 0x07,
	0x63, 0x61, 0x72, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63,
	0x61, 0x72, 0x64, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x61, 0x73,
	0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x22, 0x26, 0x0a,
	0x0b, 0x43, 0x61, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07,
	0x63, 0x61, 0x72, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63,
	0x61, 0x72, 0x64, 0x49, 0x64, 0x22, 0xa7, 0x01, 0x0a, 0x0c, 0x43, 0x61, 0x72, 0x64, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x61, 0x72, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f,
	0x64, 0x65, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x25,
	0x0a, 0x0e, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x63, 0x61, 0x72, 0x64, 0x5f, 0x69, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x61, 0x72, 0x64, 0x49, 0x64, 0x12, 0x1f,
	0x0a, 0x04, 0x63, 0x61, 0x72, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x61, 0x72, 0x64, 0x52, 0x04, 0x63, 0x61, 0x72, 0x64, 0x2a,
	0x54, 0x0a, 0x0e, 0x43, 0x61, 0x72, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64,
	0x65, 0x12, 0x06, 0x0a, 0x02, 0x4f, 0x4b, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x52, 0x52,
	0x4f, 0x52, 0x10, 0x01, 0x12, 0x12, 0x0a, 0x0e, 0x49, 0x4e, 0x54, 0x45, 0x52, 0x4e, 0x41, 0x4c,
	0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x02, 0x12, 0x0d, 0x0a, 0x09, 0x46, 0x4f, 0x52, 0x42,
	0x49, 0x44, 0x44, 0x45, 0x4e, 0x10, 0x03, 0x12, 0x0c, 0x0a, 0x08, 0x42, 0x41, 0x44, 0x5f, 0x48,
	0x41, 0x53, 0x48, 0x10, 0x04, 0x32, 0xdd, 0x01, 0x0a, 0x0b, 0x43, 0x61, 0x72, 0x64, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x30, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43,
	0x61, 0x72, 0x64, 0x12, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x61, 0x72, 0x64,
	0x1a, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x61, 0x72, 0x64, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x2d, 0x0a, 0x07, 0x53, 0x65, 0x74, 0x43, 0x61,
	0x72, 0x64, 0x12, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x61, 0x72, 0x64, 0x1a,
	0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x61, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x34, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72,
	0x64, 0x12, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x61, 0x72, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x61,
	0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x0a,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x61, 0x72, 0x64, 0x12, 0x12, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x43, 0x61, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x61, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x2a, 0x5a, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x6a, 0x61, 0x6d, 0x65, 0x73, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x74, 0x6f,
	0x6e, 0x6a, 0x31, 0x2f, 0x66, 0x6f, 0x72, 0x6c, 0x61, 0x74, 0x65, 0x72, 0x2d, 0x63, 0x6f, 0x72,
	0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_card_proto_rawDescOnce sync.Once
	file_card_proto_rawDescData = file_card_proto_rawDesc
)

func file_card_proto_rawDescGZIP() []byte {
	file_card_proto_rawDescOnce.Do(func() {
		file_card_proto_rawDescData = protoimpl.X.CompressGZIP(file_card_proto_rawDescData)
	})
	return file_card_proto_rawDescData
}

var file_card_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_card_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_card_proto_goTypes = []interface{}{
	(CardStatusCode)(0),  // 0: proto.CardStatusCode
	(*Card)(nil),         // 1: proto.Card
	(*CardRequest)(nil),  // 2: proto.CardRequest
	(*CardResponse)(nil), // 3: proto.CardResponse
}
var file_card_proto_depIdxs = []int32{
	0, // 0: proto.CardResponse.status_code:type_name -> proto.CardStatusCode
	1, // 1: proto.CardResponse.card:type_name -> proto.Card
	1, // 2: proto.CardService.CreateCard:input_type -> proto.Card
	1, // 3: proto.CardService.SetCard:input_type -> proto.Card
	2, // 4: proto.CardService.GetCard:input_type -> proto.CardRequest
	2, // 5: proto.CardService.DeleteCard:input_type -> proto.CardRequest
	3, // 6: proto.CardService.CreateCard:output_type -> proto.CardResponse
	3, // 7: proto.CardService.SetCard:output_type -> proto.CardResponse
	3, // 8: proto.CardService.GetCard:output_type -> proto.CardResponse
	3, // 9: proto.CardService.DeleteCard:output_type -> proto.CardResponse
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_card_proto_init() }
func file_card_proto_init() {
	if File_card_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_card_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Card); i {
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
		file_card_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CardRequest); i {
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
		file_card_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CardResponse); i {
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
			RawDescriptor: file_card_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_card_proto_goTypes,
		DependencyIndexes: file_card_proto_depIdxs,
		EnumInfos:         file_card_proto_enumTypes,
		MessageInfos:      file_card_proto_msgTypes,
	}.Build()
	File_card_proto = out.File
	file_card_proto_rawDesc = nil
	file_card_proto_goTypes = nil
	file_card_proto_depIdxs = nil
}