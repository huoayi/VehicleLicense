package main

import (
	"context"
	"fmt"
	gogpt "github.com/sashabaranov/go-gpt3"
)

func main() {
 	c := gogpt.NewClient("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")
	ctx := context.Background()
	req := gogpt.CompletionRequest{
		Model: "text-davinci-003",
		MaxTokens: 1234,
		Prompt:    "海商法论文800字论文" ,
		Temperature: 0.05,
	}
	resp, err := c.CreateCompletion(ctx, req)
	if err != nil {
		return
	}
	fmt.Println(resp.Choices[0].Text)


}