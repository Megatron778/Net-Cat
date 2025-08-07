package main

import (
	"fmt"
	"net"
	"netcat/functions"
)

func main() {
	listener, err := net.Listen("tcp", ":8989")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Server listening on port 8989")

	functions.ConnectionCheck(listener)

}
