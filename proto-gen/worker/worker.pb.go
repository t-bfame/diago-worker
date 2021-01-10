// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.12.3
// source: diago-idl/proto/worker.proto

package worker

import (
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Payload:
	//	*Message_Register
	//	*Message_Start
	//	*Message_Metrics
	//	*Message_Finish
	//	*Message_Stop
	//	*Message_Ack
	Payload isMessage_Payload `protobuf_oneof:"payload"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_diago_idl_proto_worker_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_diago_idl_proto_worker_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_diago_idl_proto_worker_proto_rawDescGZIP(), []int{0}
}

func (m *Message) GetPayload() isMessage_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (x *Message) GetRegister() *Register {
	if x, ok := x.GetPayload().(*Message_Register); ok {
		return x.Register
	}
	return nil
}

func (x *Message) GetStart() *Start {
	if x, ok := x.GetPayload().(*Message_Start); ok {
		return x.Start
	}
	return nil
}

func (x *Message) GetMetrics() *Metrics {
	if x, ok := x.GetPayload().(*Message_Metrics); ok {
		return x.Metrics
	}
	return nil
}

func (x *Message) GetFinish() *Finish {
	if x, ok := x.GetPayload().(*Message_Finish); ok {
		return x.Finish
	}
	return nil
}

func (x *Message) GetStop() *Stop {
	if x, ok := x.GetPayload().(*Message_Stop); ok {
		return x.Stop
	}
	return nil
}

func (x *Message) GetAck() *Ack {
	if x, ok := x.GetPayload().(*Message_Ack); ok {
		return x.Ack
	}
	return nil
}

type isMessage_Payload interface {
	isMessage_Payload()
}

type Message_Register struct {
	Register *Register `protobuf:"bytes,1,opt,name=register,proto3,oneof"`
}

type Message_Start struct {
	Start *Start `protobuf:"bytes,2,opt,name=start,proto3,oneof"`
}

type Message_Metrics struct {
	Metrics *Metrics `protobuf:"bytes,3,opt,name=metrics,proto3,oneof"`
}

type Message_Finish struct {
	Finish *Finish `protobuf:"bytes,4,opt,name=finish,proto3,oneof"`
}

type Message_Stop struct {
	Stop *Stop `protobuf:"bytes,5,opt,name=stop,proto3,oneof"`
}

type Message_Ack struct {
	Ack *Ack `protobuf:"bytes,6,opt,name=ack,proto3,oneof"`
}

func (*Message_Register) isMessage_Payload() {}

func (*Message_Start) isMessage_Payload() {}

func (*Message_Metrics) isMessage_Payload() {}

func (*Message_Finish) isMessage_Payload() {}

func (*Message_Stop) isMessage_Payload() {}

func (*Message_Ack) isMessage_Payload() {}

type Register struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Group     string `protobuf:"bytes,1,opt,name=group,proto3" json:"group,omitempty"`
	Instance  string `protobuf:"bytes,2,opt,name=instance,proto3" json:"instance,omitempty"`
	Frequency uint64 `protobuf:"varint,3,opt,name=frequency,proto3" json:"frequency,omitempty"`
}

func (x *Register) Reset() {
	*x = Register{}
	if protoimpl.UnsafeEnabled {
		mi := &file_diago_idl_proto_worker_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Register) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Register) ProtoMessage() {}

func (x *Register) ProtoReflect() protoreflect.Message {
	mi := &file_diago_idl_proto_worker_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Register.ProtoReflect.Descriptor instead.
func (*Register) Descriptor() ([]byte, []int) {
	return file_diago_idl_proto_worker_proto_rawDescGZIP(), []int{1}
}

func (x *Register) GetGroup() string {
	if x != nil {
		return x.Group
	}
	return ""
}

func (x *Register) GetInstance() string {
	if x != nil {
		return x.Instance
	}
	return ""
}

func (x *Register) GetFrequency() uint64 {
	if x != nil {
		return x.Frequency
	}
	return 0
}

type HTTPRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Method string `protobuf:"bytes,1,opt,name=method,proto3" json:"method,omitempty"`
	Url    string `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *HTTPRequest) Reset() {
	*x = HTTPRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_diago_idl_proto_worker_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HTTPRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HTTPRequest) ProtoMessage() {}

