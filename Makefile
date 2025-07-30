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
