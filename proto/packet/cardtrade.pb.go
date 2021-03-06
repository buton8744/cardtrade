// Code generated by protoc-gen-go.
// source: cardtrade.proto
// DO NOT EDIT!

/*
Package packet is a generated protocol buffer package.

It is generated from these files:
	cardtrade.proto

It has these top-level messages:
	SignUpReq
	SignUpAck
	SignInReq
	SignInAck
	WithDrawReq
	WithDrawAck
	LogoutReq
	LogoutAck
	ChatingReq
	ChatingAck
*/
package packet

import proto "code.google.com/p/goprotobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type PacketType int32

const (
	PacketType_SIGNUPREQ   PacketType = 1
	PacketType_SIGNUPACK   PacketType = 2
	PacketType_SIGNINREQ   PacketType = 3
	PacketType_SIGNINACK   PacketType = 4
	PacketType_WITHDRAWREQ PacketType = 5
	PacketType_WITHDRAWACK PacketType = 6
	PacketType_LOGOUTREQ   PacketType = 7
	PacketType_LOGOUTACK   PacketType = 8
	PacketType_CHATINGREQ  PacketType = 9
	PacketType_CHATINGACK  PacketType = 10
)

var PacketType_name = map[int32]string{
	1:  "SIGNUPREQ",
	2:  "SIGNUPACK",
	3:  "SIGNINREQ",
	4:  "SIGNINACK",
	5:  "WITHDRAWREQ",
	6:  "WITHDRAWACK",
	7:  "LOGOUTREQ",
	8:  "LOGOUTACK",
	9:  "CHATINGREQ",
	10: "CHATINGACK",
}
var PacketType_value = map[string]int32{
	"SIGNUPREQ":   1,
	"SIGNUPACK":   2,
	"SIGNINREQ":   3,
	"SIGNINACK":   4,
	"WITHDRAWREQ": 5,
	"WITHDRAWACK": 6,
	"LOGOUTREQ":   7,
	"LOGOUTACK":   8,
	"CHATINGREQ":  9,
	"CHATINGACK":  10,
}

func (x PacketType) Enum() *PacketType {
	p := new(PacketType)
	*p = x
	return p
}
func (x PacketType) String() string {
	return proto.EnumName(PacketType_name, int32(x))
}
func (x *PacketType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(PacketType_value, data, "PacketType")
	if err != nil {
		return err
	}
	*x = PacketType(value)
	return nil
}

