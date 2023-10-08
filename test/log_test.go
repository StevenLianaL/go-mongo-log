package test

import (
	"context"
	"fmt"
	mongoLog "github.com/StevenlianaL/mongo-log"
	"sync"
	"testing"
	"time"
)

var ctx = context.Background()
var collection = mongoLog.GetCollection("test", "test", "test", "test", "localhost", 27017, &ctx)
var manager = mongoLog.Manager{
	Project:    "test",
	App:        "test",
	Collection: collection,
	Ctx:        &ctx,
}

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

	t.Log(manager)
	//for i := 0; i < 10; i++ {
	//	go manager.Log("a new log out", mongoLog.INFO, i)
	//
	//}
	go func() {
		manager.Log("a new log in go func", mongoLog.INFO, 0)
		manager.Warning("a new log in go func warning ", 0)
	}()
	time.Sleep(time.Second * 1)
}
func TestFuncStack(t *testing.T) {
	go func() {
		manager.Warning("a new log in go func warning ", 0)
	}()
}
