// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.2
// source: UserService.proto

package dbserver_grpc_go

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_UserService_proto protoreflect.FileDescriptor

var file_UserService_proto_rawDesc = []byte{
	0x0a, 0x11, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x64, 0x62, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x67, 0x72,
	0x70, 0x63, 0x1a, 0x0d, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x73, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x32, 0x63, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x54, 0x0a, 0x0b, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x42, 0x79, 0x50, 0x73, 0x77, 0x64, 0x12,
	0x21, 0x2e, 0x64, 0x62, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x42, 0x79, 0x50, 0x73, 0x77, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x22, 0x2e, 0x64, 0x62, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x67, 0x72,
	0x70, 0x63, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x42, 0x79, 0x50, 0x73, 0x77, 0x64, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x14, 0x5a, 0x12, 0x2e, 0x2f, 0x64, 0x62, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x67, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var file_UserService_proto_goTypes = []interface{}{
	(*LoginByPswdRequest)(nil),  // 0: dbserver_grpc.LoginByPswdRequest
	(*LoginByPswdResponse)(nil), // 1: dbserver_grpc.LoginByPswdResponse
}
var file_UserService_proto_depIdxs = []int32{
	0, // 0: dbserver_grpc.UserService.LoginByPswd:input_type -> dbserver_grpc.LoginByPswdRequest
	1, // 1: dbserver_grpc.UserService.LoginByPswd:output_type -> dbserver_grpc.LoginByPswdResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_UserService_proto_init() }
func file_UserService_proto_init() {
	if File_UserService_proto != nil {
		return
	}
	file_UserMsg_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_UserService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_UserService_proto_goTypes,
		DependencyIndexes: file_UserService_proto_depIdxs,
	}.Build()
	File_UserService_proto = out.File
	file_UserService_proto_rawDesc = nil
	file_UserService_proto_goTypes = nil
	file_UserService_proto_depIdxs = nil
}
