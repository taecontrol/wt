package cmd_test

import (
	"context"
	"testing"
	"wt/cmd"
	"wt/pkg/core"
	"wt/pkg/core/git"
	"wt/pkg/core/utils"
)

func TestAdd(t *testing.T) {
	t.Run("add", func(t *testing.T) {
		addCmd := cmd.NewAddCmd()
		app := &core.App{
			Exec:   utils.NewCmdExecutorMock(),
			Git:    git.NewGitMock(),
			Config: core.NewConfigMock(),
		}

		app.Git.(*git.GitMock).On("GetMainWorktree", app.Exec).Return(utils.Worktree{Branch: "refs/heads/main", Path: "/home/user/main"}, nil)
		app.Git.(*git.GitMock).On("AddWorktree", "/home/user/worktrees/test_worktree", "FEAT-1", app.Exec, false).Return(nil)

		app.Config.(*core.ConfigMock).On("LoadConfig").Return(nil)
		app.Config.(*core.ConfigMock).On("GetInitCommands").Return([]string{"echo '1st command'", "echo '2nd command'"})

		app.Exec.(*utils.CmdExecutorMock).On("StdOutPipe", "echo '1st command'", "/home/user/worktrees/test_worktree").Return(nil)
		app.Exec.(*utils.CmdExecutorMock).On("StdOutPipe", "echo '2nd command'", "/home/user/worktrees/test_worktree").Return(nil)

		ctx := context.WithValue(context.Background(), core.AppKey{}, app)
		addCmd.SetContext(ctx)

		addCmd.SetArgs([]string{"test_worktree", "FEAT-1"})
		addCmd.Execute()

		app.Git.(*git.GitMock).AssertExpectations(t)
		app.Config.(*core.ConfigMock).AssertExpectations(t)
		app.Exec.(*utils.CmdExecutorMock).AssertExpectations(t)
	})

	t.Run("add with -b option", func(t *testing.T) {
		addCmd := cmd.NewAddCmd()
		app := &core.App{
			Exec:   utils.NewCmdExecutorMock(),
			Git:    git.NewGitMock(),
			Config: core.NewConfigMock(),
		}

		app.Git.(*git.GitMock).On("GetMainWorktree", app.Exec).Return(utils.Worktree{Branch: "refs/heads/main", Path: "/home/user/main"}, nil)
		app.Git.(*git.GitMock).On("AddWorktree", "/home/user/worktrees/test_worktree", "FEAT-1", app.Exec, true).Return(nil)

		app.Config.(*core.ConfigMock).On("LoadConfig").Return(nil)
		app.Config.(*core.ConfigMock).On("GetInitCommands").Return([]string{"echo '1st command'", "echo '2nd command'"})

		app.Exec.(*utils.CmdExecutorMock).On("StdOutPipe", "echo '1st command'", "/home/user/worktrees/test_worktree").Return(nil)
		app.Exec.(*utils.CmdExecutorMock).On("StdOutPipe", "echo '2nd command'", "/home/user/worktrees/test_worktree").Return(nil)

		ctx := context.WithValue(context.Background(), core.AppKey{}, app)
		addCmd.SetContext(ctx)

		addCmd.SetArgs([]string{"-b", "test_worktree", "FEAT-1"})
		addCmd.Execute()

		app.Git.(*git.GitMock).AssertExpectations(t)
		app.Config.(*core.ConfigMock).AssertExpectations(t)
		app.Exec.(*utils.CmdExecutorMock).AssertExpectations(t)
	})

	t.Run("add with -p option", func(t *testing.T) {
		addCmd := cmd.NewAddCmd()
		app := &core.App{
			Exec:   utils.NewCmdExecutorMock(),
			Git:    git.NewGitMock(),
			Config: core.NewConfigMock(),
		}

		app.Git.(*git.GitMock).On("GetMainWorktree", app.Exec).Return(utils.Worktree{Branch: "refs/heads/main", Path: "/home/user/main"}, nil)
		app.Git.(*git.GitMock).On("AddWorktree", "/home/user/different/path/test_worktree", "FEAT-1", app.Exec, false).Return(nil)

		app.Config.(*core.ConfigMock).On("LoadConfig").Return(nil)
		app.Config.(*core.ConfigMock).On("GetInitCommands").Return([]string{"echo '1st command'", "echo '2nd command'"})

		app.Exec.(*utils.CmdExecutorMock).On("StdOutPipe", "echo '1st command'", "/home/user/different/path/test_worktree").Return(nil)
		app.Exec.(*utils.CmdExecutorMock).On("StdOutPipe", "echo '2nd command'", "/home/user/different/path/test_worktree").Return(nil)

		ctx := context.WithValue(context.Background(), core.AppKey{}, app)
		addCmd.SetContext(ctx)

		addCmd.SetArgs([]string{"-p", "../different/path", "test_worktree", "FEAT-1"})
		addCmd.Execute()

		app.Git.(*git.GitMock).AssertExpectations(t)
		app.Config.(*core.ConfigMock).AssertExpectations(t)
		app.Exec.(*utils.CmdExecutorMock).AssertExpectations(t)
	})
}
