package glog

import (
	"io"
	"os"
)

type SugaredLogger struct {
	*Logger

	Formatter

	Output io.Writer

	Level Level
}

func (sg *SugaredLogger) Handle(record *Record) error {
	bts, err := sg.Formatter.Format(record)
	if err != nil {
		return err
	}
	_, err = sg.Output.Write(bts)
	return err
}

func (sg *SugaredLogger) IsHandling(level Level) bool {
	return sg.Level <= level
}

func NewSugaredLogger(output io.Writer, level Level) *SugaredLogger {
	sg := &SugaredLogger{
		Output:    output,
		Level:     level,
		Logger:    New(),
		Formatter: NewTextFormatter(),
	}

	sg.AddHandler(sg)
	return sg
}

func newStdLogger() *SugaredLogger {
	return NewSugaredLogger(os.Stdout, ErrorLevel)
}

var glog = newStdLogger()

func Debug(format string, v ...interface{}) {
	glog.Log(DebugLevel, format, v)
}

func Info(format string, v ...interface{}) {
	glog.Log(InfoLevel, format, v...)
}

func Notice(format string, v ...interface{}) {
	glog.Log(NoticeLevel, format, v...)
}

func Warn(format string, v ...interface{}) {
	glog.Log(WarnLevel, format, v...)
}

func Error(format string, v ...interface{}) {
	glog.Log(ErrorLevel, format, v)
}

func Fatal(format string, v ...interface{}) {
	glog.Log(FatalLevel, format, v...)
}
