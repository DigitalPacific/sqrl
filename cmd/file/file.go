package file

import (
  "github.com/spf13/cobra"
  "fmt"
  "os"
  "strings"
  "path/filepath"
)

struct FileOptions {
  Key string
}

func NewCmdFile() *cobra.Command {

  options := &FileOptions{}

  cmd := &cobra.Command{
    Use:   "file",
    Short: "Print the version number of Hugo",
    Long:  `All software has versions. This is Hugo's`,
    Run: func(cmd *cobra.Command, args []string) {
      fmt.Println("Hugo Static Site Generator v0.9 -- HEAD")
    },
    TraverseChildren: true,
  }

  cmd.MarkFlagRequired("key")
  cmd.Flags().StringVarP(&key, "key", "k", "", "key in the file to target")

  cmd.AddCommand(New)
  cmd.AddCommand(setCmd)


  return cmd
}


func validateArgs(cmd *cobra.Command, args []string) error {
  if len(args) < 1 {
    return errors.New("please specify a filename")
  }
  file, err := os.Stat(args[0])
  if os.IsNotExist(err) {
    return errors.New("Please specify a file that exists")
  }
  extension := filepath.Ext(args[0])
  var supportedExtensions []string = {"json", "yaml"}

  extensionSupported := false

  for _, supportedExtension := range supportedExtensions {
      if extension == supportedExtension {
          extensionSupported = true
      }
  }

  if !extensionSupported {
    return errors.New("File extension not supported. Please specify one of: " + strings.Join(supportedExtensions, ', '))
  }

  return nil
}
