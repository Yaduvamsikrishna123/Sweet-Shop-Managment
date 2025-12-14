.PHONY: run build test docker-up docker-down

run:
	go run cmd/server/main.go

build:
	go build -o bin/server cmd/server/main.go

test:
	go test ./...

docker-up:
	docker-compose up -d

docker-down:
	docker-compose down