func (x *HTTPRequest) ProtoReflect() protoreflect.Message {
	mi := &file_diago_idl_proto_worker_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HTTPRequest.ProtoReflect.Descriptor instead.
func (*HTTPRequest) Descriptor() ([]byte, []int) {
	return file_diago_idl_proto_worker_proto_rawDescGZIP(), []int{2}
}

func (x *HTTPRequest) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *HTTPRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type Start struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Each job is split into multiple workloads, each workload with the same job id
	JobId string `protobuf:"bytes,1,opt,name=job_id,json=jobId,proto3" json:"job_id,omitempty"`
	// requests / second
	Frequency uint64 `protobuf:"varint,2,opt,name=frequency,proto3" json:"frequency,omitempty"`
	// seconds
	Duration uint64       `protobuf:"varint,3,opt,name=duration,proto3" json:"duration,omitempty"`
	Request  *HTTPRequest `protobuf:"bytes,4,opt,name=request,proto3" json:"request,omitempty"`
}

func (x *Start) Reset() {
	*x = Start{}
	if protoimpl.UnsafeEnabled {
		mi := &file_diago_idl_proto_worker_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Start) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Start) ProtoMessage() {}

func (x *Start) ProtoReflect() protoreflect.Message {
	mi := &file_diago_idl_proto_worker_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Start.ProtoReflect.Descriptor instead.
func (*Start) Descriptor() ([]byte, []int) {
	return file_diago_idl_proto_worker_proto_rawDescGZIP(), []int{3}
}

func (x *Start) GetJobId() string {
	if x != nil {
		return x.JobId
	}
	return ""
}

func (x *Start) GetFrequency() uint64 {
	if x != nil {
		return x.Frequency
	}
	return 0
}

func (x *Start) GetDuration() uint64 {
	if x != nil {
		return x.Duration
	}
	return 0
}

func (x *Start) GetRequest() *HTTPRequest {
	if x != nil {
		return x.Request
	}
	return nil
}

type Finish struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JobId string `protobuf:"bytes,1,opt,name=job_id,json=jobId,proto3" json:"job_id,omitempty"`
}

func (x *Finish) Reset() {
	*x = Finish{}
	if protoimpl.UnsafeEnabled {
		mi := &file_diago_idl_proto_worker_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Finish) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Finish) ProtoMessage() {}

func (x *Finish) ProtoReflect() protoreflect.Message {
	mi := &file_diago_idl_proto_worker_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Finish.ProtoReflect.Descriptor instead.
func (*Finish) Descriptor() ([]byte, []int) {
	return file_diago_idl_proto_worker_proto_rawDescGZIP(), []int{4}
}

func (x *Finish) GetJobId() string {
	if x != nil {
		return x.JobId
	}
	return ""
}

type Metrics struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JobId    string `protobuf:"bytes,1,opt,name=job_id,json=jobId,proto3" json:"job_id,omitempty"`
	Code     uint32 `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
	BytesIn  uint64 `protobuf:"varint,3,opt,name=bytes_in,json=bytesIn,proto3" json:"bytes_in,omitempty"`
	BytesOut uint64 `protobuf:"varint,4,opt,name=bytes_out,json=bytesOut,proto3" json:"bytes_out,omitempty"`
	// Nanoseconds
	Latency int64  `protobuf:"varint,5,opt,name=latency,proto3" json:"latency,omitempty"`
	Error   string `protobuf:"bytes,6,opt,name=error,proto3" json:"error,omitempty"`
	// https://godoc.org/github.com/golang/protobuf/ptypes#TimestampProto
	Timestamp *timestamp.Timestamp `protobuf:"bytes,7,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *Metrics) Reset() {
	*x = Metrics{}
	if protoimpl.UnsafeEnabled {
		mi := &file_diago_idl_proto_worker_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Metrics) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Metrics) ProtoMessage() {}

func (x *Metrics) ProtoReflect() protoreflect.Message {
	mi := &file_diago_idl_proto_worker_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Metrics.ProtoReflect.Descriptor instead.
func (*Metrics) Descriptor() ([]byte, []int) {
	return file_diago_idl_proto_worker_proto_rawDescGZIP(), []int{5}
}

func (x *Metrics) GetJobId() string {
	if x != nil {
		return x.JobId
	}
	return ""
}

func (x *Metrics) GetCode() uint32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *Metrics) GetBytesIn() uint64 {
	if x != nil {
		return x.BytesIn
	}
	return 0
}

func (x *Metrics) GetBytesOut() uint64 {
	if x != nil {
		return x.BytesOut
	}
	return 0
}

