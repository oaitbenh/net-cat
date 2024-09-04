package main

import (
	"fmt"
	"net"

	Functions "TCP_Chat/utils"
)

func main() {
	s := new(Functions.Server)
	s.Addr = ":3030"
	var err error
	s.Listener, err = net.Listen("tcp", s.Addr)
	defer s.Listener.Close()
	if err != nil {
		fmt.Println("Error in Listen : ", err)
		return
	} else {
		fmt.Println("TCP-Chat Started at Port 3030!")
	}
	s.Messages = []string{}
	s.AcceptLoop()
}
