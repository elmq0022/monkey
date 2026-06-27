package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/elmq0022/monkey/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello: %s! Want to Monkey around?\n", user.Username)
	repl.Start(os.Stdin, os.Stdout)
}
