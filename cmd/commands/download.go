package commands

import "github.com/spf13/cobra"

var DownloadCmd = &cobra.Command{
	Use:   "download",
	Short: "Download Google Drive files to a specified path",

	Run: func(cmd *cobra.Command, args []string) {
	},
}
