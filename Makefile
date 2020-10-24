GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) gets

hello:
	echo "Hello"

build:
	go build -o bin/main ./cmd/football/football.go
