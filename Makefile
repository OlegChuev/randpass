.PHONY: build run test install help clean

# Build settings
BINARY_NAME=randpass
BUILD_DIR=bin
CMD_DIR=cmd/randpass

# Go settings
GOOS := $(shell go env GOOS)
GOARCH := $(shell go env GOARCH)

help: ## Show this help message
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Targets:'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  %-15s %s\n", $$1, $$2}' $(MAKEFILE_LIST)

build: ## Build the binary
	@echo "Building $(BINARY_NAME)..."
	@mkdir -p $(BUILD_DIR)
	@go build -ldflags="-s -w -linkmode=internal" -o $(BUILD_DIR)/$(BINARY_NAME) ./$(CMD_DIR)
	@echo "Binary built: $(BUILD_DIR)/$(BINARY_NAME)"

run: build ## Build and run the application with default settings
	@echo "Running $(BINARY_NAME)..."
	@./$(BUILD_DIR)/$(BINARY_NAME)

install: build ## Install the binary to $GOPATH/bin
	@echo "Installing $(BINARY_NAME) to $(GOPATH)/bin..."
	@cp $(BUILD_DIR)/$(BINARY_NAME) $(GOPATH)/bin/
	@echo "Installed successfully!"

test: ## Run all tests
	@echo "Running tests..."
	@go test -v ./...

test-coverage: ## Run tests with coverage report
	@echo "Running tests with coverage..."
	@go test -cover ./...

deps: ## Download dependencies
	@echo "Downloading dependencies..."
	@go mod download
	@go mod tidy

fmt: ## Format Go code
	@echo "Formatting code..."
	@go fmt ./...

vet: ## Run go vet
	@echo "Running go vet..."
	@go vet ./...

lint: fmt vet ## Run formatting and vetting

dev-setup: deps ## Set up development environment
	@echo "Setting up development environment..."
	@go mod download

# Example runs
example-default: build ## Run with default settings (16 chars, all character types)
	@echo "Example: Default password"
	@./$(BUILD_DIR)/$(BINARY_NAME)

example-no-symbols: build ## Run without symbols (24 chars)
	@echo "Example: 24-char password without symbols"
	@./$(BUILD_DIR)/$(BINARY_NAME) -l 24 --no-symbols

example-simple: build ## Run with only uppercase and digits (12 chars)
	@echo "Example: 12-char password with only uppercase and digits"
	@./$(BUILD_DIR)/$(BINARY_NAME) --length 12 --no-lower --no-symbols

example-help: build ## Show help message
	@echo "Example: Help message"
	@./$(BUILD_DIR)/$(BINARY_NAME) --help

bench: ## Run benchmarks
	@echo "Running benchmarks..."
	@go test -bench=. -benchmem ./internal/generator

bench-short: ## Run quick benchmarks
	@echo "Running quick benchmarks..."
	@go test -bench=BenchmarkGenerator_Generate -benchmem ./internal/generator

bench-scaling: ## Run scaling benchmarks
	@echo "Running scaling benchmarks..."
	@go test -bench=BenchmarkGenerator_PasswordLength -benchmem ./internal/generator

bench-comparison: ## Run comparison table benchmarks
	@echo "Running comparison benchmarks..."
	@go test -bench=BenchmarkGenerator_CharacterSetSize -benchmem ./internal/generator

bench-memory: ## Run memory allocation benchmarks
	@echo "Running memory benchmarks..."
	@go test -bench=BenchmarkGenerator_MemoryAllocation -benchmem ./internal/generator

bench-crypto: ## Run crypto/rand benchmarks
	@echo "Running crypto/rand benchmarks..."
	@go test -bench=BenchmarkCryptoRand -benchmem ./internal/generator

performance-test: ## Run performance regression test
	@echo "Running performance regression test..."
	@go test -run=TestGenerator_PerformanceBaseline ./internal/generator

build-release: ## Build release binaries for all platforms
	@echo "Building release binaries..."
	@mkdir -p dist
	@echo "Building for Linux amd64..."
	@GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o dist/randpass-linux-amd64 ./cmd/randpass
	@echo "Building for Linux arm64..."
	@GOOS=linux GOARCH=arm64 go build -ldflags="-s -w" -o dist/randpass-linux-arm64 ./cmd/randpass
	@echo "Building for macOS amd64..."
	@GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -o dist/randpass-macos-amd64 ./cmd/randpass
	@echo "Building for macOS arm64..."
	@GOOS=darwin GOARCH=arm64 go build -ldflags="-s -w" -o dist/randpass-macos-arm64 ./cmd/randpass
	@echo "Building for Windows amd64..."
	@GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o dist/randpass-windows-amd64.exe ./cmd/randpass
	@echo "All binaries built successfully!"
	@ls -la dist/

release-checksums: build-release ## Generate checksums for release binaries
	@echo "Generating checksums..."
	@cd dist && sha256sum randpass-* > checksums.txt
	@echo "Checksums generated:"
	@cat dist/checksums.txt

clean: ## Clean build artifacts
	@echo "Cleaning..."
	@rm -rf $(BUILD_DIR)
	@rm -rf dist
	@go clean
