package mongo_log

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
	apc, _, _, _ := runtime.Caller(1)
	funcName := runtime.FuncForPC(apc).Name()
	// 只保留最后一个/后的部分，即函数名
	parts := strings.Split(funcName, "/")
	fmt.Println(parts, "parts names")
	pc := make([]uintptr, 10)
	runtime.Callers(0, pc)
	var fullName = runtime.FuncForPC(pc[3]).Name()
	splitName := strings.Split(fullName, ".")
	fmt.Println(splitName, "names")
	var name = splitName[len(splitName)-2]
	return name
}
