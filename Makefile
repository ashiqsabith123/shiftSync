hello:
	echo "hello"

build:
	go build -o bin/main cmd/api/main.go

run:
	go run cmd/api/main.go

swag: 
	swag init -g /pkg/api/server.go -o ./cmd/api/docs

mockgen:
	mockgen -source=pkg/repository/interfaces/employeeinterfaces.go -destination=pkg/mock/employeeRepoMock/employeeRepoMock.go -package=mock

