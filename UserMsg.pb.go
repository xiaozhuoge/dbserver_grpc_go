// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.2
// source: UserMsg.proto

package dbserver_grpc_go

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

type LoginByPswdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LoginName string `protobuf:"bytes,1,opt,name=LoginName,proto3" json:"LoginName"`
	PswdMd5   string `protobuf:"bytes,2,opt,name=PswdMd5,proto3" json:"PswdMd5"`
}

func (x *LoginByPswdRequest) Reset() {
	*x = LoginByPswdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_UserMsg_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginByPswdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginByPswdRequest) ProtoMessage() {}

func (x *LoginByPswdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_UserMsg_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginByPswdRequest.ProtoReflect.Descriptor instead.
func (*LoginByPswdRequest) Descriptor() ([]byte, []int) {
	return file_UserMsg_proto_rawDescGZIP(), []int{0}
}

func (x *LoginByPswdRequest) GetLoginName() string {
	if x != nil {
		return x.LoginName
	}
	return ""
}

func (x *LoginByPswdRequest) GetPswdMd5() string {
	if x != nil {
		return x.PswdMd5
	}
	return ""
}

type LoginByPswdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32            `protobuf:"varint,1,opt,name=Code,proto3" json:"Code"`
	Message string           `protobuf:"bytes,2,opt,name=Message,proto3" json:"Message"`
	Data    *SysBaseOperator `protobuf:"bytes,3,opt,name=Data,proto3" json:"Data"`
}

func (x *LoginByPswdResponse) Reset() {
	*x = LoginByPswdResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_UserMsg_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginByPswdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginByPswdResponse) ProtoMessage() {}

func (x *LoginByPswdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_UserMsg_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginByPswdResponse.ProtoReflect.Descriptor instead.
func (*LoginByPswdResponse) Descriptor() ([]byte, []int) {
	return file_UserMsg_proto_rawDescGZIP(), []int{1}
}

func (x *LoginByPswdResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *LoginByPswdResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *LoginByPswdResponse) GetData() *SysBaseOperator {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_UserMsg_proto protoreflect.FileDescriptor

var file_UserMsg_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x73, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0d, 0x64, 0x62, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x1a, 0x18,
	0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x44, 0x42, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4c, 0x0a, 0x12, 0x4c, 0x6f, 0x67, 0x69,
	0x6e, 0x42, 0x79, 0x50, 0x73, 0x77, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c,
	0x0a, 0x09, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x50, 0x73, 0x77, 0x64, 0x4d, 0x64, 0x35, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x50,
	0x73, 0x77, 0x64, 0x4d, 0x64, 0x35, 0x22, 0x77, 0x0a, 0x13, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x42,
	0x79, 0x50, 0x73, 0x77, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x43, 0x6f, 0x64,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x32, 0x0a, 0x04, 0x44,
	0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x64, 0x62, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x79, 0x73, 0x42, 0x61, 0x73,
	0x65, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x04, 0x44, 0x61, 0x74, 0x61, 0x42,
	0x14, 0x5a, 0x12, 0x2e, 0x2f, 0x64, 0x62, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x67, 0x72,
	0x70, 0x63, 0x5f, 0x67, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_UserMsg_proto_rawDescOnce sync.Once
	file_UserMsg_proto_rawDescData = file_UserMsg_proto_rawDesc
)

func file_UserMsg_proto_rawDescGZIP() []byte {
	file_UserMsg_proto_rawDescOnce.Do(func() {
		file_UserMsg_proto_rawDescData = protoimpl.X.CompressGZIP(file_UserMsg_proto_rawDescData)
	})
	return file_UserMsg_proto_rawDescData
}

var file_UserMsg_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_UserMsg_proto_goTypes = []interface{}{
	(*LoginByPswdRequest)(nil),  // 0: dbserver_grpc.LoginByPswdRequest
	(*LoginByPswdResponse)(nil), // 1: dbserver_grpc.LoginByPswdResponse
	(*SysBaseOperator)(nil),     // 2: dbserver_grpc.SysBaseOperator
}
var file_UserMsg_proto_depIdxs = []int32{
	2, // 0: dbserver_grpc.LoginByPswdResponse.Data:type_name -> dbserver_grpc.SysBaseOperator
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_UserMsg_proto_init() }
func file_UserMsg_proto_init() {
	if File_UserMsg_proto != nil {
		return
	}
	file_CommonProtoDBmodel_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_UserMsg_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginByPswdRequest); i {
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
		file_UserMsg_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginByPswdResponse); i {
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
			RawDescriptor: file_UserMsg_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_UserMsg_proto_goTypes,
		DependencyIndexes: file_UserMsg_proto_depIdxs,
		MessageInfos:      file_UserMsg_proto_msgTypes,
	}.Build()
	File_UserMsg_proto = out.File
	file_UserMsg_proto_rawDesc = nil
	file_UserMsg_proto_goTypes = nil
	file_UserMsg_proto_depIdxs = nil
}
