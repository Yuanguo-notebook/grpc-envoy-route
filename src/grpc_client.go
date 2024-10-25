package main

import (
    "context"
    "log"
    // "time"
	"flag"
    "google.golang.org/grpc"
    "google.golang.org/grpc/credentials/insecure"
    "google.golang.org/grpc/metadata"
    pb "echo/github.com/Yuanguo-notebook/echo"
)

func main() {
    hostname := flag.String("hostname", "127.0.0.1:50051", "Specify the server hostname")
    originalHost:= flag.String("originalHost", "server-1", "Specify the originalHost name for routing")
    
    flag.Parse()
    log.Printf("hostname: %v, originalHost: %v\n", *hostname, *originalHost)

    conn, err := grpc.NewClient(*hostname, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewEchoServiceClient(conn)

    md := metadata.Pairs(
        "originalHost", *originalHost,
        "key2", "value2",
    )
    ctx := metadata.NewOutgoingContext(context.Background(), md)

    // ctx, cancel := context.WithTimeout(context.Background(), time.Second)
    // defer cancel()

    response, err := c.Echo(ctx, &pb.EchoRequest{Message: "hi server"})
    if err != nil {
        log.Fatalf("could not greet: %v", err)
    }

    log.Printf("Echo response: %s", response.Message)
}

