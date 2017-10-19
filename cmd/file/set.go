package file

import (
	"fmt"
	"github.com/olebedev/config"
	"github.com/spf13/cobra"
)

// SetOptions contains the options passed to the Set command
type SetOptions struct {
	Value  string
	DryRun bool
}

func NewCmdSet() *cobra.Command {

	options := &SetOptions{}

	cmd := &cobra.Command{
		Use:   "set",
		Short: "Print the version number of Hugo",
		Long:  `All software has versions. This is Hugo's`,
		Args:  validateArgs,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Hugo Static Site Generator v0.9 -- HEAD")
		},
	}

	cmd.MarkFlagRequired("value")
	cmd.Flags().StringVarP(&options.ValueFlag, "value", "v", "", "value to set in the targeted key")
	cmd.Flags().BoolVar(&options.DryRun, "dry-run", false, "print the changes to stdOut without changing the file")

	return cmd
}
