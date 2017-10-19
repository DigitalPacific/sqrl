package cmd

import (
	"ci/cmd/file"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(file.NewCmdFile())
}

// RootCmd represents the root command
var RootCmd = &cobra.Command{
	Use:   "ci",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
		and usage of using your command. For example:
		Cobra is a CLI library for Go that empowers applications.
		This application is a tool to generate the needed files
		to quickly create a Cobra application.`,
}