func (x *Metrics) GetLatency() int64 {
	if x != nil {
		return x.Latency
	}
	return 0
}

func (x *Metrics) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *Metrics) GetTimestamp() *timestamp.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

type Stop struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JobId string `protobuf:"bytes,1,opt,name=job_id,json=jobId,proto3" json:"job_id,omitempty"`
}

func (x *Stop) Reset() {
	*x = Stop{}
	if protoimpl.UnsafeEnabled {
		mi := &file_diago_idl_proto_worker_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Stop) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Stop) ProtoMessage() {}

func (x *Stop) ProtoReflect() protoreflect.Message {
	mi := &file_diago_idl_proto_worker_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Stop.ProtoReflect.Descriptor instead.
func (*Stop) Descriptor() ([]byte, []int) {
	return file_diago_idl_proto_worker_proto_rawDescGZIP(), []int{6}
}

func (x *Stop) GetJobId() string {
	if x != nil {
		return x.JobId
	}
	return ""
}

type Ack struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Ack) Reset() {
	*x = Ack{}
	if protoimpl.UnsafeEnabled {
		mi := &file_diago_idl_proto_worker_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ack) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ack) ProtoMessage() {}

func (x *Ack) ProtoReflect() protoreflect.Message {
	mi := &file_diago_idl_proto_worker_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ack.ProtoReflect.Descriptor instead.
func (*Ack) Descriptor() ([]byte, []int) {
	return file_diago_idl_proto_worker_proto_rawDescGZIP(), []int{7}
}

var File_diago_idl_proto_worker_proto protoreflect.FileDescriptor

var file_diago_idl_proto_worker_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x64, 0x69, 0x61, 0x67, 0x6f, 0x2d, 0x69, 0x64, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xdd, 0x01, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x27, 0x0a, 0x08, 0x72,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e,
	0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x48, 0x00, 0x52, 0x08, 0x72, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x65, 0x72, 0x12, 0x1e, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x06, 0x2e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x48, 0x00, 0x52, 0x05, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x12, 0x24, 0x0a, 0x07, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x48,
	0x00, 0x52, 0x07, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x12, 0x21, 0x0a, 0x06, 0x66, 0x69,
	0x6e, 0x69, 0x73, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x46, 0x69, 0x6e,
	0x69, 0x73, 0x68, 0x48, 0x00, 0x52, 0x06, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x12, 0x1b, 0x0a,
	0x04, 0x73, 0x74, 0x6f, 0x70, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x53, 0x74,
	0x6f, 0x70, 0x48, 0x00, 0x52, 0x04, 0x73, 0x74, 0x6f, 0x70, 0x12, 0x18, 0x0a, 0x03, 0x61, 0x63,
	0x6b, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x04, 0x2e, 0x41, 0x63, 0x6b, 0x48, 0x00, 0x52,
	0x03, 0x61, 0x63, 0x6b, 0x42, 0x09, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22,
	0x5a, 0x0a, 0x08, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x1c, 0x0a,
	0x09, 0x66, 0x72, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x09, 0x66, 0x72, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x79, 0x22, 0x37, 0x0a, 0x0b, 0x48,
	0x54, 0x54, 0x50, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65,
	0x74, 0x68, 0x6f, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68,
	0x6f, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x75, 0x72, 0x6c, 0x22, 0x80, 0x01, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x72, 0x74, 0x12, 0x15,
	0x0a, 0x06, 0x6a, 0x6f, 0x62, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x6a, 0x6f, 0x62, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x72, 0x65, 0x71, 0x75, 0x65, 0x6e,
	0x63, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x66, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x6e, 0x63, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x26, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0c, 0x2e, 0x48, 0x54, 0x54, 0x50, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x07,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x1f, 0x0a, 0x06, 0x46, 0x69, 0x6e, 0x69, 0x73,
	0x68, 0x12, 0x15, 0x0a, 0x06, 0x6a, 0x6f, 0x62, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x6a, 0x6f, 0x62, 0x49, 0x64, 0x22, 0xd6, 0x01, 0x0a, 0x07, 0x4d, 0x65, 0x74,
	0x72, 0x69, 0x63, 0x73, 0x12, 0x15, 0x0a, 0x06, 0x6a, 0x6f, 0x62, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6a, 0x6f, 0x62, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12,
	0x19, 0x0a, 0x08, 0x62, 0x79, 0x74, 0x65, 0x73, 0x5f, 0x69, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x07, 0x62, 0x79, 0x74, 0x65, 0x73, 0x49, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x62, 0x79,
	0x74, 0x65, 0x73, 0x5f, 0x6f, 0x75, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x62,
	0x79, 0x74, 0x65, 0x73, 0x4f, 0x75, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6c, 0x61, 0x74, 0x65, 0x6e,
	0x63, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x6c, 0x61, 0x74, 0x65, 0x6e, 0x63,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x22, 0x1d, 0x0a, 0x04, 0x53, 0x74, 0x6f, 0x70, 0x12, 0x15, 0x0a, 0x06, 0x6a, 0x6f, 0x62,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6a, 0x6f, 0x62, 0x49, 0x64,
	0x22, 0x05, 0x0a, 0x03, 0x41, 0x63, 0x6b, 0x32, 0x30, 0x0a, 0x06, 0x57, 0x6f, 0x72, 0x6b, 0x65,
	0x72, 0x12, 0x26, 0x0a, 0x0a, 0x43, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x12,
	0x08, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x08, 0x2e, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x42, 0x12, 0x5a, 0x10, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_diago_idl_proto_worker_proto_rawDescOnce sync.Once
	file_diago_idl_proto_worker_proto_rawDescData = file_diago_idl_proto_worker_proto_rawDesc
)

