# Define variables
BINARY_NAME=stupebilder
BUILD_DIR=bin
SRC_DIR=./cmd/stupebilder
PACKAGE_NAME=dist.zip

# Default target: builds the binary
all: build

# Build the Go binary
build:
	mkdir -p $(BUILD_DIR)
	GOOS=linux GOARCH=amd64 go build -o $(BUILD_DIR)/$(BINARY_NAME) $(SRC_DIR)

# Clean up build artifacts
clean:
	rm -rf $(BUILD_DIR)

package: build
	rm -rf $(PACKAGE_NAME)
	zip -r $(PACKAGE_NAME) bin/ templates/ style/

# Run the application
run: build
	bin/stupebilder

# Cross-compile for multiple platforms
build-cross:
	GOOS=linux GOARCH=amd64 go build -o $(BUILD_DIR)/$(BINARY_NAME)-linux-amd64 $(SRC_DIR)
	GOOS=darwin GOARCH=amd64 go build -o $(BUILD_DIR)/$(BINARY_NAME)-darwin-amd64 $(SRC_DIR)
	GOOS=windows GOARCH=amd64 go build -o $(BUILD_DIR)/$(BINARY_NAME)-windows-amd64.exe $(SRC_DIR)

# Run tests
test:
	go test ./...

# Format code
fmt:
	go fmt ./...

.PHONY: all build clean run build-cross test fmt
