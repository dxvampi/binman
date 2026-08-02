package cmd

import (
	"fmt"

	"github.com/dxvampi/binman/internal/version"
)

func Help() {
	fmt.Println()
	fmt.Printf("-------- Binman (%s) help --------\n", version.Version)
	fmt.Println()
	fmt.Println("'binman config' - Used to add new aliases")
	fmt.Println("'binman list' - Lists every alias that's saved on the config")
	fmt.Println("'binman remove' - Used to remove aliases")
	fmt.Println("'binman -b <alias> [args...]' - Runs a command with an alias")
	fmt.Println("'binman which <alias>' - Prints the absolute path to the binary")
	fmt.Println("'binman update' - Checks for updates and updates if possible")
	fmt.Println()
	fmt.Printf("-------- Binman (%s) help --------\n", version.Version)
}
