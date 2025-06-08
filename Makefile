.PHONY: build

VERSION=$(shell git describe --tags --always)
GOOS=$(shell go env GOOS)
GOARCH=$(shell go env GOARCH)

DATE=$(shell date -u +%Y-%m-%dT%H:%M:%SZ)

BINARY_NAME="wasmdash-$(VERSION)-$(GOOS)_$(GOARCH)"

build: ## Build the project, including Tailwind CSS and Go binary
	go mod tidy && \
	#go generate #go getailwindcss -m -o static/css/styles.css && \
	go generate && \
	go build -ldflags="-w -s -X main.buildTime=$(DATE)" -X main.commit=$(COMMIT) -o ${BINARY_NAME}

vet: ## Run go vet to check for potential issues
	go vet ./...

run: build ## Run the built binary
	./${BINARY_NAME}

dev: ## Run the development server with live reload
	templ generate --watch --cmd="go generate" &\
	templ generate --watch --cmd="go run ."

gen: ## Generate templ files
	templ generate
	# tailwindcss -o static/css/styles.css --minify

clean: ## Clean up build artifacts and *_templ-files
	go clean
	@rm ${BINARY_NAME} static/css/styles.css 2>/dev/null || echo "No build artifacts to clean."
	@rm "*_${GOOS}_${GOARCH}" 2>/dev/null || echo "No other build artifacts to clean."
	@find . -name '*_templ.go' -type d -exec rm -r {} + || echo "No *_templ-files to clean."
	@echo "\e[32mCleaned up build artifacts and *_templ-files.\e[0m"

help: ## Display available commands
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
