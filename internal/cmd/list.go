package cmd

import (
	"fmt"

	"github.com/dxvampi/binman/internal/store"
)

func List() error {
	binaries, err := store.Load()
	if err != nil {
		return err
	}

	for _, b := range binaries {
		fmt.Printf("%s (%s)\n", b.Alias, b.Path)
	}

	return nil
}
