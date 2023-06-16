package cmd

import (
	"context"
	"wt/pkg/core/utils"

	"github.com/spf13/cobra"
)

func NewRootCmd(app *utils.App) *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "wt",
		Short: "wt is a CLI tool for working with git worktrees",

		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			ctx := context.WithValue(cmd.Context(), utils.AppKey{}, app)
			cmd.SetContext(ctx)
		},

		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	rootCmd.AddCommand(NewAddCmd())

	return rootCmd
}
