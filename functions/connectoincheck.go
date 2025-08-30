package functions

import (
	"time"
)

// Notifies all connected clients that a new user has joined the chat.
func OpenConnection(user UserData) {
	for _, OtherUsers := range allUser {
		if user.Name != OtherUsers.Name {
			Now := time.Now()
			Format := Now.Format("2006-01-02 15:04:05")
			Tap := "\n\033[32m" + user.Name + " has joined our chat...\033[0m\n" + "[" + Format + "][" + OtherUsers.Name + "]:"
			OtherUsers.Connection.Write([]byte(Tap))
		}
	}
	Server.action = user.Name + " has joined"
	ServerDataPrint(Server)
}

// Notifies all clients that a user has left the chat and removes them from the list of active users.
func CloseConnection(user UserData) {
	Mutex.Lock()
	for _, OtherUsers := range allUser {
		if user.Name != OtherUsers.Name {
			Now := time.Now()
			Format := Now.Format("2006-01-02 15:04:05")
			Tap := "\n\033[31m" + user.Name + " has left our chat...\033[0m\n" + "[" + Format + "][" + OtherUsers.Name + "]:"
			OtherUsers.Connection.Write([]byte(Tap))
		}
	}

	var newAllUserSlice []UserData

	for _, OtherUsers := range allUser {
		if user.Connection != OtherUsers.Connection {
			newAllUserSlice = append(newAllUserSlice, OtherUsers)
		}
	}
	allUser = newAllUserSlice
	Server.NumberOfConnections--
	Server.action = user.Name + " has left"
	ServerDataPrint(Server)
	Mutex.Unlock()
}
