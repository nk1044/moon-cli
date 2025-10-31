package moonsystem

import (
	"fmt"
	"moon/internal"
)

func getVersion() string {
	const version = "1.0.0"
	return version
}


var ReservedWords = map[string]bool{
	// Common cobra commands
	"help":    true,
	"-h":       true,
	"--help":   true,
	"version": true,
	"-v":       true,
	"--version": true,

	// Common config/resource names
	"config": true,
	"init":   true,
	"new":    true,
	"create": true,
	"list":   true,
	"get":    true,
	"set":    true,
	"delete": true,
	"update": true,
}


func IsReserved(name string) bool {
	_, isReserved := ReservedWords[name]
	return isReserved
}


func SystemCommandHandler(command string) {
	switch command {
	case "help", "--help", "-h":
		fmt.Println(internal.PrintMoonMessage(internal.Normal, "Displaying help information..."))
	case "-v", "--version", "version":
		fmt.Println(internal.PrintMoonMessage(internal.Normal, "Moon CLI version "+ getVersion()))
	case "init":
		fmt.Println(internal.PrintMoonMessage(internal.Success, "Initializing Moon CLI..."))
	default:
		fmt.Println(internal.PrintMoonMessage(internal.Error, "Unknown system command: "+command))
	}
}