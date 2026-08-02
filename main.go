package main

import (
	"fmt"
	"os"

	"github.com/dxvampi/binman/internal/cmd"
)

func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Println("usage: binman <command>")
		return
	}

	command := args[1]

	if command == "-b" {
		if len(args) < 3 {
			fmt.Println("usage: binman -b <alias> [args...]")
			return
		}
		alias := args[2]
		extraArgs := args[3:]
		cmd.Run(alias, extraArgs)
		return
	}

	switch command {
	case "config":
		cmd.Config()
	case "list":
		cmd.List()
	case "remove":
		cmd.Remove()
	case "help":
		cmd.Help()
	default:
		fmt.Println("unknown command:", command)
	}
}
