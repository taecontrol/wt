package wt

import (
	"os/exec"
	"strings"
	"testing"
)

func TestList(t *testing.T) {
	t.Run("It returns a list of branches", func(t *testing.T) {
		ExecHandlerMock := func(name string, arg ...string) *exec.Cmd {
			return exec.Command("echo", "  master\n* develop\n  feature/branch\n+ test")
		}

		branches, err := List(ExecHandlerMock)

		if err != nil {
			t.Errorf("List should return a list of branches: %v", err)
		}

		if len(branches) != 4 {
			t.Error("List should return 4 items")
		}

		if branches[1] != "develop" {
			t.Errorf("Second item should be develop, got %s", branches[1])
		}

		if branches[3] != "test" {
			t.Errorf("Fourth item should be test, got %s", branches[3])
		}
	})
}

func TestAddWorktree(t *testing.T) {
	t.Run("It adds a worktree", func(t *testing.T) {
		var args string

		ExecHandlerMock := func(name string, arg ...string) *exec.Cmd {
			args = name + " " + strings.Join(arg, " ")
			return exec.Command("echo", "  master\n* develop\n  feature/branch\n+ test")
		}

		err := AddWorktree("test", ExecHandlerMock, false)

		if args != "git worktree add test" {
			t.Errorf("Args should be git worktree add test, got %s", args)
		}

		if err != nil {
			t.Errorf("AddWorktree should add a worktree: %v", err)
		}
	})

	t.Run("It adds a worktree with the -b option", func(t *testing.T) {
		var args string

		ExecHandlerMock := func(name string, arg ...string) *exec.Cmd {
			args = name + " " + strings.Join(arg, " ")
			return exec.Command("echo", "  master\n* develop\n  feature/branch\n+ test")
		}

		err := AddWorktree("test", ExecHandlerMock, true)

		if args != "git worktree add -b test" {
			t.Errorf("Args should be git worktree add -b test, got %s", args)
		}

		if err != nil {
			t.Errorf("AddWorktree should add a worktree: %v", err)
		}
	})
}
