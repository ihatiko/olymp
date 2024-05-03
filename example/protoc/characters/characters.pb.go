// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.12
// source: protoc/characters/characters.proto

package characters

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

type UpdateCharactersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateCharactersRequest) Reset() {
	*x = UpdateCharactersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protoc_characters_characters_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCharactersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCharactersRequest) ProtoMessage() {}

func (x *UpdateCharactersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protoc_characters_characters_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCharactersRequest.ProtoReflect.Descriptor instead.
func (*UpdateCharactersRequest) Descriptor() ([]byte, []int) {
	return file_protoc_characters_characters_proto_rawDescGZIP(), []int{0}
}

type UpdateCharactersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateCharactersResponse) Reset() {
	*x = UpdateCharactersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protoc_characters_characters_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCharactersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCharactersResponse) ProtoMessage() {}

func (x *UpdateCharactersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protoc_characters_characters_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCharactersResponse.ProtoReflect.Descriptor instead.
func (*UpdateCharactersResponse) Descriptor() ([]byte, []int) {
	return file_protoc_characters_characters_proto_rawDescGZIP(), []int{1}
}

var File_protoc_characters_characters_proto protoreflect.FileDescriptor

var file_protoc_characters_characters_proto_rawDesc = []byte{
	0x0a, 0x22, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2f, 0x63, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74,
	0x65, 0x72, 0x73, 0x2f, 0x63, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x63, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x73,
	0x22, 0x19, 0x0a, 0x17, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63,
	0x74, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x1a, 0x0a, 0x18, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x73, 0x0a, 0x11, 0x43, 0x68, 0x61, 0x72, 0x61,
	0x63, 0x74, 0x65, 0x72, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5e, 0x0a, 0x0f,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x12,
	0x23, 0x2e, 0x63, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x63, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72,
	0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65,
	0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0f, 0x5a, 0x0d,
	0x2e, 0x2e, 0x3b, 0x63, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protoc_characters_characters_proto_rawDescOnce sync.Once
	file_protoc_characters_characters_proto_rawDescData = file_protoc_characters_characters_proto_rawDesc
)

func file_protoc_characters_characters_proto_rawDescGZIP() []byte {
	file_protoc_characters_characters_proto_rawDescOnce.Do(func() {
		file_protoc_characters_characters_proto_rawDescData = protoimpl.X.CompressGZIP(file_protoc_characters_characters_proto_rawDescData)
	})
	return file_protoc_characters_characters_proto_rawDescData
}

var file_protoc_characters_characters_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_protoc_characters_characters_proto_goTypes = []interface{}{
	(*UpdateCharactersRequest)(nil),  // 0: characters.UpdateCharactersRequest
	(*UpdateCharactersResponse)(nil), // 1: characters.UpdateCharactersResponse
}
var file_protoc_characters_characters_proto_depIdxs = []int32{
	0, // 0: characters.CharactersService.UpdateCharacter:input_type -> characters.UpdateCharactersRequest
	1, // 1: characters.CharactersService.UpdateCharacter:output_type -> characters.UpdateCharactersResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protoc_characters_characters_proto_init() }
func file_protoc_characters_characters_proto_init() {
	if File_protoc_characters_characters_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protoc_characters_characters_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateCharactersRequest); i {
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
		file_protoc_characters_characters_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateCharactersResponse); i {
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
			RawDescriptor: file_protoc_characters_characters_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protoc_characters_characters_proto_goTypes,
		DependencyIndexes: file_protoc_characters_characters_proto_depIdxs,
		MessageInfos:      file_protoc_characters_characters_proto_msgTypes,
	}.Build()
	File_protoc_characters_characters_proto = out.File
	file_protoc_characters_characters_proto_rawDesc = nil
	file_protoc_characters_characters_proto_goTypes = nil
	file_protoc_characters_characters_proto_depIdxs = nil
}