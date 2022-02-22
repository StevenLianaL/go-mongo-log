package log

import (
	"context"
	"mongo-log"
	"mongo-log/driver"
	"mongo-log/level"
	"testing"
	"time"
)

func TestManager(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	collection := driver.GetCollection("record", "log_record", "test", "test", "localhost", &ctx)
	var manager = mongo_log.Manager{
		Project:    "test",
		App:        "test",
		Collection: collection,
		Ctx:        &ctx,
	}
	t.Log(manager)
	manager.Log("a new log", level.INFO, 1)
}
