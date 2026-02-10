run:
	go run ./cmd/api/main.go

# this creates the binary executable file at bin/app made by building ./cmd/api
build:
	go build -o .bin/app .cmd/api

lint:
	golangci-lint run
