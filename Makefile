build:
	go build -o ./bin/service ./cmd/service/main.go

fmt:
	gofumpt -w .

tidy:
	go mod tidy

lint: build fmt tidy
	golangci-lint run ./...

run:
	go run ./cmd/service/main.go

up:
	docker compose up -d

down:
	docker compose down

test: up
	go test -coverpkg=./... -v ./...