package file

import (
	"github.com/spf13/cobra"
)

func NewCmdFile() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "file",
		Short: "Main File Command",
		Long:  `Long file description`,
	}

	cmd.AddCommand(NewCmdSet())
	cmd.AddCommand(NewCmdGet())
	return cmd
}
