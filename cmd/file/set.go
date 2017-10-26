package file

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/olebedev/config"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

type SetOptions struct {
	Key    string
	File   string
	Value  string
	DryRun bool
}

func NewCmdSet() *cobra.Command {

	var options SetOptions

	cmd := &cobra.Command{
		Use:   "set",
		Short: "Set file command",
		Long:  `Long set description`,
		Run: func(cmd *cobra.Command, args []string) {
			options.validateFlags()

			fileType := options.getFileType()

			var cfg *config.Config

			switch fileType {
			case "json":
				cfg, _ = config.ParseJsonFile(options.File)
			case "yaml":
				cfg, _ = config.ParseYamlFile(options.File)
			case "yml":
				cfg, _ = config.ParseYamlFile(options.File)
			}

			value, _ := cfg.String(options.Key)
			fmt.Println("Current value: " + value)

			err := cfg.Set(options.Key, options.Value)
			if err != nil {
			}
			mapping, err := cfg.Map("")

			var outputString []byte
			switch fileType {
			case "json":
				outputString, _ = json.MarshalIndent(mapping, "", "  ")
			case "yaml":
				outputString, _ = yaml.Marshal(mapping)
			case "yml":
				outputString, _ = yaml.Marshal(mapping)
			}

			fmt.Println("New value:")
			fmt.Println(string(outputString))
			if options.DryRun {
				fmt.Println("-----DRY RUN-----")
				return
			}

			file, err := os.Create(options.File)
			defer file.Close()

			file.Write(outputString)
		},
	}

	cmd.MarkFlagRequired("file")
	cmd.MarkFlagRequired("key")
	cmd.MarkFlagRequired("value")

	cmd.Flags().StringVarP(&options.File, "file", "f", "", "file to target")
	cmd.Flags().StringVarP(&options.Key, "key", "k", "", "key in the file to target")
	cmd.Flags().StringVarP(&options.Value, "value", "v", "", "value to set in the targeted key")
	cmd.Flags().BoolVarP(&options.DryRun, "dry-run", "d", false, "show changes without saving")

	return cmd
}

func (options *SetOptions) validateFlags() {
	if len(options.File) == 0 {
		fmt.Println("please specify a filename")
		os.Exit(1)
	}
	extension := options.getFileType()
	supportedExtensions := []string{"json", "yaml"}
	extensionSupported := false

	for _, supportedExtension := range supportedExtensions {
		if extension == supportedExtension {
			extensionSupported = true
		}
	}

	if !extensionSupported {
		fmt.Println("File extension not supported. Please specify one of: " + strings.Join(supportedExtensions, ", "))
		os.Exit(1)
	}
}

func (options *SetOptions) getFileType() string {
	_, err := os.Stat(options.File)
	if os.IsNotExist(err) {
		fmt.Println("please specify a filename")
		os.Exit(1)
	}
	extension := filepath.Ext(options.File)
	return extension[1:]
}
