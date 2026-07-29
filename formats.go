package common

import (
	"embed"
	"encoding/json"
	"io/fs"
	"os"
	"path/filepath"
	"slices"
	"strings"
)

//go:embed formats.json
var embedFS embed.FS

type FormatID string

type Format struct {
	ID          FormatID `json:"id"`
	Description string   `json:"description"`
	Extensions  []string `json:"extensions"`
}

type Formats struct {
	Formats []Format   `json:"formats"`
	BIM     []FormatID `json:"bim"`
	CAD     []FormatID `json:"cad"`
}

func parseFormats() Formats {
	var result Formats
	f, err := embedFS.Open("formats.json")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	err = json.NewDecoder(f).Decode(&result)
	if err != nil {
		panic(err)
	}
	return result
}

var formats = parseFormats()

func GetSupportedFormats() []Format {
	return formats.Formats
}

func DetectSupportedFiles(dir string) ([]string, int64) {
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
		ok := slices.ContainsFunc(formats.Formats, func(f Format) bool {
			return slices.Contains(f.Extensions, ext)
		})
		if ok {
			mainFiles = append(mainFiles, p)
		}
		return nil
	})

	return mainFiles, totalSize
}

func IsBIM(id FormatID) bool {
	return slices.Contains(formats.BIM, id)
}

func IsCAD(id FormatID) bool {
	return slices.Contains(formats.CAD, id)
}

func GetFormatID(filename string) FormatID {
	ext := strings.ToLower(filepath.Ext(filename))
	for _, f := range formats.Formats {
		if slices.Contains(f.Extensions, ext) {
			return f.ID
		}
	}
	return ""
}
