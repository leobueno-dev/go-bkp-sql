package utils

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

func ZipFile(source, target string) error {
	zipCmd := exec.Command("zip", "-j", target, source)
	err := zipCmd.Run()
	return err
}

func DeleteOldFiles(directory string, daysAgo int) error {
	cutoff := time.Now().AddDate(0, 0, -daysAgo)

	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.ModTime().Before(cutoff) {
			err := os.Remove(path)
			if err != nil {
				log.Printf("Erro ao deletar o arquivo %s: %s\n", path, err)
			}
		}
		return nil
	})

	return err
}
