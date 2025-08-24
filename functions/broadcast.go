package functions

import (
	"time"
)

// Broadcasts the user's message to all connected clients except the sender.
func Receiver(NewUser UserData, channel chan SenderData) {
	for {
		Pack := <-channel
		Mutex.Lock()
		for _, users := range User {
			if users.Connection != Pack.Connection {
				Now := time.Now()
				Format := Now.Format("2006-01-02 15:04:05")
				Tap := "\n[" + Format + "][" + Pack.SenderName + "]:" + Pack.Message + "\n[" + Format + "][" + users.Name + "]:"
				users.Connection.Write([]byte(Tap))
			}
		}
		Mutex.Unlock()
	}
}
