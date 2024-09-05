package TCP_Chat

import (
	"bufio"
	"net"
	"sync"
)

func (s *Server) GlobalMessage(Conn net.Conn, msg []byte) {
	for _, c := range s.Conns {
		println(c.RemoteAddr().String(), Conn.LocalAddr().String())
		if c.RemoteAddr().String() != Conn.RemoteAddr().String() && len(Users[Conn.RemoteAddr().String()]) != 0 {
			c.Write(msg)
			c.Write([]byte(Format(Users[c.RemoteAddr().String()], "")))
		}
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
	} else {
		for _, CurName := range Users {
			if string(Name[:index-1]) == CurName {
				Conn.Write([]byte("Username already Exist!"))
				Conn.Close()
				return
			}
		}
	}
	Bytes := make([]byte, 1024)
	if len(s.Messages) > 0 {
		for _, msg := range s.Messages {
			Conn.Write([]byte(msg))
		}
	}
	mutex.Lock()
	Users[Conn.RemoteAddr().String()] = string(Name[:index-1])
	mutex.Unlock()
	s.GlobalMessage(Conn, []byte("\n"+Users[Conn.RemoteAddr().String()]+" Joined\n"))
	for {
		Conn.Write([]byte(Format(Users[Conn.RemoteAddr().String()], "")))
		index, err := Conn.Read(Bytes)
		if err != nil {
			if !bufio.NewScanner(Conn).Scan() {
				break
			}
			continue
		}
		mutex.Lock()
		s.Messages = append(s.Messages, Format(Users[Conn.RemoteAddr().String()], string(Bytes[:index])))
		mutex.Unlock()
		s.GlobalMessage(Conn, []byte("\n"+Format(Users[Conn.RemoteAddr().String()], string(Bytes[:index]))))
		Bytes = make([]byte, 1024)
	}
	s.GlobalMessage(Conn, []byte("\n"+Users[Conn.RemoteAddr().String()]+" Disconnected!"+"\n"))
}
