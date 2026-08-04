// Copyright (C) 2026 dxvampi
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package cmd

import (
	"bufio"
	"errors"
	"fmt"
	"net"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/dxvampi/binman/internal/updater"
)

func Update() {
	updateAvailable, latestVersion, err := updater.CheckForUpdate()
	if err != nil {
		var dnsErr *net.DNSError
		if errors.As(err, &dnsErr) {
			fmt.Println("error: no internet connection")
			return
		}
		fmt.Println("error checking for updates:", err)
		return
	}

	if !updateAvailable {
		fmt.Printf("Already on the latest version! (%s)\n", updater.Version)
		return
	}
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Found update %s\nDo you want to update? (Y/n) ", latestVersion)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	input = strings.TrimSpace(strings.ToLower(input))
	if input != "y" && input != "" {
		fmt.Println("Aborting update")
		return
	}

	switch runtime.GOOS {
	case "windows":
		cmd := exec.Command("powershell", "irm", "https://codeberg.org/dxvampi/binman/raw/branch/main/scripts/install/install.ps1", "|", "iex")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			fmt.Println("error updating:", err)
			return
		}
	default:
		cmd := exec.Command("bash", "-c", "curl -fsSL https://codeberg.org/dxvampi/binman/raw/branch/main/scripts/install/install.sh | bash")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			fmt.Println("error updating:", err)
			return
		}
	}
}
