// Code generated by protoc-gen-go. DO NOT EDIT.
// source: order_service.proto

package ecommerce

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Order struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Items                []string `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
	Description          string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Price                float32  `protobuf:"fixed32,4,opt,name=price,proto3" json:"price,omitempty"`
	Destination          string   `protobuf:"bytes,5,opt,name=destination,proto3" json:"destination,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Order) Reset()         { *m = Order{} }
func (m *Order) String() string { return proto.CompactTextString(m) }
func (*Order) ProtoMessage()    {}
func (*Order) Descriptor() ([]byte, []int) {
	return fileDescriptor_93a121d2d2ec3d32, []int{0}
}

func (m *Order) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Order.Unmarshal(m, b)
}
func (m *Order) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Order.Marshal(b, m, deterministic)
}
func (m *Order) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Order.Merge(m, src)
}
func (m *Order) XXX_Size() int {
	return xxx_messageInfo_Order.Size(m)
}
func (m *Order) XXX_DiscardUnknown() {
	xxx_messageInfo_Order.DiscardUnknown(m)
}

var xxx_messageInfo_Order proto.InternalMessageInfo

func (m *Order) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Order) GetItems() []string {
	if m != nil {
		return m.Items
	}
	return nil
}

func (m *Order) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Order) GetPrice() float32 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *Order) GetDestination() string {
	if m != nil {
		return m.Destination
	}
	return ""
}

type CombinedShipment struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Status               string   `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	OrderIDList          []string `protobuf:"bytes,3,rep,name=orderIDList,proto3" json:"orderIDList,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CombinedShipment) Reset()         { *m = CombinedShipment{} }
func (m *CombinedShipment) String() string { return proto.CompactTextString(m) }
func (*CombinedShipment) ProtoMessage()    {}
func (*CombinedShipment) Descriptor() ([]byte, []int) {
	return fileDescriptor_93a121d2d2ec3d32, []int{1}
}

func (m *CombinedShipment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CombinedShipment.Unmarshal(m, b)
}
func (m *CombinedShipment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CombinedShipment.Marshal(b, m, deterministic)
}
func (m *CombinedShipment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CombinedShipment.Merge(m, src)
}
func (m *CombinedShipment) XXX_Size() int {
	return xxx_messageInfo_CombinedShipment.Size(m)
}
func (m *CombinedShipment) XXX_DiscardUnknown() {
	xxx_messageInfo_CombinedShipment.DiscardUnknown(m)
}

var xxx_messageInfo_CombinedShipment proto.InternalMessageInfo

func (m *CombinedShipment) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *CombinedShipment) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *CombinedShipment) GetOrderIDList() []string {
	if m != nil {
		return m.OrderIDList
	}
	return nil
}

func init() {
	proto.RegisterType((*Order)(nil), "ecommerce.Order")
	proto.RegisterType((*CombinedShipment)(nil), "ecommerce.CombinedShipment")
}

func init() { proto.RegisterFile("order_service.proto", fileDescriptor_93a121d2d2ec3d32) }

