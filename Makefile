# Name of the binary to build
BINARY_NAME=releasebot

# Go source files
SRC=$(shell find . -name "*.go" -type f)

# Build the binary for the current platform
build:
	go build -ldflags="-s -w" -o $(BINARY_NAME) ./cmd/releasebot

build-race:
	go build -race -ldflags="-s -w" -o $(BINARY_NAME) ./cmd/releasebot

# Clean the project
clean:
	go clean
	rm -f $(BINARY_NAME)

# Run the tests
test:
	go test -v ./...

# Format the source code
fmt:
	gofmt -w $(SRC)