.PHONY: dev run dice-go all test clean

all: dice-go

dice-go:
	go build -o dice-go main.go

run: clean pre
	go run -race main.go

mod:
	go mod download

pre:
	go fmt ./...
	go vet ./...
