package utils

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

type CmdExecutorContract interface {
	Exec(name string, arg ...string) *exec.Cmd
	StdOutPipe(cmdString string) error
}

type CmdExecutor struct{}

func NewCmdExecutor() *CmdExecutor {
	return &CmdExecutor{}
}

func (cmdExecutor *CmdExecutor) Exec(name string, arg ...string) *exec.Cmd {
	LogInfo("[Running] %s %s \n", name, strings.Join(arg, " "))
	return exec.Command(name, arg...)
}

func (cmdExecutor *CmdExecutor) StdOutPipe(cmdString string) error {
	arr := strings.Split(cmdString, " ")
	cmd := cmdExecutor.Exec(arr[0], arr[1:]...)
	fmt.Println()

	stdoutReader, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}

	stderrReader, err := cmd.StderrPipe()
	if err != nil {
		return err
	}

	cmd.Start()

	_, err = io.Copy(os.Stdout, stdoutReader)
	if err != nil {
		return err
	}

	w, err := io.Copy(os.Stderr, stderrReader)
	if err != nil {
		return err
	}
	if w > 0 {
		os.Exit(1)
	}

	fmt.Println()
	return nil
}
