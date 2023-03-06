LIBRARY_NAME := go-toolbox

SRC_DIR := ./src/...
TEST_DIR := ./src/...
BUILD_DIR := ./build

.PHONY: all clean build test format

# The default target is to build the library.
all: build

# Clean the build output directory.
clean:
	@rm -rf $(BUILD_DIR)

# Build the library.
build: format
	@mkdir -p $(BUILD_DIR)
	@echo "Building $(LIBRARY_NAME)..."
	@go build $(GO_FLAGS) -o $(BUILD_DIR)/$(LIBRARY_NAME) $(SRC_DIR)

# Run the tests.
test: format
	@echo "Running tests for $(LIBRARY_NAME)..."
	@go test $(GO_FLAGS) $(TEST_DIR)

# Format the source code.
format:
	@echo "Formatting source code..."
	@go fmt $(SRC_DIR)