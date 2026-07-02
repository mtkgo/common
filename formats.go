package common

import (
	"embed"
	"encoding/json"
	"io/fs"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"sync"
)

//go:embed formats.json
var embedFS embed.FS

type Format struct {
	ID          string   `json:"id"`
	Description string   `json:"description"`
	Extensions  []string `json:"extensions"`
}

func GetSupportedFormats() []Format {
	var formats []Format
	f, err := embedFS.Open("formats.json")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	err = json.NewDecoder(f).Decode(&formats)
	if err != nil {
		panic(err)
	}
	return formats
}

var (
	once             sync.Once
	supportedFormats []Format
)

func DetectSupportedFiles(dir string) ([]string, int64) {
	once.Do(func() {
		supportedFormats = GetSupportedFormats()
	})

	var mainFiles []string
	var totalSize int64

	fs.WalkDir(os.DirFS(dir), ".", func(p string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		info, err := d.Info()
		if err != nil {
			return err
		}
		totalSize += info.Size()
		ext := strings.ToLower(filepath.Ext(p))
		ok := slices.ContainsFunc(supportedFormats, func(f Format) bool {
			return slices.Contains(f.Extensions, ext)
		})
		if ok {
			mainFiles = append(mainFiles, p)
		}
		return nil
	})

	return mainFiles, totalSize
}
