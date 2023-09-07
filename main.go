package main

import (
	"context"
	"flag"
	"fmt"
	"os"
)

func main() {
	// command line flags
	summaryLength := flag.Int("l", 50, "Number of tokens in summary")
	flag.Parse()

	// get current directory
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	// get git diff
	gitDiff, err := retrieveGitDiff(dir)
	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	// get openai key from environment variable
	key := os.Getenv("OPENAI_KEY")
	if key == "" {
		panic("OPENAI_KEY is not set")
	}

	// add loading animation
	loadingDone := make(chan struct{})
	go displayLoadingAnimation(loadingDone)

	// request openai
	response := requestOpenAI(ctx, key, gitDiff, *summaryLength)

	close(loadingDone)

	// display results
	fmt.Print("\n")
	fmt.Println("The generated commit message is below:")
	fmt.Print("\n")
	fmt.Println(response)
	fmt.Print("\n")
}
