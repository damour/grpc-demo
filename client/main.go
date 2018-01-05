package main

import (
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pbV1 "github.com/damour/grpc-demo/proto/v1"
	pbV2 "github.com/damour/grpc-demo/proto/v2"
	"io"
)

const (
	addressV1     = "localhost:50050"
	addressV2     = "localhost:50051"
	defaultName = "world"
)

func main() {
	log.Printf("Process client v1 & server v1:")
	processClientV1ServerV1()

	log.Printf("Process client v1 & server v2:")
	processClientV1ServerV2()

	log.Printf("Process client v2 & server v1:")
	processClientV2ServerV1()

	log.Printf("Process client v2 & server v2:")
	processClientV2ServerV2()
}

func processClientV1ServerV1() {
	conn, err := grpc.Dial(addressV1, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pbV1.NewGreeterClient(conn)

	r, err := c.SayHello(context.Background(), &pbV1.HelloRequest{Name: defaultName})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.Message)
}

func processClientV1ServerV2() {
	conn, err := grpc.Dial(addressV2, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pbV1.NewGreeterClient(conn)

	r, err := c.SayHello(context.Background(), &pbV1.HelloRequest{Name: defaultName})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.Message)
}

func processClientV2ServerV1() {
	conn, err := grpc.Dial(addressV1, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pbV2.NewGreeterClient(conn)

	cl, err := c.SayHello(context.Background(), &pbV2.HelloRequest{Name: defaultName})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	message, err := processV2Response(cl)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	log.Printf("Greeting: %s", message)
}

func processClientV2ServerV2() {
	conn, err := grpc.Dial(addressV2, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pbV2.NewGreeterClient(conn)

	cl, err := c.SayHello(context.Background(), &pbV2.HelloRequest{Name: defaultName})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	message, err := processV2Response(cl)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	log.Printf("Greeting: %s", message)
}

func processV2Response(cl pbV2.Greeter_SayHelloClient) (string, error)  {
	var message string

	for {
		result, err := cl.Recv()

		if err == io.EOF {
			return message, nil
		}

		if err != nil {
			return "", err
		}

		switch x := result.Response.(type) {
		case *pbV2.HelloReply_Message:
			message = x.Message
			break
		case *pbV2.HelloReply_Description:
			log.Printf("Description: %s", x.Description)
			break
		}
	}
}