var fileDescriptor_93a121d2d2ec3d32 = []byte{
	// 327 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0x5f, 0x4b, 0xc3, 0x30,
	0x14, 0xc5, 0x49, 0xe7, 0x86, 0x8b, 0xff, 0x46, 0x15, 0x29, 0x53, 0xa4, 0xec, 0xa9, 0x4f, 0xdd,
	0xd0, 0x37, 0x9f, 0x44, 0x7d, 0x11, 0x14, 0x61, 0x03, 0x9f, 0x04, 0xc9, 0x92, 0x6b, 0x77, 0x61,
	0x6d, 0xc2, 0x4d, 0xaa, 0x1f, 0xc1, 0xcf, 0xe7, 0x37, 0x92, 0xa4, 0xdb, 0x18, 0x0e, 0x85, 0x3d,
	0x9e, 0xd3, 0x7b, 0x7e, 0x9c, 0x9c, 0xf2, 0x63, 0x4d, 0x0a, 0xe8, 0xcd, 0x02, 0x7d, 0xa0, 0x84,
	0xdc, 0x90, 0x76, 0x3a, 0xee, 0x82, 0xd4, 0x65, 0x09, 0x24, 0xa1, 0x7f, 0x51, 0x68, 0x5d, 0xcc,
	0x61, 0x18, 0x3e, 0x4c, 0xeb, 0xf7, 0xe1, 0x27, 0x09, 0x63, 0x80, 0x6c, 0x73, 0x3a, 0xf8, 0x62,
	0xbc, 0xfd, 0xec, 0x11, 0xf1, 0x21, 0x8f, 0x50, 0x25, 0x2c, 0x65, 0x59, 0x77, 0x1c, 0xa1, 0x8a,
	0x4f, 0x78, 0x1b, 0x1d, 0x94, 0x36, 0x89, 0xd2, 0x56, 0xd6, 0x1d, 0x37, 0x22, 0x4e, 0xf9, 0x9e,
	0x02, 0x2b, 0x09, 0x8d, 0x43, 0x5d, 0x25, 0xad, 0x70, 0xbe, 0x6e, 0xf9, 0x9c, 0x21, 0x94, 0x90,
	0xec, 0xa4, 0x2c, 0x8b, 0xc6, 0x8d, 0x58, 0xe4, 0x1c, 0x56, 0x22, 0xe4, 0xda, 0xab, 0xdc, 0xd2,
	0x1a, 0xbc, 0xf2, 0xde, 0x9d, 0x2e, 0xa7, 0x58, 0x81, 0x9a, 0xcc, 0xd0, 0x94, 0x50, 0xb9, 0x8d,
	0x4e, 0xa7, 0xbc, 0x63, 0x9d, 0x70, 0xb5, 0x2f, 0xe5, 0xbd, 0x85, 0xf2, 0xf4, 0xb0, 0xc3, 0xc3,
	0xfd, 0x23, 0x5a, 0x97, 0xb4, 0x42, 0xe3, 0x75, 0xeb, 0xf2, 0x3b, 0xe2, 0x47, 0xe1, 0x9d, 0x4f,
	0xa2, 0x12, 0x05, 0x04, 0xfa, 0x35, 0xdf, 0x15, 0x4a, 0x35, 0xaf, 0xef, 0xe5, 0xab, 0xcd, 0xf2,
	0xe0, 0xf4, 0xcf, 0xf3, 0x66, 0xba, 0x7c, 0x39, 0x5d, 0x3e, 0x71, 0x84, 0x55, 0xf1, 0x22, 0xe6,
	0x35, 0xf8, 0x6c, 0x01, 0xae, 0xc9, 0xfe, 0x7b, 0xd9, 0xdf, 0x20, 0xc7, 0x37, 0x7c, 0xdf, 0x82,
	0x20, 0x39, 0x0b, 0xd2, 0x6e, 0x9b, 0x1f, 0x31, 0x4f, 0xa8, 0x8d, 0x12, 0x0e, 0x16, 0x84, 0x2d,
	0xdb, 0x67, 0x2c, 0xbe, 0xe5, 0x07, 0x86, 0xb4, 0x04, 0x6b, 0xff, 0x44, 0x9c, 0xad, 0x39, 0xbf,
	0xff, 0x4c, 0xc6, 0x46, 0x6c, 0xda, 0x09, 0xec, 0xab, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x13,
	0x02, 0xeb, 0x3f, 0x84, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// OrderManagementClient is the client API for OrderManagement service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type OrderManagementClient interface {
	AddOrder(ctx context.Context, in *Order, opts ...grpc.CallOption) (*wrappers.StringValue, error)
	GetOrder(ctx context.Context, in *wrappers.StringValue, opts ...grpc.CallOption) (*Order, error)
	SearchOrders(ctx context.Context, in *wrappers.StringValue, opts ...grpc.CallOption) (OrderManagement_SearchOrdersClient, error)
	UpdateOrders(ctx context.Context, opts ...grpc.CallOption) (OrderManagement_UpdateOrdersClient, error)
	ProcessOrders(ctx context.Context, opts ...grpc.CallOption) (OrderManagement_ProcessOrdersClient, error)
}

