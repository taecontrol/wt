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

func TestListWorktrees(t *testing.T) {
	t.Run("It returns a list of worktrees", func(t *testing.T) {
		ExecHandlerMock := func(name string, arg ...string) *exec.Cmd {
			return exec.Command("echo", "worktree /path/to/bare-source\nbare\nworktree /path/to/linked-worktree\nHEAD abcd1234abcd1234abcd1234abcd1234abcd1234\nbranch refs/heads/master\nworktree /path/to/other-linked-worktree\nHEAD 1234abc1234abc1234abc1234abc1234abc1234a\ndetached\nworktree /path/to/linked-worktree-locked-no-reason\nHEAD 5678abc5678abc5678abc5678abc5678abc5678c\nbranch refs/heads/locked-no-reason\nlocked\nworktree /path/to/linked-worktree-locked-with-reason\nHEAD 3456def3456def3456def3456def3456def3456b\nbranch refs/heads/locked-with-reason\nlocked reason why is locked\nworktree /path/to/linked-worktree-prunable\nHEAD 1233def1234def1234def1234def1234def1234b\ndetached\nprunable gitdir file points to non-existent location")
		}

		worktrees, err := ListWorktrees(ExecHandlerMock)

		if err != nil {
			t.Errorf("ListWorktrees should return a list of worktrees: %v", err)
		}

		if worktrees.Count() != 6 {
			t.Errorf("ListWorktrees should return 2 items, got %d", worktrees.Count())
		}

		if worktrees.Get(0).Path != "/path/to/bare-source" {
			t.Errorf("Second item should be /path/to/bare-source, got %s", worktrees.Get(0).Path)
		}
		if worktrees.Get(0).Bare != true {
			t.Errorf("Second item should be true, got %t", worktrees.Get(0).Bare)
		}

		if worktrees.Get(1).Path != "/path/to/linked-worktree" {
			t.Errorf("Second item should be /path/to/linked-worktree, got %s", worktrees.Get(1).Path)
		}
		if worktrees.Get(1).Head != "abcd1234abcd1234abcd1234abcd1234abcd1234" {
			t.Errorf("Second item should be /path/to/linked-worktree, got %s", worktrees.Get(1).Head)
		}
		if worktrees.Get(1).Bare != false {
			t.Errorf("Second item should be false, got %t", worktrees.Get(1).Bare)
		}
		if worktrees.Get(1).Branch != "refs/heads/master" {
			t.Errorf("Second item should be refs/heads/master, got %s", worktrees.Get(1).Branch)
		}
	})
}

func TestGetMainWorktree(t *testing.T) {
	t.Run("It returns the main worktree", func(t *testing.T) {
		ExecHandlerMock := func(name string, arg ...string) *exec.Cmd {
			return exec.Command("echo", "worktree /path/to/linked-worktree\nHEAD abcd1234abcd1234abcd1234abcd1234abcd1234\nbranch refs/heads/master\nworktree /path/to/other-linked-worktree\nHEAD 1234abc1234abc1234abc1234abc1234abc1234a\ndetached\nworktree /path/to/linked-worktree-locked-no-reason\nHEAD 5678abc5678abc5678abc5678abc5678abc5678c\nbranch refs/heads/locked-no-reason\nlocked\nworktree /path/to/linked-worktree-locked-with-reason\nHEAD 3456def3456def3456def3456def3456def3456b\nbranch refs/heads/locked-with-reason\nlocked reason why is locked\nworktree /path/to/linked-worktree-prunable\nHEAD 1233def1234def1234def1234def1234def1234b\ndetached\nprunable gitdir file points to non-existent location")
		}

		worktree, err := GetMainWorktree(ExecHandlerMock)

		if err != nil {
			t.Errorf("GetMainWorktree should return a worktrees: %v", err)
		}

		if worktree.Path != "/path/to/linked-worktree" {
			t.Errorf("Second item should be /path/to/linked-worktree, got %s", worktree.Path)
		}
		if worktree.Head != "abcd1234abcd1234abcd1234abcd1234abcd1234" {
			t.Errorf("Second item should be /path/to/linked-worktree, got %s", worktree.Head)
		}
		if worktree.Bare != false {
			t.Errorf("Second item should be false, got %t", worktree.Bare)
		}
		if worktree.Branch != "refs/heads/master" {
			t.Errorf("Second item should be refs/heads/master, got %s", worktree.Branch)
		}
	})
}
