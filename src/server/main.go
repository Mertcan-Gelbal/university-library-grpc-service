package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/university-library/grpc-service/src/server/services"
	pb "github.com/university-library/grpc-service/pb"
)

const (
	port = ":50051"
)

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()

	// Register services with constructors
	pb.RegisterBookServiceServer(s, services.NewBookServer())
	pb.RegisterStudentServiceServer(s, services.NewStudentServer())
	pb.RegisterLoanServiceServer(s, services.NewLoanServer())

	// Register reflection service for grpcurl
	reflection.Register(s)

	log.Printf("gRPC server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
} 