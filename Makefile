generate:
	protoc --proto_path=proto proto/*.proto --go_out=. --go-grpc_out=.

start:
	go run main.go && grpcui --plaintext localhost:8080