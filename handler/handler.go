package handler

import (
	"chatgpt-go/types"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
)

func Ask(c *gin.Context) {
	// 从json body中读出提问文本
	var a types.ClientRequest
	err := c.BindJSON(&a)
	handlerErr(c, err)
	log.Println(a.Content)

	chatReq := GetChatReq(a.Content)

	// 使用默认的HTTP客户端发送请求
	resp, err := http.DefaultClient.Do(chatReq)
	handlerErr(c, err)
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Println(err)
	}

	// 解析到api的响应结构体
	bytes, err := io.ReadAll(resp.Body)
	handlerErr(c, err)
	var apiResp types.ApiResponse
	err = json.Unmarshal(bytes, &apiResp)
	handlerErr(c, err)

	// 响应给前端
	answer := apiResp.Choices[0].Message.Content
	cliResp := types.ClientResponse{
		Question: a.Content,
		Answer:   answer,
	}
	add2Messages(answer, false)
	c.PureJSON(200, &cliResp)
}

func Clear(c *gin.Context) {
	// 清空对话上下文
	clearMessages()
	c.PureJSON(200, "clear chat success")
}

func handlerErr(c *gin.Context, err error) {
	if err != nil {
		c.PureJSON(400, err.Error())
	}
}
