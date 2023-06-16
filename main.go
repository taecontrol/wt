package main

import (
	"fmt"
	"os"
	"os/exec"
	"wt/cmd"
	"wt/pkg/core/utils"
)

func main() {
	app := &utils.App{
		Exec: exec.Command,
	}

	rootCmd := cmd.NewRootCmd(app)

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "There was an error while executing the command: %s\n", err)
		os.Exit(1)
	}
}
