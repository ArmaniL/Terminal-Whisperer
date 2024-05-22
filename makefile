# Define the directory for the build outputs
BUILD_DIR := build

# Find all .go files in the current directory (recursively)
GO_FILES := $(wildcard *.go) $(wildcard */*.go)

# Define the binary name
BINARY_NAME := tw

# Default target
.PHONY: all
all: clean build

# Create the build directory
.PHONY: build
build: $(GO_FILES)
	@echo "Building Go files..."
	@mkdir -p $(BUILD_DIR)
	@go build -o $(BUILD_DIR)/$(BINARY_NAME)

# Clean the build directory
.PHONY: clean
clean:
	@echo "Cleaning build directory..."
	@rm -rf $(BUILD_DIR)

# Install dependencies
.PHONY: deps
deps:
	@echo "Installing dependencies..."
	@go mod tidy

# Run tests
.PHONY: test
test:
	@echo "Running tests..."
	@go test ./...

# Run the application
.PHONY: run
run: build
	@echo "Running the application..."
	@$(BUILD_DIR)/$(BINARY_NAME)
