package functions

import (
	"unicode"
)

// Validates the username and, if valid, adds the user to the list of active clients.
func ValidateAndAddUser(NewUser UserData) bool {

	if len(NewUser.Name) == 0 {
		NewUser.Connection.Write([]byte("Empty name, Please enter at least 3 Latters.\n"))
		return false
	}
	if len(NewUser.Name) >= 15 {
		NewUser.Connection.Write([]byte("Name is too long. Please enter less than 15 Latters.\n"))
		return false
	}
	if !CorrectName(NewUser.Name) {
		NewUser.Connection.Write([]byte("Please enter at least 3 characters.\n"))
		return false
	}
	Mutex.Lock()
	defer Mutex.Unlock()
	for _, OtherUser := range User {
		if NewUser.Name == OtherUser.Name {
			NewUser.Connection.Write([]byte("This name already used.\n"))
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
