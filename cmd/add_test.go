package cmd

import (
	"context"
	"testing"
	"wt/pkg/core"
	"wt/pkg/core/git"
	"wt/pkg/core/utils"
)

func TestAdd(t *testing.T) {
	t.Run("add", func(t *testing.T) {
		addCmd := NewAddCmd()
		app := &core.App{
			Exec: utils.NewCmdExecutorMock(),

			Git: git.NewGitMock(),
		}

		app.Git.(*git.GitMock).On("GetMainWorktree", app.Exec).Return(utils.Worktree{Branch: "refs/heads/main", Path: "/home/user/main"}, nil)
		app.Git.(*git.GitMock).On("AddWorktree", "/home/user/main/test_worktree", "FEAT-1", app.Exec, false).Return(nil)

		ctx := context.WithValue(context.Background(), core.AppKey{}, app)
		addCmd.SetContext(ctx)

		addCmd.SetArgs([]string{"test_worktree", "FEAT-1"})
		addCmd.Execute()

		app.Git.(*git.GitMock).AssertExpectations(t)
	})
}
