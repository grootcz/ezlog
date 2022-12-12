package log

import (
	"time"
)

type Log struct {
	FormatType int
	Level      int
	Format     string
	Timestamp  time.Time
	Args       []interface{}
	CallerFile string
	CallerLine int
	CallerName string
	CallerPkg  string
	Stack      []byte
}
