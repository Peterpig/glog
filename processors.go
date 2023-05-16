package glog

type Processor interface {
	Process(*Record)
}
