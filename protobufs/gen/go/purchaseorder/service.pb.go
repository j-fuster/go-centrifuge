// Code generated by protoc-gen-go. DO NOT EDIT.
// source: purchaseorder/service.proto

package purchaseorderpb

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

type GetRequest struct {
	Identifier           string   `protobuf:"bytes,1,opt,name=identifier,proto3" json:"identifier,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}
func (*GetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_4c1800c623b61b4a, []int{0}
}
func (m *GetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRequest.Unmarshal(m, b)
}
func (m *GetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRequest.Marshal(b, m, deterministic)
}
func (dst *GetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRequest.Merge(dst, src)
}
func (m *GetRequest) XXX_Size() int {
	return xxx_messageInfo_GetRequest.Size(m)
}
func (m *GetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRequest proto.InternalMessageInfo

func (m *GetRequest) GetIdentifier() string {
	if m != nil {
		return m.Identifier
	}
	return ""
}

type GetVersionRequest struct {
	Identifier           string   `protobuf:"bytes,1,opt,name=identifier,proto3" json:"identifier,omitempty"`
	Version              string   `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetVersionRequest) Reset()         { *m = GetVersionRequest{} }
func (m *GetVersionRequest) String() string { return proto.CompactTextString(m) }
func (*GetVersionRequest) ProtoMessage()    {}
func (*GetVersionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_4c1800c623b61b4a, []int{1}
}
func (m *GetVersionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetVersionRequest.Unmarshal(m, b)
}
func (m *GetVersionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetVersionRequest.Marshal(b, m, deterministic)
}
func (dst *GetVersionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetVersionRequest.Merge(dst, src)
}
func (m *GetVersionRequest) XXX_Size() int {
	return xxx_messageInfo_GetVersionRequest.Size(m)
}
func (m *GetVersionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetVersionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetVersionRequest proto.InternalMessageInfo

func (m *GetVersionRequest) GetIdentifier() string {
	if m != nil {
		return m.Identifier
	}
	return ""
}

func (m *GetVersionRequest) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

type PurchaseOrderCreatePayload struct {
	Collaborators        []string           `protobuf:"bytes,1,rep,name=collaborators,proto3" json:"collaborators,omitempty"`
	Data                 *PurchaseOrderData `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *PurchaseOrderCreatePayload) Reset()         { *m = PurchaseOrderCreatePayload{} }
func (m *PurchaseOrderCreatePayload) String() string { return proto.CompactTextString(m) }
func (*PurchaseOrderCreatePayload) ProtoMessage()    {}
func (*PurchaseOrderCreatePayload) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_4c1800c623b61b4a, []int{2}
}
func (m *PurchaseOrderCreatePayload) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PurchaseOrderCreatePayload.Unmarshal(m, b)
}
func (m *PurchaseOrderCreatePayload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PurchaseOrderCreatePayload.Marshal(b, m, deterministic)
}
func (dst *PurchaseOrderCreatePayload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PurchaseOrderCreatePayload.Merge(dst, src)
}
func (m *PurchaseOrderCreatePayload) XXX_Size() int {
	return xxx_messageInfo_PurchaseOrderCreatePayload.Size(m)
}
func (m *PurchaseOrderCreatePayload) XXX_DiscardUnknown() {
	xxx_messageInfo_PurchaseOrderCreatePayload.DiscardUnknown(m)
}

var xxx_messageInfo_PurchaseOrderCreatePayload proto.InternalMessageInfo

func (m *PurchaseOrderCreatePayload) GetCollaborators() []string {
	if m != nil {
		return m.Collaborators
	}
	return nil
}

func (m *PurchaseOrderCreatePayload) GetData() *PurchaseOrderData {
	if m != nil {
		return m.Data
	}
	return nil
}

type PurchaseOrderUpdatePayload struct {
	Identifier           string             `protobuf:"bytes,1,opt,name=identifier,proto3" json:"identifier,omitempty"`
	Collaborators        []string           `protobuf:"bytes,2,rep,name=collaborators,proto3" json:"collaborators,omitempty"`
	Data                 *PurchaseOrderData `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *PurchaseOrderUpdatePayload) Reset()         { *m = PurchaseOrderUpdatePayload{} }
func (m *PurchaseOrderUpdatePayload) String() string { return proto.CompactTextString(m) }
func (*PurchaseOrderUpdatePayload) ProtoMessage()    {}
func (*PurchaseOrderUpdatePayload) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_4c1800c623b61b4a, []int{3}
}
func (m *PurchaseOrderUpdatePayload) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PurchaseOrderUpdatePayload.Unmarshal(m, b)
}
func (m *PurchaseOrderUpdatePayload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PurchaseOrderUpdatePayload.Marshal(b, m, deterministic)
}
func (dst *PurchaseOrderUpdatePayload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PurchaseOrderUpdatePayload.Merge(dst, src)
}
func (m *PurchaseOrderUpdatePayload) XXX_Size() int {
	return xxx_messageInfo_PurchaseOrderUpdatePayload.Size(m)
}
func (m *PurchaseOrderUpdatePayload) XXX_DiscardUnknown() {
	xxx_messageInfo_PurchaseOrderUpdatePayload.DiscardUnknown(m)
}

