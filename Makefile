LIBRARY_NAME := go-toolbox

SRC_DIR := ./src/...
TEST_DIR := ./src/...
BUILD_DIR := ./build

.PHONY: all test format

# The default target is to build the library.
all: test

# Run the tests.
test: format
	@echo "Running tests for $(LIBRARY_NAME)..."
	@go test $(GO_FLAGS) $(TEST_DIR)

# Format the source code.
format:
	@echo "Formatting source code..."
	@go fmt $(SRC_DIR)