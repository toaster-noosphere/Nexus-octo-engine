CMD := server

build:
	go build -o $(CMD) ./cmd/main.go

run:
	make build
	./$(CMD)

lint: 
	go get -u github.com/golangci/golangci-lint/cmd/golangci-lint
	golangci-lint run ./...

.PHONY: build run lint