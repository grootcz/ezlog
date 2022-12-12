package manager

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"unsafe"

	"github.com/grootcz/ezlog/config"
	"github.com/grootcz/ezlog/log"
)

type OutputFormat func(l *log.Log) ([]byte, error)

func (m *Manager) DefaultOutputFormat(l *log.Log) ([]byte, error) {
	total := len(config.LevelMap[l.Level]) + len(" ") +
		len(l.Timestamp.Format("2006-01-02T15:04:05.999999999Z07:00")+" ⇔ ") +
		len(l.CallerFile+":") +
		len(strconv.Itoa(l.CallerLine)+" ◆ ") +
		len(l.CallerPkg+" ▶ ") +
		len(l.CallerName+" ★  ")

	buf := make([]byte, 0, total)
	buf = append(buf, config.LevelMap[l.Level]...)
	buf = append(buf, " "...)
	buf = append(buf, l.Timestamp.Format("2006-01-02T15:04:05.999999999Z07:00")...)
	buf = append(buf, " ⇔ "...)
	buf = append(buf, l.CallerFile...)
	buf = append(buf, ":"...)
	buf = append(buf, strconv.Itoa(l.CallerLine)...)
	buf = append(buf, " ◆ "...)
	buf = append(buf, l.CallerPkg...)
	buf = append(buf, " ▶ "...)
	buf = append(buf, l.CallerName...)
	buf = append(buf, " ★  "...)

	bufPointer := unsafe.Pointer(&buf)
	bufHeader := (*reflect.SliceHeader)(bufPointer)
	strHeader := reflect.StringHeader{
		Data: bufHeader.Data,
		Len:  bufHeader.Len,
	}
	formatLog := *(*string)(unsafe.Pointer(&strHeader))

	formatLog = fmt.Sprintf(formatLog+strings.Trim(l.Format, "\n"), l.Args...)
	strPointer := unsafe.Pointer(&formatLog)
	logStrHeader := (*reflect.StringHeader)(strPointer)
	logBytes := reflect.SliceHeader{
		Data: logStrHeader.Data,
		Len:  logStrHeader.Len,
		Cap:  logStrHeader.Len,
	}

	return *(*[]byte)(unsafe.Pointer(&logBytes)), nil
}

type LogHandler func(l *log.Log) bool

func (m *Manager) isLevelFilterPassed(l *log.Log) bool {
	isPass := m.conf.LogLevels & l.Level
	if isPass == 0 {
		return false
	}

	return true
}

func (m *Manager) isModuleFilterPassed(l *log.Log) bool {
	if len(m.conf.Modules) == 0 {
		return false
	}

	var isPass bool
	for _, value := range m.conf.Modules {
		if value == config.ModulesAll {
			isPass = true
			break
		} else if value == config.ModulesNone {
			isPass = false
			break
		} else if value == l.CallerPkg {
			isPass = true
		} else {
			continue
		}
	}

	return isPass
}
