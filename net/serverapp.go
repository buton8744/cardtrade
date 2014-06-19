package net

import (
	"net"
	"bytes"
	"container/list"
)

const (
	BUFFER_SIZE = 2048
)

type ServerApp struct {
	IpAddr string
	Port string
	Chat chan string
	ClientList *list.List
}

func (s *ServerApp) Run() {
	service := s.IpAddr + ":" + s.Port
	tcpAddr, err := net.ResolveTCPAddr("tcp", service)
	if err != nil {
		Log("Error could not resolve address")
	} else {
		Log("Card Trade Server open!!", service)
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

	id := string(buffer[0:bytesRead])
	newClient := &Client{id, make(chan string), conn,  make(chan bool)}

	go s.ClientReader(newClient)
	go s.ClientSender(newClient)
	s.ClientList.PushBack(*newClient)
	s.Chat <- string(id + " has joined the chat")
}

func (s *ServerApp) ClientReader(client *Client) {
	buffer := make([]byte, BUFFER_SIZE)
	
	for client.Read(buffer) {
		// 여기서 Dispatch를 해줘야 할지도
		if bytes.Equal(buffer, []byte("/quit")) {
			client.Close()
			break
		}
		Log("ClientReader", client.ID, "> ", string(buffer))
		send := client.ID+"> " + string(buffer)
		s.Chat <- send
		for i := 0; i < BUFFER_SIZE; i++ {
			buffer[i] = 0x00
		}
	}

	s.Chat <- client.ID + " has left chat"
	Log("ClientReader stopped for ", client.ID)
}

// func (s *ServerApp) ClientReader(client *net.Client) {
// 	buffer := make([]byte, BUFFER_SIZE)

// 	packet, err := client.Read(buffer)
// 	if err != nil {
// 		// 여기서 packet에 대한 cmd 처리 해줘야 함
// 		if bytes.Equal(buffer, []byte("/quit")) {
// 			client.Close()
// 			break
// 		}
// 		Log("ClientReader", client.ID, "> ", stirng(buffer))
// 		send := client.ID+"> " + string(buffer)
// 		s.Chat <- send
// 		for i := 0; i < BUFFER_SIZE; i++ {
// 			buffer[i] = 0x00
// 		}
// 	}

// 	s.Chat <- client.ID + "has left chat"
// 	Log("ClientReader stopped for ", client.ID)
// }

func (s *ServerApp) ClientSender(client *Client) {
	for {
		select {
		case buffer := <- client.Incoming:
			count := 0
			for i := 0; i < len(buffer); i++ {
				if buffer[i] == 0x00 {
					break
				}
				count++
			}
			Log("ClientSender sending ", string(buffer), " to ", client.ID, "Send Size: ", count)
			client.Conn.Write([]byte(buffer)[0:count])
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
		case input := <- s.Chat:
			for entry := s.ClientList.Front(); entry != nil; entry = entry.Next() {
				client := entry.Value.(Client)
				client.Incoming <- input
			}
		default:
		}
	}
}
