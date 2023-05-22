package handler

import (
	"io"

	"github.com/Peterpig/glog"
)

type BaseHandler struct {
	Levels    []glog.Level
	formatter glog.Formatter
	output    io.Writer
}

func (h *BaseHandler) Formatter() glog.Formatter {
	if h.formatter == nil {
		h.formatter = glog.NewTextFormatter()
	}
	return h.formatter
}

func (h *BaseHandler) IsHandling(level glog.Level) bool {
	for _, l := range h.Levels {
		if l == level {
			return true
		}
	}
	return false
}

func (h *BaseHandler) SetFormatter(formatter glog.Formatter) {
	h.formatter = formatter
}
