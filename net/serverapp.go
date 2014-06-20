package net

import (
	"net"
	"container/list"
)

const (
	BUFFER_SIZE = 2048
)

type ServerApp struct {
	IpAddr string
	Port string
	BroadCastPacket chan *Packet
	ClientList *list.List
	db Mydb
}

func (s *ServerApp) Run() {
	service := s.IpAddr + ":" + s.Port
	tcpAddr, err := net.ResolveTCPAddr("tcp", service)
	if err != nil {
		Log("Error could not resolve address")
	} else {
		Log("Card Trade Server open!!", service)
		if !s.db.Connect("root", "asdf1234", "127.0.0.1:3366") {
			Log("Database Connect Failed")
			return
		}
		Log("Database Connected")
		netListen, err := net.Listen(tcpAddr.Network(), tcpAddr.String())
		if err != nil {
			Log(err)
		} else {
			defer netListen.Close()
			go s.Broadcast()
			for {
				Log("Waiting for clients")
				connection, err := netListen.Accept()
				if err != nil {
					Log("Client error: ", err)
				} else {
					go s.Service(connection)
				}
			}
		}
	}
}

func (s *ServerApp) Service(conn net.Conn) {
	buffer := make([]byte, 1024)
	bytesRead, error := conn.Read(buffer)
	if error != nil {
		Log("client connection error: ", error)
	}

	// 처음 오는건 무조건 Login이라 판단하고 인증 불가시 튕겨냄
	recv_packet := ReadPacket(buffer)
	if recv_packet.Type != "SIGNIN_REQ" {
		// 처음 오는 패킷이 로그인이 아니면 무시
		LOG()
		return
	}
	login_packet, err := recv_packet.RecvSigninReqPacket()
	// 여기서 디비랑 비교가 필요함
	rows, err := s.db.Query("select password from user where userid = ?", login_packet.Id)
	if err != nil{
	}
	var password string
	rows.Scan(&password)
	// 그리고 인증 됬는지 아닌지 리턴
	if password != login_packet.Password {
		send_packet := MakePacket("SIGNIN_ACK", false, login_packet.Id)
		send_packet.Byte(buffer)
		conn.Write(buffer)
		return
	}
	newClient := &Client{id, make(chan Packet), conn,  make(chan bool)}

	go s.ClientReader(newClient)
	go s.ClientSender(newClient)
	s.ClientList.PushBack(*newClient)
	send_packet := MakePacket("CHATING_REQ", string(id + " has logged in"))
	s.BroadCastPacket <- send_packet
}

func (s *ServerApp) ClientReader(client *Client) {
	buffer := make([]byte, BUFFER_SIZE)

	packet, err := client.Read(buffer)
	if err != nil {
		client.Close()
		send_packet := MakePacket("CHATING_REQ", string(client.ID + "has logged out"))
		s.BroadCastPacket <- send_packet
		Log("ClientReader stopped for ", client.ID)
	}
	// 여기서 packet에 대한 cmd 처리를 해줘야 함
	go s.CmdWoker(client, packet)
	for i := 0; i < BUFFER_SIZE; i++ {
		buffer[i] = 0x00
	}
}

func (s *ServerApp) ClientSender(client *Client) {
	for {
		select {
			case packet := <- client.Incoming:
			client.Write(packet)
		case <-client.Quit:
			Log("Client ", client.ID, "Quitting")
			client.Conn.Close()
			s.ClientRemove(client)
			break
		}
	}
}

func (s *ServerApp) ClientRemove(other *Client) {
	for entry := s.ClientList.Front(); entry != nil; entry = entry.Next() {
		client := entry.Value.(Client)
		if client.Equal(other) {
			Log("RemoveMe: ", client.ID)
			s.ClientList.Remove(entry)
		}
	}
}

func (s *ServerApp) Broadcast() {
	for {
		select {
		case packet := <- s.BroadCastPacket:
			for entry := s.ClientList.Front(); entry != nil; entry = entry.Next() {
				client := entry.Value.(Client)
				client.Incoming <- packet
			}
		default:
		}
	}
}

func (s *ServerApp) CmdWorker(client *Client, packet *Packet) {
	LOG(packet.Type, " recevied")
	switch packet.Type {
	case "WITHDRAW_REQ":
		recv_packet, err := packet.RecvWithDrawReqPacket()
		rows, err := s.db.Query("select password from user where userid = ?", Client.ID)
		var result_flag bool
		if err != nil {
			result_flag = false
		} else {
			var password string
			rows.Scan(&password)
			if recv_packet.Password != password {
				result_flag = false
			} else {
				result_flag = true
			}
		}
		send_packet := MakePacket("WITHDRAW_ACK", result_flag, Client.ID)
		client.Incoming <- send_packet
		
	case "LOGOUT_REQ":
		send_packet := MakePacket("LOGOUT_ACK", true, Client.ID)
		client.Incoming <- send_packet
		client.Close()
		
	case "CHATING_REQ":
		recv_packet, err := packet.RecvChatingReqPacket()
		send_packet := MakePacket("CHATING_ACK", client.ID, recv_packet.Message)
		s.BroadCastPacket <- send_packet
	}
}

func MakePacket(mtype PacketType, args ...interface{}) *Packet {
	send_packet := &Packet{}
	switch mtype{
	case packet.PacketType_SIGNUPREQ:
		send_packet.SendSignUpReqPacket(args[0], args[1])
	case packet.PacketType_SIGNUPACK:
		send_packet.SendSignUpAckPacket(args[0], args[1])
	case packet.PacketType_SIGNINREQ:
		send_packet.SendSignInReqPacket(args[0], args[1])
	case packet.PacketType_SIGNINACK:
		send_packet.SendSignInAckPacket(args[0], args[1])
	case packet.PacketType_WITHDRAWREQ:
		send_packet.SendWithDrawReqPacket(args[0])
	case packet.PacketType_WITHDRAWACK:
		send_packet.SendWithDrawAckPacket(args[0], args[1])
	case packet.PacketType_LOGOUTREQ:
		send_packet.SendLogoutReqPacket()
	case packet.PacketType_LOGOUTACK:
		send_packet.SendLogoutAckPacket(args[0], args[1])
	case packet.PacketType_CHATINGREQ:
		send_packet.SendChatingReqPacket(args[0], args[1])
	case packet.PacketType_CHATINGACK:
		send_packet.SendChatingAckPacket(args[0], args[1])
	}
	return send_packet
}

func ReadPacket(buffer []byte) *Packet {
	recv_packet := &Packet{}
	recv_packet.Read(buffer)
	return recv_packet
}
