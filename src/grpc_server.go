package main

import (
    "context"
    "log"
    "net"
    "flag"
    "google.golang.org/grpc"
    pb "echo/github.com/Yuanguo-notebook/echo"
)

var myname string
type server struct {
    pb.UnimplementedEchoServiceServer
}

func (s *server) Echo(ctx context.Context, req *pb.EchoRequest) (*pb.EchoResponse, error) {
    log.Printf("[%v]Got request: %s", myname, req.Message)
	return &pb.EchoResponse{Message: "what's up"}, nil
}

func main() {
    name:= flag.String("name", "", "name of cur server")
    
    flag.Parse()
    myname := *name
    log.Printf("myname:  %v\n", myname)

    
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }

    grpcServer := grpc.NewServer()
    pb.RegisterEchoServiceServer(grpcServer, &server{})

    log.Printf("[%v]Server is running on port 50051...\n", myname)
    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}