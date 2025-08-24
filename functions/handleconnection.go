package functions

import (
	"fmt"
	"net"
	"strings"
	"sync"
)

type UserData struct {
	Connection net.Conn
	Name       string
	Buffer     []byte
}

type SenderData struct {
	Connection net.Conn
	SenderName string
	Message    string
}

var (
	User    []UserData
	Mutex   sync.Mutex
	History []string
)

// This function verifies the connection and manages it.
func HandleConnection(Connection net.Conn) {
	var NewUser UserData
	var flag int

	channel := make(chan SenderData)
	
	Welcome := "Welcome to TCP-Chat!\n" +
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
		"     `-'       `--'\n"
	Connection.Write([]byte(Welcome))
	
	Buffer := make([]byte, 1024)
	
	for flag == 0 {
		flag = 0
		Connection.Write([]byte("[ENTER YOUR NAME]:"))
		n2, err := Connection.Read(Buffer)
		if err != nil {
			fmt.Println(err)
			return
		}
		UserName := strings.TrimSpace(string(Buffer[:n2-1]))
		NewUser = UserData{Connection, UserName, Buffer}
		if !ValidateAndAddUser(NewUser) {
			continue
		}
		flag = 1
		
		Mutex.Lock()
		if len(User) >= 10 {
			Connection.Write([]byte("The chat room is full."))
			Connection.Close()
			Mutex.Unlock()
			return
		}
		Mutex.Unlock()
		
		Mutex.Lock()
		User = append(User, NewUser)
		Mutex.Unlock()
	}

	Mutex.Lock()
	for _, Msg := range History {
		Connection.Write([]byte(Msg))
	}
	OpenConnection(NewUser)
	Mutex.Unlock()


	go Sender(NewUser, channel)
	go Receiver(NewUser, channel)
}