var xxx_messageInfo_PurchaseOrderUpdatePayload proto.InternalMessageInfo

func (m *PurchaseOrderUpdatePayload) GetIdentifier() string {
	if m != nil {
		return m.Identifier
	}
	return ""
}

func (m *PurchaseOrderUpdatePayload) GetCollaborators() []string {
	if m != nil {
		return m.Collaborators
	}
	return nil
}

func (m *PurchaseOrderUpdatePayload) GetData() *PurchaseOrderData {
	if m != nil {
		return m.Data
	}
	return nil
}

type PurchaseOrderResponse struct {
	Header               *ResponseHeader    `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Data                 *PurchaseOrderData `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *PurchaseOrderResponse) Reset()         { *m = PurchaseOrderResponse{} }
func (m *PurchaseOrderResponse) String() string { return proto.CompactTextString(m) }
func (*PurchaseOrderResponse) ProtoMessage()    {}
func (*PurchaseOrderResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_4c1800c623b61b4a, []int{4}
}
func (m *PurchaseOrderResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PurchaseOrderResponse.Unmarshal(m, b)
}
func (m *PurchaseOrderResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PurchaseOrderResponse.Marshal(b, m, deterministic)
}
func (dst *PurchaseOrderResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PurchaseOrderResponse.Merge(dst, src)
}
func (m *PurchaseOrderResponse) XXX_Size() int {
	return xxx_messageInfo_PurchaseOrderResponse.Size(m)
}
func (m *PurchaseOrderResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PurchaseOrderResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PurchaseOrderResponse proto.InternalMessageInfo

func (m *PurchaseOrderResponse) GetHeader() *ResponseHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *PurchaseOrderResponse) GetData() *PurchaseOrderData {
	if m != nil {
		return m.Data
	}
	return nil
}

