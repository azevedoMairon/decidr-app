.PHONY: setup compile-run run test

setup:
	docker-compose up -d mongo

compile-run: setup
	go build ./cmd/api/main.go && ./main

run: setup
	@export $$(cat .env | xargs) && go run ./cmd/api/main.go

test:
	go test ./test/... -v