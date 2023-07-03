GOCMD=go
CODE_COVERAGE=code-coverage

build:
	go build -o bin/main cmd/api/main.go

run:
	./bin/main

swag: 
	swag init -g /pkg/api/server.go -o ./cmd/api/docs

test:
	$(GOCMD) test ./... -v -cover

test-coverage: ## Run tests and generate coverage file
	$(GOCMD) test ./... -coverprofile=$(CODE_COVERAGE).out
	$(GOCMD) tool cover -html=$(CODE_COVERAGE).out

mockgen:
	mockgen -source=pkg/repository/interfaces/employeeinterfaces.go -destination=pkg/mock/employeeRepoMock/employeeRepoMock.go -package=mock

