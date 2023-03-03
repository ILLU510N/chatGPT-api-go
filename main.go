package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"time"

	"chatgpt-api-go/handler"
	"chatgpt-api-go/types"
	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v3"
)

func main() {
	conf := parseConf()
	// 可选择配置网络代理
	if len(conf.Proxy) > 0 {
		proxy(conf.Proxy)
	}
	handler.Auth = fmt.Sprintf("Bearer %s", conf.ApiKey)

	// 控制日志输出到文件
	gin.DisableConsoleColor()
	f, _ := os.OpenFile("./app.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	r := gin.Default()
	r.POST("/", handler.Ask)
	r.POST("/clear", handler.Clear)
	r.POST("/ping", handler.Ping)

	_ = r.Run(fmt.Sprintf(":%d", conf.Port))
}

func parseConf() *types.Conf {
	data, err := os.ReadFile("conf.yaml")
	if err != nil {
		panic(err)
		return nil
	}
	var conf types.Conf
	err = yaml.Unmarshal(data, &conf)
	if err != nil {
		panic(err)
		return nil
	}

	return &conf
}

func proxy(proxyUrl string) {
	// 配置代理服务器的网络地址
	proxyURL, err := url.Parse(proxyUrl)
	if err != nil {
		panic(err)
	}

	// 创建TCP连接代理，以便将HTTP请求发送到代理服务器
	http.DefaultTransport = &http.Transport{
		Proxy:               http.ProxyURL(proxyURL),
		TLSHandshakeTimeout: 10 * time.Second,
		MaxIdleConnsPerHost: 10,
	}
}
