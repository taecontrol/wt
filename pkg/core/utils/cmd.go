package utils

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

type CmdExecutorContract interface {
	Exec(name string, arg ...string) *exec.Cmd
	StdOutPipe(cmdString string, path string) error
}

type CmdExecutor struct{}

func NewCmdExecutor() *CmdExecutor {
	return &CmdExecutor{}
}

func (cmdExecutor *CmdExecutor) Exec(name string, arg ...string) *exec.Cmd {
	LogInfo("[Running] %s %s \n", name, strings.Join(arg, " "))
	return exec.Command(name, arg...)
}

func (cmdExecutor *CmdExecutor) StdOutPipe(cmdString string, path string) error {
	var shell string
	if runtime.GOOS == "windows" {
		shell = "cmd"
	} else {
		shell = "/bin/sh"
	}

	cmd := cmdExecutor.Exec(shell, "-c", os.ExpandEnv(cmdString))
	cmd.Dir = path

	fmt.Println()

	stdoutReader, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}

	stderrReader, err := cmd.StderrPipe()
	if err != nil {
		return err
	}

	err = cmd.Start()
	if err != nil {
		return err
	}

	_, err = io.Copy(os.Stdout, stdoutReader)
	if err != nil {
		return err
	}

	_, err = io.Copy(os.Stderr, stderrReader)
	if err != nil {
		return err
	}

	if err = cmd.Wait(); err != nil {
		return fmt.Errorf("error while running command")
	}

	fmt.Println()
	return nil
}
