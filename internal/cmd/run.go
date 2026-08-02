package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/dxvampi/binman/internal/store"
)

func Run(alias string, args []string) error {
	binaries, err := store.Load()
	if err != nil {
		return err
	}

	var binaryPath string
	found := false
	for _, b := range binaries {
		if b.Alias == alias {
			found = true
			binaryPath = b.Path
			break
		}
	}
	if !found {
		fmt.Printf("%s does not exist. Run 'binman list' to see available aliases.\n", alias)
		return nil
	}

	cmd := exec.Command(binaryPath, args...)

	fmt.Printf("Running command: %s %s\n", binaryPath, strings.Join(args, " "))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	return cmd.Run()
}
