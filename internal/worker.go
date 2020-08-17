package internal

import (
	"fmt"
	"io"
	"log"
	"time"

	vegeta "github.com/tsenart/vegeta/v12/lib"
)

// TODO: will have to abstract out all the functions into a class
// the class will need to store some internal state, such as a mapping between
// jobids (of ongoing workloads) and channels that can be used to communicate with the
// goroutines (to force stop for example)

func metricsFromVegetaResult(jobID uint64, res *vegeta.Result) *Metrics {
	metrics := &Metrics{
		JobId:    jobID,
		Code:     uint32(res.Code),
		BytesIn:  res.BytesIn,
		BytesOut: res.BytesOut,
		Latency:  uint64(res.Latency),
		Error:    res.Error,
	}
	return metrics
}

func handleMessageStart(stream Worker_CoordinateClient, msgRegister *Start) {
	jobID := msgRegister.Jobid

	fmt.Printf("Starting vegeta attack for job: %v\n", jobID)

	rate := vegeta.Rate{
		Freq: int(msgRegister.GetFrequency()), Per: time.Second,
	}
	duration := time.Duration(msgRegister.GetDuration()) * time.Second

	httpRequest := msgRegister.GetRequest()
	targeter := vegeta.NewStaticTargeter(vegeta.Target{
		Method: httpRequest.GetMethod(),
		URL:    httpRequest.GetUrl(),
	})
	attacker := vegeta.NewAttacker()

	// TODO: potentially consider batching results to reduce network usage
	// TODO: protect all sends with mutex
	for res := range attacker.Attack(targeter, rate, duration, "Test run") {
		stream.Send(&Message{
			Payload: &Message_Metrics{
				Metrics: metricsFromVegetaResult(jobID, res),
			},
		})
		fmt.Printf("latency: %v\n", res.Latency)
	}
	stream.Send(&Message{
		Payload: &Message_Finish{
			Finish: &Finish{JobId: jobID},
		},
	})

	fmt.Printf("Worker finished workload for job %v\n", jobID)
}

func handleMessageStop() {
	// TODO: send stop message to goroutine via channel
}

// Loop is the main event loop
func Loop(stream Worker_CoordinateClient) {
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			stream.CloseSend()
			return
		}
		if err != nil {
			log.Fatalf("Failed to receive a message: %v", err)
		}

		switch x := msg.Payload.(type) {
		case *Message_Start:
			go handleMessageStart(stream, x.Start)
		case *Message_Stop:
			handleMessageStop()
		case nil:
		default:
		}
	}
}
