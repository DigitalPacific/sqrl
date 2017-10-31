package file

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/DigitalPacific/squirrel/pkg"
	"github.com/olebedev/config"
	"github.com/spf13/cobra"
)

type FileValues struct {
	File   string
	Key    string
	Value  string
	DryRun bool
}

func NewCmdFile() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "file",
		Short: "Get and set operations for json, yaml & txt files",
		Long:  `Get and set operations for json, yaml & txt files`,
	}

	cmd.AddCommand(NewCmdSet())
	cmd.AddCommand(NewCmdGet())
	return cmd
}

func (fileValues *FileValues) ValidateArgs(args []string) {
	if len(args) < 1 {
		fmt.Println("Please specify a filename")
		os.Exit(1)
	}

	fileValues.File = args[0]

	extension := fileValues.getFileType()
	supportedExtensions := []string{"json", "yaml", ""}
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

func (fileValues *FileValues) ValidateFlags() {
	if len(fileValues.File) == 0 {
		fmt.Println("Please specify a filename")
		os.Exit(1)
	}
}

func (fileValues *FileValues) getFileType() string {
	_, err := os.Stat(fileValues.File)
	if os.IsNotExist(err) {
		fmt.Println("Please specify a filename")
		os.Exit(1)
	}
	extension := filepath.Ext(fileValues.File)

	if len(extension) == 0 {
		return extension
	}

	return extension[1:]
}

func (fileValues *FileValues) get() (*config.Config, error) {
	fileType := fileValues.getFileType()

	var cfg *config.Config
	var err error

	if fileValues.Key == "" {
		cfg, err = pkg.ParseTextFile(fileValues.File)
	}

	switch fileType {
	case "json":
		cfg, err = config.ParseJsonFile(fileValues.File)
	case "yaml":
		cfg, err = config.ParseYamlFile(fileValues.File)
	case "yml":
		cfg, err = config.ParseYamlFile(fileValues.File)
	}

	return cfg, err
}

func (fileValues *FileValues) render(cfg *config.Config) (string, error) {

	fileType := fileValues.getFileType()
	var err error
	var outputString string

	switch fileType {
	case "json":
		outputString, err = pkg.RenderJsonIndent(cfg.Root)
	case "yaml":
		outputString, err = config.RenderYaml(cfg.Root)
	case "yml":
		outputString, err = config.RenderYaml(cfg.Root)
	case "":
		outputString, err = pkg.RenderText(cfg.Root)
	}

	return outputString, err
}

func (fileValues *FileValues) set(cfg *config.Config) error {
	fileType := fileValues.getFileType()

	if fileType == "" {
		cfg.Root = fileValues.Value
		return nil
	}

	return cfg.Set(fileValues.Key, fileValues.Value)
}
