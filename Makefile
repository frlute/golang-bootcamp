COMMONENVVAR       = GOOS=linux GOARCH=amd64
BUILDENVVAR        = CGO_ENABLED=0

# 定义环境变量
export GOBIN := $(CURDIR)/bin
export GO111MODULE=on
export GOPROXY=https://goproxy.cn
export GOSUMDB=off

# 构建依赖环境
deps_update:
	go mod tidy

deps:
	go mod vendor

# 生成可执行文件
build:
	go build -o $(GOBIN)/test -ldflags "-X main.BuildTime=`date '+%Y-%m-%d_%I:%M:%S%p'` -X main.BuildGitHash=`git rev-parse HEAD` -X main.BuildGitTag=`git describe --tags 2>/dev/null || echo NO_TAG`" cmd/search

# 交叉编译出 linux 下的静态可执行文件
linux_build: deps
	$(COMMONENVVAR) $(BUILDENVVAR) make build

.PHONY: deps build linux_build all test clean	