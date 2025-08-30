package functions

import (
	"net"
	"strings"
	"sync"
)

type UserData struct {
	Connection net.Conn
	Name       string
	Buffer     []byte
}

type ServerData struct {
	NumberOfConnections int
	MessagesNumber      int
	action              string
}

var (
	allUser []UserData
	Server  ServerData
	Mutex   sync.Mutex
	History []string
)

// This function verifies the connection and manages it.
func HandleConnection(Connection net.Conn) {
	var user UserData
	var flag bool

	Mutex.Lock()
	if Server.NumberOfConnections == 10 {
		Connection.Write([]byte("The chat room is full."))
		Connection.Close()
		Mutex.Unlock()
		return
	}

	Server.NumberOfConnections++
	ServerDataPrint(Server)
	Mutex.Unlock()

	Welcome := "\033[33mWelcome to TCP-Chat!\n" +
		"         _nnnn_\n" +
		"        dGGGGMMb\n" +
		"       @p~qp~~qMb\n" +
		"       M|@||@) M|\n" +
		"       @,----.JM|\n" +
		"      JS^\\__/  qKL\n" +
		"     dZP        qKRb\n" +
		"    dZP          qKKb\n" +
		"   fZP            SMMb\n" +
		"   HZM            MMMM\n" +
		"   FqM            MMMM\n" +
		" __| \".        |\\dS\"qML\n" +
		" |    `.       | `' \\Zq\n" +
		"_)      \\.___.,|     .'\n" +
		"\\____   )MMMMMP|   .'\n" +
		"     `-'       `--'\033[0m\n"
	Connection.Write([]byte(Welcome))

	Buffer := make([]byte, 1024)

	for !flag {
		flag = false
		Connection.Write([]byte("[ENTER YOUR NAME]:"))
		n2, err := Connection.Read(Buffer)
		if err != nil {
			Server.NumberOfConnections--
			ServerDataPrint(Server)
			Connection.Close()
			return
		}
		UserName := strings.TrimSpace(string(Buffer[:n2-1]))
		user = UserData{Connection, UserName, Buffer}
		if !ValidateAndAddUser(user) {
			continue
		}
		flag = true

		Mutex.Lock()
		allUser = append(allUser, user)
		Mutex.Unlock()
	}

	Mutex.Lock()
	for _, msg := range History {
		Connection.Write([]byte("\033[30m" + msg + "\033[0m"))
	}

	OpenConnection(user)
	Mutex.Unlock()

	Sender(user)
}
