package handler

import (
	"os"

	"github.com/Peterpig/glog"
	"github.com/gookit/color"
)

type ConsoleHandler struct {
	BaseHandler
}

func NewConsoleHandler(level []glog.Level) *ConsoleHandler {
	h := &ConsoleHandler{}
	h.output = os.Stdout
	h.Levels = level

	f := glog.NewTextFormatter()
	f.EnableColor = color.IsSupportColor()

	h.SetFormatter(f)
	return h
}

func (h *ConsoleHandler) Handle(r *glog.Record) error {
	bts, err := h.Formatter().Format(r)
	if err != nil {
		return err
	}
	_, err = h.output.Write(bts)
	return err
}

func (h *ConsoleHandler) Close() error {
	return nil
}

func (h *ConsoleHandler) Flush() error {
	return nil
}
