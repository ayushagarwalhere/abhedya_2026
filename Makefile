run:
	go run ./cmd/api/main.go

migrate:
	go run ./cmd/migrate/main.go

build:
	go build -o .bin/app .cmd/api

lint:
	golangci-lint run
