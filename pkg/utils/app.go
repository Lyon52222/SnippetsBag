package utils

import (
	"log"
	"os"
	"os/exec"
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

func OpenSnippetWithEditor(path string) error {
	cmd := exec.Command("nvim", path)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}
