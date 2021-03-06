package main

import (
	"fmt"
	"github.com/femiagbabiaka/gopr"
	"os"
	"os/exec"
)

func main() {
	gogurtPullRequest()
}

func help() string {
	helpOutput := `Help docs:
	Gogurt is a git wrapper written in Golang.
	Commands:
		help: return a list of commands.

		branch: Takes a single argument of a branch name, will create a branch if one doesn't exist.

		checkout: Will take a single argument, and either switch to the branch whose name matches the input string most closely, 
		or return a nice list of branch names to then switch to.

		pull request: Takes two arguments, the first being the branch you want to open the Github.com pull request for, 
		and the second being the upstream branch you want to target. Opens a browser window so you can more easily finish your pull requests.`
	return helpOutput
}

func checkout() string {
	return "checkout"
}

func branch() string {
	return "branch"
}

func currentBranch() string {
	return "currentBranch"
}

func captureBranches() string {
	var cmdOut []byte
	var err error
	cmdName := "git"
	cmdArgs := []string{"branch"}
	if cmdOut, err = exec.Command(cmdName, cmdArgs...).Output(); err != nil {
		fmt.Fprintln(os.Stderr, "There was en error capturing branches: ", err)
		os.Exit(1)
	}

	return string(cmdOut)
}
