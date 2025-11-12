# Configuration variables

# Build the project
.PHONY: build
build:
	forge clean
	forge build

# Test the project
.PHONY: test
test:
	forge test -vvv

.PHONY: bindings
bindings:
	./bin/compile-bindings.sh

# Helper message
.PHONY: help
help:
	@echo "Available commands:"
	@echo "  make build - Build the project"
	@echo "  make test - Test the project"
	@echo "  make deploy - Deploy main contracts using environment variables as input"
	@echo "  make create-avs - Create AVS instance using CreateAVS.s.sol"
	@echo ""
	@echo "Note: Make sure to set RPC_URL and PRIVATE_KEY in your environment or .env file" 
