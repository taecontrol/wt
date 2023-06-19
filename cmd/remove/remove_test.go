package remove_test

import (
	"context"
	"os"
	"testing"
	"wt/cmd/remove"
	"wt/pkg/core"
	"wt/pkg/core/git"
	"wt/pkg/core/utils"

	"github.com/stretchr/testify/assert"
)

func TestRemove(t *testing.T) {
	t.Run("remove", func(t *testing.T) {
		addCmd := remove.NewRemoveCmd()
		app := &core.App{
			Exec:   utils.NewCmdExecutorMock(),
			Git:    git.NewGitMock(),
			Config: core.NewConfigMock(),
		}

		worktrees := utils.NewCollection[utils.Worktree]([]utils.Worktree{
			{Branch: "refs/heads/main", Path: "/home/user/path_to_main_worktree"},
			{Branch: "refs/heads/test", Path: "/home/user/test_worktree"},
		})
		app.Git.(*git.GitMock).On("ListWorktrees", app.Exec).Return(worktrees, nil)
		app.Git.(*git.GitMock).On("GetMainWorktree", app.Exec).Return(worktrees.Items[0], nil)
		app.Git.(*git.GitMock).On("RemoveWorktree", "/home/user/test_worktree", app.Exec, false).Return(nil)

		app.Config.(*core.ConfigMock).On("LoadConfig", "/home/user/path_to_main_worktree").Return(nil)
		app.Config.(*core.ConfigMock).On("GetTerminateCommands").Return([]string{"echo '1st command'", "echo '2nd command'"})

		app.Exec.(*utils.CmdExecutorMock).On("StdOutPipe", "echo '1st command'", "/home/user/test_worktree").Return(nil)
		app.Exec.(*utils.CmdExecutorMock).On("StdOutPipe", "echo '2nd command'", "/home/user/test_worktree").Return(nil)

		ctx := context.WithValue(context.Background(), core.AppKey{}, app)
		addCmd.SetContext(ctx)

		addCmd.SetArgs([]string{"test_worktree"})
		addCmd.Execute()

		assert.Equal(t, "test", os.Getenv("BRANCH_NAME"))
		assert.Equal(t, "/home/user/path_to_main_worktree", os.Getenv("MAIN_WORKTREE_PATH"))
		assert.Equal(t, "/home/user/test_worktree", os.Getenv("WORKTREE_PATH"))
		assert.Equal(t, "test_worktree", os.Getenv("WORKTREE_NAME"))

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

		worktrees := utils.NewCollection[utils.Worktree]([]utils.Worktree{
			{Branch: "refs/heads/main", Path: "/home/user/path_to_main_worktree"},
			{Branch: "refs/heads/test", Path: "/home/user/test_worktree"},
		})
		app.Git.(*git.GitMock).On("ListWorktrees", app.Exec).Return(worktrees, nil)
		app.Git.(*git.GitMock).On("GetMainWorktree", app.Exec).Return(worktrees.Items[0], nil)
		app.Git.(*git.GitMock).On("RemoveWorktree", "/home/user/test_worktree", app.Exec, true).Return(nil)

		app.Config.(*core.ConfigMock).On("LoadConfig", "/home/user/path_to_main_worktree").Return(nil)
		app.Config.(*core.ConfigMock).On("GetTerminateCommands").Return([]string{"echo '1st command'", "echo '2nd command'"})

		app.Exec.(*utils.CmdExecutorMock).On("StdOutPipe", "echo '1st command'", "/home/user/test_worktree").Return(nil)
		app.Exec.(*utils.CmdExecutorMock).On("StdOutPipe", "echo '2nd command'", "/home/user/test_worktree").Return(nil)

		ctx := context.WithValue(context.Background(), core.AppKey{}, app)
		addCmd.SetContext(ctx)

		addCmd.SetArgs([]string{"test_worktree", "--force"})
		addCmd.Execute()

		assert.Equal(t, "test", os.Getenv("BRANCH_NAME"))
		assert.Equal(t, "/home/user/path_to_main_worktree", os.Getenv("MAIN_WORKTREE_PATH"))
		assert.Equal(t, "/home/user/test_worktree", os.Getenv("WORKTREE_PATH"))
		assert.Equal(t, "test_worktree", os.Getenv("WORKTREE_NAME"))

		app.Git.(*git.GitMock).AssertExpectations(t)
		app.Config.(*core.ConfigMock).AssertExpectations(t)
		app.Exec.(*utils.CmdExecutorMock).AssertExpectations(t)
	})
}
