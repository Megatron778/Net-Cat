package functions

import (
	"fmt"
	"strings"
	"time"
)

// This function allows the user to change their name and displays this change to other users.
func ChangeName(NewUser *UserData) {
	var OldName string
	var flag2 int
	for flag2 == 0 {
		flag2 = 0
		NewUser.Connection.Write([]byte("[ENTER YOUR NEW NAME]:"))
		n2, err := NewUser.Connection.Read(NewUser.Buffer)
		if err != nil {
			fmt.Println(err)
			return
		}
		OldName = NewUser.Name
		NewUser.Name = strings.TrimSpace(string(NewUser.Buffer[:n2-1]))
		if !ValidateAndAddUser(*NewUser) {
			continue
		}
		flag2 = 1

		Mutex.Lock()
		for _, OtherUsers := range User {
			if NewUser.Connection != OtherUsers.Connection {
				Now := time.Now()
				Format := Now.Format("2006-01-02 15:04:05")
				OtherUsers.Connection.Write([]byte("\n" + OldName + " changed his name to '" + NewUser.Name + "'\n[" + Format + "][" + OtherUsers.Name + "]:"))
			}
		}
		History = append(History, OldName+" changed his name to '"+NewUser.Name+"'\n")
		Mutex.Unlock()
	}

}
