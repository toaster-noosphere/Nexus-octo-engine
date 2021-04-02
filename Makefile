CMD := engine

build:
	go build -o $(CMD) ./cmd/main.go

run:
	make build
	./$(CMD)

lint:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint
	golangci-lint run ./...

gen:
	cd pkg/model && easyjson -all models.go

.PHONY: build run lint gen
