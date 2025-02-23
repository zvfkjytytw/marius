// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        v3.21.12
// source: mulus.proto

package api

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

type PingRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Hello         string                 `protobuf:"bytes,1,opt,name=hello,proto3" json:"hello,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PingRequest) Reset() {
	*x = PingRequest{}
	mi := &file_mulus_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingRequest) ProtoMessage() {}

func (x *PingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mulus_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingRequest.ProtoReflect.Descriptor instead.
func (*PingRequest) Descriptor() ([]byte, []int) {
	return file_mulus_proto_rawDescGZIP(), []int{0}
}

func (x *PingRequest) GetHello() string {
	if x != nil {
		return x.Hello
	}
	return ""
}

type PingResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Hello         string                 `protobuf:"bytes,1,opt,name=hello,proto3" json:"hello,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PingResponse) Reset() {
	*x = PingResponse{}
	mi := &file_mulus_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingResponse) ProtoMessage() {}

func (x *PingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mulus_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingResponse.ProtoReflect.Descriptor instead.
func (*PingResponse) Descriptor() ([]byte, []int) {
	return file_mulus_proto_rawDescGZIP(), []int{1}
}

func (x *PingResponse) GetHello() string {
	if x != nil {
		return x.Hello
	}
	return ""
}

type SaveRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Data          []byte                 `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SaveRequest) Reset() {
	*x = SaveRequest{}
	mi := &file_mulus_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SaveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveRequest) ProtoMessage() {}

func (x *SaveRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mulus_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveRequest.ProtoReflect.Descriptor instead.
func (*SaveRequest) Descriptor() ([]byte, []int) {
	return file_mulus_proto_rawDescGZIP(), []int{2}
}

func (x *SaveRequest) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *SaveRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type SaveResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SaveResponse) Reset() {
	*x = SaveResponse{}
	mi := &file_mulus_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SaveResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveResponse) ProtoMessage() {}

func (x *SaveResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mulus_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveResponse.ProtoReflect.Descriptor instead.
func (*SaveResponse) Descriptor() ([]byte, []int) {
	return file_mulus_proto_rawDescGZIP(), []int{3}
}

func (x *SaveResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GetRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetRequest) Reset() {
	*x = GetRequest{}
	mi := &file_mulus_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRequest) ProtoMessage() {}

func (x *GetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mulus_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRequest.ProtoReflect.Descriptor instead.
func (*GetRequest) Descriptor() ([]byte, []int) {
	return file_mulus_proto_rawDescGZIP(), []int{4}
}

func (x *GetRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GetResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Data          []byte                 `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetResponse) Reset() {
	*x = GetResponse{}
	mi := &file_mulus_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResponse) ProtoMessage() {}

func (x *GetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mulus_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetResponse.ProtoReflect.Descriptor instead.
func (*GetResponse) Descriptor() ([]byte, []int) {
	return file_mulus_proto_rawDescGZIP(), []int{5}
}

func (x *GetResponse) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *GetResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type DeleteRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteRequest) Reset() {
	*x = DeleteRequest{}
	mi := &file_mulus_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRequest) ProtoMessage() {}

func (x *DeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mulus_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRequest.ProtoReflect.Descriptor instead.
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return file_mulus_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type DeleteResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteResponse) Reset() {
	*x = DeleteResponse{}
	mi := &file_mulus_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteResponse) ProtoMessage() {}

