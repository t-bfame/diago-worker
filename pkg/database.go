package pkg

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/t-bfame/diago-worker/pkg/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

var (
	MongoClient      *mongo.Client
	ResponseDataColl *mongo.Collection
)

const (
	dbName                 = "diago-worker"
	responseDataCollection = "responsedata"
	TTLseconds             = 10 * 24 * 60 * 60
)

func ConnectToDB(uri string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected and pinged.")
	MongoClient = client
	ResponseDataColl = MongoClient.Database(dbName).Collection(responseDataCollection)

	index := mongo.IndexModel{
		Keys:    bsonx.Doc{{Key: "created_at", Value: bsonx.Int32(1)}},
		Options: options.Index().SetExpireAfterSeconds(TTLseconds),
	}
	_, err = ResponseDataColl.Indexes().CreateOne(context.Background(), index)
	if err != nil {
		log.Printf("Index error: %s", err.Error())
	}
	return nil
}

func CreateResponseData(ctx context.Context, responseData *model.ResponseData) error {
	coll := ResponseDataColl
	data, err := bson.Marshal(responseData)
	if err != nil {
		return err
	}
	_, err = coll.InsertOne(ctx, data)
	if err != nil {
		return err
	}
	return nil
}
