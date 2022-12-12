package ezlog

import (
	"io"

	"github.com/grootcz/ezlog/config"
	"github.com/grootcz/ezlog/manager"
)

func SetConfig(c config.Config) {
	manager.GetLogManager().SetConfig(c)
}

func SetOutput(w io.Writer) {
	manager.GetLogManager().SetOutput(w)
}

func SetLevels(levels int) {
	manager.GetLogManager().SetLevels(levels)
}

func SetModules(modules []string) {
	manager.GetLogManager().SetModules(modules)
}

func SetOutputFormatHandler(handler manager.OutputFormat) {
	manager.GetLogManager().SetOutputFormatHandler(handler)
}

func AddFilterHandler(handler manager.LogHandler) {
	manager.GetLogManager().AddLogFilterHandler(handler)
}

func Debugf(format string, v ...interface{}) {
	manager.GetLogManager().Add(config.LevelDebug, format, v...)
}

func Infof(format string, v ...interface{}) {
	manager.GetLogManager().Add(config.LevelInfo, format, v...)
}

func Noticef(format string, v ...interface{}) {
	manager.GetLogManager().Add(config.LevelNotice, format, v...)
}

func Warnf(format string, v ...interface{}) {
	manager.GetLogManager().Add(config.LevelWarn, format, v...)
}

func Errorf(format string, v ...interface{}) {
	manager.GetLogManager().Add(config.LevelError, format, v...)
}

func Criticalf(format string, v ...interface{}) {
	manager.GetLogManager().Add(config.LevelCritical, format, v...)
}

func DebugfWithStack(format string, v ...interface{}) {
	manager.GetLogManager().AddWithStack(config.LevelDebug, format, v...)
}

func InfofWithStack(format string, v ...interface{}) {
	manager.GetLogManager().AddWithStack(config.LevelInfo, format, v...)
}

func NoticefWithStack(format string, v ...interface{}) {
	manager.GetLogManager().AddWithStack(config.LevelNotice, format, v...)
}

func WarnfWithStack(format string, v ...interface{}) {
	manager.GetLogManager().AddWithStack(config.LevelWarn, format, v...)
}

func ErrorfWithStack(format string, v ...interface{}) {
	manager.GetLogManager().AddWithStack(config.LevelError, format, v...)
}

func CriticalfWithStack(format string, v ...interface{}) {
	manager.GetLogManager().AddWithStack(config.LevelCritical, format, v...)
}
