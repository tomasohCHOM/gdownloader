package commands

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/tomasohCHOM/gdownloader/cmd/options"
	"github.com/tomasohCHOM/gdownloader/cmd/store"
	"github.com/tomasohCHOM/gdownloader/cmd/ui/selector"
	"github.com/tomasohCHOM/gdownloader/cmd/ui/text"
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
		header := "Choose one of the following path actions to continue:"
		for {
			_, selected, err := selector.RunSelector(header, options.PATH_CMD_OPTIONS)
			if err != nil {
				if err.Error() == selector.NO_SELECTION {
					return nil
				}
				fmt.Fprintf(os.Stderr, "Selection error: %v\n", err)
				return nil
			}
			switch selected {
			case options.ADD_PATH:
				err := pathAddCmd.RunE(pathAddCmd, args)
				if err != nil {
					return err
				}
			case options.REMOVE_PATH:
				err := pathRemoveCmd.RunE(pathRemoveCmd, args)
				if err != nil {
					return err
				}
			case options.LIST_PATHS:
				err := pathListCmd.RunE(pathListCmd, args)
				if err != nil {
					return err
				}
			case options.EXIT:
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
		if len(alias) == 0 {
			aliasInput, exited, err := text.RunTextInput("Enter the alias of the path to add")
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error: %v\n", err)
				return err
			}
			if exited {
				return nil
			}
			alias = aliasInput
		}
		dir := cmd.Flag("dir").Value.String()
		if len(dir) == 0 {
			dirInput, exited, err := text.RunTextInput("Enter the directory path to add")
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error: %v\n", err)
				return err
			}
			if exited {
				return nil
			}
			dir = dirInput
		}
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
		if len(alias) == 0 {
			aliasInput, exited, err := text.RunTextInput("Enter the alias of the path to remove")
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error: %v\n", err)
				return err
			}
			if exited {
				return nil
			}
			alias = aliasInput
		}
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
