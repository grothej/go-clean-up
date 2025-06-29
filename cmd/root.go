package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var Dir string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "go-clean-up [--dir string] [--dry-run]",
	Short: "Cli Tool for cleaning unused files on your system",
	Long:  ` "go-clean-up" helps with cleaning old, unused and redundant files from your system`,
	Run:   func(cmd *cobra.Command, args []string) { clean() },
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}

}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	var cwd, _ = os.Getwd()

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.go-clean-up.git.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.PersistentFlags().StringVar(&Dir, "dir", cwd, "--dir directory to clean up (default is working dir)")
}

func clean() {
	fmt.Printf("cleaning: %s", Dir)
}