// ResponseHeader contains a set of common fields for most documents
type ResponseHeader struct {
	DocumentId           string   `protobuf:"bytes,1,opt,name=document_id,json=documentId,proto3" json:"document_id,omitempty"`
	VersionId            string   `protobuf:"bytes,2,opt,name=version_id,json=versionId,proto3" json:"version_id,omitempty"`
	State                string   `protobuf:"bytes,3,opt,name=state,proto3" json:"state,omitempty"`
	Collaborators        []string `protobuf:"bytes,4,rep,name=collaborators,proto3" json:"collaborators,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResponseHeader) Reset()         { *m = ResponseHeader{} }
func (m *ResponseHeader) String() string { return proto.CompactTextString(m) }
func (*ResponseHeader) ProtoMessage()    {}
func (*ResponseHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_4c1800c623b61b4a, []int{5}
}
func (m *ResponseHeader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResponseHeader.Unmarshal(m, b)
}
func (m *ResponseHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResponseHeader.Marshal(b, m, deterministic)
}
func (dst *ResponseHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponseHeader.Merge(dst, src)
}
func (m *ResponseHeader) XXX_Size() int {
	return xxx_messageInfo_ResponseHeader.Size(m)
}
func (m *ResponseHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponseHeader.DiscardUnknown(m)
}

var xxx_messageInfo_ResponseHeader proto.InternalMessageInfo

func (m *ResponseHeader) GetDocumentId() string {
	if m != nil {
		return m.DocumentId
	}
	return ""
}

func (m *ResponseHeader) GetVersionId() string {
	if m != nil {
		return m.VersionId
	}
	return ""
}

func (m *ResponseHeader) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *ResponseHeader) GetCollaborators() []string {
	if m != nil {
		return m.Collaborators
	}
	return nil
}

type PurchaseOrderData struct {
	PoStatus string `protobuf:"bytes,24,opt,name=po_status,json=poStatus,proto3" json:"po_status,omitempty"`
	// purchase order number or reference number
	PoNumber string `protobuf:"bytes,1,opt,name=po_number,json=poNumber,proto3" json:"po_number,omitempty"`
	// name of the ordering company
	OrderName string `protobuf:"bytes,2,opt,name=order_name,json=orderName,proto3" json:"order_name,omitempty"`
	// street and address details of the ordering company
	OrderStreet  string `protobuf:"bytes,3,opt,name=order_street,json=orderStreet,proto3" json:"order_street,omitempty"`
	OrderCity    string `protobuf:"bytes,4,opt,name=order_city,json=orderCity,proto3" json:"order_city,omitempty"`
	OrderZipcode string `protobuf:"bytes,5,opt,name=order_zipcode,json=orderZipcode,proto3" json:"order_zipcode,omitempty"`
	// country ISO code of the ordering company of this purchase order
	OrderCountry string `protobuf:"bytes,6,opt,name=order_country,json=orderCountry,proto3" json:"order_country,omitempty"`
	// name of the recipient company
	RecipientName    string `protobuf:"bytes,7,opt,name=recipient_name,json=recipientName,proto3" json:"recipient_name,omitempty"`
	RecipientStreet  string `protobuf:"bytes,8,opt,name=recipient_street,json=recipientStreet,proto3" json:"recipient_street,omitempty"`
	RecipientCity    string `protobuf:"bytes,9,opt,name=recipient_city,json=recipientCity,proto3" json:"recipient_city,omitempty"`
	RecipientZipcode string `protobuf:"bytes,10,opt,name=recipient_zipcode,json=recipientZipcode,proto3" json:"recipient_zipcode,omitempty"`
	// country ISO code of the receipient of this purchase order
	RecipientCountry string `protobuf:"bytes,11,opt,name=recipient_country,json=recipientCountry,proto3" json:"recipient_country,omitempty"`
	// ISO currency code
	Currency string `protobuf:"bytes,12,opt,name=currency,proto3" json:"currency,omitempty"`
	// ordering gross amount including tax
	OrderAmount int64 `protobuf:"varint,13,opt,name=order_amount,json=orderAmount,proto3" json:"order_amount,omitempty"`
	// invoice amount excluding tax
	NetAmount int64  `protobuf:"varint,14,opt,name=net_amount,json=netAmount,proto3" json:"net_amount,omitempty"`
	TaxAmount int64  `protobuf:"varint,15,opt,name=tax_amount,json=taxAmount,proto3" json:"tax_amount,omitempty"`
	TaxRate   int64  `protobuf:"varint,16,opt,name=tax_rate,json=taxRate,proto3" json:"tax_rate,omitempty"`
	Recipient string `protobuf:"bytes,17,opt,name=recipient,proto3" json:"recipient,omitempty"`
	Order     string `protobuf:"bytes,18,opt,name=order,proto3" json:"order,omitempty"`
	// contact or requester or purchaser at the ordering company
	OrderContact string `protobuf:"bytes,19,opt,name=order_contact,json=orderContact,proto3" json:"order_contact,omitempty"`
	Comment      string `protobuf:"bytes,20,opt,name=comment,proto3" json:"comment,omitempty"`
	// requested delivery date
	DeliveryDate *timestamp.Timestamp `protobuf:"bytes,21,opt,name=delivery_date,json=deliveryDate,proto3" json:"delivery_date,omitempty"`
	// purchase order date
	DateCreated          *timestamp.Timestamp `protobuf:"bytes,22,opt,name=date_created,json=dateCreated,proto3" json:"date_created,omitempty"`
	ExtraData            string               `protobuf:"bytes,23,opt,name=extra_data,json=extraData,proto3" json:"extra_data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *PurchaseOrderData) Reset()         { *m = PurchaseOrderData{} }
func (m *PurchaseOrderData) String() string { return proto.CompactTextString(m) }
func (*PurchaseOrderData) ProtoMessage()    {}
func (*PurchaseOrderData) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_4c1800c623b61b4a, []int{6}
}
func (m *PurchaseOrderData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PurchaseOrderData.Unmarshal(m, b)
}
func (m *PurchaseOrderData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PurchaseOrderData.Marshal(b, m, deterministic)
}
func (dst *PurchaseOrderData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PurchaseOrderData.Merge(dst, src)
}
func (m *PurchaseOrderData) XXX_Size() int {
	return xxx_messageInfo_PurchaseOrderData.Size(m)
}
func (m *PurchaseOrderData) XXX_DiscardUnknown() {
	xxx_messageInfo_PurchaseOrderData.DiscardUnknown(m)
}

