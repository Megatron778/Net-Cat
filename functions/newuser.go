package functions

import "net"

func NewUser(user []UserData, newuser UserData, connection net.Conn) []UserData {

	for _, otherusers := range user {
		if newuser.Name == otherusers.Name {
			connection.Write([]byte("Already Used"))
			return user
		}
	}
	return append(user, newuser)
}
