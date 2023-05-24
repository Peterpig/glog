package handler

import (
	"bufio"
	"os"
	"time"

	"github.com/Peterpig/glog"
)

type RotateTime int

const (
	EveryDay    RotateTime = RotateTime(24 * time.Hour)
	EveryHour   RotateTime = RotateTime(time.Hour)
	Every30Min  RotateTime = RotateTime(30 * time.Minute)
	Every15Min  RotateTime = RotateTime(15 * time.Minute)
	EveryMinute RotateTime = RotateTime(time.Minute)
)

type RotateFileHandler struct {
	FileHandler
	RotateTime RotateTime
	written    uint64
}

func NewSizeRotateFileHandler(fpath string, maxSize int) *RotateFileHandler {
	h := &RotateFileHandler{
		FileHandler: FileHandler{
			BaseHandler: BaseHandler{
				Levels: glog.ALLlevels,
			},
			useJSON: false,
			fpath:   fpath,
			MaxSize: uint64(maxSize),
		},
	}

	return h
}

func (h *RotateFileHandler) Handle(r *glog.Record) error {
	h.mu.Lock()
	defer h.mu.Unlock()

	bts, err := h.Formatter().Format(r)
	if err != nil {
		return err
	}

	if h.file, err = createFile(h.fpath); err != nil {
		return err
	}

	if h.bufio == nil {
		h.bufio = bufio.NewWriterSize(h.file, h.BuffSize)
	}

	_, err = h.Write(bts)
	return err

}

func (h *RotateFileHandler) Write(bts []byte) (n int, err error) {
	if h.written+uint64(len(bts)) >= h.MaxSize {
		if err := h.rotateFile(time.Now()); err != nil {
			return 0, err
		}
	}

	n, err = h.file.Write(bts)
	h.written += uint64(n)
	return
}

func (h *RotateFileHandler) rotateFile(now time.Time) error {
	if h.file != nil {
		h.Close()
	}

	var err error

	err = os.Rename(h.fpath, h.fpath+"."+now.Format("20060102150405"))
	if err != nil {
		return err
	}

	h.file, err = OpenFile(h.fpath)
	h.written = 0
	return err

}
