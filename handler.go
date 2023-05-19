package glog

import (
	"io"
)

type Handler interface {
	io.Closer

	Flush() error
	IsHandling(level Level) bool

	Handle(*Record) error
	Close() error
}

type BaseHandler struct {
	Levels    []Level
	formatter Formatter
}

func (h *BaseHandler) Formatter() Formatter {
	if h.formatter == nil {
		h.formatter = NewTextFormatter()
	}
	return h.formatter
}

func (h *BaseHandler) IsHandling(level Level) bool {
	for _, l := range h.Levels {
		if l == level {
			return true
		}
	}
	return false
}
