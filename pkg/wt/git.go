package wt

import (
	"os/exec"
	"strings"
)

type (
	CommandExecHandler func(name string, arg ...string) *exec.Cmd
)

func List(Exec CommandExecHandler) ([]string, error) {
	command := Exec("git", "branch", "--list")

	out, err := command.Output()
	if err != nil {
		return []string{}, err
	}

	branches := strings.Split(string(out), "\n")
	for i := range branches {
		branches[i] = strings.ReplaceAll(branches[i], "*", "")
		branches[i] = strings.ReplaceAll(branches[i], "+", "")
		branches[i] = strings.TrimSpace(branches[i])
	}

	var filteredBranches []string

	for _, branch := range branches {
		if branch != "" {
			filteredBranches = append(filteredBranches, branch)
		}
	}

	return filteredBranches, nil
}

func AddWorktree(branch string, Exec CommandExecHandler, newBranchOption bool) error {
	args := []string{"worktree", "add"}

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
