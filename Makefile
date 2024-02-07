hello:
	@echo "Hello, World" 

format:
	@echo "Formatting code..."
	go fmt ./...
	@echo "Done!"

start_gateway:
	@echo "Starting HTTP Gateway server..."
	go run cmd/gateway/main.go
	@echo "Done!"

start_grpc:
	@echo "Starting gRPC server..."
	go run cmd/calculator/main.go
	@echo "Done!"

# Proto Commands
## gen_proto: Generate protobuf files for a given proto file
## Usage Example: make gen_proto file=proto/instructor.proto
## Reason for require_unimplemented_servers=false flag can be found here[IMPORTANT]:
## 1. https://stackoverflow.com/questions/65079032/grpc-with-mustembedunimplemented-method
## 2. https://github.com/grpc/grpc-go/issues/3669
gen_proto:
	@echo "Generating protobuf files..."
	protoc --go_out=. --go-grpc_out=require_unimplemented_servers=false:. $(file)
	@echo "Done!"
