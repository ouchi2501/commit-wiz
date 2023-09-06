package main

import (
	"fmt"
	"os/exec"
)

func retrieveGitDiff(currentDir string) (string, error) {
	cmd := exec.Command("git", "diff")
	cmd.Dir = currentDir
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Git diff command failed:", err)
		return "", err
	}
	gitDiff := string(output)
	return gitDiff, err
}
