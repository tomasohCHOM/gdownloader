package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/tomasohCHOM/gdownloader/cmd/commands"
)

func init() {
	rootCmd.AddCommand(commands.PathCmd)
	rootCmd.AddCommand(commands.DownloadCmd)
}

var rootCmd = &cobra.Command{
	Use:   "gd-downloader",
	Short: "A program to download Google Drive files from the command line",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
