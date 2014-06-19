package main

import (
	"container/list"
	"ex_cardtrade/net"
)

func main() {
	app := &net.ServerApp{"0.0.0.0", "8899", make(chan string), list.New()}
	app.Run()
}
