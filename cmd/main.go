package main

import (
	"context"
	"log"
	"os"

	worker "github.com/t-bfame/diago-worker/internal"
	"google.golang.org/grpc"
)

func register(stream worker.Worker_CoordinateClient) {

	msgRegister := &worker.Message{Payload: &worker.Message_Register{
		Register: &worker.Register{
			Group:     os.Getenv("DIAGO_WORKER_GROUP"),
			Instance:  os.Getenv("DIAGO_WORKER_GROUP_INSTANCE"),
			Frequency: 100,
		},
	}}
	if err := stream.Send(msgRegister); err != nil {
		log.Fatalf("Failed to send a register message: %v", err)
	}
}

func main() {
	log.Println("Starting worker program")

	address := os.Getenv("DIAGO_LEADER_HOST") + ":" + os.Getenv("DIAGO_LEADER_PORT")

	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Did not connect to server: %v", err)
	}
	defer conn.Close()

	client := worker.NewWorkerClient(conn)

	// TODO: if a context with a timeout is created, the program won't work
	ctx := context.Background()

	// Contact server to establish grpc stream
	stream, err := client.Coordinate(ctx)
	if err != nil {
		log.Fatalf("Failed to setup gRPC stream: %v", err)
	}

	register(stream)
	worker.Loop(stream)

	stream.CloseSend()
}
