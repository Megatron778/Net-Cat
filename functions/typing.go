package functions

import (
	"strings"
	"time"
)

// Enables the user to write messages.
func Sender(NewUser UserData) {
	
	defer CloseConnection(NewUser)
	for {
		Now := time.Now()
		Format := Now.Format("2006-01-02 15:04:05")
		Tap := "[" + Format + "][" + NewUser.Name + "]:"
		NewUser.Connection.Write([]byte(Tap))
		n, err := NewUser.Connection.Read(NewUser.Buffer)
		if err != nil {
			
			return
		}
		Message := strings.TrimSpace(string(NewUser.Buffer[:n]))

		if Message == "" {
			continue
		}

		Mutex.Lock()
		History = append(History, "["+Format+"]["+NewUser.Name+"]:"+Message + "\n")
		Server.MessageNumber = len(History)
		Server.action = NewUser.Name + " Send a Message"
		ServerDataPrint(Server)
		Mutex.Unlock()
		
		Mutex.Lock()
		for _, users := range User {
			if users.Connection != NewUser.Connection {
				Now := time.Now()
				Format := Now.Format("2006-01-02 15:04:05")
				Tap := "\n" + Format + "][" + NewUser.Name + "]:" + Message + "\n[" + Format + "][" + users.Name + "]:"
				users.Connection.Write([]byte(Tap))
			}
		}
		Mutex.Unlock()
	
	}
}
