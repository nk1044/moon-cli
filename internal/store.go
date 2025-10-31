package internal

import (
	"encoding/json"
	"fmt"
	"os"
)

const storePath = "data/store.json"

type AliasData struct {
	Commands []string `json:"commands"`
}

type Store struct {
	Aliases map[string]AliasData `json:"aliases"`
}

func LoadStore() (Store, error) {
	var s Store
	s.Aliases = make(map[string]AliasData)

	// Ensure directory exists
	os.MkdirAll("data", 0755)

	// If file doesn't exist → create empty
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
	data, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal store: %v", err)
	}
	msg := PrintMoonMessage(Success, "Saved Successfully!")
	fmt.Println(msg)
	return os.WriteFile(storePath, data, 0644)
}
