package cmd

import (
	"fmt"

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
	fmt.Printf("cleaning: %s\n", Dir)
	fmt.Printf("is dry run? -> %v\n", *IsDryRun)
}
