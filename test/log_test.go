package test

import (
	"context"
	"fmt"
	mongoLog "github.com/StevenlianaL/mongo-log"
	"sync"
	"testing"
	"time"
)

// should use like this case
func TestManagerInWg(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	fmt.Println("tag")
	collection := mongoLog.GetCollection("test", "test", "test", "test", "localhost", 27017, &ctx)
	var manager = mongoLog.Manager{
		Project:    "test",
		App:        "test",
		Collection: collection,
		Ctx:        &ctx,
	}
	t.Log(manager)
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			manager.Log("a new log", mongoLog.INFO, i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}

func TestManagerInRoutine(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	collection := mongoLog.GetCollection("test", "test", "test", "test", "localhost", 27017, &ctx)
	var manager = mongoLog.Manager{
		Project:    "test",
		App:        "test",
		Collection: collection,
		Ctx:        &ctx,
	}
	t.Log(manager)
	for i := 0; i < 10; i++ {
		go manager.Log("a new log", mongoLog.INFO, i)
	}
	time.Sleep(time.Second * 1)
}
