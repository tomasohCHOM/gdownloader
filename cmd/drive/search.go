package drive

import (
	"fmt"

	"google.golang.org/api/drive/v3"
)

func Search(srv *drive.Service, query string) *drive.FilesListCall {
	return srv.Files.List().
		Q(fmt.Sprintf("name contains '%s'", query)).
		PageSize(10).
		Fields("nextPageToken, files(id, name, mimeType)")
}
