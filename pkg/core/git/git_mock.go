package git

import (
	"wt/pkg/core/utils"

	"github.com/stretchr/testify/mock"
)

type GitMock struct {
	mock.Mock
}

func NewGitMock() *GitMock {
	return &GitMock{}
}

func (gitMock *GitMock) List(CmdExec utils.CmdExecutorContract) (utils.Collection[string], error) {
	args := gitMock.Called(CmdExec)
	return args.Get(0).(utils.Collection[string]), args.Error(1)
}

func (gitMock *GitMock) AddWorktree(path string, branch string, CmdExec utils.CmdExecutorContract, newBranchOption bool) error {
	args := gitMock.Called(path, branch, CmdExec, newBranchOption)
	return args.Error(0)
}

func (gitMock *GitMock) RemoveWorktree(path string, CmdExec utils.CmdExecutorContract, forceOption bool) error {
	args := gitMock.Called(path, CmdExec, forceOption)
	return args.Error(0)
}

func (gitMock *GitMock) ListWorktrees(CmdExec utils.CmdExecutorContract) (utils.Collection[utils.Worktree], error) {
	args := gitMock.Called(CmdExec)
	return args.Get(0).(utils.Collection[utils.Worktree]), args.Error(1)
}

func (gitMock *GitMock) GetMainWorktree(CmdExec utils.CmdExecutorContract) (utils.Worktree, error) {
	args := gitMock.Called(CmdExec)
	return args.Get(0).(utils.Worktree), args.Error(1)
}
