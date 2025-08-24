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
	fmt.Println("Server listening on port :", port)

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
