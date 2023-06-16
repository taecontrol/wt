package cmd

import (
	"context"
	"os/exec"
	"testing"
	"wt/pkg/core"
	"wt/pkg/core/git"
)

func TestAdd(t *testing.T) {
	t.Run("add", func(t *testing.T) {
		addCmd := NewAddCmd()
		app := &core.App{
			Exec: func(name string, args ...string) *exec.Cmd {
				return exec.Command("echo", "command executed")
			},

			Git: git.NewGitMock(),
		}

		ctx := context.WithValue(context.Background(), core.AppKey{}, app)
		addCmd.SetContext(ctx)

		addCmd.SetArgs([]string{"test_worktree", "FEAT-1"})
		addCmd.Execute()

		if !app.Git.(*git.GitMock).WasCalledTimes("GetMainWorktree", 1) {
			t.Errorf("Expected GetMainWorktree to be called 1 time, but it was called %d times", app.Git.(*git.GitMock).GetCount("GetMainWorktree"))
		}

		if !app.Git.(*git.GitMock).WasCalledTimes("AddWorktree", 1) {
			t.Errorf("Expected AddWorktree to be called 1 time, but it was called %d times", app.Git.(*git.GitMock).GetCount("AddWorktree"))
		}
	})
}
