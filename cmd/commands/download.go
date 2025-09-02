package commands

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tomasohCHOM/google-drive-downloader/cmd/drive"
)

var DownloadCmd = &cobra.Command{
	Use:   "download",
	Short: "Download Google Drive files to a specified path",

	RunE: func(cmd *cobra.Command, args []string) error {
		srv, err := drive.Setup()
		if err != nil {
			return err
		}
		files, err := drive.Search(srv, "MATH")
		if err != nil {
			return err
		}
		if len(files) == 0 {
			fmt.Println("No files found")
			return nil
		}
		file := files[0]
		fmt.Println(file)
		return drive.DownloadFile(srv, file.Id, file.Name)
	},
}
