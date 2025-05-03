.PHONY: setup compile-run run test

setup:
	docker-compose up -d mongo

compile-run: setup
	go build ./cmd/api/main.go && ./main

run: setup
	go run ./cmd/api/main.go

test:
	go test ./test/... -v