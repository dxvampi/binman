package main

import (
	"fmt"
	"os"

	"github.com/dxvampi/binman/internal/cmd"
	"github.com/dxvampi/binman/internal/updater"
)

func main() {
	args := os.Args

	if len(args) < 2 {
		cmd.Help()
		return
	}

	command := args[1]

	var updateChan <-chan string
	if command != "-b" && command != "update" {
		updateChan = updater.CheckAsync()
	}

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
		cmd.Config()
	case "list":
		cmd.List()
	case "remove":
		cmd.Remove()
	case "help":
		cmd.Help()
	case "update":
		cmd.Update()
	default:
		fmt.Println("unknown command:", command)
	}

	if updateChan != nil {
		select {
		case latest := <-updateChan:
			updater.PromptAndInstall(latest)
		default:
		}
	}

}
