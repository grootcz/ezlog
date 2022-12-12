package manager

import (
	"context"
	"fmt"
	"io"
	"os"
	"runtime/debug"
	"sync"
	"time"

	"github.com/grootcz/ezlog/config"
	"github.com/grootcz/ezlog/log"
	"github.com/grootcz/ezlog/utils"
)

type Manager struct {
	conf       config.Config
	cacheQueue chan *log.Log
	out        io.Writer
	mu         sync.RWMutex

	logFormatHandler  OutputFormat
	filterHandlerList []LogHandler
}

func newLogManager() *Manager {
	c := config.Config{
		LogFilePath:   config.DefaultLogFilePath,
		LogLevels:     config.LevelAll,
		LogFilePrefix: config.DefaultLogFilePrefix,
		Modules:       make([]string, 0, 64),
		IsTimeSplit:   true,
		SplitPeriod:   config.DefaultSplitPeriod,
		IsSizeSplit:   true,
		SplitSize:     config.DefaultSplitSize,
		IsClear:       true,
		SavePeriod:    config.DefaultLogFileSavePeriod,
	}
	c.Modules = append(c.Modules, config.ModulesAll)

	logManager := &Manager{
		conf:              c,
		cacheQueue:        make(chan *log.Log, config.ManagerQueueMaxNumber),
		out:               os.Stdout,
		filterHandlerList: make([]LogHandler, 0, 8),
	}

	logManager.logFormatHandler = logManager.DefaultOutputFormat

	logManager.SetConfig(c)
	go logManager.run()

	return logManager
}

var instance *Manager
var once sync.Once

// GetLogger get log impl
func GetLogManager() *Manager {
	once.Do(func() {
		instance = newLogManager()
	})

	return instance
}

func (m *Manager) Add(level int, format string, v ...interface{}) {
	l := &log.Log{
		Level:     level,
		Format:    format,
		Args:      v,
		Timestamp: time.Now(),
	}

	utils.ThirdCallerInfo(l)

	if !m.isLevelFilterPassed(l) {
		return
	}

	if !m.isModuleFilterPassed(l) {
		return
	}

	if len(m.filterHandlerList) > 0 {
		for _, handler := range m.filterHandlerList {
			if !handler(l) {
				return
			}
		}
	}

	if (cap(m.cacheQueue) - len(m.cacheQueue)) <= cap(m.cacheQueue)/100 {
		fmt.Printf("logger cache queue is nearly full!!!\n")
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(300)*time.Millisecond)
	defer cancel()

	select {
	case m.cacheQueue <- l:
		{
			return
		}
	case <-ctx.Done():
		{
			fmt.Printf("input to cache queue timeout!!!\n")
		}
	}
}

func (m *Manager) AddWithStack(level int, format string, v ...interface{}) {
	l := &log.Log{
		Level:     level,
		Format:    format,
		Args:      v,
		Timestamp: time.Now(),
		Stack:     debug.Stack(),
	}

	utils.ThirdCallerInfo(l)

	if (cap(m.cacheQueue) - len(m.cacheQueue)) <= cap(m.cacheQueue)/100 {
		fmt.Printf("logger cache queue is nearly full!!!\n")
	}

	m.cacheQueue <- l
}

func (m *Manager) output(log *log.Log) {
	if m.logFormatHandler == nil {
		fmt.Printf("log output format handler is nil\n")
		return
	}

	content, err := m.logFormatHandler(log)
	if err != nil {
		fmt.Printf("log output format failed:%s\n", err)
		return
	}

	m.mu.RLock()
	defer m.mu.RUnlock()

	var formatLogBytes []byte
	if m.out == os.Stdout || m.out == os.Stderr {
		if log.Stack == nil {
			formatLogBytes = make([]byte, 0, len(config.ColorMap[log.Level])+len(content)+
				len(config.ResetSet)+len(config.NewLineEndFlag))
		} else {
			formatLogBytes = make([]byte, 0, len(config.ColorMap[log.Level])+len(content)+
				len(config.ResetSet)+len(config.NewLineEndFlag)+len(log.Stack)+len(config.NewLineEndFlag))
		}

		formatLogBytes = append(formatLogBytes, config.ColorMap[log.Level]...)
		formatLogBytes = append(formatLogBytes, content...)
		if log.Stack != nil {
			formatLogBytes = append(formatLogBytes, config.NewLineEndFlag...)
			formatLogBytes = append(formatLogBytes, log.Stack...)
		}
		formatLogBytes = append(formatLogBytes, config.ResetSet...)
		formatLogBytes = append(formatLogBytes, config.NewLineEndFlag...)
	} else {
		if log.Stack == nil {
			formatLogBytes = make([]byte, 0, len(content)+len("\n"))
		} else {
			formatLogBytes = make([]byte, 0, len(content)+len("\n")+len(log.Stack)+len(config.NewLineEndFlag))
		}

		formatLogBytes = append(formatLogBytes, content...)
		if log.Stack != nil {
			formatLogBytes = append(formatLogBytes, config.NewLineEndFlag...)
			formatLogBytes = append(formatLogBytes, log.Stack...)
		}
		formatLogBytes = append(formatLogBytes, config.NewLineEndFlag...)
	}

	m.out.Write(formatLogBytes)
}

func (m *Manager) run() {
	for {
		log := <-m.cacheQueue
		m.output(log)
	}
}
