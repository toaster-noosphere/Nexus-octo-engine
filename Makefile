CMD := server

build:
		go build -o $(CMD) ./cmd/main.go

run:
		make build
			./$(CMD)

lint: 
		golangci-lint run ./...

gen:
	cd pkg/model && easyjson -all models.go

.PHONY: build run lint gen
