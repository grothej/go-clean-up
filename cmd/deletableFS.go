package cmd

import (
	"fmt"
	"io/fs"
	"os"
)

type DeletableFS interface {
	Remove(name string) error
	GetFsys() fs.FS
}

type DeletableFsys struct {
	fsys fs.FS
}

func (dFsys DeletableFsys) GetFsys() fs.FS {
	return dFsys.fsys
}

func (dFsys DeletableFsys) Remove(path string) error {
	err := os.Remove(path)
	if err != nil {
		fmt.Println("Couldn't remove ", path)
		return err
	}

	return nil
}
