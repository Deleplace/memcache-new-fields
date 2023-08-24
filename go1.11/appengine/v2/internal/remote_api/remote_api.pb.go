// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.12
// source: remote_api.proto

package remote_api

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

type RpcError_ErrorCode int32

const (
	RpcError_UNKNOWN             RpcError_ErrorCode = 0
	RpcError_CALL_NOT_FOUND      RpcError_ErrorCode = 1
	RpcError_PARSE_ERROR         RpcError_ErrorCode = 2
	RpcError_SECURITY_VIOLATION  RpcError_ErrorCode = 3
	RpcError_OVER_QUOTA          RpcError_ErrorCode = 4
	RpcError_REQUEST_TOO_LARGE   RpcError_ErrorCode = 5
	RpcError_CAPABILITY_DISABLED RpcError_ErrorCode = 6
	RpcError_FEATURE_DISABLED    RpcError_ErrorCode = 7
	RpcError_BAD_REQUEST         RpcError_ErrorCode = 8
	RpcError_RESPONSE_TOO_LARGE  RpcError_ErrorCode = 9
	RpcError_CANCELLED           RpcError_ErrorCode = 10
	RpcError_REPLAY_ERROR        RpcError_ErrorCode = 11
	RpcError_DEADLINE_EXCEEDED   RpcError_ErrorCode = 12
)

// Enum value maps for RpcError_ErrorCode.
var (
	RpcError_ErrorCode_name = map[int32]string{
		0:  "UNKNOWN",
		1:  "CALL_NOT_FOUND",
		2:  "PARSE_ERROR",
		3:  "SECURITY_VIOLATION",
		4:  "OVER_QUOTA",
		5:  "REQUEST_TOO_LARGE",
		6:  "CAPABILITY_DISABLED",
		7:  "FEATURE_DISABLED",
		8:  "BAD_REQUEST",
		9:  "RESPONSE_TOO_LARGE",
		10: "CANCELLED",
		11: "REPLAY_ERROR",
		12: "DEADLINE_EXCEEDED",
	}
	RpcError_ErrorCode_value = map[string]int32{
		"UNKNOWN":             0,
		"CALL_NOT_FOUND":      1,
		"PARSE_ERROR":         2,
		"SECURITY_VIOLATION":  3,
		"OVER_QUOTA":          4,
		"REQUEST_TOO_LARGE":   5,
		"CAPABILITY_DISABLED": 6,
		"FEATURE_DISABLED":    7,
		"BAD_REQUEST":         8,
		"RESPONSE_TOO_LARGE":  9,
		"CANCELLED":           10,
		"REPLAY_ERROR":        11,
		"DEADLINE_EXCEEDED":   12,
	}
)

func (x RpcError_ErrorCode) Enum() *RpcError_ErrorCode {
	p := new(RpcError_ErrorCode)
	*p = x
	return p
}

func (x RpcError_ErrorCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RpcError_ErrorCode) Descriptor() protoreflect.EnumDescriptor {
	return file_remote_api_proto_enumTypes[0].Descriptor()
}

func (RpcError_ErrorCode) Type() protoreflect.EnumType {
	return &file_remote_api_proto_enumTypes[0]
}

func (x RpcError_ErrorCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *RpcError_ErrorCode) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = RpcError_ErrorCode(num)
	return nil
}

