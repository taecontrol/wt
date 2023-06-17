package utils

import "os/exec"

type CmdExecutorContract interface {
	Exec(name string, arg ...string) *exec.Cmd
}

type CmdExecutor struct{}

func NewCmdExecutor() *CmdExecutor {
	return &CmdExecutor{}
}

func (cmdExecutor *CmdExecutor) Exec(name string, arg ...string) *exec.Cmd {
	return exec.Command(name, arg...)
}
