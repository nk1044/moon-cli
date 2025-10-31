package main

import (
	"fmt"
	"moon/internal"
	"os"
)

const version = "1.0.0"

func main() {
	if len(os.Args) <2 {
		fmt.Println(internal.PrintMoon())
		return
	}
	command := os.Args[1]
	fmt.Println("You entered:", command)
	if command == "--version" || command == "-v" {
		fmt.Println(internal.PrintMoonMessage(internal.Normal, "Moon CLI version "+version))
		return
	}


}
