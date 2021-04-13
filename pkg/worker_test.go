package pkg

import (
	"testing"
	"time"
	vegeta "github.com/tsenart/vegeta/v12/lib"
)

func TestNewWorker(t *testing.T) {
	w := NewWorker()
	if w == nil {
		t.Errorf("Expected worker to be created")
	}
	if w.attacks == nil {
		t.Errorf("Expected worker.attacks to be initialized")
	}
}

func TestMetricsFromVegetaResult(t *testing.T) {
	w := NewWorker()
	res := &vegeta.Result{
		Attack: "attack",
		Code: 200,
		Timestamp: time.Now(),
		Latency: time.Second,
		BytesOut: 0,
		BytesIn: 0,
	}
	jobID := "jobID"
	metrics := w.MetricsFromVegetaResult(jobID, res)
	if (jobID != metrics.JobId) {
		t.Errorf("Incorrect value for JobID field")
	}
	if (uint32(res.Code) != metrics.Code) {
		t.Errorf("Incorrect value for Code field")
	}
	if (res.BytesIn != metrics.BytesIn) {
		t.Errorf("Incorrect value for BytesIn field")
	}
	if (res.BytesOut != metrics.BytesOut) {
		t.Errorf("Incorrect value for BytesOut field")
	}
	if (int64(res.Latency) != metrics.Latency) {
		t.Errorf("Incorrect value for Latency field")
	}
	if (res.Error != metrics.Error) {
		t.Errorf("Incorrect value for Error field")
	}
}


