package common

import (
	"archive/zip"
	"context"
	"os"
	"path/filepath"
	"strings"

	"github.com/mholt/archives"
)

const (
	DefaultAssetName = "asset.zip"
)

const (
	pathSep = string(os.PathSeparator)
)

func BuildZip(assetFilesDir string, outputDir string, assetName string) error {
	if assetName == "" {
		assetName = DefaultAssetName
	}

	zipPath := filepath.Join(outputDir, assetName)
	if err := os.MkdirAll(filepath.Dir(zipPath), 0o755); err != nil {
		return err
	}

	zipFile, err := os.Create(zipPath)
	if err != nil {
		return err
	}
	defer zipFile.Close()

	if !strings.HasSuffix(assetFilesDir, pathSep) {
		assetFilesDir += pathSep
	}

	ctx := context.TODO()

	files, err := archives.FilesFromDisk(ctx, nil, map[string]string{
		assetFilesDir: "",
	})
	if err != nil {
		return err
	}

	format := archives.CompressedArchive{
		Archival: archives.Zip{
			Compression: zip.Deflate,
		},
	}

	return format.Archive(ctx, zipFile, files)
}
