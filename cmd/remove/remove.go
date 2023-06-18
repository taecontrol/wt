package remove

import (
	"os"
	"strings"
	"wt/pkg/core"
	"wt/pkg/core/utils"

	"github.com/spf13/cobra"
)

func NewRemoveCmd() *cobra.Command {
	removeCmd := &cobra.Command{
		Use:   "remove",
		Short: "Remove worktree and run terminate commands",
		Long: `Remove worktree and run terminate commands. For example:

wt rm <worktree_name> --force
	`,
		Run: func(cmd *cobra.Command, args []string) {
			app := cmd.Context().Value(core.AppKey{}).(*core.App)
			config := loadConfig(app)

			name := getNameArg(args)
			forceFlag := getForceFlag(cmd)

			worktrees, err := app.Git.ListWorktrees(app.Exec)
			if err != nil {
				utils.LogError("[Error] %s", err.Error())
				os.Exit(1)
			}
			worktree := worktrees.Find(func(item utils.Worktree, index int) bool {
				return strings.Contains(item.Branch, name)
			})
			if worktree == nil {
				utils.LogError("[Error] worktree %s not found", name)
				os.Exit(1)
			}

			runTerminateCommands(app, config, worktree.Path)

			err = app.Git.RemoveWorktree(worktree.Path, app.Exec, forceFlag)
			if err != nil {
				utils.LogError("[Error] %s", err.Error())
				os.Exit(1)
			}
		},
	}

	removeCmd.PersistentFlags().BoolP("force", "f", false, "force remove worktree")

	return removeCmd
}

func getNameArg(args []string) string {
	name := args[0]
	if name == "" {
		utils.LogError("[Error] worktree name is required")
		os.Exit(1)
	}
	return name
}

func getForceFlag(cmd *cobra.Command) bool {
	forceFlag, err := cmd.Flags().GetBool("force")
	if err != nil {
		utils.LogError("[Error] %s", err.Error())
		os.Exit(1)
	}

	return forceFlag
}

func loadConfig(app *core.App) core.ConfigContract {
	config := app.Config

	if err := config.LoadConfig(); err != nil {
		utils.LogError("[Error] %s", err.Error())
		os.Exit(1)
	}

	return config
}

func runTerminateCommands(app *core.App, config core.ConfigContract, path string) {
	for _, initCmd := range config.GetTerminateCommands() {
		err := app.Exec.StdOutPipe(initCmd, path)
		if err != nil {
			utils.LogError("[Error] %s", err.Error())
			os.Exit(1)
		}
	}
}
