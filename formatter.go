package glog

type Formatter interface {
	Format(record *Record) ([]byte, error)
}
