// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ctrpc.proto

/*
Package ctrpcpt is a generated protocol buffer package.

It is generated from these files:
	ctrpc.proto

It has these top-level messages:
	KeepAliveRequest
	KeepAliveReply
	LoginRequest
	LoginReply
	LogoutRequest
	LogoutReply
	ChatRequest
	ChatReply
	PushMessageRequest
	PushMessageReply
*/
package ctrpcpt

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import pbcommon "github.com/holyreaper/ggserver/rpcservice/pb/common"

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

// KeepAlive
type KeepAliveRequest struct {
	Time int64 `protobuf:"varint,1,opt,name=time" json:"time,omitempty"`
}

func (m *KeepAliveRequest) Reset()                    { *m = KeepAliveRequest{} }
func (m *KeepAliveRequest) String() string            { return proto.CompactTextString(m) }
func (*KeepAliveRequest) ProtoMessage()               {}
func (*KeepAliveRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *KeepAliveRequest) GetTime() int64 {
	if m != nil {
		return m.Time
	}
	return 0
}

type KeepAliveReply struct {
	Result int32 `protobuf:"varint,1,opt,name=result" json:"result,omitempty"`
}

func (m *KeepAliveReply) Reset()                    { *m = KeepAliveReply{} }
func (m *KeepAliveReply) String() string            { return proto.CompactTextString(m) }
func (*KeepAliveReply) ProtoMessage()               {}
func (*KeepAliveReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *KeepAliveReply) GetResult() int32 {
	if m != nil {
		return m.Result
	}
	return 0
}

// login
type LoginRequest struct {
	Uid      int64  `protobuf:"varint,1,opt,name=uid" json:"uid,omitempty"`
	ServerId int32  `protobuf:"varint,2,opt,name=serverId" json:"serverId,omitempty"`
	Uname    string `protobuf:"bytes,3,opt,name=uname" json:"uname,omitempty"`
	Figure   int32  `protobuf:"varint,4,opt,name=figure" json:"figure,omitempty"`
	Level    int32  `protobuf:"varint,5,opt,name=level" json:"level,omitempty"`
	Vip      int32  `protobuf:"varint,6,opt,name=vip" json:"vip,omitempty"`
}

func (m *LoginRequest) Reset()                    { *m = LoginRequest{} }
func (m *LoginRequest) String() string            { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()               {}
func (*LoginRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *LoginRequest) GetUid() int64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *LoginRequest) GetServerId() int32 {
	if m != nil {
		return m.ServerId
	}
	return 0
}

func (m *LoginRequest) GetUname() string {
	if m != nil {
		return m.Uname
	}
	return ""
}

func (m *LoginRequest) GetFigure() int32 {
	if m != nil {
		return m.Figure
	}
	return 0
}

func (m *LoginRequest) GetLevel() int32 {
	if m != nil {
		return m.Level
	}
	return 0
}

func (m *LoginRequest) GetVip() int32 {
	if m != nil {
		return m.Vip
	}
	return 0
}

type LoginReply struct {
	Result int32 `protobuf:"varint,1,opt,name=result" json:"result,omitempty"`
}

func (m *LoginReply) Reset()                    { *m = LoginReply{} }
func (m *LoginReply) String() string            { return proto.CompactTextString(m) }
func (*LoginReply) ProtoMessage()               {}
func (*LoginReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *LoginReply) GetResult() int32 {
	if m != nil {
		return m.Result
	}
	return 0
}

// logout
type LogoutRequest struct {
	Uid int64 `protobuf:"varint,1,opt,name=uid" json:"uid,omitempty"`
}

func (m *LogoutRequest) Reset()                    { *m = LogoutRequest{} }
func (m *LogoutRequest) String() string            { return proto.CompactTextString(m) }
func (*LogoutRequest) ProtoMessage()               {}
func (*LogoutRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *LogoutRequest) GetUid() int64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

type LogoutReply struct {
	Result int32 `protobuf:"varint,1,opt,name=result" json:"result,omitempty"`
}

func (m *LogoutReply) Reset()                    { *m = LogoutReply{} }
func (m *LogoutReply) String() string            { return proto.CompactTextString(m) }
func (*LogoutReply) ProtoMessage()               {}
func (*LogoutReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *LogoutReply) GetResult() int32 {
	if m != nil {
		return m.Result
	}
	return 0
}

// chat
type ChatRequest struct {
	ChatMsg *pbcommon.ChatMessage `protobuf:"bytes,1,opt,name=chatMsg" json:"chatMsg,omitempty"`
}