// Deprecated: Use RpcError_ErrorCode.Descriptor instead.
func (RpcError_ErrorCode) EnumDescriptor() ([]byte, []int) {
	return file_remote_api_proto_rawDescGZIP(), []int{2, 0}
}

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServiceName *string `protobuf:"bytes,2,req,name=service_name,json=serviceName" json:"service_name,omitempty"`
	Method      *string `protobuf:"bytes,3,req,name=method" json:"method,omitempty"`
	Request     []byte  `protobuf:"bytes,4,req,name=request" json:"request,omitempty"`
	RequestId   *string `protobuf:"bytes,5,opt,name=request_id,json=requestId" json:"request_id,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_remote_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_remote_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_remote_api_proto_rawDescGZIP(), []int{0}
}

func (x *Request) GetServiceName() string {
	if x != nil && x.ServiceName != nil {
		return *x.ServiceName
	}
	return ""
}

func (x *Request) GetMethod() string {
	if x != nil && x.Method != nil {
		return *x.Method
	}
	return ""
}

func (x *Request) GetRequest() []byte {
	if x != nil {
		return x.Request
	}
	return nil
}

func (x *Request) GetRequestId() string {
	if x != nil && x.RequestId != nil {
		return *x.RequestId
	}
	return ""
}

type ApplicationError struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code   *int32  `protobuf:"varint,1,req,name=code" json:"code,omitempty"`
	Detail *string `protobuf:"bytes,2,req,name=detail" json:"detail,omitempty"`
}

func (x *ApplicationError) Reset() {
	*x = ApplicationError{}
	if protoimpl.UnsafeEnabled {
		mi := &file_remote_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApplicationError) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApplicationError) ProtoMessage() {}

func (x *ApplicationError) ProtoReflect() protoreflect.Message {
	mi := &file_remote_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApplicationError.ProtoReflect.Descriptor instead.
func (*ApplicationError) Descriptor() ([]byte, []int) {
	return file_remote_api_proto_rawDescGZIP(), []int{1}
}

func (x *ApplicationError) GetCode() int32 {
	if x != nil && x.Code != nil {
		return *x.Code
	}
	return 0
}

func (x *ApplicationError) GetDetail() string {
	if x != nil && x.Detail != nil {
		return *x.Detail
	}
	return ""
}

type RpcError struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code   *int32  `protobuf:"varint,1,req,name=code" json:"code,omitempty"`
	Detail *string `protobuf:"bytes,2,opt,name=detail" json:"detail,omitempty"`
}

func (x *RpcError) Reset() {
	*x = RpcError{}
	if protoimpl.UnsafeEnabled {
		mi := &file_remote_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RpcError) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RpcError) ProtoMessage() {}

func (x *RpcError) ProtoReflect() protoreflect.Message {
	mi := &file_remote_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RpcError.ProtoReflect.Descriptor instead.
func (*RpcError) Descriptor() ([]byte, []int) {
	return file_remote_api_proto_rawDescGZIP(), []int{2}
}

func (x *RpcError) GetCode() int32 {
	if x != nil && x.Code != nil {
		return *x.Code
	}
	return 0
}

func (x *RpcError) GetDetail() string {
	if x != nil && x.Detail != nil {
		return *x.Detail
	}
	return ""
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response         []byte            `protobuf:"bytes,1,opt,name=response" json:"response,omitempty"`
	Exception        []byte            `protobuf:"bytes,2,opt,name=exception" json:"exception,omitempty"`
	ApplicationError *ApplicationError `protobuf:"bytes,3,opt,name=application_error,json=applicationError" json:"application_error,omitempty"`
	JavaException    []byte            `protobuf:"bytes,4,opt,name=java_exception,json=javaException" json:"java_exception,omitempty"`
	RpcError         *RpcError         `protobuf:"bytes,5,opt,name=rpc_error,json=rpcError" json:"rpc_error,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_remote_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_remote_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_remote_api_proto_rawDescGZIP(), []int{3}
}

func (x *Response) GetResponse() []byte {
	if x != nil {
		return x.Response
	}
	return nil
}

func (x *Response) GetException() []byte {
	if x != nil {
		return x.Exception
	}
	return nil
}

func (x *Response) GetApplicationError() *ApplicationError {
	if x != nil {
		return x.ApplicationError
	}
	return nil
}

func (x *Response) GetJavaException() []byte {
	if x != nil {
		return x.JavaException
	}
	return nil
}

func (x *Response) GetRpcError() *RpcError {
	if x != nil {
		return x.RpcError
	}
	return nil
}

var File_remote_api_proto protoreflect.FileDescriptor

