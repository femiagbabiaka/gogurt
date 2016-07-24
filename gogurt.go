package main

import "fmt"
import "os"
import "os/exec"

func main() {
	command := os.Args[1]
	branchName := os.Args[2]
	var cmdOut []byte
	var err error
	fmt.Println(command)

	m := map[string]func(string) string{
		"branch":   branch,
		"checkout": checkout,
	}

	response := m[command](branchName)
	fmt.Println(z)

}

func checkout(branchName string) string {
	return "checkout func " + branchName
}

func branch(branchName string) string {
	return "branch func " + branchName
}