var xxx_messageInfo_PurchaseOrderData proto.InternalMessageInfo

func (m *PurchaseOrderData) GetPoStatus() string {
	if m != nil {
		return m.PoStatus
	}
	return ""
}

func (m *PurchaseOrderData) GetPoNumber() string {
	if m != nil {
		return m.PoNumber
	}
	return ""
}

func (m *PurchaseOrderData) GetOrderName() string {
	if m != nil {
		return m.OrderName
	}
	return ""
}

func (m *PurchaseOrderData) GetOrderStreet() string {
	if m != nil {
		return m.OrderStreet
	}
	return ""
}

func (m *PurchaseOrderData) GetOrderCity() string {
	if m != nil {
		return m.OrderCity
	}
	return ""
}

func (m *PurchaseOrderData) GetOrderZipcode() string {
	if m != nil {
		return m.OrderZipcode
	}
	return ""
}

func (m *PurchaseOrderData) GetOrderCountry() string {
	if m != nil {
		return m.OrderCountry
	}
	return ""
}

func (m *PurchaseOrderData) GetRecipientName() string {
	if m != nil {
		return m.RecipientName
	}
	return ""
}

func (m *PurchaseOrderData) GetRecipientStreet() string {
	if m != nil {
		return m.RecipientStreet
	}
	return ""
}

func (m *PurchaseOrderData) GetRecipientCity() string {
	if m != nil {
		return m.RecipientCity
	}
	return ""
}

func (m *PurchaseOrderData) GetRecipientZipcode() string {
	if m != nil {
		return m.RecipientZipcode
	}
	return ""
}

func (m *PurchaseOrderData) GetRecipientCountry() string {
	if m != nil {
		return m.RecipientCountry
	}
	return ""
}

func (m *PurchaseOrderData) GetCurrency() string {
	if m != nil {
		return m.Currency
	}
	return ""
}

func (m *PurchaseOrderData) GetOrderAmount() int64 {
	if m != nil {
		return m.OrderAmount
	}
	return 0
}

func (m *PurchaseOrderData) GetNetAmount() int64 {
	if m != nil {
		return m.NetAmount
	}
	return 0
}

func (m *PurchaseOrderData) GetTaxAmount() int64 {
	if m != nil {
		return m.TaxAmount
	}
	return 0
}

func (m *PurchaseOrderData) GetTaxRate() int64 {
	if m != nil {
		return m.TaxRate
	}
	return 0
}

func (m *PurchaseOrderData) GetRecipient() string {
	if m != nil {
		return m.Recipient
	}
	return ""
}

func (m *PurchaseOrderData) GetOrder() string {
	if m != nil {
		return m.Order
	}
	return ""
}

func (m *PurchaseOrderData) GetOrderContact() string {
	if m != nil {
		return m.OrderContact
	}
	return ""
}

func (m *PurchaseOrderData) GetComment() string {
	if m != nil {
		return m.Comment
	}
	return ""
}

func (m *PurchaseOrderData) GetDeliveryDate() *timestamp.Timestamp {
	if m != nil {
		return m.DeliveryDate
	}
	return nil
}

func (m *PurchaseOrderData) GetDateCreated() *timestamp.Timestamp {
	if m != nil {
		return m.DateCreated
	}
	return nil
}

func (m *PurchaseOrderData) GetExtraData() string {
	if m != nil {
		return m.ExtraData
	}
	return ""
}

