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
func printStackTrace() {
	// 创建一个调用栈切片，用于存储调用栈信息
	stackTrace := make([]uintptr, 15)

	// 获取当前的调用栈信息，跳过前两层调用（printStackTrace 和 runtime.Caller）
	n := runtime.Callers(2, stackTrace)

	// 打印调用栈信息
	fmt.Println("Call Stack:")
	for i := 0; i < n; i++ {
		pc := stackTrace[i]
		funcName := runtime.FuncForPC(pc).Name()
		fmt.Printf("kkkkk: %s\n", funcName)
	}
}
func GetFuncName() string {
	printStackTrace()
	pc := make([]uintptr, 10)
	runtime.Callers(0, pc)
	var fullName = runtime.FuncForPC(pc[3]).Name()
	splitName := strings.Split(fullName, ".")
	fmt.Println(splitName, "names")
	var name = splitName[len(splitName)-2]
	return name
}
