package commands

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/tomasohCHOM/gdownloader/cmd/drive"
	"github.com/tomasohCHOM/gdownloader/cmd/options"
	"github.com/tomasohCHOM/gdownloader/cmd/store"
	"github.com/tomasohCHOM/gdownloader/cmd/ui/selector"
	"github.com/tomasohCHOM/gdownloader/cmd/ui/text"
	gdrive "google.golang.org/api/drive/v3"
)

type Page struct {
	Files     []*gdrive.File
	PageToken string
}

var DownloadCmd = &cobra.Command{
	Use:   "download",
	Short: "Download Google Drive files to a specified path",
	RunE: func(cmd *cobra.Command, args []string) error {
		srv, err := drive.Setup()
		if err != nil {
			return err
		}
		store, err := store.Load()
		if err != nil {
			return err
		}
		if len(store.Paths) == 0 {
			fmt.Println("No paths saved. Use the path command to store some paths")
			return nil
		}
		paths := make([]string, 0, len(store.Paths))
		for alias, dir := range store.Paths {
			paths = append(paths, fmt.Sprintf("%s: %s", alias, dir))
		}
		_, selected, err := selector.RunSelector("Choose the destination path", paths)
		if err != nil {
			return err
		}
		path := strings.Split(selected, ": ")[1]
		for {
			searchQuery, exited, err := text.RunTextInput("Enter a search query to list some files")
			if err != nil {
				return err
			}
			if exited {
				return nil
			}
			fmt.Println("Searching...")
			call := drive.Search(srv, searchQuery)
			resp, err := call.Do()
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error while searching: %v\n", err)
				continue
			}
			if len(resp.Files) == 0 {
				fmt.Println("No files found for that query.")
				continue
			}
			pages := []Page{{Files: resp.Files, PageToken: resp.NextPageToken}}
			pageIndex := 0
			for {
				currentPage := pages[pageIndex]
				files := currentPage.Files
				fileOptions := make([]string, 0, len(files)+3)
				for _, file := range files {
					fileOptions = append(fileOptions, fmt.Sprintf("%s", file.Name))
				}
				if pageIndex > 0 {
					fileOptions = append(fileOptions, options.PREVIOUS_PAGE_PROMPT)
				}
				if currentPage.PageToken != "" {
					fileOptions = append(fileOptions, options.NEXT_PAGE_PROMPT)
				}
				fileOptions = append(fileOptions, options.RETRY_SEARCH_PROMPT)
				fileOptions = append(fileOptions, options.EXIT)
				_, selected, err := selector.RunSelector("Select a file to download", fileOptions)
				if err != nil {
					fmt.Fprintf(os.Stderr, "Error: %v\n", err)
					return err
				}
				switch selected {
				case options.EXIT:
					return nil
				case options.RETRY_SEARCH_PROMPT:
					goto nextSearch
				case options.NEXT_PAGE_PROMPT:
					pageIndex++
					if pageIndex == len(pages) {
						if currentPage.PageToken == "" {
							fmt.Println("No more pages.")
							pageIndex--
							continue
						}
						call := drive.Search(srv, searchQuery).PageToken(currentPage.PageToken)
						resp, err := call.Do()
						if err != nil {
							fmt.Fprintf(os.Stderr, "Error fetching next page: %v\n", err)
							pageIndex--
							continue
						}
						pages = append(pages, Page{Files: resp.Files, PageToken: resp.NextPageToken})
					}
					continue
				case options.PREVIOUS_PAGE_PROMPT:
					if pageIndex > 0 {
						pageIndex--
					}
					continue
				default:
					var selectedFileId string
					for _, file := range files {
						if file.Name == selected {
							selectedFileId = file.Id
							break
						}
					}
					if selectedFileId == "" {
						fmt.Fprintf(os.Stderr, "Could not find file ID for %s\n", selected)
						continue
					}
					fmt.Printf("Downloading %s...\n", selected)
					return drive.DownloadFile(srv, selectedFileId, selected, path)
				}
			}
		nextSearch:
			continue
		}
	},
}
