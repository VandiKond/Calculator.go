package application

import (
	"bufio"
	"fmt"
	"os"
)

func CloseBy() {
	for {
		// Ads the user name on the string
		fmt.Print("$:admin> ")
		//Scans the line
		reader := bufio.NewReader(os.Stdin)
		// Read the entire line
		input, _ := reader.ReadString('\n')

		// Remove the newline character
		var val = input[:len(input)-2]

		// Checking command options
		switch val {
		case "exit":
			// Exit option
			fmt.Println("exiting")
			os.Exit(418)
		default:
			// default
			fmt.Println("unknown command")
		}
	}
}
