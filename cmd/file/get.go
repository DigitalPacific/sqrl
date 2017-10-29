package file

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewCmdGet() *cobra.Command {

	var fileValues FileValues

	cmd := &cobra.Command{
		Use:   "get",
		Short: "Get File command",
		Args: func(cmd *cobra.Command, args []string) error {
			fileValues.ValidateArgs(args)
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			fileValues.ValidateFlags()
			value, err := fileValues.get()
			if err != nil {
				panic(err)
			}
			rendered, _ := fileValues.render(value)
			fmt.Println(rendered)
		},
	}

	cmd.Flags().StringVarP(&fileValues.Key, "key", "k", "", "key in the file to target")
	return cmd
}
