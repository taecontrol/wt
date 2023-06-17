package core

import (
	"wt/pkg/core/git"
	"wt/pkg/core/utils"
)

type AppKey struct{}

type App struct {
	Exec utils.CmdExecutorContract
	Git  git.GitContract
}

func NewApp() *App {
	return &App{
		Exec: utils.NewCmdExecutor(),
		Git:  git.NewGit(),
	}
}
