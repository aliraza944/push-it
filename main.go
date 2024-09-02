package main

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func main() {

	cmd, err := exec.Command("git", "branch").Output()
	if err != nil {
		log.Fatal(err)
	}

	branches := strings.Split(string(cmd), "\n")
	var currentBranch string
	for _, branch := range branches {

		if strings.HasPrefix(branch, "*") {
			currentBranch = strings.TrimPrefix(branch, "*")
			break
		}

	}

	// executablePath := fmt.Sprintf("push -u origin %s", currentBranch)

	push, err := exec.Command("git", "push", "-u", "origin", currentBranch).Output()

	if err != nil {
		log.Fatal("falied", err)
	}

	fmt.Println("xcxcx", push)

}
