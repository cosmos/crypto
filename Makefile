WALLET_BIN := build/wallet
golangci_version=v1.61.0
golangci_installed_version=$(shell golangci-lint version --format short 2>/dev/null)

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
	@echo "--> Checking golangci-lint installation"
	@if [ "$(golangci_installed_version)" != "$(golangci_version)" ]; then \
		echo "--> Installing golangci-lint $(golangci_version)"; \
		go install github.com/golangci/golangci-lint/cmd/golangci-lint@$(golangci_version); \
	fi

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

.PHONY: lint lint-fix lint-install