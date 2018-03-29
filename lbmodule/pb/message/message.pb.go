// Code generated by protoc-gen-go. DO NOT EDIT.
// source: message.proto

package message

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Message struct {
	Command       int32             `protobuf:"varint,1,opt,name=command" json:"command,omitempty"`
	Uid           int64             `protobuf:"varint,2,opt,name=uid" json:"uid,omitempty"`
	ChatRequest   *ChatMsgRequest   `protobuf:"bytes,3,opt,name=chatRequest" json:"chatRequest,omitempty"`
	ChatReply     *ChatMsgReply     `protobuf:"bytes,4,opt,name=chatReply" json:"chatReply,omitempty"`
	LoginRequest  *LoginMsgRequest  `protobuf:"bytes,5,opt,name=loginRequest" json:"loginRequest,omitempty"`
	LoginReply    *LoginMsgReply    `protobuf:"bytes,6,opt,name=loginReply" json:"loginReply,omitempty"`
	LogoutRequest *LogOutMsgRequest `protobuf:"bytes,7,opt,name=logoutRequest" json:"logoutRequest,omitempty"`
	LogoutReply   *LogOutMsgReply   `protobuf:"bytes,8,opt,name=logoutReply" json:"logoutReply,omitempty"`
	// 服务器推送消息
	ChatmsgPush *ChatMsgPush `protobuf:"bytes,10001,opt,name=chatmsgPush" json:"chatmsgPush,omitempty"`
}

func (m *Message) Reset()                    { *m = Message{} }
func (m *Message) String() string            { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()               {}
func (*Message) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *Message) GetCommand() int32 {
	if m != nil {
		return m.Command
	}
	return 0
}

func (m *Message) GetUid() int64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *Message) GetChatRequest() *ChatMsgRequest {
	if m != nil {
		return m.ChatRequest
	}
	return nil
}

func (m *Message) GetChatReply() *ChatMsgReply {
	if m != nil {
		return m.ChatReply
	}
	return nil
}

func (m *Message) GetLoginRequest() *LoginMsgRequest {
	if m != nil {
		return m.LoginRequest
	}
	return nil
}

func (m *Message) GetLoginReply() *LoginMsgReply {
	if m != nil {
		return m.LoginReply
	}
	return nil
}

func (m *Message) GetLogoutRequest() *LogOutMsgRequest {
	if m != nil {
		return m.LogoutRequest
	}
	return nil
}

func (m *Message) GetLogoutReply() *LogOutMsgReply {
	if m != nil {
		return m.LogoutReply
	}
	return nil
}

func (m *Message) GetChatmsgPush() *ChatMsgPush {
	if m != nil {
		return m.ChatmsgPush
	}
	return nil
}

func init() {
	proto.RegisterType((*Message)(nil), "message.Message")
}

func init() { proto.RegisterFile("message.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 271 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0x4d, 0x4e, 0xc3, 0x30,
	0x10, 0x85, 0x15, 0x42, 0x1b, 0x98, 0x52, 0x09, 0x8d, 0xf8, 0x31, 0x5d, 0x45, 0xac, 0xb2, 0xea,
	0x82, 0x4a, 0x20, 0x24, 0x24, 0x16, 0x6c, 0x29, 0x20, 0xdf, 0x20, 0xb4, 0x96, 0x53, 0x29, 0xa9,
	0x43, 0x1d, 0x2f, 0x7a, 0x0c, 0x6e, 0xc6, 0x91, 0xd0, 0xd8, 0x49, 0x99, 0x86, 0xee, 0x3c, 0x9e,
	0xef, 0xcb, 0x93, 0x5f, 0x60, 0x5c, 0x29, 0x6b, 0x73, 0xad, 0xa6, 0xf5, 0xc6, 0x34, 0x06, 0x93,
	0x76, 0x9c, 0xc0, 0xa2, 0xc8, 0x9b, 0x70, 0x39, 0x01, 0x67, 0xd5, 0x26, 0x9c, 0x6f, 0x7f, 0x62,
	0x48, 0xe6, 0x81, 0x41, 0x01, 0xc9, 0xc2, 0x54, 0x55, 0xbe, 0x5e, 0x8a, 0x28, 0x8d, 0xb2, 0x81,
	0xec, 0x46, 0x3c, 0x87, 0xd8, 0xad, 0x96, 0xe2, 0x28, 0x8d, 0xb2, 0x58, 0xd2, 0x11, 0x1f, 0x61,
	0x44, 0x5f, 0x94, 0xea, 0xcb, 0x29, 0xdb, 0x88, 0x38, 0x8d, 0xb2, 0xd1, 0xdd, 0xf5, 0xb4, 0x4b,
	0x7f, 0x29, 0xf2, 0x66, 0x6e, 0x75, 0xbb, 0x96, 0x9c, 0xc5, 0x19, 0x9c, 0x86, 0xb1, 0x2e, 0xb7,
	0xe2, 0xd8, 0x8b, 0x97, 0xff, 0xc5, 0xba, 0xdc, 0xca, 0x3f, 0x0e, 0x9f, 0xe0, 0xac, 0x34, 0x7a,
	0xb5, 0xee, 0x02, 0x07, 0xde, 0x13, 0x3b, 0xef, 0x95, 0x96, 0x2c, 0x71, 0x8f, 0xc6, 0x7b, 0x80,
	0x76, 0xa6, 0xcc, 0xa1, 0x77, 0xaf, 0x0e, 0xb8, 0x14, 0xca, 0x48, 0x7c, 0x86, 0x71, 0x69, 0xb4,
	0x71, 0xbb, 0x77, 0x26, 0x5e, 0xbd, 0xe1, 0xea, 0xbb, 0xe3, 0x2f, 0xdd, 0xe7, 0xa9, 0xa6, 0xee,
	0x82, 0x92, 0x4f, 0x7a, 0x35, 0x31, 0x9d, 0xa2, 0x39, 0x8b, 0x0f, 0xa1, 0xe1, 0xca, 0xea, 0x0f,
	0x67, 0x0b, 0xf1, 0xfd, 0xe6, 0xdd, 0x8b, 0x7e, 0x53, 0xb4, 0x94, 0x9c, 0xfc, 0x1c, 0xfa, 0x3f,
	0x3b, 0xfb, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x48, 0x7c, 0x38, 0x76, 0x0b, 0x02, 0x00, 0x00,
}
