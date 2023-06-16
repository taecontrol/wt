package utils

import "os/exec"

type (
	CommandExecHandler func(name string, arg ...string) *exec.Cmd
)
