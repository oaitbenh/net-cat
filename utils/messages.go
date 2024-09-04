package TCP_Chat

import (
	"errors"
	"fmt"
	"net"
	"sync"
)

func (s *Server) GlobalMessage(Conn net.Conn, msg []byte) {
	for _, c := range s.Conns {
		if c.RemoteAddr().String() != Conn.LocalAddr().String() {
			c.Write(msg)
		}
	}
}

func (s *Server) GetMessage(Conn net.Conn, mutex *sync.Mutex) {
	defer Conn.Close()
	Conn.Write([]byte("<< Useranme Condition Min Len 4 char Max len 20 >>\n"))
	Conn.Write([]byte("Print Your Username : "))
	Name := []byte{}
	index, err := Conn.Read(Name)
	if err != nil {
		fmt.Println(err)
		return
	}
	mutex.Lock()
	Users[Conn.RemoteAddr().String()] = string(Name[:index])
	mutex.Unlock()
	Bytes := make([]byte, 1024)
	if len(s.Messages) > 0 {
		for _, msg := range s.Messages {
			Conn.Write([]byte(msg + "\n"))
		}
	}
	for {
		index, err := Conn.Read(Bytes)
		if err != nil {
			if errors.Is(err, net.ErrClosed) {
				return
			}
			continue
		}
		mutex.Lock()
		s.Messages = append(s.Messages, Conn.RemoteAddr().String()+" : "+string(Bytes[:index]))
		mutex.Unlock()
		s.GlobalMessage(Conn, []byte(Conn.RemoteAddr().String()+" : "+string(Bytes[:index])))
		Bytes = make([]byte, 1024)
	}
}
