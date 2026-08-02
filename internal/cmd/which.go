package cmd

import (
	"fmt"

	"github.com/dxvampi/binman/internal/store"
)

func Which(alias string) error {
	binaries, err := store.Load()
	if err != nil {
		return err
	}

	for _, b := range binaries {
		if b.Alias == alias {
			fmt.Println(b.Path)
			return nil
		}
	}

	fmt.Printf("%s does not exist. Run 'binman list' to see available aliases.\n", alias)
	return nil
}
