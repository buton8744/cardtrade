package net

import (
	"fmt"
	"net"
	"bytes"
)

type Client struct {
	ID string
	Incoming chan string
	Conn net.Conn
	Quit chan bool
}

// func (c *Client) Read(buffer []byte) Packet, err {
// 	bytesRead, err := c.Conn.Read(buffer)
// 	if error != nil {
// 		c.Close()
// 		Log(error)
// 		return nil, 1
// 	}
// 	Log("Read", bytesRead, "bytes")
// 	NewPacket := &Packet{}
// 	NewPacket.Read(buffer)
// 	return NewPacket, nil
// }

func (c *Client) Read(buffer []byte) bool {
	bytesRead, err := c.Conn.Read(buffer)
	if err != nil {
		c.Close()
		Log(err)
		return false
	}
	Log("Read", bytesRead, "bytes")
	return true
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
