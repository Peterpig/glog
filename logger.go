package glog

import (
	"fmt"
	"os"
	"sync"
)

type Logger struct {
	name     string
	err      error
	handlers []Handler

	mu sync.Mutex

	ReportCaller bool
	CallerFlag   uint8
	CallerSkip   int

	recordPool     sync.Pool
	LowerLevelName bool
}

func New() *Logger {
	return NewWithName("")
}

func NewWithName(name string) *Logger {
	logger := &Logger{
		name:       name,
		CallerSkip: 6,
	}

	logger.recordPool.New = func() interface{} {
		return newRecord(logger)
	}
	return logger
}

func (log *Logger) Close() error {
	return nil
}

func (log *Logger) Flush() error {
	return nil
}

func (log *Logger) AddHandler(handler ...Handler) {
	if len(handler) > 0 {
		log.handlers = append(log.handlers, handler...)
	}
}

func (log *Logger) Log(level Level, format string, args ...interface{}) {

	r := log.recordPool.Get().(*Record)
	defer log.releaseRecord(r)

	r.Log(level, format, args...)

}

func (log *Logger) releaseRecord(r *Record) {
	r.Data = map[string]interface{}{}
	r.Extra = map[string]interface{}{}
	r.Fields = map[string]interface{}{}
	log.recordPool.Put(r)
}

func (log *Logger) writeRecord(level Level, r *Record) {

	var inited bool
	for _, handler := range log.handlers {
		if handler.IsHandling(level) {
			if !inited {
				r.Init(log.LowerLevelName)

				if log.ReportCaller {
					log.mu.Lock()
					caller, ok := getCaller(r.CallerSkip)
					if ok {
						r.Caller = &caller
					}
					log.mu.Unlock()
				}
			}

			if err := handler.Handle(r); err != nil {
				log.err = err
				fmt.Fprintln(os.Stderr, "glog handler error, ", err)
			}
		}
	}
}
