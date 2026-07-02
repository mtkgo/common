package common

import (
	"archive/zip"
	"context"
	"io"
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

func BuildZip(ctx context.Context, assetFilesDir string, outputDir string, assetName string) error {
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

	inputDir := filepath.Clean(assetFilesDir)
	if !strings.HasSuffix(inputDir, pathSep) {
		inputDir += pathSep
	}

	files, err := archives.FilesFromDisk(ctx, nil, map[string]string{
		// rootOnDisk -> rootInArchive
		inputDir: "",
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

func UnZip(ctx context.Context, inputFile string, outputDir string) error {
	zipFile, err := os.Open(inputFile)
	if err != nil {
		return err
	}
	defer zipFile.Close()

	var format archives.Zip
	return format.Extract(context.TODO(), zipFile, func(ctx context.Context, info archives.FileInfo) error {
		if info.IsDir() {
			dir := filepath.Join(outputDir, info.NameInArchive)
			return os.MkdirAll(dir, 0o755)
		}

		// NOTE: only dir/file in out case

		inFile, err := info.Open()
		if err != nil {
			return err
		}
		defer inFile.Close()

		path := filepath.Join(outputDir, info.NameInArchive)
		os.MkdirAll(filepath.Dir(path), 0o755)
		outFile, err := os.Create(path)
		if err != nil {
			return err
		}
		// os.Chmod(outPath, info.Mode())
		defer outFile.Close()

		_, err = io.Copy(outFile, inFile)
		return err
	})
}