var file_remote_api_proto_rawDesc = []byte{
	0x0a, 0x10, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0d, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x32, 0x22, 0x7d, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x02,
	0x28, 0x09, 0x52, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x03, 0x20, 0x02, 0x28, 0x09, 0x52,
	0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x18, 0x04, 0x20, 0x02, 0x28, 0x0c, 0x52, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64,
	0x22, 0x3e, 0x0a, 0x10, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x02,
	0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x18, 0x02, 0x20, 0x02, 0x28, 0x09, 0x52, 0x06, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x22, 0xc5, 0x02, 0x0a, 0x08, 0x52, 0x70, 0x63, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x12, 0x0a,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x02, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x22, 0x8c, 0x02, 0x0a, 0x09, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f,
	0x57, 0x4e, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x43, 0x41, 0x4c, 0x4c, 0x5f, 0x4e, 0x4f, 0x54,
	0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x50, 0x41, 0x52, 0x53,
	0x45, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x02, 0x12, 0x16, 0x0a, 0x12, 0x53, 0x45, 0x43,
	0x55, 0x52, 0x49, 0x54, 0x59, 0x5f, 0x56, 0x49, 0x4f, 0x4c, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10,
	0x03, 0x12, 0x0e, 0x0a, 0x0a, 0x4f, 0x56, 0x45, 0x52, 0x5f, 0x51, 0x55, 0x4f, 0x54, 0x41, 0x10,
	0x04, 0x12, 0x15, 0x0a, 0x11, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x5f, 0x54, 0x4f, 0x4f,
	0x5f, 0x4c, 0x41, 0x52, 0x47, 0x45, 0x10, 0x05, 0x12, 0x17, 0x0a, 0x13, 0x43, 0x41, 0x50, 0x41,
	0x42, 0x49, 0x4c, 0x49, 0x54, 0x59, 0x5f, 0x44, 0x49, 0x53, 0x41, 0x42, 0x4c, 0x45, 0x44, 0x10,
	0x06, 0x12, 0x14, 0x0a, 0x10, 0x46, 0x45, 0x41, 0x54, 0x55, 0x52, 0x45, 0x5f, 0x44, 0x49, 0x53,
	0x41, 0x42, 0x4c, 0x45, 0x44, 0x10, 0x07, 0x12, 0x0f, 0x0a, 0x0b, 0x42, 0x41, 0x44, 0x5f, 0x52,
	0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x10, 0x08, 0x12, 0x16, 0x0a, 0x12, 0x52, 0x45, 0x53, 0x50,
	0x4f, 0x4e, 0x53, 0x45, 0x5f, 0x54, 0x4f, 0x4f, 0x5f, 0x4c, 0x41, 0x52, 0x47, 0x45, 0x10, 0x09,
	0x12, 0x0d, 0x0a, 0x09, 0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c, 0x4c, 0x45, 0x44, 0x10, 0x0a, 0x12,
	0x10, 0x0a, 0x0c, 0x52, 0x45, 0x50, 0x4c, 0x41, 0x59, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10,
	0x0b, 0x12, 0x15, 0x0a, 0x11, 0x44, 0x45, 0x41, 0x44, 0x4c, 0x49, 0x4e, 0x45, 0x5f, 0x45, 0x58,
	0x43, 0x45, 0x45, 0x44, 0x45, 0x44, 0x10, 0x0c, 0x22, 0xef, 0x01, 0x0a, 0x08, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x1c, 0x0a, 0x09, 0x65, 0x78, 0x63, 0x65, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x65, 0x78, 0x63, 0x65, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x4c, 0x0a, 0x11, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x72, 0x65, 0x6d,
	0x6f, 0x74, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x41, 0x70, 0x70, 0x6c, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x10, 0x61, 0x70, 0x70,
	0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x25, 0x0a,
	0x0e, 0x6a, 0x61, 0x76, 0x61, 0x5f, 0x65, 0x78, 0x63, 0x65, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0d, 0x6a, 0x61, 0x76, 0x61, 0x45, 0x78, 0x63, 0x65, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x34, 0x0a, 0x09, 0x72, 0x70, 0x63, 0x5f, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65,
	0x5f, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x52, 0x70, 0x63, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x52, 0x08, 0x72, 0x70, 0x63, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x42, 0x34, 0x5a, 0x32, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f,
	0x61, 0x70, 0x70, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2f, 0x76, 0x32, 0x2f, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x5f, 0x61, 0x70, 0x69,
}

var (
	file_remote_api_proto_rawDescOnce sync.Once
	file_remote_api_proto_rawDescData = file_remote_api_proto_rawDesc
)

func file_remote_api_proto_rawDescGZIP() []byte {
	file_remote_api_proto_rawDescOnce.Do(func() {
		file_remote_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_remote_api_proto_rawDescData)
	})
	return file_remote_api_proto_rawDescData
}

var file_remote_api_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_remote_api_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_remote_api_proto_goTypes = []interface{}{
	(RpcError_ErrorCode)(0),  // 0: remote_api.v2.RpcError.ErrorCode
	(*Request)(nil),          // 1: remote_api.v2.Request
	(*ApplicationError)(nil), // 2: remote_api.v2.ApplicationError
	(*RpcError)(nil),         // 3: remote_api.v2.RpcError
	(*Response)(nil),         // 4: remote_api.v2.Response
}
var file_remote_api_proto_depIdxs = []int32{
	2, // 0: remote_api.v2.Response.application_error:type_name -> remote_api.v2.ApplicationError
	3, // 1: remote_api.v2.Response.rpc_error:type_name -> remote_api.v2.RpcError
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_remote_api_proto_init() }
func file_remote_api_proto_init() {
	if File_remote_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_remote_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_remote_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApplicationError); i {
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
		file_remote_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RpcError); i {
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
		file_remote_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
			RawDescriptor: file_remote_api_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_remote_api_proto_goTypes,
		DependencyIndexes: file_remote_api_proto_depIdxs,
		EnumInfos:         file_remote_api_proto_enumTypes,
		MessageInfos:      file_remote_api_proto_msgTypes,
	}.Build()
	File_remote_api_proto = out.File
	file_remote_api_proto_rawDesc = nil
	file_remote_api_proto_goTypes = nil
	file_remote_api_proto_depIdxs = nil
}