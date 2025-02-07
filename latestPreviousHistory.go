package main

import (
	"os"
	"path/filepath"
	"time"
)

func getLatestTxtFile(dir string) (string, error) {
	var latestFile string
	var latestTime time.Time

	files, err := os.ReadDir(dir)
	if err != nil {
		return "", err
	}

	for _, file := range files {
		if filepath.Ext(file.Name()) == ".txt" {
			info, err := file.Info()
			if err != nil {
				continue
			}

			if info.ModTime().After(latestTime) {
				latestTime = info.ModTime()
				latestFile = file.Name()
			}
		}
	}

	if latestFile == "" {
		return "", nil
	}

	return latestFile, nil
}
