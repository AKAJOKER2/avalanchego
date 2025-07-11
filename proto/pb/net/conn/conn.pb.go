// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: net/conn/conn.proto

package conn

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// ErrorCode provides information for special sentinel error types
type ErrorCode int32

const (
	ErrorCode_ERROR_CODE_UNSPECIFIED              ErrorCode = 0
	ErrorCode_ERROR_CODE_EOF                      ErrorCode = 1
	ErrorCode_ERROR_CODE_OS_ERR_DEADLINE_EXCEEDED ErrorCode = 2
)

// Enum value maps for ErrorCode.
var (
	ErrorCode_name = map[int32]string{
		0: "ERROR_CODE_UNSPECIFIED",
		1: "ERROR_CODE_EOF",
		2: "ERROR_CODE_OS_ERR_DEADLINE_EXCEEDED",
	}
	ErrorCode_value = map[string]int32{
		"ERROR_CODE_UNSPECIFIED":              0,
		"ERROR_CODE_EOF":                      1,
		"ERROR_CODE_OS_ERR_DEADLINE_EXCEEDED": 2,
	}
)

func (x ErrorCode) Enum() *ErrorCode {
	p := new(ErrorCode)
	*p = x
	return p
}

func (x ErrorCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ErrorCode) Descriptor() protoreflect.EnumDescriptor {
	return file_net_conn_conn_proto_enumTypes[0].Descriptor()
}

func (ErrorCode) Type() protoreflect.EnumType {
	return &file_net_conn_conn_proto_enumTypes[0]
}

func (x ErrorCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ErrorCode.Descriptor instead.
func (ErrorCode) EnumDescriptor() ([]byte, []int) {
	return file_net_conn_conn_proto_rawDescGZIP(), []int{0}
}

type ReadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// length of the request in bytes
	Length int32 `protobuf:"varint,1,opt,name=length,proto3" json:"length,omitempty"`
}

func (x *ReadRequest) Reset() {
	*x = ReadRequest{}
	mi := &file_net_conn_conn_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ReadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadRequest) ProtoMessage() {}

func (x *ReadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_net_conn_conn_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadRequest.ProtoReflect.Descriptor instead.
func (*ReadRequest) Descriptor() ([]byte, []int) {
	return file_net_conn_conn_proto_rawDescGZIP(), []int{0}
}

func (x *ReadRequest) GetLength() int32 {
	if x != nil {
		return x.Length
	}
	return 0
}

type ReadResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// read is the payload in bytes
	Read []byte `protobuf:"bytes,1,opt,name=read,proto3" json:"read,omitempty"`
	// error is an error message
	Error *Error `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *ReadResponse) Reset() {
	*x = ReadResponse{}
	mi := &file_net_conn_conn_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ReadResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadResponse) ProtoMessage() {}

func (x *ReadResponse) ProtoReflect() protoreflect.Message {
	mi := &file_net_conn_conn_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadResponse.ProtoReflect.Descriptor instead.
func (*ReadResponse) Descriptor() ([]byte, []int) {
	return file_net_conn_conn_proto_rawDescGZIP(), []int{1}
}

func (x *ReadResponse) GetRead() []byte {
	if x != nil {
		return x.Read
	}
	return nil
}

func (x *ReadResponse) GetError() *Error {
	if x != nil {
		return x.Error
	}
	return nil
}

type WriteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// payload is the write request in bytes
	Payload []byte `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *WriteRequest) Reset() {
	*x = WriteRequest{}
	mi := &file_net_conn_conn_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WriteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WriteRequest) ProtoMessage() {}

func (x *WriteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_net_conn_conn_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WriteRequest.ProtoReflect.Descriptor instead.
func (*WriteRequest) Descriptor() ([]byte, []int) {
	return file_net_conn_conn_proto_rawDescGZIP(), []int{2}
}

func (x *WriteRequest) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

type WriteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// length of the response in bytes
	Length int32 `protobuf:"varint,1,opt,name=length,proto3" json:"length,omitempty"`
	// error is an error message
	Error *string `protobuf:"bytes,2,opt,name=error,proto3,oneof" json:"error,omitempty"`
}

func (x *WriteResponse) Reset() {
	*x = WriteResponse{}
	mi := &file_net_conn_conn_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WriteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WriteResponse) ProtoMessage() {}

func (x *WriteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_net_conn_conn_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WriteResponse.ProtoReflect.Descriptor instead.
func (*WriteResponse) Descriptor() ([]byte, []int) {
	return file_net_conn_conn_proto_rawDescGZIP(), []int{3}
}

func (x *WriteResponse) GetLength() int32 {
	if x != nil {
		return x.Length
	}
	return 0
}

func (x *WriteResponse) GetError() string {
	if x != nil && x.Error != nil {
		return *x.Error
	}
	return ""
}

type SetDeadlineRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// time represents an instant in time in bytes
	Time []byte `protobuf:"bytes,1,opt,name=time,proto3" json:"time,omitempty"`
}

