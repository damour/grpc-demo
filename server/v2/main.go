package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	pbV2 "github.com/damour/grpc-demo/proto/v2"
	"google.golang.org/grpc/reflection"
)

const (
	portV2 = ":50051"
)

// server is used to implement helloworld.GreeterServer.
type serverV2 struct{}

// SayHello implements helloworld.GreeterServer
func (s *serverV2) SayHello(in *pbV2.HelloRequest, stream pbV2.Greeter_SayHelloServer) (error) {
	sendErr := stream.Send(&pbV2.HelloReply{Response: &pbV2.HelloReply_Description{Description: "Example description"}})

	if sendErr != nil {
		return sendErr
	}

	sendErr = stream.Send(&pbV2.HelloReply{Response: &pbV2.HelloReply_Message{Message: "Hello " + in.Name}})

	if sendErr != nil {
		return sendErr
	}

	return nil
}

func main() {
	lis, err := net.Listen("tcp", portV2)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pbV2.RegisterGreeterServer(s, &serverV2{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
