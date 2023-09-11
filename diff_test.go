package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

func TestRetrieveGitDiff(t *testing.T) {
	// Create a temporary test directory and defer its removal
	tempDir, err := os.MkdirTemp("", "git_diff_test")
	if err != nil {
		t.Fatalf("Error creating temporary directory: %v", err)
	}
	defer func(path string) {
		err := os.RemoveAll(path)
		if err != nil {
			t.Fatalf("Error removing temporary directory: %v", err)
		}
	}(tempDir)

	// Initialize a test Git repository
	cmd := exec.Command("git", "init")
	cmd.Dir = tempDir
	err = cmd.Run()
	if err != nil {
		t.Fatalf("Error initializing Git repository: %v", err)
	}

	// Set Git user for the test repository
	cmd = exec.Command("git", "config", "user.email", "test@example.com")
	cmd.Dir = tempDir
	err = cmd.Run()
	if err != nil {
		t.Fatalf("Error setting Git user email: %v", err)
	}

	cmd = exec.Command("git", "config", "user.name", "Test User")
	cmd.Dir = tempDir
	err = cmd.Run()
	if err != nil {
		t.Fatalf("Error setting Git user name: %v", err)
	}

	// Create a test file and commit it
	testFile := "test.txt"
	err = os.WriteFile(filepath.Join(tempDir, testFile), []byte("Hello, World!"), 0644)
	if err != nil {
		t.Fatalf("Error creating test file: %v", err)
	}
	cmd = exec.Command("git", "add", testFile)
	cmd.Dir = tempDir
	err = cmd.Run()
	gitStatus(tempDir)
	if err != nil {
		t.Fatalf("Error adding test file to Git: %v", err)
	}

	cmd = exec.Command("git", "commit", "-m", "Initial commit")
	cmd.Dir = tempDir
	err = cmd.Run()
	if err != nil {
		gitStatus(tempDir)
		t.Fatalf("Error committing test file: %v", err)
	}

	// Test the retrieveGitDiff function
	gitDiff, err := retrieveGitDiff(tempDir)
	if err != nil {
		t.Fatalf("Error calling retrieveGitDiff: %v", err)
	}

	// Verify the test result
	expectedDiff := "" // There is no diff for the initial commit
	if gitDiff != expectedDiff {
		t.Errorf("Expected empty git diff, but got:\n%s", gitDiff)
	}

	// Append text to the test file
	appendText := "Additional content"
	err = appendToFile(filepath.Join(tempDir, testFile), appendText)
	if err != nil {
		t.Fatalf("Error appending text to test file: %v", err)
	}

	// Test the retrieveGitDiff function after appending text
	gitDiff, err = retrieveGitDiff(tempDir)
	if err != nil {
		// Verify that an error is returned when the git diff command fails
		expectedError := "Git diff command failed:"
		if !containsSubstring(err.Error(), expectedError) {
			t.Errorf("Expected error message to contain: %s\nBut it did not:\n%s", expectedError, err.Error())
		}
		return
	}

	// Verify the test result after appending text
	expectedDiff = fmt.Sprintf("diff --git a/%s b/%s\nindex", testFile, testFile)
	if !containsSubstring(gitDiff, expectedDiff) {
		t.Errorf("Expected git diff to contain: %s\nBut it did not:\n%s", expectedDiff, gitDiff)
	}

	// Test the retrieveGitDiff function error handling
	_, err = retrieveGitDiff("/tmp/test")
	if err != nil {
		// Verify that an error is returned when the git diff command fails
		expectedError := "no such file or directory"
		if !containsSubstring(err.Error(), expectedError) {
			t.Errorf("Expected error message to contain: %s\nBut it did not:\n%s", expectedError, err.Error())
		}
		return
	}

}

func appendToFile(filePath, text string) error {
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	_, err = file.WriteString(text)
	return err
}

func containsSubstring(s, substr string) bool {
	return strings.Contains(s, substr)
}

func gitStatus(dir string) {
	cmd := exec.Command("git", "status")
	cmd.Dir = dir
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Error running 'git status': %v\n%s", err, output)
	}
	fmt.Printf("Git status:\n%s\n", output)
}
