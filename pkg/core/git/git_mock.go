package git

import "wt/pkg/core/utils"

type GitMock struct {
	count map[string]int
}

func NewGitMock() *GitMock {
	return &GitMock{
		count: map[string]int{},
	}
}

func (gitMock *GitMock) WasCalledTimes(method string, count int) bool {
	return gitMock.count[method] == count
}

func (gitMock *GitMock) GetCount(method string) int {
	return gitMock.count[method]
}

func (gitMock *GitMock) List(Exec utils.CommandExecHandler) (utils.Collection[string], error) {
	gitMock.count["List"]++
	return utils.NewCollection[string]([]string{"master", "develop"}), nil
}

func (gitMock *GitMock) AddWorktree(path string, branch string, Exec utils.CommandExecHandler, newBranchOption bool) error {
	gitMock.count["AddWorktree"]++
	return nil
}

func (gitMock *GitMock) RemoveWorktree(path string, Exec utils.CommandExecHandler, forceOption bool) error {
	gitMock.count["RemoveWorktree"]++
	return nil
}

func (gitMock *GitMock) ListWorktrees(Exec utils.CommandExecHandler) (utils.Collection[utils.Worktree], error) {
	gitMock.count["ListWorktrees"]++
	return utils.NewCollection[utils.Worktree]([]utils.Worktree{
		{Branch: "refs/heads/main", Path: "/home/user/main"},
		{Branch: "refs/heads/dev", Path: "/home/user/dev"},
	}), nil
}

func (gitMock *GitMock) GetMainWorktree(Exec utils.CommandExecHandler) (utils.Worktree, error) {
	gitMock.count["GetMainWorktree"]++
	return utils.Worktree{Branch: "refs/heads/main", Path: "/home/user/main"}, nil
}