func init() {
	proto.RegisterType((*GetRequest)(nil), "purchaseorder.GetRequest")
	proto.RegisterType((*GetVersionRequest)(nil), "purchaseorder.GetVersionRequest")
	proto.RegisterType((*PurchaseOrderCreatePayload)(nil), "purchaseorder.PurchaseOrderCreatePayload")
	proto.RegisterType((*PurchaseOrderUpdatePayload)(nil), "purchaseorder.PurchaseOrderUpdatePayload")
	proto.RegisterType((*PurchaseOrderResponse)(nil), "purchaseorder.PurchaseOrderResponse")
	proto.RegisterType((*ResponseHeader)(nil), "purchaseorder.ResponseHeader")
	proto.RegisterType((*PurchaseOrderData)(nil), "purchaseorder.PurchaseOrderData")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DocumentServiceClient is the client API for DocumentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DocumentServiceClient interface {
	Create(ctx context.Context, in *PurchaseOrderCreatePayload, opts ...grpc.CallOption) (*PurchaseOrderResponse, error)
	Update(ctx context.Context, in *PurchaseOrderUpdatePayload, opts ...grpc.CallOption) (*PurchaseOrderResponse, error)
	GetVersion(ctx context.Context, in *GetVersionRequest, opts ...grpc.CallOption) (*PurchaseOrderResponse, error)
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*PurchaseOrderResponse, error)
}

type documentServiceClient struct {
	cc *grpc.ClientConn
}

func NewDocumentServiceClient(cc *grpc.ClientConn) DocumentServiceClient {
	return &documentServiceClient{cc}
}

func (c *documentServiceClient) Create(ctx context.Context, in *PurchaseOrderCreatePayload, opts ...grpc.CallOption) (*PurchaseOrderResponse, error) {
	out := new(PurchaseOrderResponse)
	err := c.cc.Invoke(ctx, "/purchaseorder.DocumentService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *documentServiceClient) Update(ctx context.Context, in *PurchaseOrderUpdatePayload, opts ...grpc.CallOption) (*PurchaseOrderResponse, error) {
	out := new(PurchaseOrderResponse)
	err := c.cc.Invoke(ctx, "/purchaseorder.DocumentService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *documentServiceClient) GetVersion(ctx context.Context, in *GetVersionRequest, opts ...grpc.CallOption) (*PurchaseOrderResponse, error) {
	out := new(PurchaseOrderResponse)
	err := c.cc.Invoke(ctx, "/purchaseorder.DocumentService/GetVersion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *documentServiceClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*PurchaseOrderResponse, error) {
	out := new(PurchaseOrderResponse)
	err := c.cc.Invoke(ctx, "/purchaseorder.DocumentService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DocumentServiceServer is the server API for DocumentService service.
type DocumentServiceServer interface {
	Create(context.Context, *PurchaseOrderCreatePayload) (*PurchaseOrderResponse, error)
	Update(context.Context, *PurchaseOrderUpdatePayload) (*PurchaseOrderResponse, error)
	GetVersion(context.Context, *GetVersionRequest) (*PurchaseOrderResponse, error)
	Get(context.Context, *GetRequest) (*PurchaseOrderResponse, error)
}

func RegisterDocumentServiceServer(s *grpc.Server, srv DocumentServiceServer) {
	s.RegisterService(&_DocumentService_serviceDesc, srv)
}

func _DocumentService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PurchaseOrderCreatePayload)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DocumentServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/purchaseorder.DocumentService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DocumentServiceServer).Create(ctx, req.(*PurchaseOrderCreatePayload))
	}
	return interceptor(ctx, in, info, handler)
}

func _DocumentService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PurchaseOrderUpdatePayload)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DocumentServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/purchaseorder.DocumentService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DocumentServiceServer).Update(ctx, req.(*PurchaseOrderUpdatePayload))
	}
	return interceptor(ctx, in, info, handler)
}

