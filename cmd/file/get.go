package file

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/olebedev/config"
	"github.com/spf13/cobra"
)

type FileOptions struct {
	Key  string
	File string
}

func NewCmdGet() *cobra.Command {

	var options FileOptions

	cmd := &cobra.Command{
		Use:   "get",
		Short: "Get File command",
		Run: func(cmd *cobra.Command, args []string) {
			options.validateFlags()

			var cfg *config.Config
			switch fileType := options.getFileType(); fileType {
			case "json":
				cfg, _ = config.ParseJsonFile(options.File)
			case "yaml":
				cfg, _ = config.ParseYamlFile(options.File)
			case "yml":
				cfg, _ = config.ParseYamlFile(options.File)
			}

			value, _ := cfg.String(options.Key)
			fmt.Println(value)
		},
	}

	cmd.MarkFlagRequired("key")
	cmd.MarkFlagRequired("file")

	cmd.Flags().StringVarP(&options.Key, "key", "k", "", "key in the file to target")
	cmd.Flags().StringVarP(&options.File, "file", "f", "", "file to target")

	return cmd
}

func (options *FileOptions) validateFlags() {
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

func (options *FileOptions) getFileType() string {
	_, err := os.Stat(options.File)
	if os.IsNotExist(err) {
		fmt.Println("please specify a filename")
		os.Exit(1)
	}
	extension := filepath.Ext(options.File)
	return extension[1:]
}
