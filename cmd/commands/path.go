package commands

import (
	"fmt"
	"log"
	"strings"

	"github.com/spf13/cobra"
	"github.com/tomasohCHOM/gdownloader/cmd/options"
	"github.com/tomasohCHOM/gdownloader/cmd/store"
	"github.com/tomasohCHOM/gdownloader/cmd/ui/selector"
	"github.com/tomasohCHOM/gdownloader/cmd/ui/styles"
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
	Run: func(cmd *cobra.Command, args []string) {
		prompt := "Choose one of the following path actions to continue:"
		for {
			selected, exited, err := selector.RunSelector(prompt, options.PATH_CMD_OPTIONS)
			if err != nil {
				log.Fatalf("Selection error: %v", err)
			}
			if exited {
				return
			}
			switch selected {
			case options.ADD_PATH:
				pathAddCmd.Run(pathAddCmd, args)
			case options.REMOVE_PATH:
				pathRemoveCmd.Run(pathRemoveCmd, args)
			case options.LIST_PATHS:
				pathListCmd.Run(pathListCmd, args)
			case options.EXIT:
				return
			default:
				log.Fatalf("Invalid command selection")
			}
		}
	},
}

var pathAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Store a new path",
	Run: func(cmd *cobra.Command, args []string) {
		alias := cmd.Flag("alias").Value.String()
		if len(alias) == 0 {
			aliasInput, exited, err := text.RunTextInput("Enter the alias of the path to add:")
			if err != nil {
				log.Fatalf("Error while processing text input: %v", err)
			}
			if exited {
				return
			}
			alias = aliasInput
		}
		dir := cmd.Flag("dir").Value.String()
		if len(dir) == 0 {
			dirInput, exited, err := text.RunTextInput("Enter the directory path to add:")
			if err != nil {
				log.Fatalf("Error while processing text input: %v", err)
			}
			if exited {
				return
			}
			dir = dirInput
		}
		pathExists, err := store.CheckPathExists(dir)
		if err != nil {
			log.Fatalf("Failed to check path: %v", err)
		}
		if !pathExists {
			fmt.Println(styles.ErrorStyle.Render("\nInvalid directory path, ensure this path exists"))
			return
		}
		store, err := store.Load()
		if err != nil {
			log.Fatalf("Failed to load path store: %v", err)
		}
		_, ok := store.Paths[alias]
		if ok {
			fmt.Println(styles.DimStyle.Render(fmt.Sprintf("\nPath already exits: %s", dir)))
			return
		}
		store.Paths[alias] = dir
		if err := store.Save(); err != nil {
			log.Fatalf("Failed to add path to store: %v", err)
		}
		fmt.Println(styles.ContrastStyle.Render(fmt.Sprintf("\nAdded path with alias: %s", alias)))
	},
}

var pathRemoveCmd = &cobra.Command{
	Use:   "remove [path]",
	Short: "Remove a path",
	Run: func(cmd *cobra.Command, args []string) {
		alias := cmd.Flag("alias").Value.String()
		store, err := store.Load()
		if err != nil {
			log.Fatalf("Failed to load path store: %v", err)
		}
		if len(store.Paths) == 0 {
			fmt.Println(styles.DimStyle.Render("\nNo paths saved."))
			return
		}
		if len(alias) == 0 {
			storedPaths := make([]string, 0, len(store.Paths))
			for a, d := range store.Paths {
				storedPaths = append(storedPaths, fmt.Sprintf("%s (%s)", a, d))
			}
			aliasInput, exited, err := selector.RunSelector("Select the path to remove:", storedPaths)
			if err != nil {
				log.Fatalf("Selection error: %v", err)
			}
			if exited {
				return
			}
			alias = strings.Split(aliasInput, " (")[0]
		}
		if err != nil {
			log.Fatalf("Selection error: %v", err)
		}
		_, ok := store.Paths[alias]
		if !ok {
			fmt.Println(styles.ErrorStyle.Render(fmt.Sprintf("\nPath alias not found: %s", alias)))
			return
		}
		delete(store.Paths, alias)
		if err := store.Save(); err != nil {
			log.Fatalf("Failed to delete path from store: %v", err)
		}
		fmt.Println(styles.ContrastStyle.Render(fmt.Sprintf("\nRemoved path with alias: %s", alias)))
	},
}

var pathListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all saved paths",
	Run: func(cmd *cobra.Command, args []string) {
		store, err := store.Load()
		if err != nil {
			log.Fatalf("Failed to load path store: %v", err)
		}
		if len(store.Paths) == 0 {
			fmt.Println(styles.DimStyle.Render("\nNo paths saved."))
			return
		}
		fmt.Printf("\n%s\n\n", styles.ContrastStyle.Render("List of stored paths:"))
		for alias, dir := range store.Paths {
			fmt.Printf("- %s: %s\n", styles.NormalTextStyle.Render(alias), styles.DimStyle.Render(dir))
		}
	},
}