func _DocumentService_GetVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetVersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DocumentServiceServer).GetVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/purchaseorder.DocumentService/GetVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DocumentServiceServer).GetVersion(ctx, req.(*GetVersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DocumentService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DocumentServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/purchaseorder.DocumentService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DocumentServiceServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DocumentService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "purchaseorder.DocumentService",
	HandlerType: (*DocumentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _DocumentService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _DocumentService_Update_Handler,
		},
		{
			MethodName: "GetVersion",
			Handler:    _DocumentService_GetVersion_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _DocumentService_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "purchaseorder/service.proto",
}

func init() {
	proto.RegisterFile("purchaseorder/service.proto", fileDescriptor_service_4c1800c623b61b4a)
}

var fileDescriptor_service_4c1800c623b61b4a = []byte{
	// 949 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x55, 0x4f, 0x6f, 0x1b, 0x45,
	0x14, 0xd7, 0xe6, 0x8f, 0x63, 0x8f, 0xed, 0xa4, 0x9e, 0xb6, 0x30, 0xdd, 0x34, 0x74, 0x31, 0xad,
	0x48, 0xdb, 0xc4, 0x96, 0x42, 0xe1, 0x80, 0x84, 0x50, 0x9a, 0x48, 0xa1, 0x07, 0x4a, 0xb4, 0x01,
	0x0e, 0x15, 0x92, 0x35, 0xde, 0x7d, 0x76, 0x56, 0xf2, 0xee, 0x2c, 0xb3, 0xe3, 0x60, 0x53, 0xf5,
	0x82, 0x38, 0x72, 0x21, 0x5c, 0x40, 0x48, 0x7c, 0x08, 0xbe, 0x0a, 0x5f, 0x81, 0xaf, 0xc0, 0x1d,
	0xcd, 0x9b, 0x59, 0xdb, 0x6b, 0x17, 0x27, 0xe5, 0xb4, 0x9a, 0xf7, 0xfb, 0xbd, 0xb7, 0xbf, 0xf7,
	0x66, 0xe6, 0x37, 0x64, 0x3b, 0x1d, 0xca, 0xe0, 0x9c, 0x67, 0x20, 0x64, 0x08, 0xb2, 0x9d, 0x81,
	0xbc, 0x88, 0x02, 0x68, 0xa5, 0x52, 0x28, 0x41, 0xeb, 0x05, 0xd0, 0xbd, 0xdb, 0x17, 0xa2, 0x3f,
	0x80, 0x36, 0x4f, 0xa3, 0x36, 0x4f, 0x12, 0xa1, 0xb8, 0x8a, 0x44, 0x92, 0x19, 0xb2, 0x7b, 0xcf,
	0xa2, 0xb8, 0xea, 0x0e, 0x7b, 0x6d, 0x15, 0xc5, 0x90, 0x29, 0x1e, 0xa7, 0x96, 0xb0, 0x87, 0x9f,
	0x60, 0xbf, 0x0f, 0xc9, 0x7e, 0xf6, 0x1d, 0xef, 0xf7, 0x41, 0xb6, 0x45, 0x8a, 0x25, 0x16, 0xcb,
	0x35, 0xf7, 0x08, 0x39, 0x01, 0xe5, 0xc3, 0xb7, 0x43, 0xc8, 0x14, 0x7d, 0x87, 0x90, 0x28, 0x84,
	0x44, 0x45, 0xbd, 0x08, 0x24, 0x73, 0x3c, 0x67, 0xb7, 0xe2, 0xcf, 0x44, 0x9a, 0x9f, 0x93, 0xc6,
	0x09, 0xa8, 0xaf, 0x41, 0x66, 0x91, 0x48, 0xae, 0x99, 0x44, 0x19, 0xd9, 0xb8, 0x30, 0x19, 0x6c,
	0x05, 0xc1, 0x7c, 0xd9, 0x1c, 0x11, 0xf7, 0xd4, 0xb6, 0xfe, 0x85, 0x6e, 0xfd, 0x48, 0x02, 0x57,
	0x70, 0xca, 0xc7, 0x03, 0xc1, 0x43, 0x7a, 0x9f, 0xd4, 0x03, 0x31, 0x18, 0xf0, 0xae, 0x90, 0x5c,
	0x09, 0x99, 0x31, 0xc7, 0x5b, 0xdd, 0xad, 0xf8, 0xc5, 0x20, 0x7d, 0x42, 0xd6, 0x42, 0xae, 0x38,
	0x96, 0xae, 0x1e, 0x78, 0xad, 0xc2, 0x2c, 0x5b, 0x85, 0xf2, 0xc7, 0x5c, 0x71, 0x1f, 0xd9, 0xcd,
	0x5f, 0x9d, 0xb9, 0x5f, 0x7f, 0x95, 0x86, 0x33, 0xbf, 0xbe, 0xaa, 0xa5, 0x05, 0x69, 0x2b, 0xcb,
	0xa4, 0xad, 0xbe, 0x91, 0xb4, 0x1f, 0x1d, 0x72, 0xbb, 0x80, 0xf9, 0x90, 0xa5, 0x22, 0xc9, 0x80,
	0x7e, 0x48, 0x4a, 0xe7, 0xc0, 0x43, 0xab, 0xa8, 0x7a, 0xb0, 0x33, 0x57, 0x31, 0x27, 0x7e, 0x86,
	0x24, 0xdf, 0x92, 0xff, 0xe7, 0x84, 0x7e, 0x72, 0xc8, 0x66, 0xb1, 0x20, 0xbd, 0x47, 0xaa, 0xa1,
	0x08, 0x86, 0x31, 0x24, 0xaa, 0x13, 0x85, 0xf9, 0x58, 0xf2, 0xd0, 0xb3, 0x90, 0xee, 0x10, 0x62,
	0xb7, 0x56, 0xe3, 0x66, 0xb3, 0x2b, 0x36, 0xf2, 0x2c, 0xa4, 0xb7, 0xc8, 0x7a, 0xa6, 0xb8, 0x02,
	0x1c, 0x48, 0xc5, 0x37, 0x8b, 0xc5, 0x59, 0xae, 0xbd, 0x66, 0x96, 0xcd, 0x7f, 0x4a, 0xa4, 0xb1,
	0x20, 0x95, 0x6e, 0x93, 0x4a, 0x2a, 0x3a, 0xba, 0xce, 0x30, 0x63, 0x0c, 0xab, 0x96, 0x53, 0x71,
	0x86, 0x6b, 0x0b, 0x26, 0xc3, 0xb8, 0x3b, 0xd9, 0xc3, 0x72, 0x2a, 0x9e, 0xe3, 0x5a, 0x4b, 0xc5,
	0xfe, 0x3b, 0x09, 0x8f, 0x21, 0x97, 0x8a, 0x91, 0xe7, 0x3c, 0x06, 0xfa, 0x2e, 0xa9, 0x19, 0x38,
	0x53, 0x12, 0x40, 0x59, 0xc5, 0x55, 0x8c, 0x9d, 0x61, 0x68, 0x5a, 0x21, 0x88, 0xd4, 0x98, 0xad,
	0xcd, 0x54, 0x38, 0x8a, 0xd4, 0x98, 0xbe, 0x47, 0xea, 0x06, 0xfe, 0x3e, 0x4a, 0x03, 0x11, 0x02,
	0x5b, 0x47, 0x86, 0x29, 0xfb, 0xc2, 0xc4, 0xa6, 0xa4, 0x40, 0x0c, 0x13, 0x25, 0xc7, 0xac, 0x34,
	0x43, 0x3a, 0x32, 0x31, 0xfa, 0x80, 0x6c, 0x4a, 0x08, 0xa2, 0x34, 0xd2, 0x73, 0x47, 0xb9, 0x1b,
	0xc8, 0xaa, 0x4f, 0xa2, 0x28, 0xf9, 0x21, 0xb9, 0x31, 0xa5, 0x59, 0xd9, 0x65, 0x24, 0x6e, 0x4d,
	0xe2, 0x56, 0x7a, 0xa1, 0x22, 0xca, 0xaf, 0xcc, 0x55, 0xc4, 0x16, 0x1e, 0x93, 0xc6, 0x94, 0x96,
	0xb7, 0x41, 0x90, 0x39, 0xfd, 0x55, 0xde, 0x4a, 0x81, 0x9c, 0xb7, 0x53, 0x9d, 0x23, 0xe7, 0x2d,
	0xb9, 0xa4, 0x1c, 0x0c, 0xa5, 0x84, 0x24, 0x18, 0xb3, 0x9a, 0xd9, 0x99, 0x7c, 0x3d, 0x1d, 0x3d,
	0x8f, 0x35, 0x9b, 0xd5, 0x3d, 0x67, 0x77, 0xd5, 0x8e, 0xfe, 0x10, 0x43, 0x7a, 0xf4, 0x09, 0xa8,
	0x9c, 0xb0, 0x89, 0x84, 0x4a, 0x02, 0x6a, 0x0a, 0x2b, 0x3e, 0xca, 0xe1, 0x2d, 0x03, 0x2b, 0x3e,
	0xb2, 0xf0, 0x1d, 0x52, 0xd6, 0xb0, 0xd4, 0x27, 0xf1, 0x06, 0x82, 0x1b, 0x8a, 0x8f, 0x7c, 0x7d,
	0x16, 0xef, 0x92, 0xca, 0x44, 0x2b, 0x6b, 0x98, 0x2d, 0x9d, 0x04, 0xf4, 0xf9, 0x45, 0x15, 0x8c,
	0x9a, 0xf3, 0x8b, 0x8b, 0xd9, 0x3d, 0x4c, 0x14, 0x0f, 0x14, 0xbb, 0x59, 0xd8, 0x43, 0x8c, 0x69,
	0x0f, 0x0c, 0x44, 0xac, 0xaf, 0x09, 0xbb, 0x65, 0x3c, 0xd0, 0x2e, 0xe9, 0xa7, 0xa4, 0x1e, 0xc2,
	0x20, 0xba, 0x00, 0x39, 0xee, 0x68, 0x0b, 0x62, 0xb7, 0xf1, 0x9a, 0xba, 0x2d, 0xe3, 0xf3, 0xad,
	0xdc, 0xe7, 0x5b, 0x5f, 0xe6, 0x3e, 0xef, 0xd7, 0xf2, 0x84, 0x63, 0xad, 0xf9, 0x13, 0x52, 0xd3,
	0x79, 0x9d, 0x00, 0xcd, 0x33, 0x64, 0x6f, 0x5d, 0x99, 0x5f, 0xd5, 0x7c, 0xe3, 0xb5, 0x78, 0x67,
	0x61, 0xa4, 0x24, 0xef, 0xa0, 0x47, 0xbc, 0x6d, 0x7a, 0xc6, 0x88, 0xbe, 0x61, 0x07, 0xbf, 0xad,
	0x93, 0xad, 0x63, 0x7b, 0xc3, 0xcf, 0xcc, 0xab, 0x45, 0x7f, 0x76, 0x48, 0xc9, 0xa4, 0xd3, 0x87,
	0xcb, 0xdc, 0xa4, 0x60, 0xe7, 0xee, 0xfd, 0x65, 0xd4, 0xdc, 0x69, 0x9a, 0x1f, 0x5d, 0x1e, 0xba,
	0x2e, 0x33, 0x99, 0x99, 0xc7, 0xbd, 0x3c, 0xc7, 0xc3, 0xa4, 0x1f, 0xfe, 0xfa, 0xfb, 0x97, 0x95,
	0x9b, 0xcd, 0xcd, 0x76, 0xa1, 0xd4, 0xc7, 0xce, 0x23, 0xfa, 0x87, 0x43, 0x4a, 0xc6, 0xc3, 0x97,
	0x6b, 0x2a, 0xf8, 0xfc, 0x35, 0x35, 0x1d, 0xa1, 0x26, 0x93, 0xf9, 0x1f, 0x9a, 0x3c, 0x77, 0xbb,
	0xa8, 0xa9, 0xfd, 0x72, 0xfa, 0x5c, 0xbc, 0xd2, 0x02, 0xff, 0x74, 0xf0, 0xa5, 0xb5, 0x6f, 0x27,
	0x9d, 0xb7, 0xe1, 0x85, 0x67, 0xf5, 0x9a, 0xda, 0xbe, 0xb9, 0x3c, 0xdc, 0x73, 0x1f, 0x9d, 0x80,
	0xf2, 0xb8, 0x97, 0xa5, 0x10, 0x44, 0xbd, 0x28, 0xf0, 0xac, 0xe5, 0x7a, 0xa2, 0xf7, 0x7a, 0xb5,
	0xef, 0xd3, 0x07, 0x4b, 0xd4, 0xb6, 0x5f, 0xda, 0xfc, 0x57, 0xf4, 0x77, 0x87, 0xac, 0x9e, 0x80,
	0xa2, 0x77, 0x16, 0xd5, 0xbe, 0x99, 0xcc, 0xb3, 0xcb, 0xc3, 0x7d, 0xf7, 0xb1, 0x96, 0xa9, 0xce,
	0xc1, 0x33, 0x77, 0x5d, 0x5d, 0xa9, 0x73, 0x87, 0x2e, 0x9b, 0xea, 0xd3, 0x27, 0xa4, 0x11, 0x88,
	0xb8, 0xf8, 0xff, 0xa7, 0x35, 0x7b, 0x4a, 0x4f, 0xf5, 0xb9, 0x3f, 0x75, 0x5e, 0x6c, 0x15, 0xe0,
	0xb4, 0xdb, 0x2d, 0xe1, 0x8d, 0xf8, 0xe0, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x2b, 0x84, 0xe7,
	0x4b, 0x93, 0x09, 0x00, 0x00,
}
