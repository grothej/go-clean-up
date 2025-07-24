package cmd

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/grothej/go-clean-up.git/cmd/clean"

	"github.com/spf13/cobra"
)

// cleanCmd represents the clean command
var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "TBD",
	Long:  `TBD`,
	Run: func(cmd *cobra.Command, args []string) {
		Clean(Dir)
	},
}

func init() {
	rootCmd.AddCommand(cleanCmd)

}

func Clean(fsys fs.FS) {
	var err error
	if dir == "" {
		dir, err = os.Getwd()
	}

	if err != nil {
		return
	}

	var clearedDiscSpace int
	err = filepath.Walk(dir, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			fmt.Println("Skip path ", path, " : is a directory")
			return nil
		}

		if isFileCleanable(info) {
			err := os.Remove(path)
			if err != nil {
				fmt.Println("Couldn't remove ", path)
			} else {
				fmt.Println("Removed file: ", path)
				clearedDiscSpace += int(info.Size())
			}

			return err
		}

		return nil
	})

	fmt.Println("Total cleaned disc space:", clearedDiscSpace)
	if err != nil {
		return
	}
}

func isFileCleanable(info fs.FileInfo) bool {
	return clean.IsFileOlderThanTTL(info) || clean.IsExtensionCleanable(info)
}
