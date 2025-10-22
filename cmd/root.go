package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/tomasohCHOM/gdownloader/cmd/commands"
	"github.com/tomasohCHOM/gdownloader/cmd/options"
	"github.com/tomasohCHOM/gdownloader/cmd/ui/selector"
)

func init() {
	rootCmd.AddCommand(commands.PathCmd)
	rootCmd.AddCommand(commands.DownloadCmd)
}

var rootCmd = &cobra.Command{
	Use:   "gd-downloader",
	Short: "A program to download Google Drive files from the command line",
	Run: func(cmd *cobra.Command, args []string) {
		header := "Select one of the following options to continue."
		for {
			selected, exited, err := selector.RunSelector(header, options.ROOT_CMD_OPTIONS)
			if err != nil {
				log.Fatalf("Selection error: %v", err)
			}
			if exited {
				return
			}
			args := make([]string, 0)
			switch selected {
			case options.DOWNLOAD:
				commands.DownloadCmd.Run(cmd, args)
			case options.PATH:
				commands.PathCmd.Run(cmd, args)
			case options.EXIT:
				return
			default:
				log.Fatalf("Invalid command selection")
			}
		}
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
