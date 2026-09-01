package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/dxvampi/binman/internal/store"
)

func Remove(args []string) error {
	binaries, err := store.Load()
	if err != nil {
		return err
	}

	if len(args) > 2 && len(args) <= 3 {
		alias := args[2]
		alias = strings.TrimSpace(alias)

		remaining := []store.Binary{}
		found := false

		for _, b := range binaries {
			if b.Alias == alias {
				found = true
				continue
			}
			remaining = append(remaining, b)
		}

		if !found {
			fmt.Printf("%s does not exist. Run 'binman list' to see available aliases.\n", alias)
			return nil
		}

		err = store.Save(remaining)
		if err != nil {
			return err
		}

		fmt.Printf("Deleted %s\n", alias)
		return nil
	} else if len(args) > 3 {
		fmt.Println("usage: binman remove <alias>")
		return fmt.Errorf("invalid command usage %s", args)
	}

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Alias to delete: ")
	alias, err := reader.ReadString('\n')
	if err != nil {
		return err
	}
	alias = strings.TrimSpace(alias)

	remaining := []store.Binary{}
	found := false

	for _, b := range binaries {
		if b.Alias == alias {
			found = true
			continue
		}
		remaining = append(remaining, b)
	}

	if !found {
		fmt.Printf("%s does not exist. Run 'binman list' to see available aliases.\n", alias)
		return nil
	}

	err = store.Save(remaining)
	if err != nil {
		return err
	}

	fmt.Printf("Deleted %s\n", alias)
	return nil
}
