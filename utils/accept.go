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
		if len(Users) == 10 {
			conn.Write([]byte("TCP-Chat Close Your Connection!\nTCP-Chat is alredy have 10 poeple!"))
			conn.Close()
			continue
		}
		s.Conns = append(s.Conns, conn)
		var mutex sync.Mutex
		go s.GetMessage(conn, &mutex)
	}
}
