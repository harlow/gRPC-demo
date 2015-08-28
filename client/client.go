package main

import (
	pb "github.com/harlow/grpc-demo/proto/demo"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata"
)

func main() {
	// gRPC client and connection
	conn, err := grpc.Dial("127.0.0.1:10000")
	if err != nil {
		grpclog.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewDemoClient(conn)

	// context and metadata
	md := metadata.Pairs("requestID", "SOME_VALUE_HERE")
	ctx := context.Background()
	ctx = metadata.NewContext(ctx, md)

	// create greeting and send request
	greeting := pb.Greeting("Yoo dawg!")
	reply, err := client.SayHello(ctx, greeting)
	if err != nil {
		grpclog.Fatalf("%v.SayHello(_) = _, %v", client, err)
	}

	grpclog.Printf("Response: %v", reply)
}
