package net

import (
	"fmt"
	"code.google.com/p/goprotobuf/proto"
)

type Packet struct {
	msg proto.Message
	msize int
}

func (p *Packet) Read(buffer []byte) {
	err := proto.Unmarshal(buffer, p.msg)
	if err != nil {
	}
}

func (p *Packet) Write(buffer []byte) {
	// 어떤 작업이 필요할까?
	data, err := proto.Marshal(p.msg)
	if err != nil {
	}
	fmt.Println(data)
}
