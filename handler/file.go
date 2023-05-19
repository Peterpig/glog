package handler

import (
	"bufio"
	"os"
	"path/filepath"
	"sync"

	"github.com/Peterpig/glog"
)

const defaultMaxSize uint64 = 1024 * 1024 * 1000

type FileHandler struct {
	glog.BaseHandler

	mu sync.Mutex

	useJSON bool

	fpath    string
	file     *os.File
	bufio    *bufio.Writer
	BuffSize int
	MaxSize  uint64
}

func NewFileHandler(fpath string, useJSON bool) *FileHandler {
	h := &FileHandler{
		fpath:   fpath,
		useJSON: useJSON,
		MaxSize: defaultMaxSize,
		BaseHandler: glog.BaseHandler{
			Levels: glog.ALLlevels,
		},
	}

	return h
}

func (h *FileHandler) Sync() error {
	return h.file.Sync()
}

func (h *FileHandler) Flush() error {
	if h.bufio != nil {
		err := h.bufio.Flush()
		if err != nil {
			return err
		}
	}
	return h.Sync()
}

func (h *FileHandler) Close() error {
	if err := h.Flush(); err != nil {
		return err
	}

	return h.file.Close()
}

func (h *FileHandler) Handle(r *glog.Record) error {
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

	_, err = h.bufio.Write(bts)
	return err

}

func createFile(fpath string) (file *os.File, err error) {
	dirPath := filepath.Dir(fpath)
	err = os.MkdirAll(dirPath, 0777)
	if err != nil {
		return
	}

	file, err = os.OpenFile(fpath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0755)
	if err != nil {
		return
	}
	return
}
