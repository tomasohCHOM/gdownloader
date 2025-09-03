package commands

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tomasohCHOM/gdownloader/cmd/store"
)

func init() {
	pathAddCmd.Flags().StringP("alias", "a", "", "Alias of the path to add")
	pathAddCmd.Flags().StringP("dir", "d", "", "Directory path to add")
	pathRemoveCmd.Flags().StringP("alias", "a", "", "Alias of the path to remove")

	PathCmd.AddCommand(pathAddCmd)
	PathCmd.AddCommand(pathRemoveCmd)
	PathCmd.AddCommand(pathListCmd)
}

var PathCmd = &cobra.Command{
	Use:   "path",
	Short: "Manage paths where you can download Google Drive files to",
}

var pathAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Store a new path",
	RunE: func(cmd *cobra.Command, args []string) error {
		alias := cmd.Flag("alias").Value.String()
		dir := cmd.Flag("dir").Value.String()
		store, err := store.Load()
		if err != nil {
			return err
		}
		_, ok := store.Paths[alias]
		if ok {
			fmt.Println("Path already exits:", dir)
			return nil
		}
		store.Paths[alias] = dir
		if err := store.Save(); err != nil {
			return err
		}
		return nil
	},
}

var pathRemoveCmd = &cobra.Command{
	Use:   "remove [path]",
	Short: "Remove a path",
	RunE: func(cmd *cobra.Command, args []string) error {
		alias := cmd.Flag("alias").Value.String()
		store, err := store.Load()
		if err != nil {
			return err
		}
		_, ok := store.Paths[alias]
		if !ok {
			fmt.Println("Path alias not found:", alias)
			return nil
		}
		delete(store.Paths, alias)
		if err := store.Save(); err != nil {
			return err
		}
		fmt.Println("Removed path with alias:", alias)
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
		for alias, dir := range store.Paths {
			fmt.Printf("%s: %s\n", alias, dir)
		}
		return nil
	},
}
