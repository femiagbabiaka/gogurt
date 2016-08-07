package main

import "fmt"
import "os"
import "os/exec"

func main() {
	if len(os.Args) < 2 {
		fmt.Println("\x1b[37;1m" + help() + "\x1b[0m")
		os.Exit(0)
	}
	command := os.Args[1]

	m := map[string]func() string{
		"branch":           branch,
		"checkout":         checkout,
		"pull-request":     pullRequest,
		"capture-branches": captureBranches,
	}

	realCommand, commandExists := m[command]
	if commandExists != true {
		fmt.Println("The command " + command + " doesn't exist.")
		fmt.Println("\x1b[37;1m" + help() + "\x1b[0m")
		os.Exit(1)
	}

	response := realCommand()

	fmt.Println(response)

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

func pullRequest() string {
	if len(os.Args) < 4 {
		fmt.Println("\x1b[37;1m" + "You need to pass both the branch that you want to pull and the one that you want to pull into." + "\x1b[0m")
		fmt.Println("\x1b[37;1m" + "Ex: `gogurt pull-request femiagbabiaka/master femiagbabiaka/bugfix-xyz`" + "\x1b[0m")
		os.Exit(1)
	}

	//branchToPull := os.Args[3]
	//branchToPullInto := os.Args[4]

	return "pullRequest"
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
