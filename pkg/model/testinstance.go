package model

import "time"

type ResponseData struct {
	CreatedAt      time.Time `bson:"created_at"`
	JobID          string    `bson:"job_id"`
	TestInstanceID string    `bson:"testinstance_id"`
	TestID         string    `bson:"test_id"`
	Response       string    `bson:"response"`
}
