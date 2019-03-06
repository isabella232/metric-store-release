// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ingress.proto

package metricstore_v1

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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

type SendRequest struct {
	Batch                *Points  `protobuf:"bytes,1,opt,name=batch,proto3" json:"batch,omitempty"`
	LocalOnly            bool     `protobuf:"varint,2,opt,name=local_only,json=localOnly,proto3" json:"local_only,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SendRequest) Reset()         { *m = SendRequest{} }
func (m *SendRequest) String() string { return proto.CompactTextString(m) }
func (*SendRequest) ProtoMessage()    {}
func (*SendRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_62b704fb0e4a291c, []int{0}
}

func (m *SendRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendRequest.Unmarshal(m, b)
}
func (m *SendRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendRequest.Marshal(b, m, deterministic)
}
func (m *SendRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendRequest.Merge(m, src)
}
func (m *SendRequest) XXX_Size() int {
	return xxx_messageInfo_SendRequest.Size(m)
}
func (m *SendRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SendRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SendRequest proto.InternalMessageInfo

func (m *SendRequest) GetBatch() *Points {
	if m != nil {
		return m.Batch
	}
	return nil
}

func (m *SendRequest) GetLocalOnly() bool {
	if m != nil {
		return m.LocalOnly
	}
	return false
}

type SendResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SendResponse) Reset()         { *m = SendResponse{} }
func (m *SendResponse) String() string { return proto.CompactTextString(m) }
func (*SendResponse) ProtoMessage()    {}
func (*SendResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_62b704fb0e4a291c, []int{1}
}

func (m *SendResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendResponse.Unmarshal(m, b)
}
func (m *SendResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendResponse.Marshal(b, m, deterministic)
}
func (m *SendResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendResponse.Merge(m, src)
}
func (m *SendResponse) XXX_Size() int {
	return xxx_messageInfo_SendResponse.Size(m)
}
func (m *SendResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SendResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SendResponse proto.InternalMessageInfo

type Point struct {
	Name                 string            `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Timestamp            int64             `protobuf:"varint,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Value                float64           `protobuf:"fixed64,3,opt,name=value,proto3" json:"value,omitempty"`
	Labels               map[string]string `protobuf:"bytes,4,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Point) Reset()         { *m = Point{} }
func (m *Point) String() string { return proto.CompactTextString(m) }
func (*Point) ProtoMessage()    {}
func (*Point) Descriptor() ([]byte, []int) {
	return fileDescriptor_62b704fb0e4a291c, []int{2}
}

func (m *Point) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Point.Unmarshal(m, b)
}
func (m *Point) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Point.Marshal(b, m, deterministic)
}
func (m *Point) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Point.Merge(m, src)
}
func (m *Point) XXX_Size() int {
	return xxx_messageInfo_Point.Size(m)
}
func (m *Point) XXX_DiscardUnknown() {
	xxx_messageInfo_Point.DiscardUnknown(m)
}

var xxx_messageInfo_Point proto.InternalMessageInfo

func (m *Point) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Point) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *Point) GetValue() float64 {
	if m != nil {
		return m.Value
	}
	return 0
}

func (m *Point) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

