package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/dxn-scrlt/protectdf/internal/commands"
)

func main() {
	var err error

	args := os.Args
	var command string

	if len(args) != 3 {
		err = errors.New("expected 3 args")
	} else {
		command = args[1]
		input := args[2]

		switch command {
		case "set":
			err = commands.SetPassword(input)
		case "remove":
			err = commands.RemovePassword(input)
		default:
			err = errors.New("unknown command")
		}
	}

	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Success: %s password", command)
	}
}
