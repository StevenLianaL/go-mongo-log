package test

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"mongo-log"
	"testing"
	"time"
)

func TestDriver(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	collection := mongo_log.GetCollection("test", "numbers", "test", "test", "localhost", &ctx)
	res, _ := collection.InsertOne(ctx, bson.D{{"name", "pie"}, {"value", 3.15}})
	t.Log(res)
}
