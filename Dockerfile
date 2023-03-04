FROM golang:alpine AS builder
WORKDIR /app

ENV GOPROXY=https://goproxy.cn GOOS=linux CGO_ENABLED=0 GOARCH=amd64

RUN sed -i "s/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g" /etc/apk/repositories \
    && apk add make \
    && apk add git

COPY . .
RUN make

FROM alpine:latest AS prod
WORKDIR /app

COPY --from=builder /app .

ENTRYPOINT /app/chatgpt-api-go
