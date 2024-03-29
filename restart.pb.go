// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.0
// source: restart.proto

package restartpb

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

type RestartRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RestartRequest) Reset() {
	*x = RestartRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_restart_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RestartRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RestartRequest) ProtoMessage() {}

func (x *RestartRequest) ProtoReflect() protoreflect.Message {
	mi := &file_restart_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RestartRequest.ProtoReflect.Descriptor instead.
func (*RestartRequest) Descriptor() ([]byte, []int) {
	return file_restart_proto_rawDescGZIP(), []int{0}
}

type RestartResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RestartResponse) Reset() {
	*x = RestartResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_restart_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RestartResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RestartResponse) ProtoMessage() {}

func (x *RestartResponse) ProtoReflect() protoreflect.Message {
	mi := &file_restart_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RestartResponse.ProtoReflect.Descriptor instead.
func (*RestartResponse) Descriptor() ([]byte, []int) {
	return file_restart_proto_rawDescGZIP(), []int{1}
}

type TerminateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *TerminateRequest) Reset() {
	*x = TerminateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_restart_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TerminateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TerminateRequest) ProtoMessage() {}

func (x *TerminateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_restart_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TerminateRequest.ProtoReflect.Descriptor instead.
func (*TerminateRequest) Descriptor() ([]byte, []int) {
	return file_restart_proto_rawDescGZIP(), []int{2}
}

type TerminateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *TerminateResponse) Reset() {
	*x = TerminateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_restart_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TerminateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TerminateResponse) ProtoMessage() {}

func (x *TerminateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_restart_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TerminateResponse.ProtoReflect.Descriptor instead.
func (*TerminateResponse) Descriptor() ([]byte, []int) {
	return file_restart_proto_rawDescGZIP(), []int{3}
}

type CrashRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CrashRequest) Reset() {
	*x = CrashRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_restart_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CrashRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CrashRequest) ProtoMessage() {}

func (x *CrashRequest) ProtoReflect() protoreflect.Message {
	mi := &file_restart_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CrashRequest.ProtoReflect.Descriptor instead.
func (*CrashRequest) Descriptor() ([]byte, []int) {
	return file_restart_proto_rawDescGZIP(), []int{4}
}

type CrashResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CrashResponse) Reset() {
	*x = CrashResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_restart_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CrashResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CrashResponse) ProtoMessage() {}

func (x *CrashResponse) ProtoReflect() protoreflect.Message {
	mi := &file_restart_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CrashResponse.ProtoReflect.Descriptor instead.
func (*CrashResponse) Descriptor() ([]byte, []int) {
	return file_restart_proto_rawDescGZIP(), []int{5}
}

var File_restart_proto protoreflect.FileDescriptor

var file_restart_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x72, 0x65, 0x73, 0x74, 0x61, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x09, 0x72, 0x65, 0x73, 0x74, 0x61, 0x72, 0x74, 0x70, 0x62, 0x22, 0x10, 0x0a, 0x0e, 0x52, 0x65,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x11, 0x0a, 0x0f,
	0x52, 0x65, 0x73, 0x74, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x12, 0x0a, 0x10, 0x54, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x22, 0x13, 0x0a, 0x11, 0x54, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x0e, 0x0a, 0x0c, 0x43, 0x72, 0x61, 0x73,
	0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x0f, 0x0a, 0x0d, 0x43, 0x72, 0x61, 0x73,
	0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xd5, 0x01, 0x0a, 0x07, 0x52, 0x65,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x12, 0x42, 0x0a, 0x07, 0x52, 0x65, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x12, 0x19, 0x2e, 0x72, 0x65, 0x73, 0x74, 0x61, 0x72, 0x74, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x72, 0x65,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x73, 0x74, 0x61, 0x72, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x48, 0x0a, 0x09, 0x54, 0x65, 0x72,
	0x6d, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x12, 0x1b, 0x2e, 0x72, 0x65, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x70, 0x62, 0x2e, 0x54, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x72, 0x65, 0x73, 0x74, 0x61, 0x72, 0x74, 0x70, 0x62, 0x2e,
	0x54, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x3c, 0x0a, 0x05, 0x43, 0x72, 0x61, 0x73, 0x68, 0x12, 0x17, 0x2e, 0x72,
	0x65, 0x73, 0x74, 0x61, 0x72, 0x74, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x61, 0x73, 0x68, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x72, 0x65, 0x73, 0x74, 0x61, 0x72, 0x74, 0x70,
	0x62, 0x2e, 0x43, 0x72, 0x61, 0x73, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0x2a, 0x5a, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x6a, 0x6f, 0x61, 0x6b, 0x69, 0x6d, 0x6f, 0x66, 0x76, 0x2f, 0x72, 0x65, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x70, 0x62, 0x3b, 0x72, 0x65, 0x73, 0x74, 0x61, 0x72, 0x74, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_restart_proto_rawDescOnce sync.Once
	file_restart_proto_rawDescData = file_restart_proto_rawDesc
)

func file_restart_proto_rawDescGZIP() []byte {
	file_restart_proto_rawDescOnce.Do(func() {
		file_restart_proto_rawDescData = protoimpl.X.CompressGZIP(file_restart_proto_rawDescData)
	})
	return file_restart_proto_rawDescData
}

var file_restart_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_restart_proto_goTypes = []interface{}{
	(*RestartRequest)(nil),    // 0: restartpb.RestartRequest
	(*RestartResponse)(nil),   // 1: restartpb.RestartResponse
	(*TerminateRequest)(nil),  // 2: restartpb.TerminateRequest
	(*TerminateResponse)(nil), // 3: restartpb.TerminateResponse
	(*CrashRequest)(nil),      // 4: restartpb.CrashRequest
	(*CrashResponse)(nil),     // 5: restartpb.CrashResponse
}
var file_restart_proto_depIdxs = []int32{
	0, // 0: restartpb.Restart.Restart:input_type -> restartpb.RestartRequest
	2, // 1: restartpb.Restart.Terminate:input_type -> restartpb.TerminateRequest
	4, // 2: restartpb.Restart.Crash:input_type -> restartpb.CrashRequest
	1, // 3: restartpb.Restart.Restart:output_type -> restartpb.RestartResponse
	3, // 4: restartpb.Restart.Terminate:output_type -> restartpb.TerminateResponse
	5, // 5: restartpb.Restart.Crash:output_type -> restartpb.CrashResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_restart_proto_init() }
func file_restart_proto_init() {
	if File_restart_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_restart_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RestartRequest); i {
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
		file_restart_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RestartResponse); i {
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
		file_restart_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TerminateRequest); i {
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
		file_restart_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TerminateResponse); i {
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
		file_restart_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CrashRequest); i {
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
		file_restart_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CrashResponse); i {
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
			RawDescriptor: file_restart_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_restart_proto_goTypes,
		DependencyIndexes: file_restart_proto_depIdxs,
		MessageInfos:      file_restart_proto_msgTypes,
	}.Build()
	File_restart_proto = out.File
	file_restart_proto_rawDesc = nil
	file_restart_proto_goTypes = nil
	file_restart_proto_depIdxs = nil
}
