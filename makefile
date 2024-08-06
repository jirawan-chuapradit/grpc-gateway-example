gen-proto:
	protoc -I api/ -I third_party \
	--go_out=pkg/example --go_opt=paths=source_relative \
	--go-grpc_out=pkg/example --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=pkg/example --grpc-gateway_opt=paths=source_relative \
	api/example.proto


start-grpc:
	go run cmd/server/main.go

start-gateway:
	go run cmd/gateway/main.go

start-client:
	go run cmd/client/main.go