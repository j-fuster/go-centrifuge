// Code generated by protoc-gen-go. DO NOT EDIT.
// source: transactions/service.proto

package transactionspb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"
import _ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type TransactionStatusRequest struct {
	TransactionId        string   `protobuf:"bytes,1,opt,name=transaction_id,json=transactionId,proto3" json:"transaction_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TransactionStatusRequest) Reset()         { *m = TransactionStatusRequest{} }
func (m *TransactionStatusRequest) String() string { return proto.CompactTextString(m) }
func (*TransactionStatusRequest) ProtoMessage()    {}
func (*TransactionStatusRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_b7b04d343851d658, []int{0}
}
func (m *TransactionStatusRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransactionStatusRequest.Unmarshal(m, b)
}
func (m *TransactionStatusRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransactionStatusRequest.Marshal(b, m, deterministic)
}
func (dst *TransactionStatusRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransactionStatusRequest.Merge(dst, src)
}
func (m *TransactionStatusRequest) XXX_Size() int {
	return xxx_messageInfo_TransactionStatusRequest.Size(m)
}
func (m *TransactionStatusRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TransactionStatusRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TransactionStatusRequest proto.InternalMessageInfo

func (m *TransactionStatusRequest) GetTransactionId() string {
	if m != nil {
		return m.TransactionId
	}
	return ""
}

type TransactionStatusResponse struct {
	TransactionId        string               `protobuf:"bytes,1,opt,name=transaction_id,json=transactionId,proto3" json:"transaction_id,omitempty"`
	Status               string               `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	Message              string               `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	LastUpdated          *timestamp.Timestamp `protobuf:"bytes,4,opt,name=last_updated,json=lastUpdated,proto3" json:"last_updated,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *TransactionStatusResponse) Reset()         { *m = TransactionStatusResponse{} }
func (m *TransactionStatusResponse) String() string { return proto.CompactTextString(m) }
func (*TransactionStatusResponse) ProtoMessage()    {}
func (*TransactionStatusResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_b7b04d343851d658, []int{1}
}
func (m *TransactionStatusResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransactionStatusResponse.Unmarshal(m, b)
}
func (m *TransactionStatusResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransactionStatusResponse.Marshal(b, m, deterministic)
}
func (dst *TransactionStatusResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransactionStatusResponse.Merge(dst, src)
}
func (m *TransactionStatusResponse) XXX_Size() int {
	return xxx_messageInfo_TransactionStatusResponse.Size(m)
}
func (m *TransactionStatusResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TransactionStatusResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TransactionStatusResponse proto.InternalMessageInfo

func (m *TransactionStatusResponse) GetTransactionId() string {
	if m != nil {
		return m.TransactionId
	}
	return ""
}

func (m *TransactionStatusResponse) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *TransactionStatusResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *TransactionStatusResponse) GetLastUpdated() *timestamp.Timestamp {
	if m != nil {
		return m.LastUpdated
	}
	return nil
}

func init() {
	proto.RegisterType((*TransactionStatusRequest)(nil), "transactions.TransactionStatusRequest")
	proto.RegisterType((*TransactionStatusResponse)(nil), "transactions.TransactionStatusResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TransactionServiceClient is the client API for TransactionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TransactionServiceClient interface {
	GetTransactionStatus(ctx context.Context, in *TransactionStatusRequest, opts ...grpc.CallOption) (*TransactionStatusResponse, error)
}

type transactionServiceClient struct {
	cc *grpc.ClientConn
}

func NewTransactionServiceClient(cc *grpc.ClientConn) TransactionServiceClient {
	return &transactionServiceClient{cc}
}

func (c *transactionServiceClient) GetTransactionStatus(ctx context.Context, in *TransactionStatusRequest, opts ...grpc.CallOption) (*TransactionStatusResponse, error) {
	out := new(TransactionStatusResponse)
	err := c.cc.Invoke(ctx, "/transactions.TransactionService/GetTransactionStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TransactionServiceServer is the server API for TransactionService service.
type TransactionServiceServer interface {
	GetTransactionStatus(context.Context, *TransactionStatusRequest) (*TransactionStatusResponse, error)
}

func RegisterTransactionServiceServer(s *grpc.Server, srv TransactionServiceServer) {
	s.RegisterService(&_TransactionService_serviceDesc, srv)
}

func _TransactionService_GetTransactionStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransactionStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServiceServer).GetTransactionStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/transactions.TransactionService/GetTransactionStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServiceServer).GetTransactionStatus(ctx, req.(*TransactionStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TransactionService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "transactions.TransactionService",
	HandlerType: (*TransactionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTransactionStatus",
			Handler:    _TransactionService_GetTransactionStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "transactions/service.proto",
}

func init() { proto.RegisterFile("transactions/service.proto", fileDescriptor_service_b7b04d343851d658) }

var fileDescriptor_service_b7b04d343851d658 = []byte{
	// 346 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0x49, 0x95, 0x8a, 0xdb, 0x2a, 0xb2, 0x48, 0x89, 0x41, 0x69, 0x28, 0xa8, 0x3d, 0xd8,
	0x2c, 0xd4, 0xb3, 0x60, 0x7b, 0x29, 0xde, 0x4a, 0xad, 0x17, 0x2f, 0x65, 0x9b, 0x8c, 0x21, 0xd0,
	0x64, 0xd7, 0xcc, 0x46, 0x0f, 0xe2, 0xc5, 0x47, 0xd0, 0xa7, 0xf0, 0xec, 0x0b, 0xf8, 0x0e, 0xbe,
	0x82, 0x0f, 0x22, 0xdd, 0x4d, 0x71, 0x4b, 0x15, 0x3d, 0x85, 0xd9, 0xff, 0xcf, 0x3f, 0x33, 0xdf,
	0x10, 0x4f, 0xe5, 0x3c, 0x43, 0x1e, 0xaa, 0x44, 0x64, 0xc8, 0x10, 0xf2, 0xbb, 0x24, 0x84, 0x40,
	0xe6, 0x42, 0x09, 0x5a, 0xb7, 0x35, 0x6f, 0x3f, 0x16, 0x22, 0x9e, 0x01, 0xe3, 0x32, 0x61, 0x3c,
	0xcb, 0x84, 0xe2, 0xfa, 0xdd, 0x78, 0xbd, 0x66, 0xa9, 0xea, 0x6a, 0x5a, 0xdc, 0x30, 0x95, 0xa4,
	0x80, 0x8a, 0xa7, 0xb2, 0x34, 0x9c, 0xe8, 0x4f, 0xd8, 0x89, 0x21, 0xeb, 0xe0, 0x3d, 0x8f, 0x63,
	0xc8, 0x99, 0x90, 0xa6, 0xed, 0x4a, 0x5c, 0xab, 0x47, 0xdc, 0xf1, 0x77, 0xf3, 0x4b, 0xc5, 0x55,
	0x81, 0x23, 0xb8, 0x2d, 0x00, 0x15, 0x3d, 0x24, 0xdb, 0xd6, 0x60, 0x93, 0x24, 0x72, 0x1d, 0xdf,
	0x69, 0x6f, 0x8e, 0xb6, 0xac, 0xd7, 0x8b, 0xa8, 0xf5, 0xe6, 0x90, 0xbd, 0x1f, 0x32, 0x50, 0x8a,
	0x0c, 0xe1, 0x9f, 0x21, 0xb4, 0x41, 0xaa, 0xa8, 0x7f, 0x74, 0x2b, 0x5a, 0x2e, 0x2b, 0xea, 0x92,
	0x8d, 0x14, 0x10, 0x79, 0x0c, 0xee, 0x9a, 0x16, 0x16, 0x25, 0x3d, 0x23, 0xf5, 0x19, 0x47, 0x35,
	0x29, 0x64, 0xc4, 0x15, 0x44, 0xee, 0xba, 0xef, 0xb4, 0x6b, 0x5d, 0x2f, 0x30, 0x7c, 0x82, 0x05,
	0x9f, 0x60, 0xbc, 0xe0, 0x33, 0xaa, 0xcd, 0xfd, 0x57, 0xc6, 0xde, 0x7d, 0x77, 0x08, 0xb5, 0xa7,
	0x36, 0x07, 0xa1, 0xaf, 0x0e, 0xd9, 0x1d, 0x80, 0x5a, 0xd9, 0x87, 0x1e, 0x05, 0xf6, 0x91, 0x82,
	0xdf, 0xa0, 0x79, 0xc7, 0x7f, 0xfa, 0x0c, 0x98, 0xd6, 0xf9, 0x73, 0xcf, 0xf5, 0x1a, 0x03, 0x50,
	0xbe, 0xe5, 0xf1, 0x8d, 0xe9, 0xe9, 0xe3, 0xf3, 0xa5, 0xd2, 0xa4, 0x07, 0xcc, 0xca, 0x62, 0x0f,
	0xcb, 0x1c, 0x1f, 0xfb, 0x5d, 0xb2, 0x13, 0x8a, 0x74, 0xa9, 0x5f, 0xbf, 0x5e, 0x2e, 0x32, 0x9c,
	0xaf, 0x3f, 0x74, 0xae, 0x6d, 0xf8, 0x28, 0xa7, 0xd3, 0xaa, 0xe6, 0x72, 0xfa, 0x15, 0x00, 0x00,
	0xff, 0xff, 0x43, 0x55, 0xde, 0xba, 0x8f, 0x02, 0x00, 0x00,
}