func (x *SetDeadlineRequest) Reset() {
	*x = SetDeadlineRequest{}
	mi := &file_net_conn_conn_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SetDeadlineRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetDeadlineRequest) ProtoMessage() {}

func (x *SetDeadlineRequest) ProtoReflect() protoreflect.Message {
	mi := &file_net_conn_conn_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetDeadlineRequest.ProtoReflect.Descriptor instead.
func (*SetDeadlineRequest) Descriptor() ([]byte, []int) {
	return file_net_conn_conn_proto_rawDescGZIP(), []int{4}
}

func (x *SetDeadlineRequest) GetTime() []byte {
	if x != nil {
		return x.Time
	}
	return nil
}

type Error struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ErrorCode ErrorCode `protobuf:"varint,1,opt,name=error_code,json=errorCode,proto3,enum=net.conn.ErrorCode" json:"error_code,omitempty"`
	Message   string    `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Error) Reset() {
	*x = Error{}
	mi := &file_net_conn_conn_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Error) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Error) ProtoMessage() {}

func (x *Error) ProtoReflect() protoreflect.Message {
	mi := &file_net_conn_conn_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Error.ProtoReflect.Descriptor instead.
func (*Error) Descriptor() ([]byte, []int) {
	return file_net_conn_conn_proto_rawDescGZIP(), []int{5}
}

func (x *Error) GetErrorCode() ErrorCode {
	if x != nil {
		return x.ErrorCode
	}
	return ErrorCode_ERROR_CODE_UNSPECIFIED
}

func (x *Error) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_net_conn_conn_proto protoreflect.FileDescriptor

var file_net_conn_conn_proto_rawDesc = []byte{
	0x0a, 0x13, 0x6e, 0x65, 0x74, 0x2f, 0x63, 0x6f, 0x6e, 0x6e, 0x2f, 0x63, 0x6f, 0x6e, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x6e, 0x65, 0x74, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x1a,
	0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x25, 0x0a, 0x0b,
	0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6c,
	0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6c, 0x65, 0x6e,
	0x67, 0x74, 0x68, 0x22, 0x49, 0x0a, 0x0c, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x65, 0x61, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x04, 0x72, 0x65, 0x61, 0x64, 0x12, 0x25, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x6e, 0x65, 0x74, 0x2e, 0x63, 0x6f, 0x6e,
	0x6e, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x28,
	0x0a, 0x0c, 0x57, 0x72, 0x69, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18,
	0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x4c, 0x0a, 0x0d, 0x57, 0x72, 0x69, 0x74,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x65, 0x6e,
	0x67, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6c, 0x65, 0x6e, 0x67, 0x74,
	0x68, 0x12, 0x19, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x00, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x88, 0x01, 0x01, 0x42, 0x08, 0x0a, 0x06,
	0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x28, 0x0a, 0x12, 0x53, 0x65, 0x74, 0x44, 0x65, 0x61,
	0x64, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65,
	0x22, 0x55, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x32, 0x0a, 0x0a, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e,
	0x6e, 0x65, 0x74, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f,
	0x64, 0x65, 0x52, 0x09, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2a, 0x64, 0x0a, 0x09, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x43, 0x6f, 0x64, 0x65, 0x12, 0x1a, 0x0a, 0x16, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x43, 0x4f,
	0x44, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00,
	0x12, 0x12, 0x0a, 0x0e, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x45,
	0x4f, 0x46, 0x10, 0x01, 0x12, 0x27, 0x0a, 0x23, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x43, 0x4f,
	0x44, 0x45, 0x5f, 0x4f, 0x53, 0x5f, 0x45, 0x52, 0x52, 0x5f, 0x44, 0x45, 0x41, 0x44, 0x4c, 0x49,
	0x4e, 0x45, 0x5f, 0x45, 0x58, 0x43, 0x45, 0x45, 0x44, 0x45, 0x44, 0x10, 0x02, 0x32, 0x88, 0x03,
	0x0a, 0x04, 0x43, 0x6f, 0x6e, 0x6e, 0x12, 0x35, 0x0a, 0x04, 0x52, 0x65, 0x61, 0x64, 0x12, 0x15,
	0x2e, 0x6e, 0x65, 0x74, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x6e, 0x65, 0x74, 0x2e, 0x63, 0x6f, 0x6e, 0x6e,
	0x2e, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a,
	0x05, 0x57, 0x72, 0x69, 0x74, 0x65, 0x12, 0x16, 0x2e, 0x6e, 0x65, 0x74, 0x2e, 0x63, 0x6f, 0x6e,
	0x6e, 0x2e, 0x57, 0x72, 0x69, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17,
	0x2e, 0x6e, 0x65, 0x74, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x2e, 0x57, 0x72, 0x69, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x05, 0x43, 0x6c, 0x6f, 0x73, 0x65,
	0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x12, 0x43, 0x0a, 0x0b, 0x53, 0x65, 0x74, 0x44, 0x65, 0x61, 0x64, 0x6c, 0x69, 0x6e, 0x65, 0x12,
	0x1c, 0x2e, 0x6e, 0x65, 0x74, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x2e, 0x53, 0x65, 0x74, 0x44, 0x65,
	0x61, 0x64, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x47, 0x0a, 0x0f, 0x53, 0x65, 0x74, 0x52, 0x65, 0x61, 0x64,
	0x44, 0x65, 0x61, 0x64, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x1c, 0x2e, 0x6e, 0x65, 0x74, 0x2e, 0x63,
	0x6f, 0x6e, 0x6e, 0x2e, 0x53, 0x65, 0x74, 0x44, 0x65, 0x61, 0x64, 0x6c, 0x69, 0x6e, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x48,
	0x0a, 0x10, 0x53, 0x65, 0x74, 0x57, 0x72, 0x69, 0x74, 0x65, 0x44, 0x65, 0x61, 0x64, 0x6c, 0x69,
	0x6e, 0x65, 0x12, 0x1c, 0x2e, 0x6e, 0x65, 0x74, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x2e, 0x53, 0x65,
	0x74, 0x44, 0x65, 0x61, 0x64, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x33, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x76, 0x61, 0x2d, 0x6c, 0x61, 0x62, 0x73, 0x2f,
	0x61, 0x76, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x68, 0x65, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x70, 0x62, 0x2f, 0x6e, 0x65, 0x74, 0x2f, 0x63, 0x6f, 0x6e, 0x6e, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_net_conn_conn_proto_rawDescOnce sync.Once
	file_net_conn_conn_proto_rawDescData = file_net_conn_conn_proto_rawDesc
)

func file_net_conn_conn_proto_rawDescGZIP() []byte {
	file_net_conn_conn_proto_rawDescOnce.Do(func() {
		file_net_conn_conn_proto_rawDescData = protoimpl.X.CompressGZIP(file_net_conn_conn_proto_rawDescData)
	})
	return file_net_conn_conn_proto_rawDescData
}

var file_net_conn_conn_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_net_conn_conn_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_net_conn_conn_proto_goTypes = []any{
	(ErrorCode)(0),             // 0: net.conn.ErrorCode
	(*ReadRequest)(nil),        // 1: net.conn.ReadRequest
	(*ReadResponse)(nil),       // 2: net.conn.ReadResponse
	(*WriteRequest)(nil),       // 3: net.conn.WriteRequest
	(*WriteResponse)(nil),      // 4: net.conn.WriteResponse
	(*SetDeadlineRequest)(nil), // 5: net.conn.SetDeadlineRequest
	(*Error)(nil),              // 6: net.conn.Error
	(*emptypb.Empty)(nil),      // 7: google.protobuf.Empty
}
var file_net_conn_conn_proto_depIdxs = []int32{
	6, // 0: net.conn.ReadResponse.error:type_name -> net.conn.Error
	0, // 1: net.conn.Error.error_code:type_name -> net.conn.ErrorCode
	1, // 2: net.conn.Conn.Read:input_type -> net.conn.ReadRequest
	3, // 3: net.conn.Conn.Write:input_type -> net.conn.WriteRequest
	7, // 4: net.conn.Conn.Close:input_type -> google.protobuf.Empty
	5, // 5: net.conn.Conn.SetDeadline:input_type -> net.conn.SetDeadlineRequest
	5, // 6: net.conn.Conn.SetReadDeadline:input_type -> net.conn.SetDeadlineRequest
	5, // 7: net.conn.Conn.SetWriteDeadline:input_type -> net.conn.SetDeadlineRequest
	2, // 8: net.conn.Conn.Read:output_type -> net.conn.ReadResponse
	4, // 9: net.conn.Conn.Write:output_type -> net.conn.WriteResponse
	7, // 10: net.conn.Conn.Close:output_type -> google.protobuf.Empty
	7, // 11: net.conn.Conn.SetDeadline:output_type -> google.protobuf.Empty
	7, // 12: net.conn.Conn.SetReadDeadline:output_type -> google.protobuf.Empty
	7, // 13: net.conn.Conn.SetWriteDeadline:output_type -> google.protobuf.Empty
	8, // [8:14] is the sub-list for method output_type
	2, // [2:8] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_net_conn_conn_proto_init() }
func file_net_conn_conn_proto_init() {
	if File_net_conn_conn_proto != nil {
		return
	}
	file_net_conn_conn_proto_msgTypes[3].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_net_conn_conn_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_net_conn_conn_proto_goTypes,
		DependencyIndexes: file_net_conn_conn_proto_depIdxs,
		EnumInfos:         file_net_conn_conn_proto_enumTypes,
		MessageInfos:      file_net_conn_conn_proto_msgTypes,
	}.Build()
	File_net_conn_conn_proto = out.File
	file_net_conn_conn_proto_rawDesc = nil
	file_net_conn_conn_proto_goTypes = nil
	file_net_conn_conn_proto_depIdxs = nil
}
