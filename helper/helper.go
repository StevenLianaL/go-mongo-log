package helper

import (
	"fmt"
	"runtime"
	"strings"
	"time"
)

func CurrentMonth() string {
	var now = time.Now()
	return fmt.Sprintf("%04d-%02d", now.Year(), now.Month())
}

func GetFuncName() string {
	pc := make([]uintptr, 5)
	runtime.Callers(0, pc)
	var fullName = runtime.FuncForPC(pc[3]).Name()
	splitName := strings.Split(fullName, ".")
	var name = splitName[len(splitName)-1]
	return name
}
