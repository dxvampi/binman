package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"

	"github.com/dxvampi/binman/internal/store"
)

func isValidAlias(s string) bool {
	if s == "" {
		return false
	}
	for _, r := range s {
		if unicode.IsControl(r) {
			return false
		}
	}
	return true
}

func Config(args []string) error {

	binaries, err := store.Load()
	if err != nil {
		return err
	}

	if len(args) > 2 && len(args) <= 4 {
		alias := args[2]
		alias = strings.TrimSpace(alias)
		if !isValidAlias(alias) {
			fmt.Println("Invalid alias, please use plain text only.")
		}

		exists := false
		for _, b := range binaries {
			if b.Alias == alias {
				exists = true
				break
			}
		}

		if exists {
			fmt.Printf("Alias '%s' already exists and will be overwritten.\n", alias)
		}

		path := args[3]
		path = strings.TrimSpace(path)

		binaries = append(binaries, store.Binary{Alias: alias, Path: path})
		err = store.Save(binaries)
		if err != nil {
			return err
		}
		fmt.Println("successfully added aliases!")
		return nil
	} else if len(args) == 3 {
		fmt.Println("usage: binman config <alias> <path>")
		return fmt.Errorf("invalid command usage %s", args)
	}

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Alias: ")
		alias, err := reader.ReadString('\n')
		if err != nil {
			return err
		}
		alias = strings.TrimSpace(alias)
		if !isValidAlias(alias) {
			fmt.Println("Invalid alias, please use plain text only.")
			continue
		}

		exists := false
		for _, b := range binaries {
			if b.Alias == alias {
				exists = true
				break
			}
		}

		if exists {
			fmt.Printf("Alias '%s' already exists and will be overwritten.\n", alias)
		}

		fmt.Print("Path: ")
		path, err := reader.ReadString('\n')
		if err != nil {
			return err
		}
		path = strings.TrimSpace(path)

		binaries = append(binaries, store.Binary{Alias: alias, Path: path})

		fmt.Print("Add another? (y/N): ")
		answer, err := reader.ReadString('\n')
		if err != nil {
			return err
		}
		answer = strings.ToLower(strings.TrimSpace(answer))

		if answer != "y" {
			break
		}
	}

	err = store.Save(binaries)
	if err != nil {
		return err
	}
	fmt.Println("successfully added aliases!")
	return nil
}
