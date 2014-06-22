package main

import (
	"container/list"
	"ex_cardtrade/net"
)

func main() {
	app := &net.ServerApp{}
	app.Run("0.0.0.0", "8899", make(chan *net.Packet), list.New())
}