func (x *DeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mulus_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteResponse.ProtoReflect.Descriptor instead.
func (*DeleteResponse) Descriptor() ([]byte, []int) {
	return file_mulus_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_mulus_proto protoreflect.FileDescriptor

var file_mulus_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x6d, 0x75, 0x6c, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x6d,
	0x61, 0x72, 0x69, 0x75, 0x73, 0x2e, 0x6d, 0x75, 0x6c, 0x75, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x31, 0x22, 0x23, 0x0a, 0x0b, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x22, 0x24, 0x0a, 0x0c, 0x50, 0x69, 0x6e, 0x67, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x68, 0x65, 0x6c, 0x6c, 0x6f,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x22, 0x35, 0x0a,
	0x0b, 0x53, 0x61, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x22, 0x22, 0x0a, 0x0c, 0x53, 0x61, 0x76, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x20, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x35, 0x0a, 0x0b, 0x47, 0x65,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x22, 0x23, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x24, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x32, 0xd5, 0x02, 0x0a,
	0x08, 0x4d, 0x75, 0x6c, 0x75, 0x73, 0x41, 0x50, 0x49, 0x12, 0x4d, 0x0a, 0x04, 0x50, 0x69, 0x6e,
	0x67, 0x12, 0x20, 0x2e, 0x6d, 0x61, 0x72, 0x69, 0x75, 0x73, 0x2e, 0x6d, 0x75, 0x6c, 0x75, 0x73,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x6d, 0x61, 0x72, 0x69, 0x75, 0x73, 0x2e, 0x6d, 0x75, 0x6c,
	0x75, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x51, 0x0a, 0x08, 0x53, 0x61, 0x76, 0x65,
	0x44, 0x61, 0x74, 0x61, 0x12, 0x20, 0x2e, 0x6d, 0x61, 0x72, 0x69, 0x75, 0x73, 0x2e, 0x6d, 0x75,
	0x6c, 0x75, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x6d, 0x61, 0x72, 0x69, 0x75, 0x73, 0x2e,
	0x6d, 0x75, 0x6c, 0x75, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x61, 0x76,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4e, 0x0a, 0x07, 0x47,
	0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1f, 0x2e, 0x6d, 0x61, 0x72, 0x69, 0x75, 0x73, 0x2e,
	0x6d, 0x75, 0x6c, 0x75, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x6d, 0x61, 0x72, 0x69, 0x75, 0x73,
	0x2e, 0x6d, 0x75, 0x6c, 0x75, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x57, 0x0a, 0x0a, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x22, 0x2e, 0x6d, 0x61, 0x72, 0x69,
	0x75, 0x73, 0x2e, 0x6d, 0x75, 0x6c, 0x75, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e,
	0x6d, 0x61, 0x72, 0x69, 0x75, 0x73, 0x2e, 0x6d, 0x75, 0x6c, 0x75, 0x73, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x42, 0x35, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x7a, 0x76, 0x66, 0x6b, 0x6a, 0x79, 0x74, 0x79, 0x74, 0x77, 0x2f, 0x6d, 0x61,
	0x72, 0x69, 0x75, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x75, 0x6c, 0x75, 0x73,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x3b, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_mulus_proto_rawDescOnce sync.Once
	file_mulus_proto_rawDescData = file_mulus_proto_rawDesc
)

func file_mulus_proto_rawDescGZIP() []byte {
	file_mulus_proto_rawDescOnce.Do(func() {
		file_mulus_proto_rawDescData = protoimpl.X.CompressGZIP(file_mulus_proto_rawDescData)
	})
	return file_mulus_proto_rawDescData
}

var file_mulus_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_mulus_proto_goTypes = []any{
	(*PingRequest)(nil),    // 0: marius.mulus.api.v1.PingRequest
	(*PingResponse)(nil),   // 1: marius.mulus.api.v1.PingResponse
	(*SaveRequest)(nil),    // 2: marius.mulus.api.v1.SaveRequest
	(*SaveResponse)(nil),   // 3: marius.mulus.api.v1.SaveResponse
	(*GetRequest)(nil),     // 4: marius.mulus.api.v1.GetRequest
	(*GetResponse)(nil),    // 5: marius.mulus.api.v1.GetResponse
	(*DeleteRequest)(nil),  // 6: marius.mulus.api.v1.DeleteRequest
	(*DeleteResponse)(nil), // 7: marius.mulus.api.v1.DeleteResponse
}
var file_mulus_proto_depIdxs = []int32{
	0, // 0: marius.mulus.api.v1.MulusAPI.Ping:input_type -> marius.mulus.api.v1.PingRequest
	2, // 1: marius.mulus.api.v1.MulusAPI.SaveData:input_type -> marius.mulus.api.v1.SaveRequest
	4, // 2: marius.mulus.api.v1.MulusAPI.GetData:input_type -> marius.mulus.api.v1.GetRequest
	6, // 3: marius.mulus.api.v1.MulusAPI.DeleteData:input_type -> marius.mulus.api.v1.DeleteRequest
	1, // 4: marius.mulus.api.v1.MulusAPI.Ping:output_type -> marius.mulus.api.v1.PingResponse
	3, // 5: marius.mulus.api.v1.MulusAPI.SaveData:output_type -> marius.mulus.api.v1.SaveResponse
	5, // 6: marius.mulus.api.v1.MulusAPI.GetData:output_type -> marius.mulus.api.v1.GetResponse
	7, // 7: marius.mulus.api.v1.MulusAPI.DeleteData:output_type -> marius.mulus.api.v1.DeleteResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_mulus_proto_init() }
func file_mulus_proto_init() {
	if File_mulus_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_mulus_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_mulus_proto_goTypes,
		DependencyIndexes: file_mulus_proto_depIdxs,
		MessageInfos:      file_mulus_proto_msgTypes,
	}.Build()
	File_mulus_proto = out.File
	file_mulus_proto_rawDesc = nil
	file_mulus_proto_goTypes = nil
	file_mulus_proto_depIdxs = nil
}
