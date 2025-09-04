package cmd

import (
	"fmt"
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
	Run: func(cmd *cobra.Command, args []string) {
		for {
			fmt.Println("\nMain Menu")
			fmt.Println("1. Download files")
			fmt.Println("2. Manage paths")
			fmt.Println("3. Exit")

			var choice int
			fmt.Print("Choose an option: ")
			fmt.Scan(&choice)

			switch choice {
			case 1:
				if err := commands.DownloadCmd.RunE(cmd, []string{}); err != nil {
					fmt.Println("Error:", err)
				}
			case 2:
				if err := commands.PathCmd.RunE(cmd, []string{}); err != nil {
					fmt.Println("Error:", err)
				}
			case 3:
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
