WALLET_BIN := build/wallet
GOLANGCI_VERSION := v1.60.3

.PHONY: demo run clean

.DEFAULT_GOAL := build-wallet

# Build the wallet command
build-wallet:
	go build -o $(WALLET_BIN) ./cmd

# Run the demo
demo:
	cd demo; \
	go run main.go

# Clean built binaries
clean:
	rm -f $(WALLET_BIN)


# Install golangci-lint
lint-install:
	@echo "--> Installing golangci-lint $(GOLANGCI_VERSION)"
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@$(GOLANGCI_VERSION)

# Run golangci-lint
lint:
	@echo "--> Running linter"
	$(MAKE) lint-install
	@golangci-lint run --timeout=15m

# Run golangci-lint and fix
lint-fix:
	@echo "--> Running linter with fix"
	$(MAKE) lint-install
	@golangci-lint run --fix