package handler

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"chatgpt-api-go/types"
)

var (
	msgs []types.Message
	Auth string
)

// Send2ChatGPT 返回chatGPT的回答
func Send2ChatGPT(content string) (string, error) {
	add2Messages(content, true)

	apiReq := types.ApiRequest{
		Model:    "gpt-3.5-turbo",
		Messages: msgs,
	}

	b, err := json.Marshal(apiReq)
	if err != nil {
		return "", err
	}

	// todo 使用http连接池
	req, err := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", bytes.NewBuffer(b))
	if err != nil {
		return "", err
	}
	req.Header.Set("Authorization", Auth)
	req.Header.Set("Content-Type", "application/json")

	// 使用默认的HTTP客户端发送请求
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// 解析到api的响应结构体
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != http.StatusOK {
		log.Println("openai resp :", string(data))
		return string(data), err
	}

	var apiResp types.ApiResponse
	err = json.Unmarshal(data, &apiResp)
	if err != nil {
		// 打印openai返回的错误信息
		return string(data), err
	}

	return apiResp.Choices[0].Message.Content, nil
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
