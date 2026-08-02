package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/dxvampi/binman/internal/store"
)

func Config() error {
	binaries, err := store.Load()
	if err != nil {
		return err
	}
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Alias: ")
		alias, err := reader.ReadString('\n')
		if err != nil {
			return err
		}
		alias = strings.TrimSpace(alias)

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
