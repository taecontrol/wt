package core

import (
	"os/exec"
	"strings"
	"wt/pkg/core/utils"
)

type (
	CommandExecHandler func(name string, arg ...string) *exec.Cmd
)

func List(Exec CommandExecHandler) (utils.Collection[string], error) {
	command := Exec("git", "branch", "--list")

	out, err := command.Output()
	if err != nil {
		return utils.Collection[string]{}, err
	}

	branches := utils.NewCollection[string](strings.Split(string(out), "\n"))

	branches = branches.Map(func(branch string, _ int) string {
		branch = strings.ReplaceAll(branch, "*", "")
		branch = strings.ReplaceAll(branch, "+", "")
		branch = strings.TrimSpace(branch)

		return branch
	})

	branches = branches.Filter(func(branch string, _ int) bool {
		return branch != ""
	})

	return branches, nil
}

func AddWorktree(path string, branch string, Exec CommandExecHandler, newBranchOption bool) error {
	args := []string{"worktree", "add", path}

	if newBranchOption {
		args = append(args, "-b")
	}

	command := Exec("git", append(args, branch)...)

	err := command.Run()
	if err != nil {
		return err
	}

	return nil
}

func RemoveWorktree(path string, Exec CommandExecHandler, forceOption bool) error {
	args := []string{"worktree", "remove", path}

	if forceOption {
		args = append(args, "--force")
	}

	command := Exec("git", args...)

	err := command.Run()
	if err != nil {
		return err
	}

	return nil
}

func ListWorktrees(Exec CommandExecHandler) (utils.Collection[utils.Worktree], error) {
	command := Exec("git", "worktree", "list", "--porcelain")

	out, err := command.Output()
	if err != nil {
		return utils.Collection[utils.Worktree]{}, err
	}

	worktrees := utils.NewCollection[string](strings.Split(string(out), "\n"))
	worktrees = worktrees.Filter(func(worktree string, _ int) bool {
		return worktree != ""
	})

	groupedWorktrees := utils.Reduce[string, [][]string](worktrees, func(acc [][]string, worktree string, index int) [][]string {
		line := strings.Split(worktree, " ")

		if line[0] == "worktree" {
			startIndex := index
			var endIndex int

			for i := index + 1; i < worktrees.Count(); i++ {
				endIndex = i
				if strings.Split(worktrees.Get(i), " ")[0] == "worktree" {
					break
				}
			}

			group := worktrees.Slice(startIndex, endIndex)

			if group.Count() > 0 {
				acc = append(acc, group.Items)
			}

			return acc
		}

		return acc
	}, [][]string{})

	worktreesCollection := utils.ToCollection[utils.Worktree, []string](groupedWorktrees, func(group []string) utils.Worktree {
		return utils.NewWorktreeFromGroupArray(group)
	})

	return worktreesCollection, nil
}

func GetMainWorktree(Exec CommandExecHandler) (utils.Worktree, error) {
	worktrees, err := ListWorktrees(Exec)
	if err != nil {
		return utils.Worktree{}, err
	}

	return worktrees.Items[0], nil
}
