package driver

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"testing"
	"time"
)

func TestDriver(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	collection := GetCollection("test", "numbers", "test", "test", "localhost", &ctx)
	res, _ := collection.InsertOne(ctx, bson.D{{"name", "pie"}, {"value", 3.15}})
	t.Log(res)
}
