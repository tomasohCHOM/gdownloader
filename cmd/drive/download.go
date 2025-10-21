package drive

import (
	"io"
	"os"
	"os/user"
	"path/filepath"
	"strings"

	"google.golang.org/api/drive/v3"
)

func DownloadFile(srv *drive.Service, fileId, fileName string, path string) error {
	if strings.HasPrefix(path, "~") {
		usr, _ := user.Current()
		homeDir := usr.HomeDir
		path = filepath.Join(homeDir, strings.TrimPrefix(path, "~"))
	}
	resp, err := srv.Files.Export(fileId, "application/pdf").Download()
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	outPath := filepath.Join(path, fileName)
	out, err := os.Create(outPath)
	if err != nil {
		return err
	}
	defer out.Close()
	_, err = io.Copy(out, resp.Body)
	return err
}
