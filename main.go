package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"sync"
)

type Data struct {
	Conn net.Conn
	Name string
}

var (
	Clients []*Data
	Mutex sync.Mutex
)

func main() {
	Port := "8989"
	if len(os.Args) == 2 {
		Port = os.Args[1] 
	} else if len(os.Args) > 2 {
		fmt.Println("Usage : go run main.go <port>")
		return
	}

	Listener, err := net.Listen("tcp", ":"+Port) 
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Server Listening on port :", Port)

	defer Listener.Close()

	for {
		Connection, err := Listener.Accept()
		if err != nil {
			log.Println(err)
		}
		go HandlerConnection(Connection)
	}
}

func HandlerConnection(Connection net.Conn) {
	fmt.Println("The Clienr Connected")
	defer Connection.Close()

	_,err := Connection.Write([]byte("Server Connected \n Entre your name :"))
	if err != nil {
		log.Println(err)
	}

	reader := bufio.NewReader(Connection)
	UserName, err := reader.ReadString('\n')
	fmt.Println("Her name is" ,UserName)
	_,err = Connection.Write([]byte("Welcome " + UserName))
	if err != nil {
		log.Println(err)
	}

	Client := &Data {
		Conn: Connection,
		Name: UserName,
	}

	Mutex.Lock()
	Clients = append(Clients, Client)
	Mutex.Unlock()
}