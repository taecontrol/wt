package utils

import (
	"os/exec"

	"github.com/stretchr/testify/mock"
)

type CmdExecutorMock struct {
	mock.Mock
}

func NewCmdExecutorMock() *CmdExecutorMock {
	return &CmdExecutorMock{}
}

func (cmdExecutorMock *CmdExecutorMock) Exec(name string, arg ...string) *exec.Cmd {
	args := cmdExecutorMock.Called(name, arg)
	return args.Get(0).(*exec.Cmd)
}

func (cmdExecutorMock *CmdExecutorMock) StdOutPipe(cmdString string, path string) error {
	args := cmdExecutorMock.Called(cmdString, path)

	return args.Error(0)
}
