package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
	"time"

	worker "github.com/t-bfame/diago-worker/internal"
	pb "github.com/t-bfame/diago-worker/proto-gen/worker"
	"google.golang.org/grpc"
)

func register(stream pb.Worker_CoordinateClient) {

	msgRegister := &pb.Message{Payload: &pb.Message_Register{
		Register: &pb.Register{
			Group:    os.Getenv("DIAGO_WORKER_GROUP"),
			Instance: os.Getenv("DIAGO_WORKER_GROUP_INSTANCE"),
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

	client := pb.NewWorkerClient(conn)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Contact server to establish grpc stream
	stream, err := client.Coordinate(ctx)
	if err != nil {
		log.Fatalf("Failed to setup gRPC stream: %v", err)
	}

	wg := &sync.WaitGroup{}
	lastProcessedTime := time.Now()
	timeMutex := &sync.Mutex{}
	streamMutex := &sync.Mutex{}
	gracePeriod, err := strconv.ParseFloat(os.Getenv("TERMINATION_GRACE_PERIOD_SECONDS"), 32)

	// TODO: do i have to do graceful shutdown or can i just kill the program?
	go func() {
		for {
			wg.Wait()

			timeMutex.Lock()
			diff := time.Now().Sub(lastProcessedTime)
			timeMutex.Unlock()
			if diff.Seconds() > gracePeriod {
				fmt.Printf("It's been more than %v seconds\n", gracePeriod)
				streamMutex.Lock()
				// TODO: this is super flaky, sometimes it doesn't trigger an io.EOF on the server side
				// maybe the server actually needs to be currently blocked on Recv() in order to get the io.eof?
				stream.CloseSend()
				streamMutex.Unlock()
				cancel()
				return
			}
			time.Sleep(time.Second)
		}
	}()

	register(stream)
	w := worker.NewWorker()
	w.Loop(streamMutex, stream, wg, timeMutex, &lastProcessedTime)
}
