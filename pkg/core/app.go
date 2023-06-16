package core

import (
	"os/exec"
	"wt/pkg/core/git"
	"wt/pkg/core/utils"
)

type AppKey struct{}

type App struct {
	Exec utils.CommandExecHandler
	Git  git.GitContract
}

func NewApp() *App {
	return &App{
		Exec: exec.Command,
		Git:  git.NewGit(),
	}
}
