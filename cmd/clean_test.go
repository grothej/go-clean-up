package cmd

import (
	"testing"
	"testing/fstest"
)

func TestDefaultClean(t *testing.T) {
	fsys := fstest.MapFS{
		"keep.txt":   {Data: []byte("important")},
		"delete.log": {Data: []byte("remove this")},
		".DS_Store":  {Data: []byte("system file")},
	}

	Clean()

	fstest.TestFS(fsys, "keep.txt")
}

func TestCleanByExtension(t *testing.T) {
	filesToDelete := []string{"old-log.log", ".DS_Store"}
	fileToKeep := []string{"important.txt", "secret.yaml"}

	Clean(dir)

	fstest.TestFS(fsys)
}
