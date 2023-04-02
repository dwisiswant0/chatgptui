package chat

import (
	"context"
	"strings"
)

func (m model) sendChat(prompt string) (string, error) {
	ctx := context.Background()

	m.openaiRequest.Prompt = prompt

	resp, err := m.openaiClient.CreateCompletion(ctx, m.openaiRequest)
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(resp.Choices[0].Text), nil
}
