package functions

import (
	"time"
	"strings"
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
			return
		}
		Message := string(NewUser.Buffer[:n])

		if strings.TrimSpace(Message) == "" {
			continue
		}
		if Message == "--name\n" {
			ChangeName(&NewUser)
			continue
		}
		Mutex.Lock()
		History = append(History, "["+Format+"]["+NewUser.Name+"]:"+Message)
		Mutex.Unlock()
		

		Pack := SenderData{NewUser.Connection, NewUser.Name, Message}

		channel <- Pack
		time.Sleep(100 * time.Millisecond)
	}
}
