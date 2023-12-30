package jobs

import (
	"api/src/config/logger"
	"archive/zip"
	"io"
	"os"
	"path/filepath"
)

func Compress(folderPath string, zipName string) {
	zipFile, err := os.Create(zipName)
	if err != nil {
		logger.Error("Error create file zip", err)
		return
	}
	defer zipFile.Close()

	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	err = filepath.Walk(folderPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()

			relativePath, err := filepath.Rel(folderPath, path)
			if err != nil {
				return err
			}

			zipFile, err := zipWriter.Create(relativePath)
			if err != nil {
				return err
			}

			_, err = io.Copy(zipFile, file)
			if err != nil {
				return err
			}

			if err := os.Remove(path); err != nil {
				logger.Error("Error remove file compacted", err)
			}
		}

		return nil
	})

	if err != nil {
		logger.Error("Error compact files", err)
		return
	}

	logger.Info("Successfully compressed files")
}
