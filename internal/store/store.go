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

func configPath() (string, error) {
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
	data, err := json.MarshalIndent(binaries, "", "  ")
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
