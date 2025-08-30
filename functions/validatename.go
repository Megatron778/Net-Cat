package functions

import (
	"unicode"
)

// Validates the username and, if valid, adds the user to the list of active clients.
func ValidateAndAddUser(user UserData) bool {
	if len(user.Name) == 0 {
		user.Connection.Write([]byte("Empty name, Please enter at least 3 Latters.\n"))
		return false
	}
	if len(user.Name) >= 15 {
		user.Connection.Write([]byte("Name is too long. Please enter less than 15 Latters.\n"))
		return false
	}
	if !CorrectName(user.Name) {
		user.Connection.Write([]byte("Please enter at least 3 Latters.\n"))
		return false
	}
	Mutex.Lock()
	defer Mutex.Unlock()
	for _, OtherUser := range allUser {
		if user.Name == OtherUser.Name {
			user.Connection.Write([]byte("This name already used.\n"))
			return false
		}
	}

	return true
}

// This function checks if the username is logical.
func CorrectName(Name string) bool {
	var count int
	for _, r := range Name {
		if unicode.IsLetter(r) {
			count++
		}
		if count == 3 {
			return true
		}
	}
	return false
}
