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
		Short: "Add new worktree and run initialization commands",
		Long: `Add new worktree and run initialization commands. For example:

wt add <worktree_name> <branch>

This command will create a new worktree in ../worktrees/<worktree_name> and will checkout the <branch> branch.
	`,
		Run: func(cmd *cobra.Command, args []string) {
			app := cmd.Context().Value(core.AppKey{}).(*core.App)
			mainWorktree := getMainWorktree(app)
			config := loadConfig(app)

			name := getNameArg(args)
			branch := getBranchArg(args)
			newBranchFlag := getNewBranchFlag(cmd)
			pathFlag := getPathFlag(cmd, mainWorktree)

			var path string
			if pathFlag != "" {
				path = filepath.Clean(pathFlag + "/" + name)
			} else {
				path = filepath.Clean(mainWorktree.Path + "/../worktrees/" + name)
			}

			os.Setenv("BRANCH_NAME", branch)
			os.Setenv("MAIN_WORKTREE_PATH", mainWorktree.Path)
			os.Setenv("WORKTREE_PATH", path)
			os.Setenv("WORKTREE_NAME", name)

			addWorktree(app, path, branch, newBranchFlag)
			runInitCommands(app, config, path)
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
		utils.LogError("[Error] branch name is required")
		os.Exit(1)
	}
	return branch
}

func getPathFlag(cmd *cobra.Command, mainWorktree *utils.Worktree) string {
	pathFlag, err := cmd.Flags().GetString("path")
	if err != nil {
		utils.LogError("[Error] %s", err.Error())
		os.Exit(1)
	}

	if pathFlag == "" {
		return pathFlag
	}

	if !filepath.IsAbs(pathFlag) {
		return filepath.Clean(mainWorktree.Path + "/" + pathFlag)
	}

	return pathFlag
}

func getNewBranchFlag(cmd *cobra.Command) bool {
	newBranchFlag, err := cmd.Flags().GetBool("newBranch")
	if err != nil {
		utils.LogError("[Error] %s", err.Error())
		os.Exit(1)
	}

	return newBranchFlag
}

func getMainWorktree(app *core.App) *utils.Worktree {
	mainWorktree, err := app.Git.GetMainWorktree(app.Exec)
	if err != nil {
		utils.LogError("[Error] %s", err.Error())
		os.Exit(1)
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

func addWorktree(app *core.App, path string, branch string, newBranchFlag bool) {
	err := app.Git.AddWorktree(path, branch, app.Exec, newBranchFlag)
	if err != nil {
		utils.LogError("[Error] %s", err.Error())
		os.Exit(1)
	}
}

func runInitCommands(app *core.App, config core.ConfigContract, path string) {
	for _, initCmd := range config.GetInitCommands() {
		err := app.Exec.StdOutPipe(initCmd, path)
		if err != nil {
			utils.LogError("[Error] %s", err.Error())
			os.Exit(1)
		}
	}
}
