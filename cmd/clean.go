package cmd

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

// cleanCmd represents the clean command
var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
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
			fmt.Println("Skip path ", path, ". Is a directory")
			return nil
		}
		if isFileCleanable(info) {
			fmt.Printf("File: %s would be removed\n", info.Name())
			clearedDiscSpace += int(info.Size())
		}

		return nil
	})
	fmt.Println("Cleaned disc space:", clearedDiscSpace)

	if err != nil {
		return
	}
}

func isFileCleanable(info fs.FileInfo) bool {
	ext := filepath.Ext(info.Name())
	return isExtensionCleanable(ext)

}

func isExtensionCleanable(ext string) bool {
	return true

}
