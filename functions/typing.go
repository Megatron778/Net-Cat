package functions

import (
	"strings"
	"time"
)

// Enables the user to write messages.
func Sender(user UserData) {
	defer CloseConnection(user)
	for {
		Now := time.Now()
		Format := Now.Format("2006-01-02 15:04:05")
		Tap := "[" + Format + "][" + user.Name + "]:"
		user.Connection.Write([]byte(Tap))
		n, err := user.Connection.Read(user.Buffer)
		if err != nil {
			return
		}
		Message := strings.TrimSpace(string(user.Buffer[:n]))
		if Message == "" {
			continue
		}
		if !inPrintable(Message) {
			continue
		}

		Mutex.Lock()
		History = append(History, "["+Format+"]["+user.Name+"]:"+Message+"\n")
		Server.MessagesNumber = len(History)
		Server.action = user.Name + " Send a Message"
		ServerDataPrint(Server)
		Mutex.Unlock()

		Mutex.Lock()
		for _, otherUsers := range allUser {
			if  user.Connection != otherUsers.Connection {
				Now := time.Now()
				Format := Now.Format("2006-01-02 15:04:05")
				Tap := "\n" + Format + "][" + user.Name + "]:" + Message + "\n[" + Format + "][" + otherUsers.Name + "]:"
				otherUsers.Connection.Write([]byte(Tap))
			}
		}
		Mutex.Unlock()

	}
}
