IMPORT_PATH		:= chatgpt-api-go
BUILD_VERSION   := 1.0.0
BUILD_BRANCH	:= `git symbolic-ref HEAD 2>/dev/null | cut -d"/" -f 3`
LAST_COMMIT		:= `git log -1`
BUILD_TIME      := $(shell date "+%F %T")
BUILD_NAME		:= chatgpt-api-go
GO_VERSION		:= `go version`
SOURCE          := ./*.go
TARGET_DIR      := ./bin
DEPLOY_VERSION  := default

linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
    -o ${BUILD_NAME} ${SOURCE}

clean:
	rm ${BUILD_NAME} -f

install:
	mkdir -p ${TARGET_DIR}
	cp ${BUILD_NAME} ${TARGET_DIR}/${BUILD_NAME}.${BUILD_VERSION} -f

deploy:
	cp ${BUILD_NAME} deploy/base/bin/${BUILD_NAME}.${DEPLOY_VERSION} -f

.PHONY : all clean install deploy ${BUILD_NAME}