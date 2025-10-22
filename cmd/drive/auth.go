package drive

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"golang.org/x/oauth2"
)

func tokenFilePath() string {
	home, _ := os.UserHomeDir()
	return filepath.Join(home, ".gdownloader", "token.json")
}

func credentialsFilePath() string {
	home, _ := os.UserHomeDir()
	return filepath.Join(home, ".gdownloader", "credentials.json")
}

func getClient(config *oauth2.Config) (*oauth2.Token, error) {
	tok, err := tokenFromFile()
	if err != nil {
		tok, err = getTokenFromWeb(config)
		if err != nil {
			return nil, fmt.Errorf("failed to get token from web: %w", err)
		}
		if err := saveToken(tok); err != nil {
			return nil, fmt.Errorf("failed to save token: %w", err)
		}
	}
	return tok, nil
}

func getTokenFromWeb(config *oauth2.Config) (*oauth2.Token, error) {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("\nGo to the following link in your browser and enter the code:\n%v\n", authURL)
	var code string
	if _, err := fmt.Scan(&code); err != nil {
		return nil, fmt.Errorf("unable to read authorization code: %w", err)
	}
	tok, err := config.Exchange(context.Background(), code)
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve token from web: %w", err)
	}
	return tok, nil
}

func tokenFromFile() (*oauth2.Token, error) {
	path := tokenFilePath()
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	tok := &oauth2.Token{}
	if err := json.NewDecoder(f).Decode(tok); err != nil {
		return nil, fmt.Errorf("failed to decode token file: %w", err)
	}
	return tok, nil
}

func saveToken(token *oauth2.Token) error {
	path := tokenFilePath()
	if err := os.MkdirAll(filepath.Dir(path), 0700); err != nil {
		return fmt.Errorf("failed to create token directory: %w", err)
	}
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return fmt.Errorf("failed to open token file: %w", err)
	}
	defer f.Close()
	if err := json.NewEncoder(f).Encode(token); err != nil {
		return fmt.Errorf("failed to encode token: %w", err)
	}
	fmt.Printf("\nSaved OAuth token to %s\n", path)
	return nil
}
