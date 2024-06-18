.PHONY: run-backend
run-backend: ## It runs the main app for backend
	go run cmd/app/main.go --config local-env/config.yaml

.PHONY: run-client
run-client: ## It runs the main app for client
	go run main.go

.PHONY: get-quiz
get-quiz: ## It runs quiz command on client
	go run main.go get-quiz

.PHONY: select-choice
select-choice:
	go run main.go select-choice $(filter-out $@,$(MAKECMDGOALS))

.PHONY: submit-quiz
submit-quiz:
	go run main.go submit-quiz

.PHONY: lint
lint: ## It starts the linter report
	@golangci-lint run --color always ./...

.PHONY: test
test: ## It runs the tests
	go test ./...


%:
	@: