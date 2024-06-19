# Variables
PKG := ./...
GOFILES := $(shell find . -name '*.go' | grep -v _test.go)
TESTFILES := $(shell find . -name '*_test.go')

all: build

build:
	@echo "Building..."
	@go build $(PKG)

# Run tests
test:
	@echo "Running tests..."
	@go test -v $(PKG)