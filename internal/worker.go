package internal

import (
	"fmt"
	"io"
	"log"
	"sync"
	"time"

	vegeta "github.com/tsenart/vegeta/v12/lib"
)

// TODO: will have to abstract out all the functions into a class
// the class will need to store some internal state, such as a mapping between
// jobids (of ongoing workloads) and channels that can be used to communicate with the
// goroutines (to force stop for example)

func metricsFromVegetaResult(jobID string, res *vegeta.Result) *Metrics {
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

func handleMessageStart(stream Worker_CoordinateClient, msgRegister *Start, mutex *sync.Mutex) {
	jobID := msgRegister.JobId

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

	// TODO: potentially consider batching results to reduce network usage as this is definitely a bottleneck
	for res := range attacker.Attack(targeter, rate, duration, "Test run") {
		mutex.Lock()
		stream.Send(&Message{
			Payload: &Message_Metrics{
				Metrics: metricsFromVegetaResult(jobID, res),
			},
		})
		mutex.Unlock()
		fmt.Printf("latency: %v\n", res.Latency)
	}
	mutex.Lock()
	stream.Send(&Message{
		Payload: &Message_Finish{
			Finish: &Finish{JobId: jobID},
		},
	})
	mutex.Unlock()

	fmt.Printf("Worker finished workload for job %v\n", jobID)
}

func handleMessageStop() {
	// TODO: send stop message to goroutine via channel
}

// Loop is the main event loop
func Loop(streamMutex *sync.Mutex, stream Worker_CoordinateClient, wg *sync.WaitGroup, timeMutex *sync.Mutex, lastProcessedTime *time.Time) {
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			return
		}
		if err != nil {
			log.Fatalf("Failed to receive a message: %v", err)
		}

		switch x := msg.Payload.(type) {
		case *Message_Start:
			go func() {
				wg.Add(1)
				handleMessageStart(stream, x.Start, streamMutex)
				timeMutex.Lock()
				*lastProcessedTime = time.Now()
				timeMutex.Unlock()
				wg.Done()
			}()
		case *Message_Stop:
			handleMessageStop()
		case nil:
		default:
		}
	}
}
