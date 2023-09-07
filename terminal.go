package main

import (
	"fmt"
	"time"
)

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
