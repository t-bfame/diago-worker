package pkg

import (
	"strconv"
	"time"
	"log"

	ptypes "github.com/golang/protobuf/ptypes"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	vegeta "github.com/tsenart/vegeta/v12/lib"
)

type Metrics struct {
	// Latencies holds computed request latency metrics.
	Latencies []*durationpb.Duration `json:"latencies"`

	// Histogram, only if requested
	// Histogram *Histogram `json:"buckets,omitempty"`

	// BytesIn holds computed incoming byte metrics.
	BytesIn uint64 `json:"bytes_in"`

	// BytesOut holds computed outgoing byte metrics.
	BytesOut uint64 `json:"bytes_out"`

	// Earliest is the earliest timestamp in a Result set.
	Earliest time.Time `json:"earliest"`

	// Latest is the latest timestamp in a Result set.
	Latest time.Time `json:"latest"`

	// End is the latest timestamp in a Result set plus its latency.
	End time.Time `json:"end"`

	// Requests is the total number of requests executed.
	Requests uint64 `json:"requests"`

	// StatusCodes is a histogram of the responses' status codes.
	StatusCodes map[string]uint64 `json:"status_codes"`

	// Errors is a set of unique errors returned by the targets during the attack.
	Errors []string `json:"errors"`

	// Used for fast lookup of errors in Errors
	errors  map[string]struct{}

	// Finish message has been received
	Finished bool 
}


// Add implements the Add method of the Report interface by adding the given
// Result to Metrics.
func (m *Metrics) AddVegetaResult(jobID string, res *vegeta.Result) {
	timestampProto, err := ptypes.TimestampProto(res.Timestamp)
	if err != nil {
		log.Fatal(err)
	}
	timestamp, err := ptypes.Timestamp(timestampProto)
	if err != nil {
		log.Fatal(err)
	}

	m.Requests ++
	m.StatusCodes[strconv.Itoa(int(res.Code))]++
	m.BytesOut += res.BytesOut
	m.BytesIn += res.BytesIn


	m.Latencies = append(m.Latencies, ptypes.DurationProto(res.Latency))

	if m.Earliest.IsZero() || m.Earliest.After(timestamp) {
		m.Earliest = timestamp
	}

	if timestamp.After(m.Latest) {
		m.Latest = timestamp
	}

	if end := timestamp.Add(res.Latency); end.After(m.End) {
		m.End = end
	}

	if res.Error != "" {
		if _, ok := m.errors[res.Error]; !ok {
			m.errors[res.Error] = struct{}{}
			m.Errors = append(m.Errors, res.Error)
		}
	}
}

func (m *Metrics) Finish() {
	m.Finished = true
}

// NewMetrics Creates a new Metrics
func NewMetrics() *Metrics {
	var m Metrics

	if m.StatusCodes == nil {
		m.StatusCodes = map[string]uint64{}
	}

	if m.errors == nil {
		m.errors = map[string]struct{}{}
	}

	if m.Errors == nil {
		m.Errors = make([]string, 0)
	}

	return &m
}
