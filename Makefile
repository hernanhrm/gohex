build:
	@go build -o ./bin/api ./cmd/restapi

run: build
	@./bin/api

