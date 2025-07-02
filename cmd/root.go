package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var Dir string
var IsDryRun bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "go-clean-up [--dir string] [--dry-run -d string]",
	Short: "Cli Tool for cleaning unused files on your system",
	Long:  ` "go-clean-up" helps with cleaning old, unused and redundant files from your system`,
	Run:   func(cmd *cobra.Command, args []string) { Clean() },
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	var cwd, _ = os.Getwd()

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.go-clean-up.git.yaml)")

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.PersistentFlags().BoolVarP(&IsDryRun, "dry-run", "d", true, "-d, --dry-run only output files that would be deleted (default false)")
	rootCmd.PersistentFlags().StringVar(&Dir, "dir", cwd, "--dir directory to clean up (default is working dir)")
}
