package main

import (
	"bytes"
	"os"
	"testing"
	"time"
)

func TestDisplayLoadingAnimation(t *testing.T) {
	// Create a channel to signal when the animation should stop
	done := make(chan struct{})

	// Capture the current stdout
	oldStdout := os.Stdout

	// Create a pipe to capture stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	var output string

	go func() {
		displayLoadingAnimation(done)
	}()

	// Let the animation run for a while (e.g., 500 milliseconds)
	time.Sleep(500 * time.Millisecond)

	// Close the channel to stop the animation
	close(done)

	// Restore the original stdout
	os.Stdout = oldStdout
	_ = w.Close()

	// Read the captured output
	buf := new(bytes.Buffer)
	_, _ = buf.ReadFrom(r)
	output = buf.String()

	// Verify that the captured output contains "Loading..."
	if !containsSubstring(output, "Loading...") {
		t.Error("Expected 'Loading...' to be present in the captured output, but it was not.")
	}

	// Allow some time for the animation to stop and check again
	time.Sleep(200 * time.Millisecond)
	buf.Reset()
	_, _ = buf.ReadFrom(r)
	output = buf.String()

	if containsSubstring(output, "Loading...") {
		t.Error("Expected 'Loading...' to be absent in the captured output after done channel is closed, but it was still present.")
	}
}
