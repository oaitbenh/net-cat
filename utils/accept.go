package TCP_Chat

import (
	"fmt"
	"net"
	"sync"
)

func (s *Server) AcceptLoop() {
	s.Conns = []net.Conn{}
	for {
		conn, err := s.Listener.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		s.Conns = append(s.Conns, conn)
		var mutex sync.Mutex
		fmt.Println(conn.RemoteAddr().String())
		go s.GetMessage(conn, &mutex)
	}
}
