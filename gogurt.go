package main

import "fmt"
import "os"
import "os/exec"

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Must supply a branch name.")
		os.Exit(1)
	}

	command := os.Args[1]
	branchName := os.Args[2]

	m := map[string]func(string) string{
		"branch":   branch,
		"checkout": checkout,
	}

	realCommand, commandExists := m[command]
	if commandExists != true {
		fmt.Println("The command " + command + " doesn't exist.")
		os.Exit(1)
	}

	response := realCommand(branchName)

	fmt.Println(response)

}

func checkout(branchName string) string {
	return captureBranches()
}

func branch(branchName string) string {
	return captureBranches()
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
