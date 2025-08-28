package main

import (
	"fmt"
	"net"
	"netcat/functions"
	"os"
)

func main() {
	port := "8989"
	if len(os.Args) == 2 {
		if !functions.ValidPort(os.Args[1]) {
			fmt.Println("Please entre a valid port [1024 --> 49151]")
			return
		}
		port = os.Args[1]
	} else if len(os.Args) != 1 {
		fmt.Println("[USAGE]: ./TCPChat $port")
		return
	}
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("\033[37mServer listening on port :\033[0m", port)
	fmt.Println("\033[37m Group Mode :\033[0m", "\033[31mIdle\033[0m")
	fmt.Println("\033[37m Connections Number :\033[0m", 0)
	fmt.Println("\033[37m Number of Messages :\033[0m", 0)
	fmt.Println("\033[37m Action :\033[0m")
	for {
		Connection, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		defer Connection.Close()
		go functions.HandleConnection(Connection)
	}
	
}
