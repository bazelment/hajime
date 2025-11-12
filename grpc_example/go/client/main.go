package main

import (
	"context"
	"flag"
	"log"
	"time"

	pb "github.com/bazelment/hajime/grpc_example/proto/greeter"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = flag.String("addr", "localhost:50052", "the address to connect to")
)

func main() {
	flag.Parse()

	log.Printf("Go client connecting to Rust server at %s", *addr)

	conn, err := grpc.NewClient(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: "Go Client"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Response from Rust server: %s", r.GetMessage())
}
