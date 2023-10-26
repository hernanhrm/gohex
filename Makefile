build:
	go build -o bin/api cmd/restapi/main.go

run: build
	go run cmd/restapi/main.go

