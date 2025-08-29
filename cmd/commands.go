package cmd

import (
	"fmt"
	"slices"

	"github.com/spf13/cobra"
	"github.com/tomasohCHOM/google-drive-downloader/cmd/store"
)

func init() {
	rootCmd.AddCommand(pathCmd)
	pathCmd.AddCommand(pathAddCmd)
	pathCmd.AddCommand(pathRemoveCmd)
	pathCmd.AddCommand(pathListCmd)

	rootCmd.AddCommand(downloadCmd)
}

var pathCmd = &cobra.Command{
	Use:   "path",
	Short: "Manage paths where you can download Google Drive files to",
}

var pathAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Store a new path",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		dir := args[0]
		store, err := store.Load()
		if err != nil {
			return err
		}
		if slices.Contains(store.Paths, dir) {
			fmt.Println("Path already exits:", dir)
			return nil
		}
		store.Paths = append(store.Paths, dir)
		if err := store.Save(); err != nil {
			return err
		}
		return nil
	},
}

var pathRemoveCmd = &cobra.Command{
	Use:   "remove [path]",
	Short: "Remove a path",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		dir := args[0]
		store, err := store.Load()
		if err != nil {
			return err
		}
		newPaths := []string{}
		found := false
		for _, p := range store.Paths {
			if p != dir {
				newPaths = append(newPaths, p)
			} else {
				found = true
			}
		}
		if !found {
			fmt.Println("Path not found:", dir)
			return nil
		}
		store.Paths = newPaths
		if err := store.Save(); err != nil {
			return err
		}
		fmt.Println("Removed path:", dir)
		return nil
	},
}

var pathListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all saved paths",
	RunE: func(cmd *cobra.Command, args []string) error {
		store, err := store.Load()
		if err != nil {
			return err
		}

		if len(store.Paths) == 0 {
			fmt.Println("No paths saved.")
			return nil
		}

		for i, p := range store.Paths {
			fmt.Printf("%d. %s\n", i+1, p)
		}
		return nil
	},
}

var downloadCmd = &cobra.Command{
	Use:   "download",
	Short: "Download Google Drive files to a specified path",

	Run: func(cmd *cobra.Command, args []string) {
	},
}
