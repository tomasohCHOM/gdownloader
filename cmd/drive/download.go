package drive

import (
	"io"
	"os"
	"path/filepath"

	"google.golang.org/api/drive/v3"
)

func DownloadFile(srv *drive.Service, fileId, fileName string) error {
	resp, err := srv.Files.Export(fileId, "application/pdf").Download()
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	outPath := filepath.Join(fileName)
	out, err := os.Create(outPath)
	if err != nil {
		return err
	}
	defer out.Close()
	_, errr := io.Copy(out, resp.Body)
	return errr
}
