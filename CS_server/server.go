package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "DS02/CS_Proto"

	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 10000, "The server port")
)

type criticalServiceServer struct {
	pb.UnimplementedCriticalServiceServer
}

func (s *criticalServiceServer) GetCriticalAccess(ctx context.Context, in *pb.ClientInfo) (*pb.UserResponse, error) {
	// do stuff
	return &pb.UserResponse{Message: "message"}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)
	pb.RegisterCriticalServiceServer(grpcServer, &criticalServiceServer{})
	grpcServer.Serve(lis)
}
