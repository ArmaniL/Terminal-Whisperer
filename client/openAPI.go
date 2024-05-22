package client

import (
	"context"
	"tw/util"

	openai "github.com/sashabaranov/go-openai"
)

type OpenAPIClient struct {
}

func (OpenAPIClient) SendRequest(question string) (string, error) {

	token := util.GetEnvVariable("TOKEN")
	client := openai.NewClient(token)
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: question,
				},
			},
		},
	)
	if err != nil {
		return "", err
	}

	return resp.Choices[0].Message.Content, nil

}
