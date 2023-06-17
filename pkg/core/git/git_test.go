package git_test

import (
	"os/exec"
	"testing"
	"wt/pkg/core/git"
	"wt/pkg/core/utils"
)

func TestList(t *testing.T) {
	t.Run("It returns a list of branches", func(t *testing.T) {
		cmdExecMock := utils.NewCmdExecutorMock()
		cmdExecMock.On("Exec", "git", []string{"branch", "--list"}).Return(exec.Command("echo", "  master\n* develop\n  feature/branch\n+ test"))

		git := &git.Git{}
		branches, err := git.List(cmdExecMock)

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

		cmdExecMock.AssertExpectations(t)
	})
}

func TestAddWorktree(t *testing.T) {
	t.Run("It adds a worktree", func(t *testing.T) {
		cmdExecMock := utils.NewCmdExecutorMock()
		cmdExecMock.On("Exec", "git", []string{"worktree", "add", "/path/to/worktree", "test"}).Return(exec.Command("echo", "Command executed"))

		git := &git.Git{}
		err := git.AddWorktree("/path/to/worktree", "test", cmdExecMock, false)

		if err != nil {
			t.Errorf("AddWorktree should add a worktree: %v", err)
		}

		cmdExecMock.AssertExpectations(t)
	})

	t.Run("It adds a worktree with the -b option", func(t *testing.T) {
		cmdExecMock := utils.NewCmdExecutorMock()
		cmdExecMock.On("Exec", "git", []string{"worktree", "add", "/path/to/worktree", "-b", "test"}).Return(exec.Command("echo", "Command executed"))

		git := &git.Git{}
		err := git.AddWorktree("/path/to/worktree", "test", cmdExecMock, true)

		if err != nil {
			t.Errorf("AddWorktree should add a worktree: %v", err)
		}

		cmdExecMock.AssertExpectations(t)
	})
}

func TestRemoveWorktree(t *testing.T) {
	t.Run("It removes a worktree", func(t *testing.T) {
		cmdExecMock := utils.NewCmdExecutorMock()
		cmdExecMock.On("Exec", "git", []string{"worktree", "remove", "/path/to/worktree"}).Return(exec.Command("echo", "Command executed"))

		git := &git.Git{}
		err := git.RemoveWorktree("/path/to/worktree", cmdExecMock, false)

		if err != nil {
			t.Errorf("RemoveWorktree should remove a worktree: %v", err)
		}

		cmdExecMock.AssertExpectations(t)
	})

	t.Run("It removes a worktree with the --force option", func(t *testing.T) {
		cmdExecMock := utils.NewCmdExecutorMock()
		cmdExecMock.On("Exec", "git", []string{"worktree", "remove", "/path/to/worktree", "--force"}).Return(exec.Command("echo", "Command executed"))

		git := &git.Git{}
		err := git.RemoveWorktree("/path/to/worktree", cmdExecMock, true)

		if err != nil {
			t.Errorf("RemoveWorktree should remove a worktree: %v", err)
		}

		cmdExecMock.AssertExpectations(t)
	})
}

func TestListWorktrees(t *testing.T) {
	t.Run("It returns a list of worktrees", func(t *testing.T) {
		cmdExecMock := utils.NewCmdExecutorMock()
		cmdExecMock.On("Exec", "git", []string{"worktree", "list", "--porcelain"}).Return(exec.Command("echo", "worktree /path/to/bare-source\nbare\nworktree /path/to/linked-worktree\nHEAD abcd1234abcd1234abcd1234abcd1234abcd1234\nbranch refs/heads/master\nworktree /path/to/other-linked-worktree\nHEAD 1234abc1234abc1234abc1234abc1234abc1234a\ndetached\nworktree /path/to/linked-worktree-locked-no-reason\nHEAD 5678abc5678abc5678abc5678abc5678abc5678c\nbranch refs/heads/locked-no-reason\nlocked\nworktree /path/to/linked-worktree-locked-with-reason\nHEAD 3456def3456def3456def3456def3456def3456b\nbranch refs/heads/locked-with-reason\nlocked reason why is locked\nworktree /path/to/linked-worktree-prunable\nHEAD 1233def1234def1234def1234def1234def1234b\ndetached\nprunable gitdir file points to non-existent location"))

		git := &git.Git{}
		worktrees, err := git.ListWorktrees(cmdExecMock)

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

		cmdExecMock.AssertExpectations(t)
	})
}

func TestGetMainWorktree(t *testing.T) {
	t.Run("It returns the main worktree", func(t *testing.T) {
		cmdExecMock := utils.NewCmdExecutorMock()
		cmdExecMock.On("Exec", "git", []string{"worktree", "list", "--porcelain"}).Return(exec.Command("echo", "worktree /path/to/linked-worktree\nHEAD abcd1234abcd1234abcd1234abcd1234abcd1234\nbranch refs/heads/master\nworktree /path/to/other-linked-worktree\nHEAD 1234abc1234abc1234abc1234abc1234abc1234a\ndetached\nworktree /path/to/linked-worktree-locked-no-reason\nHEAD 5678abc5678abc5678abc5678abc5678abc5678c\nbranch refs/heads/locked-no-reason\nlocked\nworktree /path/to/linked-worktree-locked-with-reason\nHEAD 3456def3456def3456def3456def3456def3456b\nbranch refs/heads/locked-with-reason\nlocked reason why is locked\nworktree /path/to/linked-worktree-prunable\nHEAD 1233def1234def1234def1234def1234def1234b\ndetached\nprunable gitdir file points to non-existent location"))

		git := &git.Git{}
		worktree, err := git.GetMainWorktree(cmdExecMock)

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

		cmdExecMock.AssertExpectations(t)
	})
}
