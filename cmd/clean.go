package cmd

import (
	"fmt"
	"io/fs"
	"os"

	"github.com/grothej/go-clean-up.git/cmd/clean"

	"github.com/spf13/cobra"
)

// cleanCmd represents the clean command
var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "TBD",
	Long:  `TBD`,
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		if Dir == "" {
			Dir, err = os.Getwd()
		}
		if err != nil {
			return
		}
		root := Dir
		fsys := os.DirFS(root)

		Clean(fsys)
	},
}

func init() {
	rootCmd.AddCommand(cleanCmd)

}

func Clean(fsys fs.FS) {
	var clearedDiscSpace int
	err := fs.WalkDir(fsys, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return fs.SkipDir
		}

		info, err := fs.Stat(fsys, path)
		if err != nil {
			fmt.Println("Coudn't read info from ", path)
			return nil
		}
		if info.IsDir() {
			return nil
		}

		if isFileCleanable(info) {
			err := os.Remove(path)
			if err != nil {
				fmt.Printf("Couldn't remove %s \n%s\n", path, err)
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
		fmt.Println(err)
	}
}

func isFileCleanable(info fs.FileInfo) bool {
	return clean.IsFileOlderThanTTL(info) || clean.IsExtensionCleanable(info)
}
