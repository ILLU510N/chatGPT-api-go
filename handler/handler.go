package handler

import (
	"chatgpt-api-go/types"
	"github.com/gin-gonic/gin"
)

func Ask(c *gin.Context) {
	// 从json body中读出提问文本
	var req types.ClientRequest
	err := c.BindJSON(&req)
	handlerErr(c, err)

	answer, err := Send2ChatGPT(req.Content)
	handlerErr(c, err)

	// 响应给前端
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
