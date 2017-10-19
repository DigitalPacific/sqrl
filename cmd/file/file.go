package file

import (
	"fmt"

	"github.com/spf13/cobra"
)

type FileOptions struct {
	Key string
}

func NewCmdFile() *cobra.Command {

	options := new(FileOptions)

	cmd := &cobra.Command{
		Use:   "file",
		Short: "Main File Command",
		Long:  `Long file description`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("file called")
		},
	}

	cmd.MarkFlagRequired("key")
	cmd.Flags().StringVarP(&options.Key, "key", "k", "", "key in the file to target")

	cmd.AddCommand(NewCmdSet())
	cmd.AddCommand(NewCmdGet())

	return cmd
}
