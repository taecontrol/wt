package utils

import (
	"strings"
)

type Worktree struct {
	Path   string
	Head   string
	Branch string
	Bare   bool
}

func NewWorktreeFromGroupArray(group []string) Worktree {
	return Worktree{
		Path:   getPathFromGroupArray(group),
		Head:   getHeadFromGroupArray(group),
		Branch: getBranchFromGroupArray(group),
		Bare:   getBareFromGroupArray(group),
	}
}

func getPathFromGroupArray(group []string) string {
	selectedItem := getItemThatContains(group, "worktree")

	if len(selectedItem) == 0 {
		return ""
	}

	return strings.Split(selectedItem, " ")[1]
}

func getHeadFromGroupArray(group []string) string {
	selectedItem := getItemThatContains(group, "HEAD")

	if len(selectedItem) == 0 {
		return ""
	}

	return strings.Split(selectedItem, " ")[1]
}

func getBranchFromGroupArray(group []string) string {
	selectedItem := getItemThatContains(group, "branch")

	if len(selectedItem) == 0 {
		return ""
	}

	return strings.Split(selectedItem, " ")[1]
}

func getBareFromGroupArray(group []string) bool {
	selectedItem := getItemThatContains(group, "bare")
	return len(selectedItem) != 0
}

func getItemThatContains(group []string, substring string) string {
	selectedItem := ""

	for _, item := range group {
		if strings.Contains(item, substring) {
			selectedItem = item
			break
		}
	}

	return selectedItem
}
