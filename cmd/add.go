package cmd

import (
	"os"
	"path/filepath"
	"wt/pkg/core"
	"wt/pkg/core/utils"

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
			name := getNameArg(args)
			branch := getBranchArg(args)
			newBranchFlag := getNewBranchFlag(cmd)
			config := loadConfig(app)

			addWorktree(app, name, branch, newBranchFlag)
			runInitCommands(app, config)
		},
	}

	addCmd.PersistentFlags().StringP("path", "p", "", "Path to the worktree")
	addCmd.PersistentFlags().BoolP("newBranch", "b", false, "Create a new branch")

	return addCmd
}

func getNameArg(args []string) string {
	name := args[0]
	if name == "" {
		panic("worktree name is required")
	}
	return name
}

func getBranchArg(args []string) string {
	branch := args[1]
	if branch == "" {
		panic("branch name is required")
	}
	return branch
}

func getNewBranchFlag(cmd *cobra.Command) bool {
	newBranchFlag, err := cmd.Flags().GetBool("newBranch")
	if err != nil {
		panic(err)
	}

	return newBranchFlag
}

func getMainWorktree(app *core.App) *utils.Worktree {
	mainWorktree, err := app.Git.GetMainWorktree(app.Exec)
	if err != nil {
		panic(err)
	}

	return &mainWorktree
}

func loadConfig(app *core.App) core.ConfigContract {
	config := app.Config

	if err := config.LoadConfig(); err != nil {
		utils.LogError("[Error] %s", err.Error())
		os.Exit(1)
	}

	return config
}

func addWorktree(app *core.App, name string, branch string, newBranchFlag bool) {
	mainWorktree := getMainWorktree(app)
	path := filepath.Clean(mainWorktree.Path + "/../worktrees/" + name)

	err := app.Git.AddWorktree(path, branch, app.Exec, newBranchFlag)
	if err != nil {
		utils.LogError("[Error] %s", err.Error())
		os.Exit(1)
	}
}

func runInitCommands(app *core.App, config core.ConfigContract) {
	for _, initCmd := range config.GetInitCommands() {
		err := app.Exec.StdOutPipe(initCmd)
		if err != nil {
			utils.LogError("[Error] %s", err.Error())
			os.Exit(1)
		}
	}
}
