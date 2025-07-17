.PHONY: build clean run test deps help prepare-frontend

# Build variables
BINARY_NAME=openwebui-go
VERSION=$(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")
BUILD_TIME=$(shell date -u '+%Y-%m-%d_%H:%M:%S')
LDFLAGS=-ldflags "-X main.Version=${VERSION} -X main.BuildTime=${BUILD_TIME}"

# Default target
all: build

# Install dependencies
deps:
	go mod download
	go mod tidy

# Prepare OpenWebUI frontend
prepare-frontend:
	@echo "Preparing OpenWebUI frontend..."
	./scripts/prepare_frontend.sh

# Build the application
build: deps prepare-frontend
	@echo "Building ${BINARY_NAME}..."
	go build ${LDFLAGS} -o bin/${BINARY_NAME} .

# Build for multiple platforms
build-all: deps prepare-frontend
	@echo "Building for multiple platforms..."
	
	# Linux
	GOOS=linux GOARCH=amd64 go build ${LDFLAGS} -o bin/${BINARY_NAME}-linux-amd64 ./cmd
	GOOS=linux GOARCH=arm64 go build ${LDFLAGS} -o bin/${BINARY_NAME}-linux-arm64 ./cmd
	
	# macOS
	GOOS=darwin GOARCH=amd64 go build ${LDFLAGS} -o bin/${BINARY_NAME}-darwin-amd64 ./cmd
	GOOS=darwin GOARCH=arm64 go build ${LDFLAGS} -o bin/${BINARY_NAME}-darwin-arm64 ./cmd
	
	# Windows
	GOOS=windows GOARCH=amd64 go build ${LDFLAGS} -o bin/${BINARY_NAME}-windows-amd64.exe ./cmd

# Run the application
run: build
	@echo "Running ${BINARY_NAME}..."
	./bin/${BINARY_NAME}

# Run with debug mode
run-debug: build
	@echo "Running ${BINARY_NAME} in debug mode..."
	./bin/${BINARY_NAME} --debug

# Run tests
test:
	go test -v ./...

# Clean build artifacts
clean:
	@echo "Cleaning..."
	rm -rf bin/
	rm -rf temp_*/
	rm -rf assets/frontend/*
	@mkdir -p assets/frontend
	@echo "# This file exists to preserve the directory structure in git" > assets/frontend/.gitkeep
	@echo "# The actual frontend files will be pulled and built at build time" >> assets/frontend/.gitkeep
	go clean

# Install the binary
install: build
	@echo "Installing ${BINARY_NAME}..."
	cp bin/${BINARY_NAME} /usr/local/bin/

# Development mode (with hot reload)
dev:
	@echo "Starting development mode..."
	@if command -v air >/dev/null 2>&1; then \
		air; \
	else \
		echo "Air not found. Install with: go install github.com/cosmtrek/air@latest"; \
		echo "Running without hot reload..."; \
		go run ./cmd; \
	fi

# Show help
help:
	@echo "Available targets:"
	@echo "  build      - Build the application"
	@echo "  build-all  - Build for multiple platforms"
	@echo "  run        - Build and run the application"
	@echo "  run-debug  - Build and run with debug mode"
	@echo "  test       - Run tests"
	@echo "  clean      - Clean build artifacts"
	@echo "  install    - Install the binary"
	@echo "  dev        - Run in development mode"
	@echo "  deps       - Install dependencies"
	@echo "  prepare-frontend - Prepare OpenWebUI frontend"
	@echo "  help       - Show this help" 