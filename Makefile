.PHONY: build

VERSION=$(shell git describe --tags --always)
GOOS=$(shell go env GOOS)
GOARCH=$(shell go env GOARCH)

DATE=$(shell date -u +%Y-%m-%dT%H:%M:%SZ)

BINARY_NAME="wasmdash-$(VERSION)-$(GOOS)_$(GOARCH)"

build: ## Build the project, including theme and assets
	go mod tidy && \
	go run scripts/theme-config.go && \
	go generate && \
	tailwindcss -i assets/css/base.css -o static/css/styles.css -m && \
	go build -ldflags="-w -s -X main.buildTime=$(DATE) -X main.commit=$(VERSION)" -o ${BINARY_NAME}

vet: ## Run go vet to check for potential issues
	go vet ./...

run: build ## Run the built binary
	./${BINARY_NAME}

dev: ## Run the development server with live reload
	templ generate --watch --cmd="go generate" &\
	templ generate --watch --cmd="go run ."

gen: ## Generate templ files
	templ generate

theme: ## Generate dashboard theme
	go run scripts/theme-config.go
	@echo "Theme generated successfully"

clean: ## Clean up build artifacts and *_templ-files
	go clean
	# @rm ${BINARY_NAME} static/css/styles.css 2>/dev/null || echo "No build artifacts to clean."
	@rm "*_${GOOS}_${GOARCH}" 2>/dev/null || echo "No other build artifacts to clean."
	@find . -name '*_templ.go' -type f -exec rm -r {} + || echo "No templ files to clean."
	@echo "\e[32mCleaned up build artifacts and templ files.\e[0m"

help: ## Display available commands
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
