.PHONY: run docs dev test db-local-start db-local-stop setup

build:
	go build -o veritas ./cmd/api

run: build
	./veritas

docs:
	$(shell go env GOPATH)/bin/swag init -g cmd/api/main.go

dev:
	go run ./cmd/api/main.go

test:
	go test -v ./...

db-local-start:
	docker compose up -d veritas-db

db-local-stop:
	docker compose stop veritas-db

setup:
	@echo "Installing Go tools..."
	go install github.com/swaggo/swag/cmd/swag@latest
	@echo "Setup complete. Make sure $(go env GOPATH)/bin is in your PATH"
