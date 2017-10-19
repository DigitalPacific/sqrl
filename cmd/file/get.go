package file

import (
	"fmt"
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Print the version number of Hugo",
	Long:  `All software has versions. This is Hugo's`,
	Args:  validateArgs,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hugo Static Site Generator v0.9 -- HEAD")
	},
}
