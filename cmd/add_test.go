package cmd

import (
	"context"
	"os/exec"
	"strings"
	"testing"
	"wt/pkg/core/utils"
)

func TestAdd(t *testing.T) {
	t.Run("add", func(t *testing.T) {
		addWorktreeMock := utils.NewMock()

		addCmd := NewAddCmd()
		app := &utils.App{
			Exec: func(name string, args ...string) *exec.Cmd {
				fullArgs := strings.Join(append([]string{name}, args...), " ")
				if fullArgs == "git worktree add /path/to/worktree/test_worktree FEAT-1" {
					addWorktreeMock.Call()
				}
				return exec.Command("echo", "command executed")
			},
		}

		ctx := context.WithValue(context.Background(), utils.AppKey{}, app)
		addCmd.SetContext(ctx)

		addCmd.SetArgs([]string{"test_worktree", "FEAT-1"})
		addCmd.Execute()

		if !addWorktreeMock.WasCalledTimes(1) {
			t.Errorf("Expected addWorktreeMock to be called 1 time, but it was called %d times", addWorktreeMock.Count)
		}
	})
}
