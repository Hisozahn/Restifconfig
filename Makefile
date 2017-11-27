GO=go
CLI_BINARY=./bin/cli
SERVICE_BINARY=./bin/service

CLI_SRC=./ifconfig-cli/main
SERVICE_SRC=./ifconfig-service/main

all: get service cli

get: 
	$(GO) get -v ./...

cli:
	$(GO) build -o $(CLI_BINARY) $(CLI_SRC)/main.go

service:
	$(GO) build -o $(SERVICE_BINARY) $(SERVICE_SRC)/main.go
