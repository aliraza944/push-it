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
			currentBranch = strings.TrimSpace(strings.TrimPrefix(branch, "*"))
			break
		}

	}
	push := exec.Command("git", "push", "-u", "origin", currentBranch)
	output, err := push.CombinedOutput() // Capture both stdout and stderr

	if err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			// Handle specific exit codes if needed
			log.Fatalf("git push failed with exit code %d: %s", exitError.ExitCode(), output)
		} else {
			log.Fatalf("git push failed: %s", err)
		}
	}

	fmt.Println(string(output))

}
