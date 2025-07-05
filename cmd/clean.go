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
		Clean()
	},
}

func init() {
	rootCmd.AddCommand(cleanCmd)

}

func Clean() {
	var err error
	if Dir == "" {
		Dir, err = os.Getwd()
	}

	if err != nil {
		return
	}

	var clearedDiscSpace int
	err = filepath.Walk(Dir, func(path string, info fs.FileInfo, err error) error {
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
