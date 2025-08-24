package functions

import (
	"fmt"
	"strings"
	"time"
)

// Enables the user to write messages.
func Sender(NewUser UserData, channel chan SenderData) {
	
	for {
		Now := time.Now()
		Format := Now.Format("2006-01-02 15:04:05")
		Tap := "[" + Format + "][" + NewUser.Name + "]:"
		NewUser.Connection.Write([]byte(Tap))
		n, err := NewUser.Connection.Read(NewUser.Buffer)
		if err != nil {
			CloseConnection(NewUser)
			fmt.Println("Error: ",err)
			return
		}
		Message := strings.TrimSpace(string(NewUser.Buffer[:n]))

		if Message == "" {
			continue
		}

		Mutex.Lock()
		History = append(History, "["+Format+"]["+NewUser.Name+"]:"+Message + "\n")
		Mutex.Unlock()

		Pack := SenderData{NewUser.Connection, NewUser.Name, Message}

		channel <- Pack
	}
}
