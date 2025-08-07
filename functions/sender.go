package functions

import (
	"fmt"
	"time"
)

func Sender(newuser UserData) {
	buffer := make([]byte, 1024)
	for {
		Tap := "\nWrite msg [" + newuser.Name + "] : "
		newuser.Connection.Write([]byte(Tap))
		n, err := newuser.Connection.Read(buffer)
		if err != nil {
			fmt.Println(err)
			return
		}
		data := string(buffer[:n])
		fmt.Print("Client : ", data)

		fmt.Println(newuser.Name)
		Message := SenderData{newuser.Name, data}

		channel <- Message
		time.Sleep(100 * time.Millisecond)
	}
}
