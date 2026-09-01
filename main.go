package main

import (
	"fmt"
	"os"

	"github.com/dxvampi/binman/internal/cmd"
)

func main() {
	args := os.Args

	if len(args) < 2 {
		cmd.Help()
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
	case "which":
		if len(args) < 3 {
			fmt.Println("usage: binman which <alias>")
			return
		}
		cmd.Which(args[2])
	case "config":
		cmd.Config(args)
	case "list":
		cmd.List()
	case "remove":
		cmd.Remove(args)
	case "help":
		cmd.Help()
	case "update", "-U", "--update":
		cmd.Update()
	default:
		fmt.Println("unknown command:", command)
	}
}
