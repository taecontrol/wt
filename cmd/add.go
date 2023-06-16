package cmd

import (
	"wt/pkg/core"

	"github.com/spf13/cobra"
)

func NewAddCmd() *cobra.Command {
	addCmd := &cobra.Command{
		Use:   "add",
		Short: "Add new worktree and run initialization command",
		Long: `Add new worktree and run initialization command. For example:

	This command will create a new wortree in ../worktrees/<worktree_name> and will checkout the <branch> branch.
	wt add <worktree_name> <branch>
	`,
		Run: func(cmd *cobra.Command, args []string) {
			app := cmd.Context().Value(core.AppKey{}).(*core.App)
			name := args[0]
			branch := args[1]
			mainWorktree, err := app.Git.GetMainWorktree(app.Exec)
			if err != nil {
				panic(err)
			}

			path := mainWorktree.Path + "/" + name

			app.Git.AddWorktree(path, branch, app.Exec, false)
		},
	}

	addCmd.PersistentFlags().StringP("path", "p", "", "Path to the worktree")
	addCmd.PersistentFlags().BoolP("branch", "b", true, "Create a new branch")

	return addCmd
}
