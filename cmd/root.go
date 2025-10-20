package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/tomasohCHOM/gdownloader/cmd/commands"
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
		options := []string{"Download files", "Manage paths", "Exit"}

		for {
			idx, _, err := selector.RunSelector(header, options)
			if err != nil {
				if err.Error() == selector.NO_SELECTION {
					return
				}
				log.Fatalf("Selection error: %v\n", err)
			}
			switch idx {
			case 0:
				if err := commands.DownloadCmd.RunE(cmd, []string{}); err != nil {
					log.Fatalf("Error: %v\n", err)
				}
			case 1:
				if err := commands.PathCmd.RunE(cmd, []string{}); err != nil {
					log.Fatalf("Error: %v\n", err)
				}
			case 2:
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
