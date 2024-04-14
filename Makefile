build:
	@go build -o bin/GOAPI cmd/main.go

test:
	@go test -v ./...

run: build
	@./bin/GOAPI
