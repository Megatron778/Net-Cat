package functions

import "time"

// Notifies all connected clients that a new user has joined the chat.
func OpenConnection(NewUser UserData) {
	for _, OtherUsers := range User {
		if NewUser.Name != OtherUsers.Name {
			Now := time.Now()
			Format := Now.Format("2006-01-02 15:04:05")
			Tap := "\n" + NewUser.Name + " has joined our chat...\n" + "[" + Format + "][" + OtherUsers.Name + "]:"
			OtherUsers.Connection.Write([]byte(Tap))
		}
	}
	History = append(History, NewUser.Name + " has joined our chat...\n")
}

// Notifies all clients that a user has left the chat and removes them from the list of active users.
func CloseConnection(NewUser UserData) {
	for _, OtherUsers := range User {
		if NewUser.Name != OtherUsers.Name {
			Now := time.Now()
			Format := Now.Format("2006-01-02 15:04:05")
			Tap := "\n" + NewUser.Name + " has left our chat...\n" + "[" + Format + "][" + OtherUsers.Name + "]:"
			OtherUsers.Connection.Write([]byte(Tap))
		}
	}
	History = append(History, NewUser.Name + " has left our chat...\n")

	var NewData []UserData

	for _, OtherUsers := range User {
		if NewUser.Connection != OtherUsers.Connection {
			NewData = append(NewData, OtherUsers)
		}
	}
    User = NewData
}
