// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api.proto

package api

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type EchoMessage struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EchoMessage) Reset()         { *m = EchoMessage{} }
func (m *EchoMessage) String() string { return proto.CompactTextString(m) }
func (*EchoMessage) ProtoMessage()    {}
func (*EchoMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{0}
}

func (m *EchoMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EchoMessage.Unmarshal(m, b)
}
func (m *EchoMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EchoMessage.Marshal(b, m, deterministic)
}
func (m *EchoMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EchoMessage.Merge(m, src)
}
func (m *EchoMessage) XXX_Size() int {
	return xxx_messageInfo_EchoMessage.Size(m)
}
func (m *EchoMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_EchoMessage.DiscardUnknown(m)
}

var xxx_messageInfo_EchoMessage proto.InternalMessageInfo

func (m *EchoMessage) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*EchoMessage)(nil), "api.echoMessage")
}

func init() { proto.RegisterFile("api.proto", fileDescriptor_00212fb1f9d3bf1c) }

var fileDescriptor_00212fb1f9d3bf1c = []byte{
	// 109 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4c, 0x2c, 0xc8, 0xd4,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4e, 0x2c, 0xc8, 0x54, 0x52, 0xe7, 0xe2, 0x4e, 0x4d,
	0xce, 0xc8, 0xf7, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0x15, 0x92, 0xe0, 0x62, 0xcf, 0x85, 0x30,
	0x25, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0x60, 0x5c, 0x23, 0x6b, 0x2e, 0x16, 0xd7, 0xe4, 0x8c,
	0x7c, 0x21, 0x63, 0x2e, 0x6e, 0x10, 0x1d, 0x94, 0x5a, 0x58, 0x9a, 0x5a, 0x5c, 0x22, 0x24, 0xa0,
	0x07, 0x32, 0x10, 0xc9, 0x08, 0x29, 0x0c, 0x11, 0x25, 0x86, 0x24, 0x36, 0xb0, 0x8d, 0xc6, 0x80,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xa5, 0xb0, 0x1c, 0xb9, 0x7e, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// EchoClient is the client API for Echo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EchoClient interface {
	EchoRequest(ctx context.Context, in *EchoMessage, opts ...grpc.CallOption) (*EchoMessage, error)
}

type echoClient struct {
	cc *grpc.ClientConn
}

func NewEchoClient(cc *grpc.ClientConn) EchoClient {
	return &echoClient{cc}
}

func (c *echoClient) EchoRequest(ctx context.Context, in *EchoMessage, opts ...grpc.CallOption) (*EchoMessage, error) {
	out := new(EchoMessage)
	err := c.cc.Invoke(ctx, "/api.Echo/EchoRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EchoServer is the server API for Echo service.
type EchoServer interface {
	EchoRequest(context.Context, *EchoMessage) (*EchoMessage, error)
}

// UnimplementedEchoServer can be embedded to have forward compatible implementations.
type UnimplementedEchoServer struct {
}

func (*UnimplementedEchoServer) EchoRequest(ctx context.Context, req *EchoMessage) (*EchoMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EchoRequest not implemented")
}

func RegisterEchoServer(s *grpc.Server, srv EchoServer) {
	s.RegisterService(&_Echo_serviceDesc, srv)
}

func _Echo_EchoRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EchoMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EchoServer).EchoRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Echo/EchoRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EchoServer).EchoRequest(ctx, req.(*EchoMessage))
	}
	return interceptor(ctx, in, info, handler)
}

var _Echo_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.Echo",
	HandlerType: (*EchoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "EchoRequest",
			Handler:    _Echo_EchoRequest_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}