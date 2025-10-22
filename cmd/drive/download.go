package drive

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"google.golang.org/api/drive/v3"
)

func DownloadFile(srv *drive.Service, fileId, fileName string, path string) error {
	if strings.HasPrefix(path, "~") {
		home, err := os.UserHomeDir()
		if err != nil {
			return err
		}
		path = filepath.Join(home, strings.TrimPrefix(path, "~"))
	}
	file, err := srv.Files.Get(fileId).Fields("mimeType", "name").Do()
	if err != nil {
		return fmt.Errorf("failed to get file metadata: %w", err)
	}
	mime := file.MimeType
	exportMime, extension := getPreferredExportFormat(mime)
	var resp *http.Response
	if strings.HasPrefix(mime, "application/vnd.google-apps") {
		resp, err = srv.Files.Export(fileId, exportMime).Download()
	} else {
		resp, err = srv.Files.Get(fileId).Download()
	}
	if err != nil {
		return fmt.Errorf("failed to download file: %w", err)
	}
	defer resp.Body.Close()
	outPath := filepath.Join(path, fileName+extension)
	out, err := os.Create(outPath)
	if err != nil {
		return err
	}
	defer out.Close()
	if _, err := io.Copy(out, resp.Body); err != nil {
		return fmt.Errorf("failed to save file: %w", err)
	}
	return nil
}

func getPreferredExportFormat(mimeType string) (exportMime, extension string) {
	switch mimeType {
	case "application/vnd.google-apps.document":
		return "application/pdf", ".pdf"
	case "application/vnd.google-apps.spreadsheet":
		return "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet", ".xlsx"
	case "application/vnd.google-apps.presentation":
		return "application/vnd.openxmlformats-officedocument.presentationml.presentation", ".pptx"
	case "application/vnd.google-apps.drawing":
		return "image/png", ".png"
	case "application/vnd.google-apps.script":
		return "application/vnd.google-apps.script+json", ".json"
	case "application/vnd.google-apps.form":
		return "text/plain", ".txt"
	default:
		return "", ""
	}
}
