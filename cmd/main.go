// Package main is the entrypoint of the Diago worker process
// It performs the following functions:
// 1. Initializes the gRPC stub and connects to the leader, establishing a gRPC stream
// 2. Registers the worker with the leader
// 3. Starts a goroutine to check if a graceful shutdown should be done after a period of inactivity
// 4. Starts the main worker event loop
package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
	"time"

	worker "github.com/t-bfame/diago-worker/pkg"
	pb "github.com/t-bfame/diago-worker/proto-gen/worker"
	"google.golang.org/grpc"
)

func createRegisterMessage(group string, instance string, frequency uint64) *pb.Message {
	return &pb.Message{Payload: &pb.Message_Register{
		Register: &pb.Register{
			Group:     group,
			Instance:  instance,
			Frequency: frequency,
		},
	}}
}

func register(stream pb.Worker_CoordinateClient) {
	cap, _ := strconv.ParseUint(os.Getenv("DIAGO_WORKER_GROUP_INSTANCE_CAPACITY"), 10, 64)

	msgRegister := createRegisterMessage(
		os.Getenv("DIAGO_WORKER_GROUP"),
		os.Getenv("DIAGO_WORKER_GROUP_INSTANCE"),
		cap,
	)
	if err := stream.Send(msgRegister); err != nil {
		log.Fatalf("Failed to send a register message: %v", err)
	}
}

func getAddress(host string, port string) string {
	return host + ":" + port;
}

func main() {
	log.Println("Starting worker program")

	address := getAddress(os.Getenv("DIAGO_LEADER_HOST"), os.Getenv("DIAGO_LEADER_PORT"))

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

	log.Println("Connected to leader")

	wg := &sync.WaitGroup{}
	lastProcessedTime := time.Now()
	timeMutex := &sync.Mutex{}
	streamMutex := &sync.Mutex{}
	gracePeriod, err := strconv.ParseFloat(os.Getenv("ALLOWED_INACTIVITY_PERIOD_SECONDS"), 32)

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
