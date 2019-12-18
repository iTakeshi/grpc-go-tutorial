package main

import (
    "context"
    "log"
    "net"

    "google.golang.org/grpc"
    pb "github.com/iTakeshi/grpc-go-tutorial/helloworld"
)

const (
    port = ":50001"
)

type server struct {
}

func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
    log.Printf("Received: %v", req.GetName())
    return &pb.HelloReply{Message: "Hello " + req.GetName()}, nil
}

func main() {
    lis, err := net.Listen("tcp", port)
    if err != nil {
        log.Fatalf("failed to listen :%v", err)
    }
    s := grpc.NewServer()
    pb.RegisterGreeterServer(s, &server{})
    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}
