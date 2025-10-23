GOOS ?=
GOARCH ?=

GOCMD := go
GOCLEAN := $(GOCMD) clean
GOTEST := $(GOCMD) test
GOMOD := $(GOCMD) mod
GOFMT := $(GOCMD) fmt
GOVET := $(GOCMD) vet

PKGS := ./...
GOFLAGS ?=
TESTFLAGS ?= -race
LDFLAGS ?= -s -w

GOLANGCI_LINT_VER := latest

RED := \033[0;31m
GREEN := \033[0;32m
YELLOW := \033[0;33m
BLUE := \033[0;34m
NC := \033[0m

.DEFAULT_GOAL := help

.PHONY: help clean fmt vet lint lint-fix test check-tokens deps deps-update check version install-tools dev-setup

help: ## Show available commands
	@echo "$(BLUE)Scaleway Go SDK - Available Commands$(NC)\n"
	@grep -E '^[a-zA-Z0-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort \
	| awk 'BEGIN {FS = ":.*?## "}; {printf "  $(GREEN)%-22s$(NC) %s\n", $$1, $$2}'
	@echo ""

clean: ## Clean build artifacts
	@echo "$(BLUE)Cleaning...$(NC)"
	$(GOCLEAN)

fmt: ## Format Go code
	@echo "$(BLUE)Formatting...$(NC)"
	$(GOFMT) $(PKGS)

vet: ## Run go vet
	@echo "$(BLUE)Vetting...$(NC)"
	$(GOVET) $(PKGS)

lint: ## Run linters
	@echo "$(BLUE)Linting...$(NC)"
	@./scripts/lint.sh

lint-fix: ## Run linters with auto-fix
	@echo "$(BLUE)Linting with fix...$(NC)"
	@./scripts/lint.sh --write

test: ## Run tests with race detector
	@echo "$(BLUE)Testing...$(NC)"
	$(GOTEST) $(GOFLAGS) $(TESTFLAGS) $(PKGS)

check-tokens: ## Check for exposed tokens
	@echo "$(BLUE)Checking for exposed tokens...$(NC)"
	@./scripts/check_for_tokens.sh

deps: ## Download dependencies
	@echo "$(BLUE)Downloading deps...$(NC)"
	$(GOMOD) download

deps-update: ## Tidy and update dependencies
	@echo "$(BLUE)Updating dependencies...$(NC)"
	$(GOMOD) tidy
	$(GOMOD) download

check: deps check-tokens fmt vet lint test ## Run full local checks
	@echo "$(GREEN)All checks passed!$(NC)"

version: ## Show tool versions
	@echo "$(BLUE)Version Information:$(NC)"
	@echo "Go version: $$($(GOCMD) version)"
	@echo "Go modules: $$($(GOCMD) env GOMOD)"
	@if command -v golangci-lint >/dev/null 2>&1; then \
		echo "golangci-lint: $$(golangci-lint version)"; \
	fi

install-tools: ## Install dev tools
	@echo "$(BLUE)Installing tools...$(NC)"
	@echo "$(YELLOW)Installing golangci-lint $(GOLANGCI_LINT_VER)$(NC)"
	@$(GOCMD) install github.com/golangci/golangci-lint/cmd/golangci-lint@$(GOLANGCI_LINT_VER)

dev-setup: install-tools deps ## Setup dev environment
	@echo "$(GREEN)Development environment setup complete!$(NC)"
	@echo "$(YELLOW)Run 'make help' to see available commands$(NC)"
