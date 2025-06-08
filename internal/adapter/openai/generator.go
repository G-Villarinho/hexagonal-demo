package openai

import (
	"context"

	"github.com/g-villarinho/hexagonal-demo/internal/core/port"
	"github.com/sashabaranov/go-openai"
)

type openAIGenerator struct {
	client *openai.Client
}

func NewOpenAiGenerator(client *openai.Client) port.AIGenerator {
	return &openAIGenerator{
		client: client,
	}
}

func (g *openAIGenerator) GenerateText(ctx context.Context, prompt string) (string, error) {
	resp, err := g.client.CreateChatCompletion(
		ctx,
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: prompt,
				},
			},
		},
	)

	if err != nil {
		return "", err
	}

	return resp.Choices[0].Message.Content, nil
}
