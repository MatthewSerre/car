auth:
	go run internal/server/authentication/main.go

cli:
	go run cmd/bluelink-cli/main.go

information:
	go run internal/server/information/main.go

protos:
	buf generate protos