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
	RunE: func(cmd *cobra.Command, args []string) error {
		// interactive menu here
		for {
			fmt.Println("\nPath Menu")
			fmt.Println("1. Add path")
			fmt.Println("2. Remove path")
			fmt.Println("3. List paths")
			fmt.Println("4. Exit")
			var choice int
			fmt.Print("Choose an option: ")
			fmt.Scan(&choice)

			switch choice {
			case 1:
				return pathAddCmd.RunE(pathAddCmd, args)
			case 2:
				return pathRemoveCmd.RunE(pathRemoveCmd, args)
			case 3:
				return pathListCmd.RunE(pathListCmd, args)
			case 4:
				return nil
			default:
				fmt.Println("Invalid choice")
			}
		}
	},
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
