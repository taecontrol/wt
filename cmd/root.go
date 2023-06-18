package cmd

import (
	"context"
	"wt/cmd/add"
	"wt/pkg/core"

	"github.com/spf13/cobra"
)

func NewRootCmd(app *core.App) *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "wt",
		Short: "wt is a CLI tool for working with git worktrees",

		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			ctx := context.WithValue(cmd.Context(), core.AppKey{}, app)
			cmd.SetContext(ctx)
		},

		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	rootCmd.AddCommand(add.NewAddCmd())

	return rootCmd
}
