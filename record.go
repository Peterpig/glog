package glog

import (
	"fmt"
	"time"
)

type Record struct {
	logger *Logger

	Time time.Time

	Level     Level
	LevelName string
	Message   string

	Data   M
	Extra  M
	Fields M
}

func newRecord(logger *Logger) *Record {
	return &Record{
		logger: logger,
		Data:   make(M, 3),
		Extra:  make(M, 3),
		Fields: make(M, 3),
	}
}

func (r *Record) Log(level Level, format string, args ...interface{}) {
	r.Level = level
	r.Message = fmt.Sprintf(format, args...)
	r.logger.writeRecord(level, r)
}

func (r *Record) Init(lowerLevelName bool) {
	if lowerLevelName {
		r.LevelName = r.Level.LowerName()
	} else {
		r.LevelName = r.Level.Name()
	}

	if r.Time.IsZero() {
		r.Time = time.Now()
	}
}

// func formatArgsWithSpace(args ...interface{}) string {
// 	// ln := len(args)
// 	// if ln == 0 {
// 	// 	return ""
// 	// }
// 	// if ln == 1 {
// 	// 	return
// 	// }
// }
