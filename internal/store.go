package internal

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

func getStorePath() string {
	var baseDir string

	switch runtime.GOOS {
	case "windows":
		baseDir = os.Getenv("APPDATA")
		if baseDir == "" {
			baseDir = filepath.Join(os.Getenv("USERPROFILE"), "AppData", "Roaming")
		}
		baseDir = filepath.Join(baseDir, "Moon")
	default:
		home, _ := os.UserHomeDir()
		baseDir = filepath.Join(home, ".moon")
	}

	os.MkdirAll(baseDir, 0755)
	return filepath.Join(baseDir, "store.json")
}

type AliasData struct {
	Commands []string `json:"commands"`
}

type Store struct {
	Aliases map[string]AliasData `json:"aliases"`
}

func LoadStore() (Store, error) {
	var s Store
	s.Aliases = make(map[string]AliasData)

	storePath := getStorePath()

	if _, err := os.Stat(storePath); os.IsNotExist(err) {
		if err := SaveStore(s); err != nil {
			return s, fmt.Errorf("failed to create new store: %v", err)
		}
		return s, nil
	}

	data, err := os.ReadFile(storePath)
	if err != nil {
		return s, fmt.Errorf("cannot read store file: %v", err)
	}

	if len(data) == 0 {
		SaveStore(s)
		return s, nil
	}

	if err := json.Unmarshal(data, &s); err != nil {
		fmt.Println("⚠️  store.json corrupted, rebuilding a new one...")
		SaveStore(s)
	}

	return s, nil
}

func SaveStore(s Store) error {
	storePath := getStorePath()

	data, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal store: %v", err)
	}
	return os.WriteFile(storePath, data, 0644)
}
