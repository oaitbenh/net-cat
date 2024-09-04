package TCP_Chat

import "net"

type MsgStruct struct {
	Addr    string
	Message []byte
}

type Server struct {
	Addr     string
	Listener net.Listener
	Conns    []net.Conn
	Messages []string
}

var Users = map[string]string{}
