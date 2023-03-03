package handler

import (
	"bytes"
	"encoding/json"
	"net/http"

	"chatgpt-api-go/types"
)

var (
	msgs   []types.Message
	ApiKey string
)

func GetChatReq(content string) (*http.Request, error) {
	add2Messages(content, true)

	apiReq := types.ApiRequest{
		Model:    "gpt-3.5-turbo",
		Messages: msgs,
	}

	b, err := json.Marshal(apiReq)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+ApiKey)
	req.Header.Set("Content-Type", "application/json")

	return req, nil
}

// 添加对话到对话上下文
func add2Messages(content string, isUser bool) {
	if len(msgs) == 0 {
		first := types.Message{
			Role:    "system",
			Content: "You are a helpful assistant.",
		}
		msgs = append(msgs, first)
	}

	role := "assistant"
	if isUser {
		role = "user"
	}

	m := types.Message{
		Role:    role,
		Content: content,
	}

	msgs = append(msgs, m)
}

func clearMessages() {
	msgs = make([]types.Message, 0)
}
