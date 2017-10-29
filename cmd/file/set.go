package file

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func NewCmdSet() *cobra.Command {

	var fileValues FileValues

	cmd := &cobra.Command{
		Use:   "set",
		Short: "Set file command",
		Long:  `Long set description`,
		Args: func(cmd *cobra.Command, args []string) error {
			fileValues.ValidateArgs(args)
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			fileValues.ValidateFlags()

			cfg, err := fileValues.get()

			if err != nil {
				panic(err)
			}

			newCfg, err := cfg.Get(fileValues.Key)
			if err == nil {
				value, err := fileValues.render(newCfg)
				if err != nil {
					panic(err)
				}
				fmt.Println("Current value: " + value)
			} else {
				fmt.Println("Creating new key")
			}

			err = fileValues.set(cfg)

			if err != nil {
				panic(err)
			}

			newValue, err := fileValues.render(cfg)
			if err != nil {
				panic(err)
			}

			fmt.Println("New value:")
			fmt.Println(newValue)
			if fileValues.DryRun {
				fmt.Println("-----DRY RUN-----")
				return
			}

			file, _ := os.Create(fileValues.File)
			defer file.Close()

			file.Write([]byte(newValue))
		},
	}

	cmd.MarkFlagRequired("key")
	cmd.MarkFlagRequired("value")

	cmd.Flags().StringVarP(&fileValues.Key, "key", "k", "", "key in the file to target")
	cmd.Flags().StringVarP(&fileValues.Value, "value", "v", "", "value to set in the targeted key")
	cmd.Flags().BoolVarP(&fileValues.DryRun, "dry-run", "d", false, "show changes without saving")

	return cmd
}
