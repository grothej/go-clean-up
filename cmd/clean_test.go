package cmd

import (
	"io/fs"
	"testing"
	"testing/fstest"
	"time"
)

type deletableMapFS struct {
	mapFS fstest.MapFS
}

func (d deletableMapFS) GetFsys() fs.FS {
	return d.mapFS
}
func (d deletableMapFS) Remove(path string) error {
	info, err := d.mapFS.Stat(path)
	if err != nil {
		return err
	}
	delete(d.mapFS, info.Name())

	return nil
}

func TestDefaultClean(t *testing.T) {
	fsys := fstest.MapFS{
		"keep.txt":   {Data: []byte("important"), ModTime: time.Now().AddDate(0, -1, 0)},
		"delete.log": {Data: []byte("remove this")},
		".DS_Store":  {Data: []byte("system file")},
	}

	deletableFsys := deletableMapFS{
		mapFS: fsys,
	}

	Clean(deletableFsys)

	if err := fstest.TestFS(deletableFsys.GetFsys(), "keep.txt"); err != nil {
		t.Fatal(err)
	}
}

func TestCleanByExtension(t *testing.T) {
	fsys := fstest.MapFS{
		"old-log.log":   {Data: []byte("remove this")},
		".DS_Store":     {Data: []byte("system file")},
		"important.txt": {Data: []byte("system file"), ModTime: time.Now().AddDate(0, -1, 0)},
		"secret.yaml":   {Data: []byte("apiVersion: secret/v1"), ModTime: time.Now().AddDate(0, -1, 0)},
	}
	deletableFsys := deletableMapFS{
		mapFS: fsys,
	}

	Clean(deletableFsys)

	if err := fstest.TestFS(deletableFsys.GetFsys(), "secret.yaml", "important.txt"); err != nil {
		t.Fatal(err)
	}
}

func TestCleanByModTime(t *testing.T) {
	fsys := fstest.MapFS{
		"old-file.txt": {Data: []byte("remove this"), ModTime: time.Now().AddDate(0, -8, 0)},
		"new-file.txt": {Data: []byte("keep this"), ModTime: time.Now().AddDate(0, -1, 0)},
	}

	deletableFsys := deletableMapFS{
		mapFS: fsys,
	}

	Clean(deletableFsys)

	if err := fstest.TestFS(fsys, "new-file.txt"); err != nil {
		t.Fatal(err)
	}
}
