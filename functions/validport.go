package functions

// Validating the port coming from the inputs
func ValidPort(Port string) bool {

	IntPort := Atoi(Port)

	if IntPort >= 1024 && IntPort <= 49151 {
		return true
	}
	return false
}

// Converting the string to int
func Atoi(Port string) int {
	var FinalPort int
	for _, r := range Port {
		if r < '0' || r > '9' {
			return 0
		}
		FinalPort = FinalPort*10 + int(r-'0')

	}
	return FinalPort
}