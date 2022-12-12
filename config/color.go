package config

var (
	// StartSet color start chars
	StartSet = []byte("\x1b[")
	// ResetSet color close all properties.
	ResetSet = []byte("\x1b[0m")

	NewLineEndFlag = []byte("\n")
)

type Color uint8

const (
	FgBlack Color = iota + 30
	FgRed
	FgGreen
	FgYellow
	FgBlue
	FgMagenta
	FgCyan
	FgWhite
	// FgDefault revert default FG
	FgDefault Color = 39
)

var ColorMap = make(map[int][]byte, 8)

func init() {
	ColorMap[LevelDebug] = debugColorPrefix
	ColorMap[LevelInfo] = infoColorPrefix
	ColorMap[LevelNotice] = noticeColorPrefix
	ColorMap[LevelWarn] = warnColorPrefix
	ColorMap[LevelError] = errorColorPrefix
	ColorMap[LevelCritical] = criticalColorPrefix
}
