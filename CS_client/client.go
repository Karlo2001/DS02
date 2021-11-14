package main

import (
	"context"
	"flag"
	"log"

	pb "DS02/CS_Proto"

	"google.golang.org/grpc"
)

var (
	serverAddr = flag.String("server_addr", "localhost:10000", "The server address in the format of host:port")
)

var (
	id   int32
	name string
)

func main() {
	flag.Parse()
	var opts []grpc.DialOption

	opts = append(opts, grpc.WithInsecure())
	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewCriticalServiceClient(conn)
	// Client connected

	// Make grpc call
	GetCriticalAccess(client)
}

func GetCriticalAccess(client pb.CriticalServiceClient) {
	response, err := client.GetCriticalAccess(context.Background(), &pb.ClientInfo{Id: id, Name: name})
	if err != nil {
		// do something with response
		log.Println(response)
	}
}
