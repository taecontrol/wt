package core

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

		if branches.Count() != 4 {
			t.Error("List should return 4 items")
		}

		if branches.Get(1) != "develop" {
			t.Errorf("Second item should be develop, got %s", branches.Get(1))
		}

		if branches.Get(3) != "test" {
			t.Errorf("Fourth item should be test, got %s", branches.Get(3))
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

		err := AddWorktree("/path/to/worktree", "test", ExecHandlerMock, false)

		if args != "git worktree add /path/to/worktree test" {
			t.Errorf("Args should be git worktree add /path/to/worktree test, got %s", args)
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

		err := AddWorktree("/path/to/worktree", "test", ExecHandlerMock, true)

		if args != "git worktree add /path/to/worktree -b test" {
			t.Errorf("Args should be git worktree add /path/to/worktree -b test, got %s", args)
		}

		if err != nil {
			t.Errorf("AddWorktree should add a worktree: %v", err)
		}
	})
}

func TestRemoveWorktree(t *testing.T) {
	t.Run("It removes a worktree", func(t *testing.T) {
		var args string

		ExecHandlerMock := func(name string, arg ...string) *exec.Cmd {
			args = name + " " + strings.Join(arg, " ")
			return exec.Command("echo", "  master\n* develop\n  feature/branch\n+ test")
		}

		err := RemoveWorktree("/path/to/worktree", ExecHandlerMock, false)

		if args != "git worktree remove /path/to/worktree" {
			t.Errorf("Args should be git worktree remove /path/to/worktree, got %s", args)
		}

		if err != nil {
			t.Errorf("RemoveWorktree should remove a worktree: %v", err)
		}
	})

	t.Run("It removes a worktree with the --force option", func(t *testing.T) {
		var args string

		ExecHandlerMock := func(name string, arg ...string) *exec.Cmd {
			args = name + " " + strings.Join(arg, " ")
			return exec.Command("echo", "  master\n* develop\n  feature/branch\n+ test")
		}

		err := RemoveWorktree("/path/to/worktree", ExecHandlerMock, true)

		if args != "git worktree remove /path/to/worktree --force" {
			t.Errorf("Args should be git worktree remove /path/to/worktree --force, got %s", args)
		}

		if err != nil {
			t.Errorf("RemoveWorktree should remove a worktree: %v", err)
		}
	})
}
