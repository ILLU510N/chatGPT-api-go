# ChatGPT api for golang

simple demo for golang chatGPT api server.

---

powered by gin. 

### Install
```shell
git clone https://github.com/ILLU510N/chatGPT-api-go.git
go build
./chatgpt-api-go
```

### Usage
client use http post to send message to server.
```shell
curl -H "Content-Type: application/json" -X POST -d "{\"content\":\"what is chatGPT\"}" "http://localhost:8080"
```

### Config
edit conf.yaml

```yaml
# example
port: 8080
# use proxy to visit openai
proxy: http://127.0.0.1:1234
api_key: yourkey
```
