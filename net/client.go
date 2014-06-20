package net

import (
	"fmt"
	"net"
	"bytes"
)

type Client struct {
	ID string
	Incoming chan *Packet
	Conn net.Conn
	Quit chan bool
}

func (c *Client) Read(buffer []byte) (*Packet, error) {
	bytesRead, err := c.Conn.Read(buffer)
	if err != nil {
		c.Close()
		Log(err)
		return nil, err
	}
	Log("Read", bytesRead, "bytes")
	NewPacket := &Packet{}
	err = NewPacket.Read(buffer)
	if err != nil {
		return nil, err 
	}
	return NewPacket, nil
}

func (c *Client) Write(packet *Packet) (int , error) {
	buffer := make([]byte, BUFFER_SIZE)
	err := packet.Byte(buffer)
	if err != nil {
		return 0, err
	}
	
	count := 0
	for i := 0; i < len(buffer); i++ {
		if buffer[i] == 0x00 {
			break
		}
		count++
	}	
	BytesWrite, err := c.Conn.Write([]byte(buffer)[0:count])
	Log("Send to ", c.ID, "Bytes Size: ", BytesWrite, "Send Size: ", count)
	return BytesWrite, nil
}


func (c *Client) Close() {
	c.Quit <- true
	c.Conn.Close()
	Log("Client Close: ", c.ID)
}

func (c *Client) Equal(other *Client) bool {
	if bytes.Equal([]byte(c.ID), []byte(other.ID)){
		if c.Conn == other.Conn {
			return true
		}
	}
	return false
}

func Log(v ...interface{}) {
	fmt.Println(v...)
}
