WALLET_BIN := build/wallet

.PHONY: demo run clean

.DEFAULT_GOAL := build

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