package functions

import (
	"time"
)

// Notifies all connected clients that a new user has joined the chat.
func OpenConnection(NewUser UserData) {
	for _, OtherUsers := range User {
		if NewUser.Name != OtherUsers.Name {
			Now := time.Now()
			Format := Now.Format("2006-01-02 15:04:05")
			Tap := "\n\033[32m" + NewUser.Name + " has joined our chat...\033[0m\n" + "[" + Format + "][" + OtherUsers.Name + "]:"
			OtherUsers.Connection.Write([]byte(Tap))
		}
	}
	Server.action = NewUser.Name + " has joined"
	ServerDataPrint(Server)
}

// Notifies all clients that a user has left the chat and removes them from the list of active users.
func CloseConnection(NewUser UserData) {
	Mutex.Lock()
	for _, OtherUsers := range User {
		if NewUser.Name != OtherUsers.Name {
			Now := time.Now()
			Format := Now.Format("2006-01-02 15:04:05")
			Tap := "\n\033[31m" + NewUser.Name + " has left our chat...\033[0m\n" + "[" + Format + "][" + OtherUsers.Name + "]:"
			OtherUsers.Connection.Write([]byte(Tap))
		}
	}

	var NewData []UserData

	for _, OtherUsers := range User {
		if NewUser.Connection != OtherUsers.Connection {
			NewData = append(NewData, OtherUsers)
		}
	}
	User = NewData
	Server.NumberOfConnection--
	Server.action = NewUser.Name + " has left" 
	ServerDataPrint(Server)
	Mutex.Unlock()
}
