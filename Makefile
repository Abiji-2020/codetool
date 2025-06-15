APP_NAME = codetool

.PHONY: build run clean test lint fmt 

build:
	go build -o $(APP_NAME) main.go

run:
	go run main.go 

clean:
	rm -f $(APP_NAME)

test:
	go test ./...

lint:
	golangci-lint run --config .golangci.yml 

fmt:
	go fmt ./...

