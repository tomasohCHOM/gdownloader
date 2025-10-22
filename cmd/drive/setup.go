package drive

import (
	"context"
	"os"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
)

func Setup() (*drive.Service, error) {
	credPath := credentialsFilePath()
	b, err := os.ReadFile(credPath)
	if err != nil {
		return nil, err
	}
	config, err := google.ConfigFromJSON(b, drive.DriveReadonlyScope)
	if err != nil {
		return nil, err
	}
	token, err := getClient(config)
	if err != nil {
		return nil, err
	}
	ctx := context.Background()
	client := config.Client(ctx, token)
	return drive.NewService(ctx, option.WithHTTPClient(client))
}