func (m *ChatRequest) Reset()                    { *m = ChatRequest{} }
func (m *ChatRequest) String() string            { return proto.CompactTextString(m) }
func (*ChatRequest) ProtoMessage()               {}
func (*ChatRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *ChatRequest) GetChatMsg() *pbcommon.ChatMessage {
	if m != nil {
		return m.ChatMsg
	}
	return nil
}

type ChatReply struct {
	Result int32 `protobuf:"varint,1,opt,name=result" json:"result,omitempty"`
}

func (m *ChatReply) Reset()                    { *m = ChatReply{} }
func (m *ChatReply) String() string            { return proto.CompactTextString(m) }
func (*ChatReply) ProtoMessage()               {}
func (*ChatReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *ChatReply) GetResult() int32 {
	if m != nil {
		return m.Result
	}
	return 0
}

// push rpc
type PushMessageRequest struct {
	ServerId int32 `protobuf:"varint,1,opt,name=serverId" json:"serverId,omitempty"`
}

func (m *PushMessageRequest) Reset()                    { *m = PushMessageRequest{} }
func (m *PushMessageRequest) String() string            { return proto.CompactTextString(m) }
func (*PushMessageRequest) ProtoMessage()               {}
func (*PushMessageRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *PushMessageRequest) GetServerId() int32 {
	if m != nil {
		return m.ServerId
	}
	return 0
}

type PushMessageReply struct {
	Type    int32                 `protobuf:"varint,1,opt,name=type" json:"type,omitempty"`
	Time    int32                 `protobuf:"varint,2,opt,name=time" json:"time,omitempty"`
	Uid     int32                 `protobuf:"varint,3,opt,name=uid" json:"uid,omitempty"`
	ChatMsg *pbcommon.ChatMessage `protobuf:"bytes,4,opt,name=chatMsg" json:"chatMsg,omitempty"`
}

func (m *PushMessageReply) Reset()                    { *m = PushMessageReply{} }
func (m *PushMessageReply) String() string            { return proto.CompactTextString(m) }
func (*PushMessageReply) ProtoMessage()               {}
func (*PushMessageReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *PushMessageReply) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *PushMessageReply) GetTime() int32 {
	if m != nil {
		return m.Time
	}
	return 0
}

func (m *PushMessageReply) GetUid() int32 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *PushMessageReply) GetChatMsg() *pbcommon.ChatMessage {
	if m != nil {
		return m.ChatMsg
	}
	return nil
}

