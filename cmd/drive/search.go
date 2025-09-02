package drive

import (
	"fmt"

	"google.golang.org/api/drive/v3"
)

func Search(srv *drive.Service, query string) ([]*drive.File, error) {
	r, err := srv.Files.List().
		Q(fmt.Sprintf("name contains '%s'", query)).
		PageSize(10).
		Fields("files(id, name, mimeType)").Do()
	if err != nil {
		return nil, err
	}
	return r.Files, nil
}
