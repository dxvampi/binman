package cmd

import "fmt"

func Help() {
	fmt.Println()
	fmt.Println("-------- Binman (v1.0) help --------")
	fmt.Println()
	fmt.Println("'binman config' - Used to add new aliases")
	fmt.Println("'binman list' - Lists every alias that's saved on the config")
	fmt.Println("'binman remove' - Used to remove aliases")
	fmt.Println("'binman -b <alias> [args...]' - Runs a command with an alias")
	fmt.Println()
	fmt.Println("-------- Binman (v1.0) help --------")
}
