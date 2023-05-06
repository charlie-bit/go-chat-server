PROJECT = "chat_socket"
OS=$(shell uname -s)
VERSION ?= $(shell git describe --tags --always || git rev-parse --short HEAD)
BUILD_TIME ?= $(shell date +%Y-%m-%dT%T%z)

GO ?= go

.PHONY:
default:
	@echo Project: ${PROJECT}

clean:
	@rm -f bin/*

build:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 $(GO) build -o bin/tcp_server cmd/tcp/server/server.go
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 $(GO) build -o bin/tcp_client cmd/tcp/client/client.go