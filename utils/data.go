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

var (
	Users = map[string]string{}
	Art   = "\n	_nnnn_\n	dGGGGMMb\n	@p~qp~~qMb\n	M|@||@) M|\n	@,----.JM|\n	JS^\\__/  qKL\n	dZP        qKRb\n	dZP          qKKb\n	fZP            SMMb\n	HZM            MMMM\n	FqM            MMMM\n	__| \".         |\\dS\"qML\n	|    `.       | `' \\Zq\n	_)      \\.___.,|     .'\n	\\____   )MMMMMP|   .'\n	    `-'       `--'\n"
)
