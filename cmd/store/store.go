package store

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Store struct {
	Paths map[string]string `json:"paths"`
}

func getStorePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, ".local", "share", "gdownloader", "store.json"), nil
}

func Load() (*Store, error) {
	path, err := getStorePath()
	if err != nil {
		return nil, err
	}
	data, err := os.ReadFile(path)
	if os.IsNotExist(err) {
		return &Store{Paths: map[string]string{}}, nil
	} else if err != nil {
		return nil, err
	}
	var store Store
	if err := json.Unmarshal(data, &store); err != nil {
		return nil, err
	}
	return &store, nil
}

func (store *Store) Save() error {
	path, err := getStorePath()
	if err != nil {
		return err
	}
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return err
	}
	data, err := json.MarshalIndent(store, "", "	")
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0644)
}
