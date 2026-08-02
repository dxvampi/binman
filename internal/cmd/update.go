package cmd

import (
	"fmt"
	"strings"

	"github.com/dxvampi/binman/internal/updater"
	"github.com/dxvampi/binman/internal/version"
)

func Update() error {
	fmt.Printf("Running binman v%s\n", version.Version)
	fmt.Println("Checking for updates...")

	latest, err := updater.FetchLatestVersion()
	if err != nil {
		return err
	}
	latest = strings.TrimPrefix(latest, "v")

	if latest == version.Version {
		fmt.Println("Already on the latest version!")
		return nil
	}

	return updater.PromptAndInstall(latest)
}
