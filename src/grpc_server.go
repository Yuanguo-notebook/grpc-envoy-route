package main

import (
    "context"
    "log"
    "net"
    "flag"
    "fmt"
    "google.golang.org/grpc"
    pb "echo/github.com/Yuanguo-notebook/echo"
)

type server struct {
    pb.UnimplementedEchoServiceServer
    myname string
}

func (s *server) Echo(ctx context.Context, req *pb.EchoRequest) (*pb.EchoResponse, error) {
    log.Printf("[%v]Got request: %s", s.myname, req.Message)
    msg := fmt.Sprintf("[%s]what's up", s.myname)
	return &pb.EchoResponse{Message: msg}, nil
}

func main() {
    name:= flag.String("name", "", "name of cur server")
    
    flag.Parse()
    log.Printf("name:  %v\n", *name)

    
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }

    grpcServer := grpc.NewServer()
    pb.RegisterEchoServiceServer(grpcServer, &server{myname: *name})

    log.Printf("[%v]Server is running on port 50051...\n", *name)
    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}