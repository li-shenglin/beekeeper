.PHONY: service web lint test

all: build

build:
	go build -o bin/bacn cmd/main.go

mock-fetch:
	go get github.com/golang/mock/gomock
	go get github.com/golang/mock/mockgen

lint:
	#golangci-lint run ./cmd/...
	golangci-lint run ./...

fix:
	golangci-lint run --fix ./cmd/...
	golangci-lint run --fix ./pkg/...

test:
	go test -v -coverprofile=coverage.out ./pkg/...

