package utils

import (
	"os/exec"
)

type AppKey struct{}

type App struct {
	Exec            CommandExecHandler
	AddWorktree     func(path string, branch string, Exec CommandExecHandler, newBranchOption bool) error
	GetMainWorktree func(Exec CommandExecHandler) (Worktree, error)
}

func NewApp() *App {
	return &App{
		Exec: exec.Command,
	}
}
