package commands

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
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
		header := "Choose which path actions you would like to execute:"
		options := []string{"Add path", "Remove path", "List paths", "Exit"}
		for {
			idx, _, err := selector.RunSelector(header, options)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Selection error: %v\n", err)
				return nil
			}
			switch idx {
			case 0:
				return pathAddCmd.RunE(pathAddCmd, args)
			case 1:
				return pathRemoveCmd.RunE(pathRemoveCmd, args)
			case 2:
				return pathListCmd.RunE(pathListCmd, args)
			case 3:
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
		var p *tea.Program
		alias := cmd.Flag("alias").Value.String()
		if len(alias) == 0 {
			p = tea.NewProgram(text.InitialTextModel("Enter the alias of the path to add", ""))
			if _, err := p.Run(); err != nil {
				fmt.Fprintf(os.Stderr, "Error: %v\n", err)
				return err
			}
		}
		dir := cmd.Flag("dir").Value.String()
		if len(dir) == 0 {
			p = tea.NewProgram(text.InitialTextModel("Enter the directory path to add", ""))
			if _, err := p.Run(); err != nil {
				fmt.Fprintf(os.Stderr, "Error: %v\n", err)
				return err
			}
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
