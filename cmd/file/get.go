package file

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewCmdGet() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get",
		Short: "Get File command",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("File get called")
		},
	}

	return cmd
}
