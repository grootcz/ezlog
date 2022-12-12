package config

import "strconv"

const (
	LevelDebug    = 1
	LevelInfo     = 1 << 1
	LevelNotice   = 1 << 2
	LevelWarn     = 1 << 3
	LevelError    = 1 << 4
	LevelCritical = 1 << 5

	LevelAll = LevelDebug | LevelInfo | LevelNotice |
		LevelWarn | LevelError | LevelCritical
)

var debugColorPrefix = append(StartSet, []byte(strconv.FormatInt(int64(FgWhite), 10)+"m")...)
var infoColorPrefix = append(StartSet, []byte(strconv.FormatInt(int64(FgGreen), 10)+"m")...)
var noticeColorPrefix = append(StartSet, []byte(strconv.FormatInt(int64(FgCyan), 10)+"m")...)
var warnColorPrefix = append(StartSet, []byte(strconv.FormatInt(int64(FgYellow), 10)+"m")...)
var errorColorPrefix = append(StartSet, []byte(strconv.FormatInt(int64(FgRed), 10)+"m")...)
var criticalColorPrefix = append(StartSet, []byte(strconv.FormatInt(int64(FgMagenta), 10)+"m")...)

var LevelMap = make(map[int][]byte, 8)

func init() {
	LevelMap[LevelDebug] = []byte("[ DEBUG  ]")
	LevelMap[LevelInfo] = []byte("[  INFO  ]")
	LevelMap[LevelNotice] = []byte("[ NOTICE ]")
	LevelMap[LevelWarn] = []byte("[  WARN  ]")
	LevelMap[LevelError] = []byte("[ ERROR  ]")
	LevelMap[LevelCritical] = []byte("[CRITICAL]")
}
