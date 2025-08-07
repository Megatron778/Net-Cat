package functions

import (
	"fmt"
)

func Receiver(newuser UserData) {

	for {
		msg := <-channel
		fmt.Println(newuser.Name)
		for _,users := range user {
			if users.Name != newuser.Name {
				
				Tap := "\nReceive msg [" + newuser.Name + "] : " + msg + "\n send msg [" + users.Name + "] : "
				users.Connection.Write([]byte(Tap))
			}
		}
	}

}
