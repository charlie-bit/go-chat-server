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
	go test -v ./...
	go build -v ./...