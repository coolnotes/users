run:
	@go run cmd/main.go

build-protos:
	@rm -rf ./src/rpc/*
	@protoc -I=protos --go_out=. --go-grpc_out=. protos/*