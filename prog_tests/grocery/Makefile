export GOOS?=$(shell go env GOOS)
export GOARCH?=$(shell go env GOARCH)

build: test
	GOARCH=$(GOARCH) GOOS=$(GOOS) go build -v -o bin/grocery main.go

test:
	go test -race ./...