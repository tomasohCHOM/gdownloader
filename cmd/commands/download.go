package commands

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tomasohCHOM/gdownloader/cmd/drive"
)

var DownloadCmd = &cobra.Command{
	Use:   "download",
	Short: "Download Google Drive files to a specified path",
	RunE: func(cmd *cobra.Command, args []string) error {
		srv, err := drive.Setup()
		if err != nil {
			return err
		}
		var searchQuery string
		fmt.Print("Enter a search query to list some files: ")
		fmt.Scan(&searchQuery)
		fmt.Println("Searching...")
		files, err := drive.Search(srv, searchQuery)
		if err != nil {
			return err
		}
		if len(files) == 0 {
			fmt.Println("No files found")
			return nil
		}
		for _, file := range files {
			fmt.Printf("%s (%s)\n", file.Name, file.Id)
		}
		file := files[0]
		return drive.DownloadFile(srv, file.Id, file.Name)
	},
}
