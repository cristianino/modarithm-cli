BINARY=modarithm

all: fmt lint test build

build:
	go build -o $(BINARY) ./main.go

run:
	go run ./main.go $(CMD)

test:
	go test ./...

lint:
	golangci-lint run

fmt:
	go fmt ./...
	go vet ./...
