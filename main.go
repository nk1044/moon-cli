package main

import (
	"fmt"
	"moon/internal"
	"moon/moonsystem"
	"moon/moonuser"
	"os"
)

func main() {
	args := os.Args

	// No arguments → show moon banner
	if len(args) < 2 {
		fmt.Println(internal.PrintMoon())
		fmt.Println("\nUsage: moon <command> [options]")
		fmt.Println("Try:  moon help")
		return
	}

	command := args[1]

	// Handle reserved system commands like help, version, etc.
	if moonsystem.IsReserved(command) {
		moonsystem.SystemCommandHandler(command)
		return
	}

	// Handle user-defined commands safely
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(internal.PrintMoonMessage(internal.Error, fmt.Sprintf("Unexpected error: %v", r)))
		}
	}()

	err := moonuser.UserCommandHandler(args)
	if err != nil {
		fmt.Println(internal.PrintMoonMessage(internal.Error, err.Error()))
	}
}
