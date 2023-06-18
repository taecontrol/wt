package remove_test

import (
	"context"
	"testing"
	"wt/cmd/remove"
	"wt/pkg/core"
	"wt/pkg/core/git"
	"wt/pkg/core/utils"
)

func TestAdd(t *testing.T) {
	t.Run("remove", func(t *testing.T) {
		addCmd := remove.NewRemoveCmd()
		app := &core.App{
			Exec:   utils.NewCmdExecutorMock(),
			Git:    git.NewGitMock(),
			Config: core.NewConfigMock(),
		}

		app.Git.(*git.GitMock).On("ListWorktrees", app.Exec).Return(utils.NewCollection[utils.Worktree]([]utils.Worktree{{Branch: "refs/heads/test", Path: "/home/user/test_worktree"}}), nil)
		app.Git.(*git.GitMock).On("RemoveWorktree", "/home/user/test_worktree", app.Exec, false).Return(nil)

		app.Config.(*core.ConfigMock).On("LoadConfig").Return(nil)
		app.Config.(*core.ConfigMock).On("GetTerminateCommands").Return([]string{"echo '1st command'", "echo '2nd command'"})

		app.Exec.(*utils.CmdExecutorMock).On("StdOutPipe", "echo '1st command'", "/home/user/test_worktree").Return(nil)
		app.Exec.(*utils.CmdExecutorMock).On("StdOutPipe", "echo '2nd command'", "/home/user/test_worktree").Return(nil)

		ctx := context.WithValue(context.Background(), core.AppKey{}, app)
		addCmd.SetContext(ctx)

		addCmd.SetArgs([]string{"test_worktree"})
		addCmd.Execute()

		app.Git.(*git.GitMock).AssertExpectations(t)
		app.Config.(*core.ConfigMock).AssertExpectations(t)
		app.Exec.(*utils.CmdExecutorMock).AssertExpectations(t)
	})

	t.Run("remove with force flag", func(t *testing.T) {
		addCmd := remove.NewRemoveCmd()
		app := &core.App{
			Exec:   utils.NewCmdExecutorMock(),
			Git:    git.NewGitMock(),
			Config: core.NewConfigMock(),
		}

		app.Git.(*git.GitMock).On("ListWorktrees", app.Exec).Return(utils.NewCollection[utils.Worktree]([]utils.Worktree{{Branch: "refs/heads/test_worktree", Path: "/home/user/test_worktree"}}), nil)
		app.Git.(*git.GitMock).On("RemoveWorktree", "/home/user/test_worktree", app.Exec, true).Return(nil)

		app.Config.(*core.ConfigMock).On("LoadConfig").Return(nil)
		app.Config.(*core.ConfigMock).On("GetTerminateCommands").Return([]string{"echo '1st command'", "echo '2nd command'"})

		app.Exec.(*utils.CmdExecutorMock).On("StdOutPipe", "echo '1st command'", "/home/user/test_worktree").Return(nil)
		app.Exec.(*utils.CmdExecutorMock).On("StdOutPipe", "echo '2nd command'", "/home/user/test_worktree").Return(nil)

		ctx := context.WithValue(context.Background(), core.AppKey{}, app)
		addCmd.SetContext(ctx)

		addCmd.SetArgs([]string{"test_worktree", "--force"})
		addCmd.Execute()

		app.Git.(*git.GitMock).AssertExpectations(t)
		app.Config.(*core.ConfigMock).AssertExpectations(t)
		app.Exec.(*utils.CmdExecutorMock).AssertExpectations(t)
	})
}
