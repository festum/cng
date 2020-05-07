PROJECT_NAME := "cng"
PKG := "github.com/festum/$(PROJECT_NAME)"

.PHONY: all gen dep build race clean test coverage lint dl

.DEFAULT_GOAL := build

all: dep test coverage build

lint: ## Lint and security checks
	@golangci-lint run

test: ## Run unittests
	@go test ./... -v -tags=test

race: dep ## Run data race detector
	@go test ./... -race -tags=test

coverage: ## Generate global code coverage report
	@go test ./... -cover -tags=test

dep: ## Get the dependencies
	@go get -v -d ./...

build: dep ## Build the binary file
	@go build -o $(PKG) .

clean: ## Remove previous built binnay and keep latest lean packages
	@go mod tidy
	@rm -f $(PROJECT_NAME)

help: ## Display this help screen
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

install: dep ## Build and install the binary file
	@go install -v ./...

dl: ## Download golang related command line tools for GitLab-CI only!
	@curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b ${GOPATH}/bin latest
