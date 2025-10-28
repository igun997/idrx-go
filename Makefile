
# Clean up go.mod and go.sum files
tidy: ## Clean up go.mod and go.sum files
	go mod tidy

# Install dependencies
deps: ## Install dependencies
	go mod download
	go mod tidy

# Run tests
test: ## Run all tests
	@echo "Running tests..."
	go test -v ./...

# Run tests with coverage report
test-coverage: ## Run tests with coverage report
	@echo "Running tests with coverage..."
	go test -v -race -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

# Run tests with race condition detection
test-race: ## Run tests with race condition detection
	@echo "Running tests with race detection..."
	go test -v -race ./...

# Run benchmark tests
test-bench: ## Run benchmark tests with memory profiling
	@echo "Running benchmark tests..."
	go test -v -bench=. -benchmem ./...

# Format code using gofumpt and goimports-reviser
fmt: tidy ## Format Go code using gofumpt and goimports-reviser
	@echo "Formatting code..."
	@find . -type f -name '*.go' \
		-not -path './vendor/*' \
		-not -path './scripts/*' \
		-not -path './tools/*' \
		-not -path './bin/*' \
		| xargs gofumpt -w
	@find . -type f -name '*.go' \
		-not -path './vendor/*' \
		-not -path './scripts/*' \
		-not -path './tools/*' \
		-not -path './bin/*' \
		| xargs goimports-reviser -rm-unused -set-alias -format -output file

# Lint code
lint: ## Run golangci-lint on the codebase
	@echo "Linting code..."
	@golangci-lint run -v

# Run security scan using gosec
security: ## Run security scan using gosec
	@echo "Running security scan..."
	gosec ./...

# Run pre-commit hooks on all files
pre-commit: ## Run pre-commit hooks on all files
	@echo "Running pre-commit on all files..."
	pre-commit run --all-files

# Run complete code quality checks
quality: fmt lint test-coverage security ## Run complete code quality checks
	@echo "Quality check completed!"

# Generate smart contract Go bindings from ABIs
generate-contracts: ## Generate Go bindings for IDRX smart contracts
	@echo "Generating smart contract bindings..."
	@mkdir -p contracts
	abigen --abi abis/idrx.json --pkg contracts --type IDRX --out contracts/idrx.go --alias _totalSupply=InternalTotalSupply
	abigen --abi abis/proxy.json --pkg contracts --type ERC1967Proxy --out contracts/proxy.go
	@echo "Contract bindings generated successfully!"

.PHONY: dev-setup
dev-setup: tools setup-hooks ## Set up development environment
	@echo "Development environment setup completed!"

tools: ## Install all development tools via mise
	@echo "Installing tools via mise..."
	mise install

# Install pre-commit hooks
setup-hooks: ## Install pre-commit hooks
	@echo "Setting up pre-commit hooks..."
	pre-commit install --install-hooks
