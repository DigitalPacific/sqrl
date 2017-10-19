package main

import (
	"fmt"
	"github.com/DigitalPacific/ci/cmd"
	"os"
)

func main() {

	cmd.RootCmd.AddCommand(file.NewCmdFile())

	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
