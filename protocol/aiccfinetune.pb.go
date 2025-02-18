// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.3
// source: aiccfinetune.proto

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

type AICCFinetuneInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	User          string `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	Status        string `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	LogPath       string `protobuf:"bytes,4,opt,name=log_path,json=logPath,proto3" json:"log_path,omitempty"`
	OutputZipPath string `protobuf:"bytes,5,opt,name=output_zip_path,json=outputZipPath,proto3" json:"output_zip_path,omitempty"`
	Duration      int32  `protobuf:"varint,6,opt,name=duration,proto3" json:"duration,omitempty"`
	Model         string `protobuf:"bytes,7,opt,name=model,proto3" json:"model,omitempty"`
}

func (x *AICCFinetuneInfo) Reset() {
	*x = AICCFinetuneInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aiccfinetune_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AICCFinetuneInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AICCFinetuneInfo) ProtoMessage() {}

func (x *AICCFinetuneInfo) ProtoReflect() protoreflect.Message {
	mi := &file_aiccfinetune_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AICCFinetuneInfo.ProtoReflect.Descriptor instead.
func (*AICCFinetuneInfo) Descriptor() ([]byte, []int) {
	return file_aiccfinetune_proto_rawDescGZIP(), []int{0}
}

func (x *AICCFinetuneInfo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *AICCFinetuneInfo) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *AICCFinetuneInfo) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *AICCFinetuneInfo) GetLogPath() string {
	if x != nil {
		return x.LogPath
	}
	return ""
}

func (x *AICCFinetuneInfo) GetOutputZipPath() string {
	if x != nil {
		return x.OutputZipPath
	}
	return ""
}

func (x *AICCFinetuneInfo) GetDuration() int32 {
	if x != nil {
		return x.Duration
	}
	return 0
}

func (x *AICCFinetuneInfo) GetModel() string {
	if x != nil {
		return x.Model
	}
	return ""
}

type AICCFinetuneResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AICCFinetuneResult) Reset() {
	*x = AICCFinetuneResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aiccfinetune_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AICCFinetuneResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AICCFinetuneResult) ProtoMessage() {}

func (x *AICCFinetuneResult) ProtoReflect() protoreflect.Message {
	mi := &file_aiccfinetune_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AICCFinetuneResult.ProtoReflect.Descriptor instead.
func (*AICCFinetuneResult) Descriptor() ([]byte, []int) {
	return file_aiccfinetune_proto_rawDescGZIP(), []int{1}
}

var File_aiccfinetune_proto protoreflect.FileDescriptor

var file_aiccfinetune_proto_rawDesc = []byte{
	0x0a, 0x12, 0x61, 0x69, 0x63, 0x63, 0x66, 0x69, 0x6e, 0x65, 0x74, 0x75, 0x6e, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x61, 0x69, 0x63, 0x63, 0x66, 0x69, 0x6e, 0x65, 0x74, 0x75,
	0x6e, 0x65, 0x22, 0xc3, 0x01, 0x0a, 0x10, 0x41, 0x49, 0x43, 0x43, 0x46, 0x69, 0x6e, 0x65, 0x74,
	0x75, 0x6e, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x6c, 0x6f, 0x67, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6c, 0x6f, 0x67, 0x50, 0x61, 0x74, 0x68, 0x12, 0x26,
	0x0a, 0x0f, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x5f, 0x7a, 0x69, 0x70, 0x5f, 0x70, 0x61, 0x74,
	0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x5a,
	0x69, 0x70, 0x50, 0x61, 0x74, 0x68, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x22, 0x14, 0x0a, 0x12, 0x41, 0x49, 0x43, 0x43,
	0x46, 0x69, 0x6e, 0x65, 0x74, 0x75, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x32, 0x69,
	0x0a, 0x0c, 0x41, 0x49, 0x43, 0x43, 0x46, 0x69, 0x6e, 0x65, 0x74, 0x75, 0x6e, 0x65, 0x12, 0x59,
	0x0a, 0x13, 0x53, 0x65, 0x74, 0x41, 0x49, 0x43, 0x43, 0x46, 0x69, 0x6e, 0x65, 0x74, 0x75, 0x6e,
	0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1e, 0x2e, 0x61, 0x69, 0x63, 0x63, 0x66, 0x69, 0x6e, 0x65,
	0x74, 0x75, 0x6e, 0x65, 0x2e, 0x41, 0x49, 0x43, 0x43, 0x46, 0x69, 0x6e, 0x65, 0x74, 0x75, 0x6e,
	0x65, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x20, 0x2e, 0x61, 0x69, 0x63, 0x63, 0x66, 0x69, 0x6e, 0x65,
	0x74, 0x75, 0x6e, 0x65, 0x2e, 0x41, 0x49, 0x43, 0x43, 0x46, 0x69, 0x6e, 0x65, 0x74, 0x75, 0x6e,
	0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x42, 0x0d, 0x5a, 0x0b, 0x2e, 0x2e, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_aiccfinetune_proto_rawDescOnce sync.Once
	file_aiccfinetune_proto_rawDescData = file_aiccfinetune_proto_rawDesc
)

func file_aiccfinetune_proto_rawDescGZIP() []byte {
	file_aiccfinetune_proto_rawDescOnce.Do(func() {
		file_aiccfinetune_proto_rawDescData = protoimpl.X.CompressGZIP(file_aiccfinetune_proto_rawDescData)
	})
	return file_aiccfinetune_proto_rawDescData
}

var file_aiccfinetune_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_aiccfinetune_proto_goTypes = []interface{}{
	(*AICCFinetuneInfo)(nil),   // 0: aiccfinetune.AICCFinetuneInfo
	(*AICCFinetuneResult)(nil), // 1: aiccfinetune.AICCFinetuneResult
}
var file_aiccfinetune_proto_depIdxs = []int32{
	0, // 0: aiccfinetune.AICCFinetune.SetAICCFinetuneInfo:input_type -> aiccfinetune.AICCFinetuneInfo
	1, // 1: aiccfinetune.AICCFinetune.SetAICCFinetuneInfo:output_type -> aiccfinetune.AICCFinetuneResult
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_aiccfinetune_proto_init() }
func file_aiccfinetune_proto_init() {
	if File_aiccfinetune_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_aiccfinetune_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AICCFinetuneInfo); i {
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
		file_aiccfinetune_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AICCFinetuneResult); i {
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
			RawDescriptor: file_aiccfinetune_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_aiccfinetune_proto_goTypes,
		DependencyIndexes: file_aiccfinetune_proto_depIdxs,
		MessageInfos:      file_aiccfinetune_proto_msgTypes,
	}.Build()
	File_aiccfinetune_proto = out.File
	file_aiccfinetune_proto_rawDesc = nil
	file_aiccfinetune_proto_goTypes = nil
	file_aiccfinetune_proto_depIdxs = nil
}
