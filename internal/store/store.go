package store

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Binary struct {
	Alias string `json:"alias"`
	Path  string `json:"path"`
}

var configDir = ""

func configPath() (string, error) {

	if configDir != "" {
		return filepath.Join(configDir, "config.json"), nil
	}

	dir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(dir, "binman", "config.json"), nil
}

func Load() ([]Binary, error) {
	path, err := configPath()
	if err != nil {
		return nil, err
	}

	configFile, err := os.ReadFile(path)
	if os.IsNotExist(err) {
		return []Binary{}, nil
	}
	if err != nil {
		return nil, err
	}

	var binaries []Binary
	err = json.Unmarshal(configFile, &binaries)
	if err != nil {
		return nil, err
	}
	return binaries, nil
}

func Save(binaries []Binary) error {

	deduped := []Binary{}

	for _, b := range binaries {
		found := false
		for i, d := range deduped {
			if d.Alias == b.Alias {
				deduped[i].Path = b.Path
				found = true
				break
			}
		}

		if !found {
			deduped = append(deduped, b)
		}
	}

	data, err := json.MarshalIndent(deduped, "", "  ")
	if err != nil {
		return err
	}

	path, err := configPath()
	if err != nil {
		return err
	}

	dir := filepath.Dir(path)
	err = os.MkdirAll(dir, 0755)
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0644)
}
