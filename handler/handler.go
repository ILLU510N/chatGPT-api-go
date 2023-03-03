package handler

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"chatgpt-api-go/types"
	"github.com/gin-gonic/gin"
)

func Ask(c *gin.Context) {
	// 从json body中读出提问文本
	var a types.ClientRequest
	err := c.BindJSON(&a)
	handlerErr(c, err)
	log.Println(a.Content)

	chatReq, err := GetChatReq(a.Content)
	handlerErr(c, err)

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
	add2Messages(answer, false)
	c.String(200, answer)
}

func Clear(c *gin.Context) {
	// 清空对话上下文
	clearMessages()
	c.PureJSON(200, "clear chat success")
}

func Ping(c *gin.Context) {
	c.String(200, "Pong!")
}

func handlerErr(c *gin.Context, err error) {
	if err != nil {
		c.PureJSON(400, err.Error())
	}
}
