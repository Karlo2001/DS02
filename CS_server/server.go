package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"strconv"
	"time"

	pb "DS02/CS_Proto"

	"google.golang.org/grpc"
)

var (
	port  = flag.Int("port", 10000, "The server port")
	queue = make([]client, 0)
	cs    = 0
)

type client struct {
	name string
	id   int32
}

type criticalServiceServer struct {
	pb.UnimplementedCriticalServiceServer
}

func (s *criticalServiceServer) GetCriticalAccess(ctx context.Context, in *pb.ClientInfo) (*pb.UserResponse, error) {
	queue = append(queue, client{id: in.Id, name: in.Name})
	for {
		if queue[0].id == in.Id {
			cs++
			//Simulate some heavy function
			time.Sleep(5 * time.Second)
			queue = queue[1:]
			return &pb.UserResponse{Message: "The critical section have now been accesed " + strconv.Itoa(cs) + " times"}, nil
		}
	}
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
