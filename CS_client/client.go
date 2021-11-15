package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"

	pb "DS02/CS_Proto"

	"google.golang.org/grpc"
)

var (
	serverAddr = flag.String("server_addr", "localhost:10000", "The server address in the format of host:port")
)

var (
	id       int32
	name     = "1"
	actionid int
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
	if len(os.Args) < 2 {
		fmt.Println("Please specify the id in the command line arguments")
		os.Exit(1)
	}
	tid, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("The specified id is not of type int. Please try again")
		os.Exit(1)
	}
	id = int32(tid)
	for {
		fmt.Println("---------------------------------------------")
		fmt.Println("Type 1 to request access to the critical section")
		fmt.Scanln(&actionid)
		if actionid == 1 {
			GetCriticalAccess(client)
		}
	}
}

func GetCriticalAccess(client pb.CriticalServiceClient) {
	response, err := client.GetCriticalAccess(context.Background(), &pb.ClientInfo{Id: id, Name: name})
	if err != nil {
		// do something with response
		log.Println(err)
	} else {
		fmt.Println(response.Message)
	}
}
