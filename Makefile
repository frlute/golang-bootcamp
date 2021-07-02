# 定义项目基本信息
CURRENT_GIT_GROUP := github.com/frlute
CURRENT_GIT_REPO  := golang-bootcamp

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
	mkdir -p coverage
	make test_DSA

coverage:
	find ./coverage -name "coverage.*.out.*" -exec tail +2 {} >> ./coverage/coverage.out \;
	go tool cover -html coverage.out -o coverage.html
	go tool cover -func=coverage.out

# DSA
test_DSA:
	go test -gcflags=-l -v $(CURRENT_GIT_GROUP)/$(CURRENT_GIT_REPO)/DSA/array -coverprofile=./coverage/coverage.out -covermode=count -coverpkg=./...
	go test -gcflags=-l -v $(CURRENT_GIT_GROUP)/$(CURRENT_GIT_REPO)/DSA/string -coverprofile=./coverage/coverage.dsa.out.1 -covermode=count -coverpkg=./...
	go test -gcflags=-l -v $(CURRENT_GIT_GROUP)/$(CURRENT_GIT_REPO)/DSA/linkedlist -coverprofile=./coverage/coverage.dsa.out.2 -covermode=count -coverpkg=./...
	go test -gcflags=-l -v $(CURRENT_GIT_GROUP)/$(CURRENT_GIT_REPO)/DSA/queue -coverprofile=./coverage/coverage.dsa.out.3 -covermode=count -coverpkg=./...
	go test -gcflags=-l -v $(CURRENT_GIT_GROUP)/$(CURRENT_GIT_REPO)/DSA/stack -coverprofile=./coverage/coverage.dsa.out.4 -covermode=count -coverpkg=./...
	go test -gcflags=-l -v $(CURRENT_GIT_GROUP)/$(CURRENT_GIT_REPO)/DSA/tree -coverprofile=./coverage/coverage.dsa.out.5 -covermode=count -coverpkg=./...
	go test -gcflags=-l -v $(CURRENT_GIT_GROUP)/$(CURRENT_GIT_REPO)/DSA/search -coverprofile=./coverage/coverage.dsa.out.6 -covermode=count -coverpkg=./...
	go test -gcflags=-l -v $(CURRENT_GIT_GROUP)/$(CURRENT_GIT_REPO)/DSA/hash -coverprofile=./coverage/coverage.dsa.out.7 -covermode=count -coverpkg=./...
	go test -gcflags=-l -v $(CURRENT_GIT_GROUP)/$(CURRENT_GIT_REPO)/DSA/graph -coverprofile=./coverage/coverage.dsa.out.8 -covermode=count -coverpkg=./...

clean:
	rm -rf coverage

linux_test:
	$(COMMONENVVAR) $(BUILDENVVAR) make test

.PHONY: deps build linux_build all test clean	