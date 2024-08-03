package main

import (
	"fmt"
	"monkey-interpreter/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()

	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! WHAT Monkey DO?", user.Username)
	fmt.Printf("TYPE thing DO!\n")

	repl.Start(os.Stdin, os.Stdout)
}
