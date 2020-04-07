.PHONY: build
build:
	go build -o bin/server cmd/**/*.go

.PHONY: start
start:
	go run main.go