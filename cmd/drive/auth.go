package drive

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"golang.org/x/oauth2"
)

func tokenFilePath() string {
	home, _ := os.UserHomeDir()
	return filepath.Join(home, ".gdownloader", "token.json")
}

func credentialsFilePath() string {
	return filepath.Join("credentials.json")
}

func getClient(config *oauth2.Config) *oauth2.Token {
	tok, err := tokenFromFile()
	if err != nil {
		tok = getTokenFromWeb(config)
		saveToken(tok)
	}
	return tok
}

func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser and enter the code:\n%v\n", authURL)

	var code string
	if _, err := fmt.Scan(&code); err != nil {
		log.Fatalf("Unable to read authorization code: %v", err)
	}
	tok, err := config.Exchange(context.Background(), code)
	if err != nil {
		log.Fatalf("Unable to retrieve token: %v", err)
	}
	return tok
}

func tokenFromFile() (*oauth2.Token, error) {
	path := tokenFilePath()
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	tok := &oauth2.Token{}
	return tok, json.NewDecoder(f).Decode(tok)
}

func saveToken(token *oauth2.Token) {
	path := tokenFilePath()
	os.MkdirAll(filepath.Dir(path), 0700)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("Unable to save token: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
	fmt.Printf("Saved OAuth token to %s\n", path)
}
