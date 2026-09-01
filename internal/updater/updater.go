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

package updater

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Tag struct {
	Name string `json:"name"`
}

const apiURL = "https://codeberg.org/api/v1/repos/dxvampi/binman/tags"

func CheckForUpdate() (bool, string, error) {
	resp, err := http.Get(apiURL)
	if err != nil {
		return false, "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return false, "", fmt.Errorf("unexpected status: %d", resp.StatusCode)
	}

	var tags []Tag

	if err = json.NewDecoder(resp.Body).Decode(&tags); err != nil {
		return false, "", err
	}

	if len(tags) == 0 {
		return false, Version, nil
	}

	latestTag := tags[0].Name

	if latestTag != Version {
		return true, latestTag, nil
	}
	return false, Version, nil
}
