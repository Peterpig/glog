package glog

import "io"

type Handler interface {
	io.Closer

	Flush() error
	IsHandling(level Level) bool

	Handle(*Record) error
}
