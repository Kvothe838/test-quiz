.PHONY: run-backend
run-backend: ## It runs the main app for backend
	go run cmd/app/main.go --config local-env/config.yaml

.PHONY: run-client
run-client: ## It runs the main app for client
	go run main.go

.PHONY: get-quiz
get-quiz:
	go run main.go quiz

.PHONY: lint
lint: ## It starts the linter report
	@golangci-lint run --color always ./...

.PHONY: test
test: ## It runs the tests
	go test ./...