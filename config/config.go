package config

import "time"

type Config struct {
	LogFilePath   string        `json:"logFilePath"`   // log file path, default use DefaultLogFilePath
	LogLevels     int           `json:"logLevels"`     // effect log level, default LevelAll, means all
	LogFilePrefix string        `json:"logFilePrefix"` // log file prefix, default use DefaultLogFilePrefix
	Modules       []string      `json:"modules"`       // effect log modules, default ModulesAll
	IsTimeSplit   bool          `json:"isTimeSplit"`   // is split log file depend on time, default true
	SplitPeriod   time.Duration `json:"splitPeriod"`   // the period of split log file on time，per - second, default DefaultSplitPeriod
	IsSizeSplit   bool          `json:"isSizeSplit"`   // is split log file depend on file size, default true
	SplitSize     int64         `json:"splitSize"`     // the size of split log file on size，per - byte, default DefaultSplitSize
	IsClear       bool          `json:"isClear"`       // is clear expired file, default true
	SavePeriod    time.Duration `json:"savePeriod"`    // log file save period, per - day, default DefaultLogFileSavePeriod
}

// ManagerQueueMaxNumber logger cache queue max length
var ManagerQueueMaxNumber = 500 * 10000 // 500w

// default config
const (
	ModulesAll  = "all"  // display all modules log
	ModulesNone = "none" // hide all modules log

	DefaultLogFilePath   = "./"
	DefaultLogFilePrefix = "default"

	DefaultSplitPeriod       = time.Duration(24) * time.Hour   // 24 hour
	DefaultSplitSize         = 10 << 20                        // 10 MiB
	DefaultLogFileSavePeriod = time.Duration(24*7) * time.Hour // 7 day
)