type Points struct {
	Points               []*Point `protobuf:"bytes,1,rep,name=points,proto3" json:"points,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Points) Reset()         { *m = Points{} }
func (m *Points) String() string { return proto.CompactTextString(m) }
func (*Points) ProtoMessage()    {}
func (*Points) Descriptor() ([]byte, []int) {
	return fileDescriptor_62b704fb0e4a291c, []int{3}
}

func (m *Points) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Points.Unmarshal(m, b)
}
func (m *Points) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Points.Marshal(b, m, deterministic)
}
func (m *Points) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Points.Merge(m, src)
}
func (m *Points) XXX_Size() int {
	return xxx_messageInfo_Points.Size(m)
}
func (m *Points) XXX_DiscardUnknown() {
	xxx_messageInfo_Points.DiscardUnknown(m)
}

var xxx_messageInfo_Points proto.InternalMessageInfo

func (m *Points) GetPoints() []*Point {
	if m != nil {
		return m.Points
	}
	return nil
}

func init() {
	proto.RegisterType((*SendRequest)(nil), "metricstore.v1.SendRequest")
	proto.RegisterType((*SendResponse)(nil), "metricstore.v1.SendResponse")
	proto.RegisterType((*Point)(nil), "metricstore.v1.Point")
	proto.RegisterMapType((map[string]string)(nil), "metricstore.v1.Point.LabelsEntry")
	proto.RegisterType((*Points)(nil), "metricstore.v1.Points")
}

func init() { proto.RegisterFile("ingress.proto", fileDescriptor_62b704fb0e4a291c) }

var fileDescriptor_62b704fb0e4a291c = []byte{
	// 300 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x51, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x75, 0x9b, 0x36, 0xda, 0x89, 0x16, 0x19, 0x54, 0x42, 0xad, 0x10, 0x73, 0xca, 0x41, 0x03,
	0xd6, 0x83, 0xd6, 0xab, 0x78, 0x10, 0x44, 0x65, 0xbd, 0x79, 0x91, 0x6d, 0x1d, 0x34, 0xb8, 0xd9,
	0x8d, 0xd9, 0x6d, 0x21, 0x1f, 0xe8, 0x7f, 0x49, 0x77, 0x03, 0xad, 0xd2, 0xdb, 0xec, 0xdb, 0x37,
	0xef, 0x0d, 0xef, 0xc1, 0x5e, 0xa1, 0x3e, 0x6a, 0x32, 0x26, 0xaf, 0x6a, 0x6d, 0x35, 0x0e, 0x4a,
	0xb2, 0x75, 0x31, 0x33, 0x56, 0xd7, 0x94, 0x2f, 0x2e, 0xd2, 0x57, 0x88, 0x5e, 0x48, 0xbd, 0x73,
	0xfa, 0x9e, 0x93, 0xb1, 0x78, 0x06, 0xbd, 0xa9, 0xb0, 0xb3, 0xcf, 0x98, 0x25, 0x2c, 0x8b, 0xc6,
	0x47, 0xf9, 0x5f, 0x7a, 0xfe, 0xac, 0x0b, 0x65, 0x0d, 0xf7, 0x24, 0x3c, 0x01, 0x90, 0x7a, 0x26,
	0xe4, 0x9b, 0x56, 0xb2, 0x89, 0x3b, 0x09, 0xcb, 0x76, 0x78, 0xdf, 0x21, 0x4f, 0x4a, 0x36, 0xe9,
	0x00, 0x76, 0xbd, 0xb6, 0xa9, 0xb4, 0x32, 0x94, 0xfe, 0x30, 0xe8, 0x39, 0x01, 0x44, 0xe8, 0x2a,
	0x51, 0x92, 0x73, 0xe9, 0x73, 0x37, 0xe3, 0x08, 0xfa, 0xb6, 0x28, 0xc9, 0x58, 0x51, 0x56, 0x4e,
	0x2b, 0xe0, 0x2b, 0x00, 0x0f, 0xa0, 0xb7, 0x10, 0x72, 0x4e, 0x71, 0x90, 0xb0, 0x8c, 0x71, 0xff,
	0xc0, 0x09, 0x84, 0x52, 0x4c, 0x49, 0x9a, 0xb8, 0x9b, 0x04, 0x59, 0x34, 0x3e, 0xdd, 0x78, 0x6f,
	0xfe, 0xe0, 0x38, 0x77, 0xca, 0xd6, 0x0d, 0x6f, 0x17, 0x86, 0x13, 0x88, 0xd6, 0x60, 0xdc, 0x87,
	0xe0, 0x8b, 0x9a, 0xf6, 0xa0, 0xe5, 0xb8, 0x72, 0xec, 0x38, 0xcc, 0x3f, 0x6e, 0x3a, 0xd7, 0x2c,
	0xbd, 0x82, 0xd0, 0xe7, 0x80, 0xe7, 0x10, 0x56, 0x6e, 0x8a, 0x99, 0xf3, 0x3f, 0xdc, 0xe8, 0xcf,
	0x5b, 0xd2, 0xf8, 0x11, 0xb6, 0xef, 0x7d, 0x1b, 0x78, 0x0b, 0xdd, 0x65, 0x36, 0x78, 0xfc, 0x7f,
	0x63, 0xad, 0x8d, 0xe1, 0x68, 0xf3, 0x67, 0x1b, 0xe7, 0xd6, 0x34, 0x74, 0x9d, 0x5e, 0xfe, 0x06,
	0x00, 0x00, 0xff, 0xff, 0x04, 0xdf, 0xff, 0x7e, 0xe4, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// IngressClient is the client API for Ingress service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type IngressClient interface {
	// Send is used to emit Point batches into MetricStore. The RPC function
	// will not return until the data has been stored.
	Send(ctx context.Context, in *SendRequest, opts ...grpc.CallOption) (*SendResponse, error)
}

type ingressClient struct {
	cc *grpc.ClientConn
}

func NewIngressClient(cc *grpc.ClientConn) IngressClient {
	return &ingressClient{cc}
}

func (c *ingressClient) Send(ctx context.Context, in *SendRequest, opts ...grpc.CallOption) (*SendResponse, error) {
	out := new(SendResponse)
	err := c.cc.Invoke(ctx, "/metricstore.v1.Ingress/Send", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IngressServer is the server API for Ingress service.
type IngressServer interface {
	// Send is used to emit Point batches into MetricStore. The RPC function
	// will not return until the data has been stored.
	Send(context.Context, *SendRequest) (*SendResponse, error)
}

func RegisterIngressServer(s *grpc.Server, srv IngressServer) {
	s.RegisterService(&_Ingress_serviceDesc, srv)
}

func _Ingress_Send_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IngressServer).Send(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/metricstore.v1.Ingress/Send",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IngressServer).Send(ctx, req.(*SendRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Ingress_serviceDesc = grpc.ServiceDesc{
	ServiceName: "metricstore.v1.Ingress",
	HandlerType: (*IngressServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Send",
			Handler:    _Ingress_Send_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ingress.proto",
}