func file_diago_idl_proto_worker_proto_rawDescGZIP() []byte {
	file_diago_idl_proto_worker_proto_rawDescOnce.Do(func() {
		file_diago_idl_proto_worker_proto_rawDescData = protoimpl.X.CompressGZIP(file_diago_idl_proto_worker_proto_rawDescData)
	})
	return file_diago_idl_proto_worker_proto_rawDescData
}

var file_diago_idl_proto_worker_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_diago_idl_proto_worker_proto_goTypes = []interface{}{
	(*Message)(nil),             // 0: Message
	(*Register)(nil),            // 1: Register
	(*HTTPRequest)(nil),         // 2: HTTPRequest
	(*Start)(nil),               // 3: Start
	(*Finish)(nil),              // 4: Finish
	(*Metrics)(nil),             // 5: Metrics
	(*Stop)(nil),                // 6: Stop
	(*Ack)(nil),                 // 7: Ack
	(*timestamp.Timestamp)(nil), // 8: google.protobuf.Timestamp
}
var file_diago_idl_proto_worker_proto_depIdxs = []int32{
	1, // 0: Message.register:type_name -> Register
	3, // 1: Message.start:type_name -> Start
	5, // 2: Message.metrics:type_name -> Metrics
	4, // 3: Message.finish:type_name -> Finish
	6, // 4: Message.stop:type_name -> Stop
	7, // 5: Message.ack:type_name -> Ack
	2, // 6: Start.request:type_name -> HTTPRequest
	8, // 7: Metrics.timestamp:type_name -> google.protobuf.Timestamp
	0, // 8: Worker.Coordinate:input_type -> Message
	0, // 9: Worker.Coordinate:output_type -> Message
	9, // [9:10] is the sub-list for method output_type
	8, // [8:9] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_diago_idl_proto_worker_proto_init() }
func file_diago_idl_proto_worker_proto_init() {
	if File_diago_idl_proto_worker_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_diago_idl_proto_worker_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
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
		file_diago_idl_proto_worker_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Register); i {
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
		file_diago_idl_proto_worker_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HTTPRequest); i {
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
		file_diago_idl_proto_worker_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Start); i {
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
		file_diago_idl_proto_worker_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Finish); i {
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
		file_diago_idl_proto_worker_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Metrics); i {
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
		file_diago_idl_proto_worker_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Stop); i {
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
		file_diago_idl_proto_worker_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ack); i {
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
	file_diago_idl_proto_worker_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Message_Register)(nil),
		(*Message_Start)(nil),
		(*Message_Metrics)(nil),
		(*Message_Finish)(nil),
		(*Message_Stop)(nil),
		(*Message_Ack)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_diago_idl_proto_worker_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_diago_idl_proto_worker_proto_goTypes,
		DependencyIndexes: file_diago_idl_proto_worker_proto_depIdxs,
		MessageInfos:      file_diago_idl_proto_worker_proto_msgTypes,
	}.Build()
	File_diago_idl_proto_worker_proto = out.File
	file_diago_idl_proto_worker_proto_rawDesc = nil
	file_diago_idl_proto_worker_proto_goTypes = nil
	file_diago_idl_proto_worker_proto_depIdxs = nil
}