package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	cheese "github.com/t-bfame/diago-worker/cheese"
	pb "github.com/t-bfame/diago-worker/internal"
	"google.golang.org/grpc"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func main() {
	fmt.Printf("cheesing %v\n", cheese.Cheesing())
	/*
		rate := vegeta.Rate{Freq: 5, Per: time.Second}
		duration := 3 * time.Second
		targeter := vegeta.NewStaticTargeter(vegeta.Target{
			Method: "GET",
			URL:    "http://localhost:3000",
		})
		attacker := vegeta.NewAttacker()

		var metrics vegeta.Metrics
		for res := range attacker.Attack(targeter, rate, duration, "Test run") {
			metrics.Add(res)
		}
		metrics.Close()

		fmt.Printf("99th percentile: %s\n", metrics.Latencies.P99)
	*/

	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	// Contact the server and print out its response.
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())

	r, err = c.SayHelloAgain(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())
}
