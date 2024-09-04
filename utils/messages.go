package TCP_Chat

import (
	"errors"
	"net"
	"sync"
)

func (s *Server) GlobalMessage(Conn net.Conn, msg []byte) {
	for _, c := range s.Conns {
		if Users[c.RemoteAddr().String()] == Users[Conn.LocalAddr().String()] {
			continue
		}
		c.Write(msg)
		c.Write([]byte(Format(Users[c.RemoteAddr().String()], "")))
	}
}

func (s *Server) GetMessage(Conn net.Conn, mutex *sync.Mutex) {
	defer Conn.Close()
	Conn.Write([]byte("<< Useranme Condition Min Len 4 char Max len 20 >>\n"))
	Conn.Write([]byte("Print Your Username : "))
	Name := make([]byte, 20)
	index, err := Conn.Read(Name)
	if err != nil || index < 4 || index > 20 {
		Conn.Write([]byte("Invalid Name!"))
		Conn.Close()
		return
	}
	mutex.Lock()
	Users[Conn.RemoteAddr().String()] = string(Name[:index-1])
	mutex.Unlock()
	Bytes := make([]byte, 1024)
	if len(s.Messages) > 0 {
		for _, msg := range s.Messages {
			Conn.Write([]byte(msg))
		}
	}
	s.GlobalMessage(Conn, []byte(Users[Conn.RemoteAddr().String()]+" Joined"))
	for {
		Conn.Write([]byte(Format(Users[Conn.RemoteAddr().String()], "")))
		index, err := Conn.Read(Bytes)
		if err != nil {
			if errors.Is(err, net.ErrClosed) {
				s.GlobalMessage(Conn, []byte(Users[Conn.RemoteAddr().String()]+" Disconnected!"))
				return
			}
			continue
		}
		mutex.Lock()
		s.Messages = append(s.Messages, Format(Users[Conn.RemoteAddr().String()], string(Bytes[:index])))
		mutex.Unlock()
		s.GlobalMessage(Conn, []byte("\n"+Format(Users[Conn.RemoteAddr().String()], string(Bytes[:index]))))
		Bytes = make([]byte, 1024)
	}
}
