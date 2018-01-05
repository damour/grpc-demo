package main

import (
	"log"
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pbV1 "github.com/damour/grpc-demo/proto/v1"
	"google.golang.org/grpc/reflection"
)

const (
	portV1 = ":50050"
)

// server is used to implement helloworld.GreeterServer.
type serverV1 struct{}

// SayHello implements helloworld.GreeterServer
func (s *serverV1) SayHello(ctx context.Context, in *pbV1.HelloRequest) (*pbV1.HelloReply, error) {
	return &pbV1.HelloReply{Message: "Hello " + in.Name}, nil
}

func main() {
	lis, err := net.Listen("tcp", portV1)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pbV1.RegisterGreeterServer(s, &serverV1{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
