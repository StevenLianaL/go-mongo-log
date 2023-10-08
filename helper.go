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
	if n <= 0 {
		return "Unknown"
	}

	// 获取最后一个调用栈信息
	pc := stackTrace[n-1]
	funcName := runtime.FuncForPC(pc).Name()

	// 使用 "." 分割字符串
	parts := strings.Split(funcName, ".")
	if len(parts) < 2 {
		return "Unknown"
	}

	// 提取倒数第二部分
	return parts[len(parts)-2]
}
