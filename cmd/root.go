package cmd

import (
	"fmt"
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
			_, selected, err := selector.RunSelector(header, options.ROOT_CMD_OPTIONS)
			if err != nil {
				if err.Error() == selector.NO_SELECTION {
					return
				}
				log.Fatalf("Selection error: %v\n", err)
			}
			switch selected {
			case options.DOWNLOAD:
				if err := commands.DownloadCmd.RunE(cmd, []string{}); err != nil {
					log.Fatalf("Error: %v\n", err)
				}
			case options.PATH:
				if err := commands.PathCmd.RunE(cmd, []string{}); err != nil {
					log.Fatalf("Error: %v\n", err)
				}
			case options.EXIT:
				return
			default:
				fmt.Println("Invalid choice")
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
