package commands

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/tomasohCHOM/gdownloader/cmd/drive"
	"github.com/tomasohCHOM/gdownloader/cmd/options"
	"github.com/tomasohCHOM/gdownloader/cmd/ui/selector"
	"github.com/tomasohCHOM/gdownloader/cmd/ui/text"
	gdrive "google.golang.org/api/drive/v3"
)

var DownloadCmd = &cobra.Command{
	Use:   "download",
	Short: "Download Google Drive files to a specified path",
	RunE: func(cmd *cobra.Command, args []string) error {
		srv, err := drive.Setup()
		if err != nil {
			return err
		}
		searchQuery, exited, err := text.RunTextInput("Enter a search query to list some files")
		if err != nil {
			return err
		}
		if exited {
			return nil
		}
		fmt.Println("Searching...")
		files, err := drive.Search(srv, searchQuery)
		if err != nil {
			return err
		}
		if len(files) == 0 {
			fmt.Println("No files found")
			return nil
		}
		header := "Choose which file to download from"
		fileOptions := []string{}
		for _, file := range files {
			fileOptions = append(fileOptions, fmt.Sprintf("%s", file.Name))
		}
		fileOptions = append(fileOptions, options.DOWNlOAD_QUERY_MORE_PROMPT)
		var selectedFile string
		for {
			_, selected, err := selector.RunSelector(header, fileOptions)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error: %v\n", err)
				return err
			}
			if selected != options.DOWNlOAD_QUERY_MORE_PROMPT {
				selectedFile = selected
				break
			}
		}
		start := strings.LastIndex(selectedFile, "(")
		end := strings.LastIndex(selectedFile, ")")
		if start == -1 || end == -1 || start >= end {
			return fmt.Errorf("invalid file selection format: %q", selectedFile)
		}
		fileID := selectedFile[start+1 : end]
		var chosenFile *gdrive.File
		for _, f := range files {
			if f.Id == fileID {
				chosenFile = f
				break
			}
		}
		if chosenFile == nil {
			return fmt.Errorf("could not find file with id %s", fileID)
		}
		return drive.DownloadFile(srv, chosenFile.Id, chosenFile.Name)
	},
}

