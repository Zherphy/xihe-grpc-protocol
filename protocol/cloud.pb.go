// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v5.27.2
// source: cloud.proto

package protocol

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

type PodInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PodId     string `protobuf:"bytes,1,opt,name=pod_id,json=podId,proto3" json:"pod_id,omitempty"`
	Error     string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	AccessUrl string `protobuf:"bytes,3,opt,name=access_url,json=accessUrl,proto3" json:"access_url,omitempty"`
}

func (x *PodInfo) Reset() {
	*x = PodInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PodInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PodInfo) ProtoMessage() {}

func (x *PodInfo) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PodInfo.ProtoReflect.Descriptor instead.
func (*PodInfo) Descriptor() ([]byte, []int) {
	return file_cloud_proto_rawDescGZIP(), []int{0}
}

func (x *PodInfo) GetPodId() string {
	if x != nil {
		return x.PodId
	}
	return ""
}

func (x *PodInfo) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *PodInfo) GetAccessUrl() string {
	if x != nil {
		return x.AccessUrl
	}
	return ""
}

type PodInfoResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PodInfoResult) Reset() {
	*x = PodInfoResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PodInfoResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PodInfoResult) ProtoMessage() {}

func (x *PodInfoResult) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PodInfoResult.ProtoReflect.Descriptor instead.
func (*PodInfoResult) Descriptor() ([]byte, []int) {
	return file_cloud_proto_rawDescGZIP(), []int{1}
}

type ReleasePodMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PodId string `protobuf:"bytes,1,opt,name=pod_id,json=podId,proto3" json:"pod_id,omitempty"`
}

func (x *ReleasePodMessage) Reset() {
	*x = ReleasePodMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReleasePodMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReleasePodMessage) ProtoMessage() {}

func (x *ReleasePodMessage) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReleasePodMessage.ProtoReflect.Descriptor instead.
func (*ReleasePodMessage) Descriptor() ([]byte, []int) {
	return file_cloud_proto_rawDescGZIP(), []int{2}
}

func (x *ReleasePodMessage) GetPodId() string {
	if x != nil {
		return x.PodId
	}
	return ""
}

type ReleasePodMessageResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ReleasePodMessageResult) Reset() {
	*x = ReleasePodMessageResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReleasePodMessageResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReleasePodMessageResult) ProtoMessage() {}

func (x *ReleasePodMessageResult) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReleasePodMessageResult.ProtoReflect.Descriptor instead.
func (*ReleasePodMessageResult) Descriptor() ([]byte, []int) {
	return file_cloud_proto_rawDescGZIP(), []int{3}
}

var File_cloud_proto protoreflect.FileDescriptor

var file_cloud_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x22, 0x55, 0x0a, 0x07, 0x50, 0x6f, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x15, 0x0a, 0x06, 0x70, 0x6f, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x70, 0x6f, 0x64, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x1d, 0x0a, 0x0a,
	0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x55, 0x72, 0x6c, 0x22, 0x0f, 0x0a, 0x0d, 0x50,
	0x6f, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x2a, 0x0a, 0x11,
	0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x50, 0x6f, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x15, 0x0a, 0x06, 0x70, 0x6f, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x70, 0x6f, 0x64, 0x49, 0x64, 0x22, 0x19, 0x0a, 0x17, 0x52, 0x65, 0x6c, 0x65,
	0x61, 0x73, 0x65, 0x50, 0x6f, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x32, 0x87, 0x01, 0x0a, 0x05, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x12, 0x34, 0x0a,
	0x0a, 0x53, 0x65, 0x74, 0x50, 0x6f, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x50, 0x6f, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x14, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x50, 0x6f, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x22, 0x00, 0x12, 0x48, 0x0a, 0x0a, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x50, 0x6f,
	0x64, 0x12, 0x18, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73,
	0x65, 0x50, 0x6f, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x1e, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x50, 0x6f, 0x64, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x42, 0x0d, 0x5a,
	0x0b, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cloud_proto_rawDescOnce sync.Once
	file_cloud_proto_rawDescData = file_cloud_proto_rawDesc
)

func file_cloud_proto_rawDescGZIP() []byte {
	file_cloud_proto_rawDescOnce.Do(func() {
		file_cloud_proto_rawDescData = protoimpl.X.CompressGZIP(file_cloud_proto_rawDescData)
	})
	return file_cloud_proto_rawDescData
}

var file_cloud_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_cloud_proto_goTypes = []interface{}{
	(*PodInfo)(nil),                 // 0: cloud.PodInfo
	(*PodInfoResult)(nil),           // 1: cloud.PodInfoResult
	(*ReleasePodMessage)(nil),       // 2: cloud.ReleasePodMessage
	(*ReleasePodMessageResult)(nil), // 3: cloud.ReleasePodMessageResult
}
var file_cloud_proto_depIdxs = []int32{
	0, // 0: cloud.Cloud.SetPodInfo:input_type -> cloud.PodInfo
	2, // 1: cloud.Cloud.ReleasePod:input_type -> cloud.ReleasePodMessage
	1, // 2: cloud.Cloud.SetPodInfo:output_type -> cloud.PodInfoResult
	3, // 3: cloud.Cloud.ReleasePod:output_type -> cloud.ReleasePodMessageResult
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_cloud_proto_init() }
func file_cloud_proto_init() {
	if File_cloud_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cloud_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PodInfo); i {
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
		file_cloud_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PodInfoResult); i {
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
		file_cloud_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReleasePodMessage); i {
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
		file_cloud_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReleasePodMessageResult); i {
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
			RawDescriptor: file_cloud_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cloud_proto_goTypes,
		DependencyIndexes: file_cloud_proto_depIdxs,
		MessageInfos:      file_cloud_proto_msgTypes,
	}.Build()
	File_cloud_proto = out.File
	file_cloud_proto_rawDesc = nil
	file_cloud_proto_goTypes = nil
	file_cloud_proto_depIdxs = nil
}
