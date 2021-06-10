package utils

import (
	"log"
	"os"
	"path/filepath"
)

func SetupApp(snippetsDir string) error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Panicln(err)
	} else {
		snippetsDir = filepath.Join(homeDir, ".snippets")
		if !IsExist(snippetsDir) {
			NewDir(snippetsDir)
		}
	}
	return nil
}
