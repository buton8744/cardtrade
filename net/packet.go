package net

import (
	"strings"
	"code.google.com/p/goprotobuf/proto"
	"ex_cardtrade/proto/packet"
)

type Packet struct {
	Type int32
	Pb proto.Message
}

func (p *Packet) Read(buffer []byte) error {
	err = proto.Unmarshal(buffer, p.Pb)
	if err != nil {
		return err
	}
	p.Type = packet.PacketType_value[strings.ToUpper(p.Pb.String())]
	return nil
}

func (p *Packet) MakeRecv(pb proto.Message) {
	data, err := proto.Marshal(p.Pb)
	if err != nil {
	}
	err = proto.Unmarshal(data, pb)
}

func (p *Packet) RecvSignUpReqPacket() (pb *packet.SignUpReq, err error) {
	pb := &packet.SignUpReq{}
	err := p.MakeRecv(pb)
	if err != nil {
	}
	return
}

func (p *Packet) RecvSignUpAckPacket() (pb *packet.SignUpAck, err error) {
	pb := &packet.SignUpAck{}
	err := p.MakeRecv(pb)
	if err != nil {
	}
	return
}

func (p *Packet) RecvSignInReqPacket() (pb *packet.SignInReq, err error) {
	pb = &packet.SignInReq{}
	err := p.MakeRecv(pb)
	if err != nil {
	}
	return
}

func (p *Packet) RecvSignInAckPacket() (pb *packet.SignInAck, err error) {
	pb = &packet.SignInAck{}
	err := p.MakeRecv(pb)
	if err != nil {
	}
	return
}

func (p *Packet) RecvWithDrawReq() (pb *packet.WithDrawReq, err error) {
	pb = &packet.WithDrawReq{}
	err := p.MakeRecv(pb)
	if err != nil {
	}
	return
}

func (p *Packet) RecvWithDrawAck() (pb *packet.WithDrawAck, err error) {
	pb = &packet.WithDrawAck{}
	err := p.MakeRecv(pb)
	if err != nil {
	}
	return
}

func (p *Packet) RecvLogoutReq() (pb *packet.LogoutReq, err error){
	pb = &packet.LogoutReq{}
	err := p.MakeRecv(pb)
	if err != nil {
	}
	return
}

func (p *Packet) RecvLogoutAck() (pb *packet.LogoutAck, err error) {
	pb = &packet.LogoutAck{}
	err := p.MakeRecv(pb)
	if err != nil {
	}
	return
}

func (p *Packet) RecvChatingReq() (pb *packet.ChatingReq, err error) {
	pb = &packet.ChatingReq{}
	err := p.MakeRecv(pb)
	if err != nil {
	}
	return
}

func (p *Packet) RecvChatingAck() (pb *packet.ChatingAck, err error) {
	pb = &packet.ChatingAck{}
	err := p.MakeRecv(pb)
	if err != nil {
	}
	return
}

////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////

func (p *Packet) Byte(buffer []byte) error {
	for i := 0; i < len(buffer); i++ {
		buffer[i] = 0x00
	}	
	data, err := proto.Marshal(p.Pb)
	if err != nil {
		Log("packet byte error")
		return err
	}
	buffer = []byte(data)
	return nil
}

func (p *Packet) SendSignUpReqPacket(id string, password string) {
	p.Pb = &packet.SignUpReq{
		Id: proto.String(id),
		Password: proto.String(password),
	}
	p.Type = packet.PacketType_value[strings.ToUpper(p.Pb.String())]
}

func (p *Packet) SendSignUpAckPacket(result bool, userids []string) {
	p.Pb = &packet.SignUpAck{
		Result: proto.Bool(result),
		Userids: proto.Repeatedstring(useruids),
	}
	p.Type = packet.PacketType_value[strings.ToUpper(p.Pb.String())]
}

func (p *Packet) SendSignInReqPacket(id string, password string) {
	p.Pb = &packet.SignInReq{
		Id: proto.String(id),
		Password: proto.String(password),
	}
	p.Type = packet.PacketType_value[strings.ToUpper(p.Pb.String())]
}

func (p *Packet) SendSignInAckPacket(result bool, userids []string) {
	p.Pb = &packet.SignInAck{
		Result: proto.Bool(result),
		Userids: proto.Repeatedstring(useruids),
	}
	p.Type = packet.PacketType_value[strings.ToUpper(p.Pb.String())]
}

func (p *Packet) SendWithDrawReq(password string) {
	p.Pb = &packet.WithDrawReq{
		Password: proto.String(password),
	}
	p.Type = packet.PacketType_value[strings.ToUpper(p.Pb.String())]
}

func (p *Packet) SendWithDrawAck(result bool, userid string) {
	p.Pb = &packet.WithDrawAck{
		Result: proto.Bool(result),
		Userid: proto.String(userid),
	}
	p.Type = packet.PacketType_value[strings.ToUpper(p.Pb.String())]
}

func (p *Packet) SendLogoutReq() {
	p.Pb = &packet.LogoutReq{
	}
	p.Type = packet.PacketType_value[strings.ToUpper(p.Pb.String())]
}

func (p *Packet) SendLogoutAck(result bool, userid string) {
	p.Pb = &packet.LogoutAck{
		Result: proto.Bool(result),
		Userid: proto.String(userid),
	}
	p.Type = packet.PacketType_value[strings.ToUpper(p.Pb.String())]
}

func (p *Packet) SendChatingReq(userid string, message string) {
	p.Pb = &packet.ChatingReq{
		Userid: proto.String(userid),
		Message: proto.String(message),
	}
	p.Type = packet.PacketType_value[strings.ToUpper(p.Pb.String())]
}

func (p *Packet) SendChatingAck(userid string, message string) {
	p.Pb = &packet.ChatingAck{
		Userid: proto.String(userid),
		Message: proto.String(message),
	}
	p.Type = packet.PacketType_value[strings.ToUpper(p.Pb.String())]
}
