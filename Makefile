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

# Deploy main contracts from using config/deploy.json as input
.PHONY: deploy
deploy:
	forge script script/Deploy.s.sol \
		--rpc-url $(RPC_URL) \
		--broadcast \
		--private-key $(PRIVATE_KEY) \
		$(if $(VERIFY),--verify --verifier etherscan --etherscan-api-key $(ETHERSCAN_API_KEY)) \
		$(if $(RESUME),--resume) \
		--slow \
		-vvvv \
		--sig "run(string)" -- $(ENVIRONMENT)

# Create AVS instance
.PHONY: create-app
create-app:
	forge script script/CreateApp.s.sol \
		--rpc-url $(RPC_URL) \
		--broadcast \
		--sig "run(string)" $(ENVIRONMENT) \
		--private-key $(PRIVATE_KEY) \
		--slow \
		-vvvv

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
