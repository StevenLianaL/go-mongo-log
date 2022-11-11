## Example

```go
package main

import (
	"context"
	mongolog "github.com/StevenlianaL/mongo-log"
)

func init() {
	var Ctx = context.Background()

	Collection = mongolog.GetCollection("db_name", "collection_name", "user", "pass", "localhost",27017, &Ctx)
	Manager = mongolog.Manager{
		Project:    "project",
		App:        "app",
		Collection: Collection,
		Ctx:        &Ctx,
	}
}

// use
func AFunction() {
	var user_id = 3
	go func() {
		log.Manager.Info("run a function", user_id)
	}()
}
```