package net

import (
	"strings"
	"code.google.com/p/goprotobuf/proto"
	"ex_cardtrade/proto/packet"
)

type Packet struct {
	Type packet.PacketType
	Pb proto.Message
}

func (p *Packet) Read(buffer []byte) error {
	err = proto.Unmarshal(buffer, p.Pb)
	if err != nil {
		return err
	}
	p.Type = packet.PacketType(packet.PacketType_value[strings.ToUpper(p.Pb.String())])
	return nil
}

func (p *Packet) RecvSignUpReq() (pb *packet.SignUpReq, err error) {
	pb = &packet.SignUpReq{}
	proto.Merge(pb, p.Pb)
	return
}

func (p *Packet) RecvSignUpAck() (pb *packet.SignUpAck, err error) {
	pb = &packet.SignUpAck{}
	proto.Merge(pb, p.Pb)
	return
}

func (p *Packet) RecvSignInReq() (pb *packet.SignInReq, err error) {
	pb = &packet.SignInReq{}
	proto.Merge(pb, p.Pb)
	return
}

func (p *Packet) RecvSignInAck() (pb *packet.SignInAck, err error) {
	pb = &packet.SignInAck{}
	proto.Merge(pb, p.Pb)
	return
}

func (p *Packet) RecvWithDrawReq() (pb *packet.WithDrawReq, err error) {
	pb = &packet.WithDrawReq{}
	proto.Merge(pb, p.Pb)
	return
}

func (p *Packet) RecvWithDrawAck() (pb *packet.WithDrawAck, err error) {
	pb = &packet.WithDrawAck{}
	proto.Merge(pb, p.Pb)
	return
}

func (p *Packet) RecvLogoutReq() (pb *packet.LogoutReq, err error){
	pb = &packet.LogoutReq{}
	proto.Merge(pb, p.Pb)
	return
}

func (p *Packet) RecvLogoutAck() (pb *packet.LogoutAck, err error) {
	pb = &packet.LogoutAck{}
	proto.Merge(pb, p.Pb)
	return
}

func (p *Packet) RecvChatingReq() (pb *packet.ChatingReq, err error) {
	pb = &packet.ChatingReq{}
	proto.Merge(pb, p.Pb)
	return
}

func (p *Packet) RecvChatingAck() (pb *packet.ChatingAck, err error) {
	pb = &packet.ChatingAck{}
	proto.Merge(pb, p.Pb)
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

func (p *Packet) SetType() {
	p.Type = packet.PacketType(packet.PacketType_value[strings.ToUpper(p.Pb.String())])
}

func (p *Packet) SendSignUpReq(id string, password string) {
	p.Pb = &packet.SignUpReq{
		Id: proto.String(id),
		Password: proto.String(password),
	}
	p.SetType()
}

func (p *Packet) SendSignUpAck(result bool, userids []string) {
	p.Pb = &packet.SignUpAck{
		Result: proto.Bool(result),
		Userids: userids,
	}
	p.SetType()
}

func (p *Packet) SendSignInReq(id string, password string) {
	p.Pb = &packet.SignInReq{
		Id: proto.String(id),
		Password: proto.String(password),
	}
	p.SetType()
}

func (p *Packet) SendSignInAck(result bool, userids []string) {
	p.Pb = &packet.SignInAck{
		Result: proto.Bool(result),
		Userids: userids,
	}
	p.SetType()
}

func (p *Packet) SendWithDrawReq(password string) {
	p.Pb = &packet.WithDrawReq{
		Password: proto.String(password),
	}
	p.SetType()
}

func (p *Packet) SendWithDrawAck(result bool, userid string) {
	p.Pb = &packet.WithDrawAck{
		Result: proto.Bool(result),
		Userid: proto.String(userid),
	}
	p.SetType()
}

func (p *Packet) SendLogoutReq() {
	p.Pb = &packet.LogoutReq{
	}
	p.SetType()
}

func (p *Packet) SendLogoutAck(result bool, userid string) {
	p.Pb = &packet.LogoutAck{
		Result: proto.Bool(result),
		Userid: proto.String(userid),
	}
	p.SetType()
}

func (p *Packet) SendChatingReq(userid string, message string) {
	p.Pb = &packet.ChatingReq{
		Userid: proto.String(userid),
		Message: proto.String(message),
	}
	p.SetType()
}

func (p *Packet) SendChatingAck(userid string, message string) {
	p.Pb = &packet.ChatingAck{
		Userid: proto.String(userid),
		Message: proto.String(message),
	}
	p.SetType()
}
