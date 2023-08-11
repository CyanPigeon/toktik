package util

import "fmt"

const (
	RED    = "\\033[31m"
	YELLOW = "\\033[33m"
	GREEN  = "\\033[32m"
	BLUE   = "\\033[34m"
)

type Logger interface {
	Log(code int, msg string)
	LogLevel(level int, code int, msg string)
}

type LogDetail struct {
	level int
	code  int
	msg   string
}

func (ld LogDetail) Log(code int, msg string) {
	fmt.Printf("[INFO] code: %d, msg: %s", code, msg)
}

// LogLevel level参数
//
// 1 表示 WARNING
//
// 2 表示 ERROR
//
// 其他表示 INFO
func (ld LogDetail) LogLevel(level int, code int, msg string) {
	switch level {
	case 1:
		fmt.Printf("%s [WARNING] code: %d, msg: %s\n", YELLOW, code, msg)
	case 2:
		fmt.Printf("%s [ERROR] code: %d, msg: %s\n", RED, code, msg)
	default:
		fmt.Printf("%s [INFO] code: %d, msg: %s\n", GREEN, code, msg)
	}
}
