package main

import (
	"context"
	"fmt"

	"github.com/sashabaranov/go-openai"
)

const systemContent = `
[SUMMARIZE_LENGTH=%d]
You are an assistant that thinks about git commit messages. Once you have entered the git diff, come up with a commit message from it.

considerations
- Only reply with the content of the commit message
- Lines that start with + are added lines, lines that start with - are deleted lines
- Reply with the Token length specified in SUMMARIZE_LENGTH
- Please mention all changes in principle
`

func requestOpenAI(ctx context.Context, key string, gitDiff string, length int) string {
	client := openai.NewClient(key)
	response, err := client.CreateChatCompletion(ctx, openai.ChatCompletionRequest{
		Model: openai.GPT3Dot5Turbo,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleSystem,
				Content: fmt.Sprintf(systemContent, length),
			},
			{
				Role:    openai.ChatMessageRoleUser,
				Content: gitDiff,
			},
		},
	})
	if err != nil {
		panic(err)
	}
	return response.Choices[0].Message.Content
}
