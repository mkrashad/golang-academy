APP_NAME := golang-academy
MAIN := ./cmd/server/main.go

.PHONY: all build run clean test

all: build

build:
	go build -o bin/$(APP_NAME) $(MAIN)

run:
	go run $(MAIN)

clean:
	rm -rf bin
	rm -rf ./certs

test:
	go test ./internal/db/...