type orderManagementClient struct {
	cc *grpc.ClientConn
}

func NewOrderManagementClient(cc *grpc.ClientConn) OrderManagementClient {
	return &orderManagementClient{cc}
}

func (c *orderManagementClient) AddOrder(ctx context.Context, in *Order, opts ...grpc.CallOption) (*wrappers.StringValue, error) {
	out := new(wrappers.StringValue)
	err := c.cc.Invoke(ctx, "/ecommerce.OrderManagement/addOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderManagementClient) GetOrder(ctx context.Context, in *wrappers.StringValue, opts ...grpc.CallOption) (*Order, error) {
	out := new(Order)
	err := c.cc.Invoke(ctx, "/ecommerce.OrderManagement/getOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderManagementClient) SearchOrders(ctx context.Context, in *wrappers.StringValue, opts ...grpc.CallOption) (OrderManagement_SearchOrdersClient, error) {
	stream, err := c.cc.NewStream(ctx, &_OrderManagement_serviceDesc.Streams[0], "/ecommerce.OrderManagement/searchOrders", opts...)
	if err != nil {
		return nil, err
	}
	x := &orderManagementSearchOrdersClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type OrderManagement_SearchOrdersClient interface {
	Recv() (*Order, error)
	grpc.ClientStream
}

type orderManagementSearchOrdersClient struct {
	grpc.ClientStream
}

func (x *orderManagementSearchOrdersClient) Recv() (*Order, error) {
	m := new(Order)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *orderManagementClient) UpdateOrders(ctx context.Context, opts ...grpc.CallOption) (OrderManagement_UpdateOrdersClient, error) {
	stream, err := c.cc.NewStream(ctx, &_OrderManagement_serviceDesc.Streams[1], "/ecommerce.OrderManagement/updateOrders", opts...)
	if err != nil {
		return nil, err
	}
	x := &orderManagementUpdateOrdersClient{stream}
	return x, nil
}

type OrderManagement_UpdateOrdersClient interface {
	Send(*Order) error
	CloseAndRecv() (*wrappers.StringValue, error)
	grpc.ClientStream
}

type orderManagementUpdateOrdersClient struct {
	grpc.ClientStream
}

func (x *orderManagementUpdateOrdersClient) Send(m *Order) error {
	return x.ClientStream.SendMsg(m)
}

func (x *orderManagementUpdateOrdersClient) CloseAndRecv() (*wrappers.StringValue, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(wrappers.StringValue)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *orderManagementClient) ProcessOrders(ctx context.Context, opts ...grpc.CallOption) (OrderManagement_ProcessOrdersClient, error) {
	stream, err := c.cc.NewStream(ctx, &_OrderManagement_serviceDesc.Streams[2], "/ecommerce.OrderManagement/processOrders", opts...)
	if err != nil {
		return nil, err
	}
	x := &orderManagementProcessOrdersClient{stream}
	return x, nil
}

type OrderManagement_ProcessOrdersClient interface {
	Send(*Order) error
	Recv() (*CombinedShipment, error)
	grpc.ClientStream
}

type orderManagementProcessOrdersClient struct {
	grpc.ClientStream
}

func (x *orderManagementProcessOrdersClient) Send(m *Order) error {
	return x.ClientStream.SendMsg(m)
}

func (x *orderManagementProcessOrdersClient) Recv() (*CombinedShipment, error) {
	m := new(CombinedShipment)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// OrderManagementServer is the server API for OrderManagement service.
type OrderManagementServer interface {
	AddOrder(context.Context, *Order) (*wrappers.StringValue, error)
	GetOrder(context.Context, *wrappers.StringValue) (*Order, error)
	SearchOrders(*wrappers.StringValue, OrderManagement_SearchOrdersServer) error
	UpdateOrders(OrderManagement_UpdateOrdersServer) error
	ProcessOrders(OrderManagement_ProcessOrdersServer) error
}

// UnimplementedOrderManagementServer can be embedded to have forward compatible implementations.
type UnimplementedOrderManagementServer struct {
}

func (*UnimplementedOrderManagementServer) AddOrder(ctx context.Context, req *Order) (*wrappers.StringValue, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddOrder not implemented")
}
func (*UnimplementedOrderManagementServer) GetOrder(ctx context.Context, req *wrappers.StringValue) (*Order, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrder not implemented")
}
func (*UnimplementedOrderManagementServer) SearchOrders(req *wrappers.StringValue, srv OrderManagement_SearchOrdersServer) error {
	return status.Errorf(codes.Unimplemented, "method SearchOrders not implemented")
}
func (*UnimplementedOrderManagementServer) UpdateOrders(srv OrderManagement_UpdateOrdersServer) error {
	return status.Errorf(codes.Unimplemented, "method UpdateOrders not implemented")
}
func (*UnimplementedOrderManagementServer) ProcessOrders(srv OrderManagement_ProcessOrdersServer) error {
	return status.Errorf(codes.Unimplemented, "method ProcessOrders not implemented")
}

func RegisterOrderManagementServer(s *grpc.Server, srv OrderManagementServer) {
	s.RegisterService(&_OrderManagement_serviceDesc, srv)
}

func _OrderManagement_AddOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Order)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderManagementServer).AddOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ecommerce.OrderManagement/AddOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderManagementServer).AddOrder(ctx, req.(*Order))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderManagement_GetOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(wrappers.StringValue)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderManagementServer).GetOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ecommerce.OrderManagement/GetOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderManagementServer).GetOrder(ctx, req.(*wrappers.StringValue))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderManagement_SearchOrders_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(wrappers.StringValue)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(OrderManagementServer).SearchOrders(m, &orderManagementSearchOrdersServer{stream})
}

