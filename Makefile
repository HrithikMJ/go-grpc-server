generate:
	protoc --proto_path=proto proto/*.proto --go_out=. --go-grpc_out=.

build:
	go build -o bin/server server/main.go