func init() {
	proto.RegisterType((*KeepAliveRequest)(nil), "ctrpcpt.KeepAliveRequest")
	proto.RegisterType((*KeepAliveReply)(nil), "ctrpcpt.KeepAliveReply")
	proto.RegisterType((*LoginRequest)(nil), "ctrpcpt.LoginRequest")
	proto.RegisterType((*LoginReply)(nil), "ctrpcpt.LoginReply")
	proto.RegisterType((*LogoutRequest)(nil), "ctrpcpt.LogoutRequest")
	proto.RegisterType((*LogoutReply)(nil), "ctrpcpt.LogoutReply")
	proto.RegisterType((*ChatRequest)(nil), "ctrpcpt.ChatRequest")
	proto.RegisterType((*ChatReply)(nil), "ctrpcpt.ChatReply")
	proto.RegisterType((*PushMessageRequest)(nil), "ctrpcpt.PushMessageRequest")
	proto.RegisterType((*PushMessageReply)(nil), "ctrpcpt.PushMessageReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for CTRPC service

type CTRPCClient interface {
	// KeepAlive rpc
	KeepAlive(ctx context.Context, in *KeepAliveRequest, opts ...grpc.CallOption) (*KeepAliveReply, error)
	// Login
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginReply, error)
	// Logout
	Logout(ctx context.Context, in *LogoutRequest, opts ...grpc.CallOption) (*LogoutReply, error)
	// chat
	Chat(ctx context.Context, in *ChatRequest, opts ...grpc.CallOption) (*ChatReply, error)
	// PushStream
	PushStream(ctx context.Context, in *PushMessageRequest, opts ...grpc.CallOption) (CTRPC_PushStreamClient, error)
}

type cTRPCClient struct {
	cc *grpc.ClientConn
}

func NewCTRPCClient(cc *grpc.ClientConn) CTRPCClient {
	return &cTRPCClient{cc}
}

func (c *cTRPCClient) KeepAlive(ctx context.Context, in *KeepAliveRequest, opts ...grpc.CallOption) (*KeepAliveReply, error) {
	out := new(KeepAliveReply)
	err := grpc.Invoke(ctx, "/ctrpcpt.CTRPC/KeepAlive", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cTRPCClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginReply, error) {
	out := new(LoginReply)
	err := grpc.Invoke(ctx, "/ctrpcpt.CTRPC/Login", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cTRPCClient) Logout(ctx context.Context, in *LogoutRequest, opts ...grpc.CallOption) (*LogoutReply, error) {
	out := new(LogoutReply)
	err := grpc.Invoke(ctx, "/ctrpcpt.CTRPC/Logout", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cTRPCClient) Chat(ctx context.Context, in *ChatRequest, opts ...grpc.CallOption) (*ChatReply, error) {
	out := new(ChatReply)
	err := grpc.Invoke(ctx, "/ctrpcpt.CTRPC/Chat", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cTRPCClient) PushStream(ctx context.Context, in *PushMessageRequest, opts ...grpc.CallOption) (CTRPC_PushStreamClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_CTRPC_serviceDesc.Streams[0], c.cc, "/ctrpcpt.CTRPC/PushStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &cTRPCPushStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CTRPC_PushStreamClient interface {
	Recv() (*PushMessageReply, error)
	grpc.ClientStream
}

type cTRPCPushStreamClient struct {
	grpc.ClientStream
}

func (x *cTRPCPushStreamClient) Recv() (*PushMessageReply, error) {
	m := new(PushMessageReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for CTRPC service

type CTRPCServer interface {
	// KeepAlive rpc
	KeepAlive(context.Context, *KeepAliveRequest) (*KeepAliveReply, error)
	// Login
	Login(context.Context, *LoginRequest) (*LoginReply, error)
	// Logout
	Logout(context.Context, *LogoutRequest) (*LogoutReply, error)
	// chat
	Chat(context.Context, *ChatRequest) (*ChatReply, error)
	// PushStream
	PushStream(*PushMessageRequest, CTRPC_PushStreamServer) error
}

func RegisterCTRPCServer(s *grpc.Server, srv CTRPCServer) {
	s.RegisterService(&_CTRPC_serviceDesc, srv)
}

func _CTRPC_KeepAlive_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KeepAliveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CTRPCServer).KeepAlive(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ctrpcpt.CTRPC/KeepAlive",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CTRPCServer).KeepAlive(ctx, req.(*KeepAliveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CTRPC_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CTRPCServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ctrpcpt.CTRPC/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CTRPCServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CTRPC_Logout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogoutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CTRPCServer).Logout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ctrpcpt.CTRPC/Logout",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CTRPCServer).Logout(ctx, req.(*LogoutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CTRPC_Chat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CTRPCServer).Chat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ctrpcpt.CTRPC/Chat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CTRPCServer).Chat(ctx, req.(*ChatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CTRPC_PushStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PushMessageRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CTRPCServer).PushStream(m, &cTRPCPushStreamServer{stream})
}

type CTRPC_PushStreamServer interface {
	Send(*PushMessageReply) error
	grpc.ServerStream
}

type cTRPCPushStreamServer struct {
	grpc.ServerStream
}

func (x *cTRPCPushStreamServer) Send(m *PushMessageReply) error {
	return x.ServerStream.SendMsg(m)
}

var _CTRPC_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ctrpcpt.CTRPC",
	HandlerType: (*CTRPCServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "KeepAlive",
			Handler:    _CTRPC_KeepAlive_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _CTRPC_Login_Handler,
		},
		{
			MethodName: "Logout",
			Handler:    _CTRPC_Logout_Handler,
		},
		{
			MethodName: "Chat",
			Handler:    _CTRPC_Chat_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PushStream",
			Handler:       _CTRPC_PushStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "ctrpc.proto",
}

func init() { proto.RegisterFile("ctrpc.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 438 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0x4b, 0x8e, 0xd3, 0x40,
	0x10, 0xc5, 0xf1, 0x67, 0x48, 0x79, 0x40, 0x51, 0x33, 0x33, 0x78, 0xcc, 0x66, 0x68, 0x3e, 0xca,
	0xca, 0x8c, 0x82, 0x90, 0x58, 0x21, 0x8d, 0xb2, 0x01, 0x91, 0x48, 0x91, 0xe1, 0x02, 0x8e, 0x53,
	0x38, 0x96, 0xec, 0xb8, 0xb1, 0xdb, 0x96, 0xb2, 0xe0, 0x12, 0xdc, 0x86, 0xdb, 0xa1, 0xfe, 0xd8,
	0x71, 0x3e, 0x44, 0xec, 0xaa, 0x5e, 0xbf, 0x57, 0xd5, 0x5d, 0xaf, 0x1a, 0xdc, 0x98, 0x97, 0x2c,
	0x0e, 0x58, 0x59, 0xf0, 0x82, 0x5c, 0xc8, 0x84, 0x71, 0xff, 0x32, 0x2e, 0xf2, 0xbc, 0xd8, 0x28,
	0x98, 0xbe, 0x85, 0xd1, 0x57, 0x44, 0xf6, 0x90, 0xa5, 0x0d, 0x86, 0xf8, 0xb3, 0xc6, 0x8a, 0x13,
	0x02, 0x16, 0x4f, 0x73, 0xf4, 0x8c, 0x3b, 0x63, 0x6c, 0x86, 0x32, 0xa6, 0x63, 0x78, 0xda, 0xe3,
	0xb1, 0x6c, 0x4b, 0x6e, 0xc0, 0x29, 0xb1, 0xaa, 0x33, 0x2e, 0x79, 0x76, 0xa8, 0x33, 0xfa, 0xdb,
	0x80, 0xcb, 0x59, 0x91, 0xa4, 0x9b, 0xb6, 0xdc, 0x08, 0xcc, 0x3a, 0x5d, 0xe9, 0x6a, 0x22, 0x24,
	0x3e, 0x3c, 0xae, 0xb0, 0x6c, 0xb0, 0xfc, 0xb2, 0xf2, 0x06, 0x52, 0xdc, 0xe5, 0xe4, 0x0a, 0xec,
	0x7a, 0x13, 0xe5, 0xe8, 0x99, 0x77, 0xc6, 0x78, 0x18, 0xaa, 0x44, 0x34, 0xfb, 0x91, 0x26, 0x75,
	0x89, 0x9e, 0xa5, 0x9a, 0xa9, 0x4c, 0xb0, 0x33, 0x6c, 0x30, 0xf3, 0x6c, 0x09, 0xab, 0x44, 0x74,
	0x6c, 0x52, 0xe6, 0x39, 0x12, 0x13, 0x21, 0x7d, 0x0d, 0xa0, 0xef, 0x74, 0xee, 0xea, 0x2f, 0xe1,
	0xc9, 0xac, 0x48, 0x8a, 0x9a, 0xff, 0xf3, 0xea, 0xf4, 0x0d, 0xb8, 0x2d, 0xe5, 0x5c, 0xa5, 0x4f,
	0xe0, 0x4e, 0xd7, 0x51, 0x57, 0xe7, 0x1d, 0x5c, 0xc4, 0xeb, 0x88, 0xcf, 0xab, 0x44, 0xf2, 0xdc,
	0xc9, 0x75, 0xc0, 0x96, 0xda, 0x07, 0xc1, 0x9b, 0x63, 0x55, 0x45, 0x09, 0x86, 0x2d, 0x8b, 0xbe,
	0x82, 0xa1, 0xd2, 0x9f, 0x6b, 0x72, 0x0f, 0x64, 0x51, 0x57, 0xeb, 0x56, 0xac, 0x7b, 0xf5, 0x87,
	0x6b, 0xec, 0x0f, 0x97, 0xfe, 0x82, 0xd1, 0x9e, 0x42, 0x54, 0x17, 0x6e, 0x6f, 0x19, 0x6a, 0xae,
	0x8c, 0xbb, 0x0d, 0x18, 0x68, 0x2c, 0xcd, 0xb1, 0x9d, 0x85, 0xa9, 0x86, 0x2a, 0x6c, 0xec, 0xbd,
	0xca, 0xfa, 0x9f, 0x57, 0x4d, 0xfe, 0x0c, 0xc0, 0x9e, 0x7e, 0x0f, 0x17, 0x53, 0xf2, 0x00, 0xc3,
	0x6e, 0x9d, 0xc8, 0x6d, 0xa0, 0x77, 0x33, 0x38, 0x5c, 0x45, 0xff, 0xf9, 0xa9, 0x23, 0x96, 0x6d,
	0xe9, 0x23, 0xf2, 0x01, 0x6c, 0x69, 0x29, 0xb9, 0xee, 0x38, 0xfd, 0xb5, 0xf3, 0x9f, 0x1d, 0xc2,
	0x4a, 0xf6, 0x11, 0x1c, 0x65, 0x20, 0xb9, 0xe9, 0x13, 0x76, 0xa6, 0xfb, 0x57, 0x47, 0xb8, 0x52,
	0x4e, 0xc0, 0x12, 0xaf, 0x22, 0xbb, 0xf3, 0x9e, 0xc5, 0x3e, 0x39, 0x40, 0x95, 0xe6, 0x33, 0x80,
	0x18, 0xf8, 0x37, 0x5e, 0x62, 0x94, 0x93, 0x17, 0x1d, 0xe7, 0xd8, 0x37, 0xff, 0xf6, 0xf4, 0xa1,
	0xac, 0x73, 0x6f, 0x2c, 0x1d, 0xf9, 0x5f, 0xdf, 0xff, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x2b, 0xbf,
	0x05, 0x66, 0xd5, 0x03, 0x00, 0x00,
}
