BINARY=modarithm
VERSION ?= $(shell git describe --tags --always --dirty)
LDFLAGS=-ldflags "-s -w -X main.version=$(VERSION)"

all: fmt lint test build

build:
	go build $(LDFLAGS) -o $(BINARY) ./main.go

build-all:
	@echo "Building for all platforms..."
	GOOS=linux GOARCH=amd64 go build $(LDFLAGS) -o $(BINARY)-linux-amd64 ./main.go
	GOOS=linux GOARCH=arm64 go build $(LDFLAGS) -o $(BINARY)-linux-arm64 ./main.go
	GOOS=darwin GOARCH=amd64 go build $(LDFLAGS) -o $(BINARY)-darwin-amd64 ./main.go
	GOOS=darwin GOARCH=arm64 go build $(LDFLAGS) -o $(BINARY)-darwin-arm64 ./main.go
	GOOS=windows GOARCH=amd64 go build $(LDFLAGS) -o $(BINARY)-windows-amd64.exe ./main.go

run:
	go run ./main.go $(CMD)

test:
	go test ./...

lint:
	golangci-lint run

fmt:
	go fmt ./...
	go vet ./...

clean:
	rm -f $(BINARY)*

# Release helpers
tag:
	@echo "Current tags:"
	@git tag -l | tail -5
	@echo ""
	@echo "Create a new tag with: make tag VERSION=v1.0.0"
ifdef VERSION
	git tag $(VERSION)
	git push origin $(VERSION)
	@echo "Tag $(VERSION) created and pushed!"
else
	@echo "Usage: make tag VERSION=v1.0.0"
endif

release-check:
	@echo "Pre-release checklist:"
	@echo "✓ Run tests: make test"
	@echo "✓ Run linting: make lint"
	@echo "✓ Check build: make build-all"
	@echo "✓ Update version and create tag: make tag VERSION=vX.Y.Z"

.PHONY: all build build-all run test lint fmt clean tag release-check
