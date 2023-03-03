package main

import (
	"chatgpt-go/handler"
	"chatgpt-go/types"
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v3"
	"io"
	"net/http"
	"net/url"
	"os"
	"time"
)

func main() {
	conf := parseConf()
	proxy(conf.Proxy)
	handler.ApiKey = conf.ApiKey

	r := gin.Default()
	r.POST("/", handler.Ask)
	r.POST("/clear", handler.Clear)

	_, err := os.Stat("log")
	if err != nil {
		return
	}
	var f *os.File
	if os.IsNotExist(err) {
		f, _ = os.Create("log")
	} else {
		f, _ = os.Open("log")
	}
	gin.DefaultWriter = io.MultiWriter(f)

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
