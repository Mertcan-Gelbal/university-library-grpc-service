.PHONY: proto build run-server run-client test clean

# Generate protobuf code
proto:
	protoc --go_out=. --go-grpc_out=. university.proto

# Install dependencies
deps:
	go mod tidy

# Build all binaries
build: proto
	go build -o bin/server src/server/main.go
	go build -o bin/client src/client/main.go

# Run server
run-server: proto
	go run src/server/main.go

# Run client
run-client: proto
	go run src/client/main.go

# Test with grpcurl
test-grpcurl:
	@echo "Testing BookService..."
	grpcurl -plaintext localhost:50051 university.library.BookService/ListBooks
	@echo "\nTesting StudentService..."
	grpcurl -plaintext localhost:50051 university.library.StudentService/ListStudents
	@echo "\nTesting LoanService..."
	grpcurl -plaintext localhost:50051 university.library.LoanService/ListLoans

# List all services
list-services:
	grpcurl -plaintext localhost:50051 list

# Clean generated files
clean:
	rm -rf pb/
	rm -rf bin/

# Setup everything
setup: deps proto

# Help
help:
	@echo "Available commands:"
	@echo "  proto         - Generate protobuf code"
	@echo "  deps          - Install dependencies"
	@echo "  build         - Build server and client binaries"
	@echo "  run-server    - Run gRPC server"
	@echo "  run-client    - Run gRPC client"
	@echo "  test-grpcurl  - Test services with grpcurl"
	@echo "  list-services - List available services"
	@echo "  clean         - Clean generated files"
	@echo "  setup         - Setup project (deps + proto)"
	@echo "  help          - Show this help" 