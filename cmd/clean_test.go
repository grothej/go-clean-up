package cmd

import (
	"os"
	"path/filepath"
	"testing"
)

func setupTestDir(fileNames []string, t *testing.T) (path string) {
	dir, err := os.MkdirTemp("", "test-dir")
	if err != nil {
		t.Fatalf("Failed to create dir: '%s'\nerr: %v", dir, err)
	}

	for _, file := range fileNames {
		path := filepath.Join(dir, file)
		_, err = os.Create(path)
		if err != nil {
			t.Fatalf("Failed to create file: '%s'\nerr: %v", dir, err)
		}
	}

	t.Cleanup(func() {
		err = os.RemoveAll(dir)
		if err != nil {
			t.Fatalf("Couldn't remove dir and files: %v", err)
		}
	})

	return dir
}

func assertFilesAreDeleted(filenames []string, t *testing.T) {
	for _, pattern := range filenames {
		wd, _ := os.Getwd()
		path := filepath.Join(wd, pattern)
		if matches, _ := filepath.Glob(path); len(matches) > 0 {
			t.Errorf("Expected '%s' to be deleted but was still found in '%s'", pattern, wd)
		}
	}
}

func TestDefaultClean(t *testing.T) {
	filesToDelete := []string{"old-log.log", ".DS_Store"}
	dir := setupTestDir(filesToDelete, t)

	Clean(dir)

	assertFilesAreDeleted(filesToDelete, t)
}

func TestCleanByExtension(t *testing.T) {
	filesToDelete := []string{"old-log.log", ".DS_Store"}
	fileToKeep := []string{"important.txt", "secret.yaml"}
	filesToSetup := append(fileToKeep, filesToDelete...)
	dir := setupTestDir(filesToSetup, t)

	Clean(dir)

	assertFilesAreDeleted(filesToDelete, t)
}
