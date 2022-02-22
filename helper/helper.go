package helper

import (
	"fmt"
	"time"
)

func CurrentMonth() string {
	var now = time.Now()
	return fmt.Sprintf("%04d-%02d", now.Year(), now.Month())
}
