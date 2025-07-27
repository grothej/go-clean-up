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

	Clean(fsys)

	if err := fstest.TestFS(fsys, "keep.txt"); err != nil {
		t.Fatal(err)
	}
}

func TestCleanByExtension(t *testing.T) {
	fsys := fstest.MapFS{
		"old-log.log":   {Data: []byte("remove this")},
		".DS_Store":     {Data: []byte("system file")},
		"important.txt": {Data: []byte("system file")},
		"secret.yaml":   {Data: []byte("apiVersion: secret/v1")},
	}

	Clean(fsys)

	if err := fstest.TestFS(fsys, "secret.yaml", "important.txt"); err != nil {
		t.Fatal(err)
	}
}