type OrderManagement_SearchOrdersServer interface {
	Send(*Order) error
	grpc.ServerStream
}

type orderManagementSearchOrdersServer struct {
	grpc.ServerStream
}

func (x *orderManagementSearchOrdersServer) Send(m *Order) error {
	return x.ServerStream.SendMsg(m)
}

func _OrderManagement_UpdateOrders_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(OrderManagementServer).UpdateOrders(&orderManagementUpdateOrdersServer{stream})
}

type OrderManagement_UpdateOrdersServer interface {
	SendAndClose(*wrappers.StringValue) error
	Recv() (*Order, error)
	grpc.ServerStream
}

type orderManagementUpdateOrdersServer struct {
	grpc.ServerStream
}

func (x *orderManagementUpdateOrdersServer) SendAndClose(m *wrappers.StringValue) error {
	return x.ServerStream.SendMsg(m)
}

func (x *orderManagementUpdateOrdersServer) Recv() (*Order, error) {
	m := new(Order)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _OrderManagement_ProcessOrders_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(OrderManagementServer).ProcessOrders(&orderManagementProcessOrdersServer{stream})
}

type OrderManagement_ProcessOrdersServer interface {
	Send(*CombinedShipment) error
	Recv() (*Order, error)
	grpc.ServerStream
}

type orderManagementProcessOrdersServer struct {
	grpc.ServerStream
}

func (x *orderManagementProcessOrdersServer) Send(m *CombinedShipment) error {
	return x.ServerStream.SendMsg(m)
}

func (x *orderManagementProcessOrdersServer) Recv() (*Order, error) {
	m := new(Order)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _OrderManagement_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ecommerce.OrderManagement",
	HandlerType: (*OrderManagementServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "addOrder",
			Handler:    _OrderManagement_AddOrder_Handler,
		},
		{
			MethodName: "getOrder",
			Handler:    _OrderManagement_GetOrder_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "searchOrders",
			Handler:       _OrderManagement_SearchOrders_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "updateOrders",
			Handler:       _OrderManagement_UpdateOrders_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "processOrders",
			Handler:       _OrderManagement_ProcessOrders_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "order_service.proto",
}