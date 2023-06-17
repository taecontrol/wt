package main

import (
	"fmt"
	"os"
	"wt/cmd"
	"wt/pkg/core"
)

func main() {
	app := core.NewApp()

	rootCmd := cmd.NewRootCmd(app)

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "There was an error while executing the command: %s\n", err)
		os.Exit(1)
	}
}
