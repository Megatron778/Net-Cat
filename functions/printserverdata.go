package functions

import "fmt"

func ServerDataPrint(Server ServerData) {
	var mode string
	for i := 0; i < 4; i++ {
		fmt.Print("\033[A")
		fmt.Print("\033[2K")
	}
	if Server.NumberOfConnection == 10 {
		mode = "\033[33mFull\033[0m"
	} else if Server.NumberOfConnection == 0 {
		mode = "\033[31mIdle\033[0m"
	} else if Server.NumberOfConnection > 0 && Server.NumberOfConnection < 3 {
		mode = "\033[92mActive\033[0m"
	}

	fmt.Println("\033[37m Group Mode :\033[0m", mode)
	fmt.Println("\033[37m Connections Number :\033[0m", Server.NumberOfConnection)
	fmt.Println("\033[37m Number of Messages :\033[0m", Server.MessageNumber)
	fmt.Println("\033[37m Action :\033[0m", Server.action)
}
