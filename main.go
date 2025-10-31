package main

import (
	"fmt"
	"moon/internal"
	"moon/moonsystem"
	"moon/moonuser"
	"os"
)


func main() {
	if len(os.Args) < 2 {
		fmt.Println(internal.PrintMoon())
		return
	}
	command := os.Args[1]
	fmt.Println("You entered:", command)

	if moonsystem.IsReserved(command) {
		moonsystem.SystemCommandHandler(command)
	} else {
		moonuser.UserCommandHandler(command)
	}

}
