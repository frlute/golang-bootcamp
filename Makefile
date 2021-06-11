# 定义项目基本信息
CURRENT_GIT_GROUP := github.com/frlute
CURRENT_GIT_REPO  := go-playground

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

test:
	make test_DSA
	make test_search

coverage:
	find . -name "coverage.*.out.*" -exec tail +2 {} >> coverage.out \;
	go tool cover -html coverage.out -o coverage.html
	go tool cover -func=coverage.out

# DSA
test_DSA:
	go test -gcflags=-l -v $(CURRENT_GIT_GROUP)/$(CURRENT_GIT_REPO)/DSA/array -coverprofile=coverage.out -covermode=count -coverpkg=./...
	go test -gcflags=-l -v $(CURRENT_GIT_GROUP)/$(CURRENT_GIT_REPO)/DSA/string -coverprofile=coverage.dsa.out.1 -covermode=count -coverpkg=./...
	go test -gcflags=-l -v $(CURRENT_GIT_GROUP)/$(CURRENT_GIT_REPO)/DSA/linkedlist -coverprofile=coverage.dsa.out.2 -covermode=count -coverpkg=./...
	go test -gcflags=-l -v $(CURRENT_GIT_GROUP)/$(CURRENT_GIT_REPO)/DSA/queue -coverprofile=coverage.dsa.out.3 -covermode=count -coverpkg=./...
	go test -gcflags=-l -v $(CURRENT_GIT_GROUP)/$(CURRENT_GIT_REPO)/DSA/stack -coverprofile=coverage.dsa.out.4 -covermode=count -coverpkg=./...
	go test -gcflags=-l -v $(CURRENT_GIT_GROUP)/$(CURRENT_GIT_REPO)/DSA/tree -coverprofile=coverage.dsa.out.5 -covermode=count -coverpkg=./...
	go test -gcflags=-l -v $(CURRENT_GIT_GROUP)/$(CURRENT_GIT_REPO)/DSA/search -coverprofile=coverage.dsa.out.6 -covermode=count -coverpkg=./...
	go test -gcflags=-l -v $(CURRENT_GIT_GROUP)/$(CURRENT_GIT_REPO)/DSA/hash -coverprofile=coverage.dsa.out.7 -covermode=count -coverpkg=./...
	go test -gcflags=-l -v $(CURRENT_GIT_GROUP)/$(CURRENT_GIT_REPO)/DSA/graph -coverprofile=coverage.dsa.out.8 -covermode=count -coverpkg=./...

test_search:
	go test -gcflags=-l -v $(CURRENT_GIT_GROUP)/$(CURRENT_GIT_REPO)/search/BinarySearch -coverprofile=coverage.search.out -covermode=count -coverpkg=./...
	go test -gcflags=-l -v $(CURRENT_GIT_GROUP)/$(CURRENT_GIT_REPO)/search/StringSearch -coverprofile=coverage.search.out.1 -covermode=count -coverpkg=./...


linux_test:
	$(COMMONENVVAR) $(BUILDENVVAR) make test

.PHONY: deps build linux_build all test clean	