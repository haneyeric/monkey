package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/haneyeric/monkey/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		os.Exit(1)
	}

	fmt.Printf("Hello %s! Let's Monkey around.\n", user.Username)
	fmt.Println("Feel free to type is commands.")

	repl.Start(os.Stdin, os.Stdout)
}
