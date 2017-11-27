export GOROOT=/usr/lib/go-1.9
export GOPATH=$(PWD)
GO=go
CLI_BINARY=./pkg/cli
SERVICE_BINARY=./pkg/service
TEST_BINARY=./pkg/sys-test
CLI_SRC=src/ifconfig-cli/main
SERVICE_SRC=src/ifconfig-service/main
TEST_SRC=src/sys-test/main

cli:
	cd $(CLI_SRC) && $(GO) get
	$(GO) build -o $(CLI_BINARY) $(CLI_SRC)/main.go

service:
	cd $(SERVICE_SRC) && $(GO) get
	$(GO) build -o $(SERVICE_BINARY) $(SERVICE_SRC)/main.go

test:
	cd src/sys-test/main/ && $(GO) get
	cd ./pkg && $(GO) test -c  ./../$(TEST_SRC)/main_test.go

all: service cli test

run_tests:
	cd $(CLI_SRC)/.. && $(GO) test ./...
	cd $(SERVICE_SRC)/.. && $(GO) test ./...
	
