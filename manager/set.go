package manager

import (
	"io"
	"time"

	"github.com/grootcz/ezlog/config"
)

func (m *Manager) SetConfig(c config.Config) {
	if c.LogFilePath != "" {
		m.conf.LogFilePath = c.LogFilePrefix
	}

	if c.LogLevels > 0 {
		m.conf.LogLevels = c.LogLevels
	}

	if c.LogFilePrefix != "" {
		m.conf.LogFilePrefix = c.LogFilePrefix
	}

	if len(c.Modules) > 0 {
		m.conf.Modules = c.Modules
	}

	if c.SplitPeriod <= time.Duration(10)*time.Minute {
		c.SplitPeriod = time.Duration(10) * time.Minute
	}

	if c.SplitSize <= 1<<20 {
		c.SplitPeriod = 1 << 20
	}

	m.conf = c
}

func (m *Manager) SetOutput(w io.Writer) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.out = w
}

func (m *Manager) SetLevels(levels int) {
	m.conf.LogLevels = levels
}

func (m *Manager) SetModules(modules []string) {
	m.conf.Modules = modules
}

func (m *Manager) SetOutputFormatHandler(handler OutputFormat) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.logFormatHandler = handler
}

func (m *Manager) AddLogFilterHandler(handler LogHandler) {
	m.filterHandlerList = append(m.filterHandlerList, handler)
}