type SignUpReq struct {
	Id               *string `protobuf:"bytes,1,req,name=id" json:"id,omitempty"`
	Password         *string `protobuf:"bytes,2,req,name=password" json:"password,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *SignUpReq) Reset()         { *m = SignUpReq{} }
func (m *SignUpReq) String() string { return proto.CompactTextString(m) }
func (*SignUpReq) ProtoMessage()    {}

func (m *SignUpReq) GetId() string {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return ""
}

func (m *SignUpReq) GetPassword() string {
	if m != nil && m.Password != nil {
		return *m.Password
	}
	return ""
}

type SignUpAck struct {
	Result           *bool    `protobuf:"varint,1,req,name=result" json:"result,omitempty"`
	Userids          []string `protobuf:"bytes,2,rep,name=userids" json:"userids,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *SignUpAck) Reset()         { *m = SignUpAck{} }
func (m *SignUpAck) String() string { return proto.CompactTextString(m) }
func (*SignUpAck) ProtoMessage()    {}

func (m *SignUpAck) GetResult() bool {
	if m != nil && m.Result != nil {
		return *m.Result
	}
	return false
}

func (m *SignUpAck) GetUserids() []string {
	if m != nil {
		return m.Userids
	}
	return nil
}

type SignInReq struct {
	Id               *string `protobuf:"bytes,1,req,name=id" json:"id,omitempty"`
	Password         *string `protobuf:"bytes,2,req,name=password" json:"password,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *SignInReq) Reset()         { *m = SignInReq{} }
func (m *SignInReq) String() string { return proto.CompactTextString(m) }
func (*SignInReq) ProtoMessage()    {}

func (m *SignInReq) GetId() string {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return ""
}

func (m *SignInReq) GetPassword() string {
	if m != nil && m.Password != nil {
		return *m.Password
	}
	return ""
}

type SignInAck struct {
	Result           *bool    `protobuf:"varint,1,req,name=result" json:"result,omitempty"`
	Userids          []string `protobuf:"bytes,2,rep,name=userids" json:"userids,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *SignInAck) Reset()         { *m = SignInAck{} }
func (m *SignInAck) String() string { return proto.CompactTextString(m) }
func (*SignInAck) ProtoMessage()    {}

func (m *SignInAck) GetResult() bool {
	if m != nil && m.Result != nil {
		return *m.Result
	}
	return false
}

func (m *SignInAck) GetUserids() []string {
	if m != nil {
		return m.Userids
	}
	return nil
}

type WithDrawReq struct {
	Password         *string `protobuf:"bytes,1,req,name=password" json:"password,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *WithDrawReq) Reset()         { *m = WithDrawReq{} }
func (m *WithDrawReq) String() string { return proto.CompactTextString(m) }
func (*WithDrawReq) ProtoMessage()    {}

func (m *WithDrawReq) GetPassword() string {
	if m != nil && m.Password != nil {
		return *m.Password
	}
	return ""
}

type WithDrawAck struct {
	Result           *bool   `protobuf:"varint,1,req,name=result" json:"result,omitempty"`
	Userid           *string `protobuf:"bytes,2,opt,name=userid" json:"userid,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *WithDrawAck) Reset()         { *m = WithDrawAck{} }
func (m *WithDrawAck) String() string { return proto.CompactTextString(m) }
func (*WithDrawAck) ProtoMessage()    {}

func (m *WithDrawAck) GetResult() bool {
	if m != nil && m.Result != nil {
		return *m.Result
	}
	return false
}

func (m *WithDrawAck) GetUserid() string {
	if m != nil && m.Userid != nil {
		return *m.Userid
	}
	return ""
}

type LogoutReq struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *LogoutReq) Reset()         { *m = LogoutReq{} }
func (m *LogoutReq) String() string { return proto.CompactTextString(m) }
func (*LogoutReq) ProtoMessage()    {}

type LogoutAck struct {
	Result           *bool   `protobuf:"varint,1,req,name=result" json:"result,omitempty"`
	Userid           *string `protobuf:"bytes,2,opt,name=userid" json:"userid,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *LogoutAck) Reset()         { *m = LogoutAck{} }
func (m *LogoutAck) String() string { return proto.CompactTextString(m) }
func (*LogoutAck) ProtoMessage()    {}

func (m *LogoutAck) GetResult() bool {
	if m != nil && m.Result != nil {
		return *m.Result
	}
	return false
}

func (m *LogoutAck) GetUserid() string {
	if m != nil && m.Userid != nil {
		return *m.Userid
	}
	return ""
}

type ChatingReq struct {
	Userid           *string `protobuf:"bytes,1,req,name=userid" json:"userid,omitempty"`
	Message          *string `protobuf:"bytes,2,req,name=message" json:"message,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ChatingReq) Reset()         { *m = ChatingReq{} }
func (m *ChatingReq) String() string { return proto.CompactTextString(m) }
func (*ChatingReq) ProtoMessage()    {}

func (m *ChatingReq) GetUserid() string {
	if m != nil && m.Userid != nil {
		return *m.Userid
	}
	return ""
}

func (m *ChatingReq) GetMessage() string {
	if m != nil && m.Message != nil {
		return *m.Message
	}
	return ""
}

type ChatingAck struct {
	Userid           *string `protobuf:"bytes,1,req,name=userid" json:"userid,omitempty"`
	Message          *string `protobuf:"bytes,2,req,name=message" json:"message,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ChatingAck) Reset()         { *m = ChatingAck{} }
func (m *ChatingAck) String() string { return proto.CompactTextString(m) }
func (*ChatingAck) ProtoMessage()    {}

func (m *ChatingAck) GetUserid() string {
	if m != nil && m.Userid != nil {
		return *m.Userid
	}
	return ""
}

func (m *ChatingAck) GetMessage() string {
	if m != nil && m.Message != nil {
		return *m.Message
	}
	return ""
}

func init() {
	proto.RegisterEnum("packet.PacketType", PacketType_name, PacketType_value)
}
