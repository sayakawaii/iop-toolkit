/*
# ------------------------------------------------------------
# -- trace.go
# --
# -- Huang Minghe
# -- 2022-4-30
# ------------------------------------------------------------
*/

package utils

import (
	"fmt"
	"path"
	"runtime"
	"time"
)

var logEnable = true

func Log(a ...any) {
	if logEnable {
		timeLayoutStr := "2006-01-02 15:04:05"
		pc, file, lineNo, ok := runtime.Caller(1)
		if ok {
			funcName := runtime.FuncForPC(pc).Name()
			fileName := path.Base(file)
			info := fmt.Sprintf("FuncName:%s, file:%s, line:%d ", funcName, fileName, lineNo)
			fmt.Println("[" + time.Now().Format(timeLayoutStr) + "]" + "[" + info + "]" + fmt.Sprintln(a...))
		}
	}
}
