package net

import (
	"net"
	"container/list"
	"ex_cardtrade/proto/packet"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

const (
	BUFFER_SIZE = 2048
)

type ServerApp struct {
	IpAddr string
	Port string
	BroadCastPacket chan *Packet
	ClientList *list.List
	db *sql.DB
}

func (s *ServerApp) Run(ipaddr string, port string, bcp chan *Packet, cl *list.List) {
	s.IpAddr = ipaddr
	s.Port = port
	s.BroadCastPacket = bcp
	s.ClientList = cl
	
	service := s.IpAddr + ":" + s.Port

	tcpAddr, err := net.ResolveTCPAddr("tcp", service)
	if err != nil {
		Log("Error could not resolve address")
	} else {
		Log("Card Trade Server open!!", service)
		s.db, err = DBConnect("root", "18qntdhd", "cardtrade")
		if err != nil {
			Log("Database Connect Failed")
			return 	
		}
		if s.db.Ping() != nil {
			Log("Database Ping Failed")
			return
		}
		Log("Database Connected")
		defer s.db.Close()
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
	_, error := conn.Read(buffer)
	if error != nil {
		Log("client connection error: ", error)
	}

	// 처음 오는건 무조건 Login이라 판단하고 인증 불가시 튕겨냄
	recv_packet := ReadPacket(buffer)
	if recv_packet.Type != packet.PacketType_SIGNINREQ {
		// 처음 오는 패킷이 로그인이 아니면 무시
		Log("first packet must be the type of SignInReq")
		return
	}
	login_packet, err := recv_packet.RecvSignInReq()
	// 여기서 디비랑 비교가 필요함
	rows, err := s.db.Query("select password from user where userid = ?", login_packet.Id)
	if err != nil{
	}
	var password string
	rows.Scan(&password)
	// 그리고 인증 됬는지 아닌지 리턴
	if password != login_packet.GetPassword() {
		send_packet := &Packet{}
		send_packet.SendSignInAck(false, s.GetClientIds())
		send_packet.Byte(buffer)
		conn.Write(buffer)
		return
	}
	newClient := &Client{login_packet.GetId(), make(chan *Packet), conn,  make(chan bool)}

	go s.ClientReader(newClient)
	go s.ClientSender(newClient)
	s.ClientList.PushBack(*newClient)
	send_packet := &Packet{}
	send_packet.SendChatingAck(login_packet.GetId(),  string(login_packet.GetId() + " has logged in"))
	s.BroadCastPacket <- send_packet
}

func (s *ServerApp) ClientReader(client *Client) {
	buffer := make([]byte, BUFFER_SIZE)

	recv_packet, err := client.Read(buffer)
	if err != nil {
		client.Close()
		send_packet := &Packet{}
		send_packet.SendChatingAck(client.ID, string(client.ID + "has logged out"))
		s.BroadCastPacket <- send_packet
		Log("ClientReader stopped for ", client.ID)
	}
	// 여기서 packet에 대한 cmd 처리를 해줘야 함
	go s.CmdWorker(client, recv_packet)
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

func (s *ServerApp) CmdWorker(client *Client, recv_packet *Packet) {
	Log(recv_packet.Type.String(), "PacketType recevied")
	switch recv_packet.Type {
	case packet.PacketType_WITHDRAWREQ:
		result_packet, err := recv_packet.RecvWithDrawReq()
		rows, err := s.db.Query("select password from user where userid = ?", client.ID)
		var result_flag bool
		if err != nil {
			result_flag = false
		} else {
			var password string
			rows.Scan(&password)
			if result_packet.GetPassword() != password {
				result_flag = false
			} else {
				result_flag = true
			}
		}
		send_packet := &Packet{}
		send_packet.SendWithDrawAck(result_flag, client.ID)
		client.Incoming <- send_packet
		
	case packet.PacketType_LOGOUTREQ:
		send_packet := &Packet{}
		send_packet.SendLogoutAck(true, client.ID)
		client.Incoming <- send_packet
		client.Close()
		
	case packet.PacketType_CHATINGREQ:
		result_packet, _ := recv_packet.RecvChatingReq()
		send_packet := &Packet{}
		send_packet.SendChatingAck(client.ID, result_packet.GetMessage())
		s.BroadCastPacket <- send_packet
	}
}

func (s *ServerApp) GetClientIds() []string {
	useruids := make([]string, 0, s.ClientList.Len())
	count := 0
	for entry := s.ClientList.Front(); entry != nil; entry = entry.Next() {
		client := entry.Value.(Client)
		useruids[count] = client.ID
		count++
	}
	return useruids
}

func DBConnect(user string, password string, addr string) (*sql.DB, error) {
	service := user + ":" + password + "@/" + addr
	db, err := sql.Open("mysql", service)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func ReadPacket(buffer []byte) *Packet {
	recv_packet := &Packet{}
	recv_packet.Read(buffer)
	return recv_packet
}
