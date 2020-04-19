package ui

import (
	"path/filepath"
)

func completeFilePath(t string) []string {
	if t == "" {
		return nil
	}

	matches, _ := filepath.Glob(t + "*")
	return matches
}
