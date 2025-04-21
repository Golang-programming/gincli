.PHONY: build clean release

# Version information
VERSION ?= $(shell git describe --tags --always --dirty)
BUILD_DATE = $(shell date +%Y-%m-%dT%H:%M:%S%z)
LDFLAGS = -ldflags "-X main.version=$(VERSION) -X main.buildDate=$(BUILD_DATE)"

# Build for current platform
build:
	go build $(LDFLAGS) -o gincli

# Build for all platforms
release: clean
	GOOS=linux GOARCH=amd64 go build $(LDFLAGS) -o bin/gincli-linux-amd64
	GOOS=linux GOARCH=arm64 go build $(LDFLAGS) -o bin/gincli-linux-arm64
	GOOS=darwin GOARCH=amd64 go build $(LDFLAGS) -o bin/gincli-darwin-amd64
	GOOS=darwin GOARCH=arm64 go build $(LDFLAGS) -o bin/gincli-darwin-arm64
	GOOS=windows GOARCH=amd64 go build $(LDFLAGS) -o bin/gincli-windows-amd64.exe
	GOOS=windows GOARCH=386 go build $(LDFLAGS) -o bin/gincli-windows-386.exe

# Clean build artifacts
clean:
	rm -rf bin/
	mkdir -p bin 