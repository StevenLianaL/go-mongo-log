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
	// 创建一个调用栈切片，用于存储调用栈信息
	stackTrace := make([]uintptr, 15)

	// 获取当前的调用栈信息，跳过前两层调用（printStackTrace 和 runtime.Caller）
	n := runtime.Callers(2, stackTrace)

	// 打印调用栈信息
	var theFuncNames []string
	for i := 0; i < n-1; i++ {
		pc := stackTrace[i]
		funcName := runtime.FuncForPC(pc).Name()
		theFuncNames = append(theFuncNames, funcName)
	}
	finalFuncName := strings.Split(theFuncNames[len(theFuncNames)-1], ".")
	theFinalFuncName := finalFuncName[len(finalFuncName)-2]
	return theFinalFuncName

}
