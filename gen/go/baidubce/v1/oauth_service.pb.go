// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: baidubce/v1/oauth_service.proto

package baidubcev1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

var File_baidubce_v1_oauth_service_proto protoreflect.FileDescriptor

var file_baidubce_v1_oauth_service_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x62, 0x61, 0x69, 0x64, 0x75, 0x62, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x61,
	0x75, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0b, 0x62, 0x61, 0x69, 0x64, 0x75, 0x62, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x62, 0x61, 0x69, 0x64, 0x75, 0x62, 0x63, 0x65, 0x2f,
	0x76, 0x31, 0x2f, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x72,
	0x0a, 0x0c, 0x4f, 0x61, 0x75, 0x74, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x62,
	0x0a, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x19, 0x2e, 0x62, 0x61, 0x69, 0x64, 0x75, 0x62,
	0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x62, 0x61, 0x69, 0x64, 0x75, 0x62, 0x63, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x22,
	0xda, 0x41, 0x01, 0x2a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18, 0x3a, 0x01, 0x2a, 0x22, 0x13, 0x2f,
	0x76, 0x31, 0x3a, 0x70, 0x61, 0x73, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x73, 0x69, 0x67, 0x6e,
	0x75, 0x70, 0x42, 0xb1, 0x01, 0x0a, 0x0f, 0x63, 0x6f, 0x6d, 0x2e, 0x62, 0x61, 0x69, 0x64, 0x75,
	0x62, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x42, 0x11, 0x4f, 0x61, 0x75, 0x74, 0x68, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3e, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x41,
	0x49, 0x2d, 0x45, 0x2f, 0x67, 0x6f, 0x2d, 0x77, 0x65, 0x6e, 0x78, 0x69, 0x6e, 0x2f, 0x67, 0x65,
	0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x62, 0x61, 0x69, 0x64, 0x75, 0x62, 0x63, 0x65, 0x2f, 0x76, 0x31,
	0x3b, 0x62, 0x61, 0x69, 0x64, 0x75, 0x62, 0x63, 0x65, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x42, 0x58,
	0x58, 0xaa, 0x02, 0x0b, 0x42, 0x61, 0x69, 0x64, 0x75, 0x62, 0x63, 0x65, 0x2e, 0x56, 0x31, 0xca,
	0x02, 0x0b, 0x42, 0x61, 0x69, 0x64, 0x75, 0x62, 0x63, 0x65, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x17,
	0x42, 0x61, 0x69, 0x64, 0x75, 0x62, 0x63, 0x65, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0c, 0x42, 0x61, 0x69, 0x64, 0x75, 0x62,
	0x63, 0x65, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_baidubce_v1_oauth_service_proto_goTypes = []interface{}{
	(*TokenRequest)(nil),  // 0: baidubce.v1.TokenRequest
	(*TokenResponse)(nil), // 1: baidubce.v1.TokenResponse
}
var file_baidubce_v1_oauth_service_proto_depIdxs = []int32{
	0, // 0: baidubce.v1.OauthService.Token:input_type -> baidubce.v1.TokenRequest
	1, // 1: baidubce.v1.OauthService.Token:output_type -> baidubce.v1.TokenResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_baidubce_v1_oauth_service_proto_init() }
func file_baidubce_v1_oauth_service_proto_init() {
	if File_baidubce_v1_oauth_service_proto != nil {
		return
	}
	file_baidubce_v1_oauth_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_baidubce_v1_oauth_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_baidubce_v1_oauth_service_proto_goTypes,
		DependencyIndexes: file_baidubce_v1_oauth_service_proto_depIdxs,
	}.Build()
	File_baidubce_v1_oauth_service_proto = out.File
	file_baidubce_v1_oauth_service_proto_rawDesc = nil
	file_baidubce_v1_oauth_service_proto_goTypes = nil
	file_baidubce_v1_oauth_service_proto_depIdxs = nil
}