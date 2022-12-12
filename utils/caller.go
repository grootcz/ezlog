package utils

import (
	"runtime"
	"strings"

	"github.com/grootcz/ezlog/log"
)

// ThirdCallerInfo third layout caller
func ThirdCallerInfo(log *log.Log) {
	level := 3
	pc, file, line, ok := runtime.Caller(level)
	if !ok {
		return
	}

	pcInfoStr := runtime.FuncForPC(pc).Name()

	var pkgName string
	var funcName string

	lastPathSplitIndex := strings.LastIndex(pcInfoStr, "/")
	if lastPathSplitIndex <= 0 {
		firstFuncSplitIndex := strings.Index(pcInfoStr, ".")
		pkgName = pcInfoStr[:firstFuncSplitIndex]
		funcName = pcInfoStr[firstFuncSplitIndex+1:]
	} else {
		pkgStr := pcInfoStr[:lastPathSplitIndex]
		funcStr := pcInfoStr[lastPathSplitIndex+1:]

		firstFuncSplitIndex := strings.Index(funcStr, ".")

		pkgName = pkgStr + "/" + funcStr[:firstFuncSplitIndex]
		funcName = funcStr[firstFuncSplitIndex+1:]
	}

	if !strings.Contains(funcName, "()") {
		funcName = funcName + "()"
	}

	log.CallerFile = file
	log.CallerLine = line
	log.CallerName = funcName
	log.CallerPkg = pkgName

	return
}
