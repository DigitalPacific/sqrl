package cmd

import (
	"github.com/DigitalPacific/squirrel/cmd/file"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(file.NewCmdFile())
}

// RootCmd represents the root command
var RootCmd = &cobra.Command{
	Use:   "squirrel",
	Short: "A colleciton of CLI tools for doing CI work",
	Long: `Squirrel is a collection of CLI tools which help simplify the
	work required for CI tasks.`,
}
