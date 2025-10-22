package store

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
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

func CheckPathExists(path string) (bool, error) {
	if strings.HasPrefix(path, "~") {
		home, err := os.UserHomeDir()
		if err != nil {
			return false, err
		}
		path = filepath.Join(home, strings.TrimPrefix(path, "~"))
	}
	absPath, err := filepath.Abs(path)
	if err != nil {
		return false, err
	}
	_, err = os.Stat(absPath)
	if os.IsNotExist(err) {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, nil
}
