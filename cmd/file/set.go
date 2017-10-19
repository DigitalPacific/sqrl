package file

import (
	"fmt"

	"github.com/spf13/cobra"
)

type SetOptions struct {
	ValueFlag string
	DryRun    bool
}

func NewCmdSet() *cobra.Command {

	options := new(SetOptions)

	cmd := &cobra.Command{
		Use:   "set",
		Short: "Set file command",
		Long:  `Long set description`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("file set called")
		},
	}

	cmd.MarkFlagRequired("value")
	cmd.Flags().StringVarP(&options.ValueFlag, "value", "v", "", "value to set in the targeted key")

	return cmd
}
