package functions

import (
	"fmt"
)

func Receiver(newuser UserData) {

	for {
		Message := <-channel
		fmt.Println(newuser.Name)
		for _,users := range user {
			if users.Name != Message.SenderName {
				
				Tap := "\nReceive msg [" + Message.SenderName + "] : " + Message.Content + "\nsend msg [" + users.Name + "] : "
				users.Connection.Write([]byte(Tap))
			}
		}
	}

}
