package internal

import (
	"fmt"
	"io"
	"log"
	"sync"
	"time"

	pytypes "github.com/golang/protobuf/ptypes"
	pb "github.com/t-bfame/diago-worker/proto-gen/worker"
	vegeta "github.com/tsenart/vegeta/v12/lib"
)

// Worker that manages ongoing attacks
type Worker struct {
	attacks map[string]chan struct{}
}

// NewWorker instantiates and returns a Worker
func NewWorker() *Worker {
	w := &Worker{
		attacks: make(map[string]chan struct{}),
	}
	return w
}

func (w *Worker) metricsFromVegetaResult(jobID string, res *vegeta.Result) *pb.Metrics {
	timestampProto, err := pytypes.TimestampProto(res.Timestamp)
	if err != nil {
		log.Fatal(err)
	}
	metrics := &pb.Metrics{
		JobId:     jobID,
		Code:      uint32(res.Code),
		BytesIn:   res.BytesIn,
		BytesOut:  res.BytesOut,
		Latency:   int64(res.Latency),
		Error:     res.Error,
		Timestamp: timestampProto,
	}
	return metrics
}

func (w *Worker) handleMessageStart(stream pb.Worker_CoordinateClient, msgRegister *pb.Start, mutex *sync.Mutex) {
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
Loop:
	for res := range attacker.Attack(targeter, rate, duration, "Test run") {
		select {
		case <-w.attacks[jobID]:
			fmt.Printf("Prematurely ending job with id %v\n", jobID)
			close(w.attacks[jobID])
			delete(w.attacks, jobID)
			break Loop
		default:
			mutex.Lock()
			stream.Send(&pb.Message{
				Payload: &pb.Message_Metrics{
					Metrics: w.metricsFromVegetaResult(jobID, res),
				},
			})
			mutex.Unlock()
			fmt.Printf("latency: %v\n", res.Latency)
		}
	}
	mutex.Lock()
	stream.Send(&pb.Message{
		Payload: &pb.Message_Finish{
			Finish: &pb.Finish{JobId: jobID},
		},
	})
	mutex.Unlock()

	fmt.Printf("Worker finished workload for job %v\n", jobID)
}

func (w *Worker) handleMessageStop(msgStop *pb.Stop) {
	fmt.Printf("Attempting to stop job with id %v\n", msgStop.GetJobId())
	channel, ok := w.attacks[msgStop.GetJobId()]
	if !ok {
		log.Printf("This is a noop, as job %v is not currently executing", msgStop.GetJobId())
		return
	}
	channel <- struct{}{}
}

// Loop is the main event loop
func (w *Worker) Loop(streamMutex *sync.Mutex, stream pb.Worker_CoordinateClient, wg *sync.WaitGroup, timeMutex *sync.Mutex, lastProcessedTime *time.Time) {
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			return
		}
		if err != nil {
			log.Fatalf("Failed to receive a message: %v", err)
		}

		switch x := msg.Payload.(type) {
		case *pb.Message_Start:
			wg.Add(1)
			w.attacks[x.Start.GetJobId()] = make(chan struct{}, 1)
			go func() {
				w.handleMessageStart(stream, x.Start, streamMutex)
				timeMutex.Lock()
				*lastProcessedTime = time.Now()
				timeMutex.Unlock()
				wg.Done()
			}()
		case *pb.Message_Stop:
			w.handleMessageStop(x.Stop)
		default:
		}
	}
}
