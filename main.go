package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/codypenta/ape/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! This is the Ape programming language!\n", user.Username)
	fmt.Printf("Feel free to type in commmands\n")
	repl.Start(os.Stdin, os.Stdout)
}
