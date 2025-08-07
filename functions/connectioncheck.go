package functions

import (
	"fmt"
	"net"
	"sync"
)

type UserData struct {
	Connection net.Conn
	Name       string
}

type SenderData struct {
	SenderName string
	Content string
}

var (
	channel  chan SenderData
	user      []UserData
	Mutex     sync.Mutex
	flag      int
)

func ConnectionCheck(listener net.Listener) {
    var newuser UserData
	for {
		connection, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("Connection now")
		buffer2 := make([]byte, 1024)

		flag = len(user)
		for flag == len(user) {
			connection.Write([]byte("\nEntre your name : "))
			n2, err := connection.Read(buffer2)
			if err != nil {
				fmt.Println(err)
				return
			}

			UserName := string(buffer2[:n2-1])
			newuser = UserData{connection,UserName}

			user = NewUser(user, newuser, connection)

		}

		connection.Write([]byte("Correct Name"))
		channel = make(chan SenderData)

		go Sender(newuser)
		go Receiver(newuser)
	}
}
