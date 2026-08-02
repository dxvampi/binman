package updater

import (
	"encoding/json"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/dxvampi/binman/internal/version"
)

type Cache struct {
	LastChecked time.Time `json:"last_checked"`
	Version     string    `json:"version"`
}

func FetchLatestVersion() (string, error) {
	client := http.Client{
		Timeout: 2 * time.Second,
	}

	resp, err := client.Get("https://api.github.com/repos/dxvampi/binman/releases/latest")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var release struct {
		TagName string `json:"tag_name"`
	}

	err = json.NewDecoder(resp.Body).Decode(&release)
	if err != nil {
		return "", err
	}

	return release.TagName, nil
}

func cachePath() (string, error) {
	dir, err := os.UserCacheDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(dir, "binman", "last_check.json"), nil
}

func checkAndCache() (string, error) {

	cache, err := loadCache()
	if err != nil {
		return "", err
	}

	if time.Since(cache.LastChecked) < 24*time.Hour {
		return cache.Version, nil
	}

	latest, err := FetchLatestVersion()
	if err != nil {
		return "", err
	}

	latest = strings.TrimPrefix(latest, "v")

	_ = saveCache(Cache{
		LastChecked: time.Now(),
		Version:     latest,
	})
	return latest, nil
}

func loadCache() (Cache, error) {
	path, err := cachePath()
	if err != nil {
		return Cache{}, err
	}

	data, err := os.ReadFile(path)
	if os.IsNotExist(err) {
		return Cache{}, nil
	}
	if err != nil {
		return Cache{}, err
	}

	var cache Cache
	err = json.Unmarshal(data, &cache)
	if err != nil {
		return Cache{}, err
	}

	return cache, nil
}

func saveCache(cache Cache) error {
	path, err := cachePath()
	if err != nil {
		return err
	}

	data, err := json.MarshalIndent(cache, "", "  ")
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

func CheckAsync() <-chan string {
	resultChan := make(chan string, 1)

	go func() {
		latest, err := checkAndCache()
		if err != nil {
			return
		}
		if latest != "" && latest != strings.TrimPrefix(version.Version, "v") {
			resultChan <- latest
		}
	}()

	return resultChan
}

func ClearCache() error {
	path, err := cachePath()
	if err != nil {
		return err
	}

	err = os.Remove(path)
	if err != nil && !os.IsNotExist(err) {
		return err
	}

	return nil
}
