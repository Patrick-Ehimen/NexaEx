# Variables
APP_NAME := cex-exchange
BUILD_DIR := bin
COVERAGE_FILE := coverage.out
COVERAGE_HTML := coverage.html

# Declare phony targets (so they’re always executed)
.PHONY: all build run test test-coverage fmt vet lint tidy vendor cross clean fresh help

# Default target: build the application
all: build

# Build the binary into the bin/ directory
build:
	@echo "Building $(APP_NAME)..."
	mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(APP_NAME) .

# Run the application (builds first if necessary)
run: build
	@echo "Running $(APP_NAME)..."
	./$(BUILD_DIR)/$(APP_NAME)

# Run tests for all packages
test:
	@echo "Running tests..."
	go test -v ./...

# Run tests with coverage and generate an HTML report
test-coverage:
	@echo "Running tests with coverage..."
	go test -coverprofile=$(COVERAGE_FILE) ./...
	@echo "Generating coverage report..."
	go tool cover -html=$(COVERAGE_FILE) -o $(COVERAGE_HTML)
	@echo "Coverage report generated at $(COVERAGE_HTML)"

# Format source code
fmt:
	@echo "Formatting source code..."
	go fmt ./...

# Run go vet to check for potential issues
vet:
	@echo "Running go vet..."
	go vet ./...

# Run linting (requires golangci-lint installed)
lint:
	@echo "Running lint..."
	golangci-lint run

# Tidy up module dependencies
tidy:
	@echo "Tidying up module dependencies..."
	go mod tidy

# Vendor dependencies into the vendor/ directory
vendor:
	@echo "Vendoring dependencies..."
	go mod vendor

# Cross compile for Linux, Darwin (macOS), and Windows
cross:
	@echo "Cross compiling for Linux..."
	GOOS=linux GOARCH=amd64 go build -o $(BUILD_DIR)/$(APP_NAME)-linux .
	@echo "Cross compiling for Darwin..."
	GOOS=darwin GOARCH=amd64 go build -o $(BUILD_DIR)/$(APP_NAME)-darwin .
	@echo "Cross compiling for Windows..."
	GOOS=windows GOARCH=amd64 go build -o $(BUILD_DIR)/$(APP_NAME)-windows.exe .

# Clean build artifacts and coverage files
clean:
	@echo "Cleaning up..."
	go clean
	rm -rf $(BUILD_DIR)
	rm -f $(COVERAGE_FILE) $(COVERAGE_HTML)

# Fresh: clean, then build and run the application
fresh: clean build run

# Help: list available targets
help:
	@echo ""
	@echo "Makefile for $(APP_NAME)"
	@echo ""
	@echo "Available targets:"
	@echo "  all             - Default target: builds the application"
	@echo "  build           - Build the binary"
	@echo "  run             - Run the binary (builds if needed)"
	@echo "  test            - Run tests"
	@echo "  test-coverage   - Run tests with coverage and generate an HTML report"
	@echo "  fmt             - Format the source code"
	@echo "  vet             - Run go vet to analyze code"
	@echo "  lint            - Run linting (requires golangci-lint)"
	@echo "  tidy            - Tidy up go.mod dependencies"
	@echo "  vendor          - Vendor dependencies"
	@echo "  cross           - Cross compile for Linux, Darwin, and Windows"
	@echo "  clean           - Remove build artifacts and coverage files"
	@echo "  fresh           - Clean, build, and run the application"
	@echo "  help            - Display this help message"
	@echo ""
