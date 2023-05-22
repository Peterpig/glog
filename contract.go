package glog

import "io"

type Handler interface {
	io.Closer

	Flush() error
	IsHandling(Level) bool

	Handle(*Record) error
	Close() error
}
