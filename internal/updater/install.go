package updater

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func RunInstallScript() error {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("powershell", "-Command",
			"irm https://raw.githubusercontent.com/dxvampi/binman/main/scripts/install/install.ps1 | iex")
	default:
		cmd = exec.Command("bash", "-c",
			"curl -fsSL https://raw.githubusercontent.com/dxvampi/binman/main/scripts/install/install.sh | bash")
	}

	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr

	return cmd.Run()

}

func PromptAndInstall(latest string) error {
	fmt.Printf("Update %s found!\n", latest)
	fmt.Print("Do you want to update? (Y/n): ")

	reader := bufio.NewReader(os.Stdin)
	answer, _ := reader.ReadString('\n')
	answer = strings.ToLower(strings.TrimSpace(answer))

	if answer != "" && answer != "y" {
		fmt.Println("Update cancelled")
		return nil
	}
	err := RunInstallScript()
	if err != nil {
		return err
	}

	_ = ClearCache()

	return nil
}
