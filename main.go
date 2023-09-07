package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {
	// command line flags
	summaryLength := flag.Int("l", 50, "Number of tokens in summary")
	path := flag.String("p", "", "Git repository path(default: current directory)")
	flag.Parse()

	// get current directory
	var dir string
	if path != nil && *path != "" {
		cwd, err := os.Getwd()
		if err != nil {
			panic(err)
		}
		dir = cwd
	} else {
		dir = *path
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

func displayLoadingAnimation(done chan struct{}) {
	animationChars := `|/-\`
	i := 0
	for {
		select {
		case <-done:
			fmt.Print("\n")
			return
		default:
			fmt.Printf("\rLoading... %c", animationChars[i])
			i = (i + 1) % len(animationChars)
			time.Sleep(100 * time.Millisecond)
		}
	}
}
