GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) gets



.PHONY: build
build:
	go build -v ./cmd/football

.DEFAULT_GOAL := build
