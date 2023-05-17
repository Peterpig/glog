package glog

import "strings"

type Level uint32

func (level Level) String() string {
	return LevelName[level]
}

func (level Level) Name() string {
	return LevelName[level]
}

func (level Level) LowerName() string {
	return lowerLevelName[level]
}

type M map[string]interface{}

type StringMap map[string]string

var (
	DefaultTimeFormat = "2006-01-02 15:04:05"
)

const (
	// PanicLevel level, highest level of severity. will call panic() if the logging level <= PanicLevel.
	_          Level = iota
	PanicLevel Level = 100 * iota
	// FatalLevel level. Logs and then calls `logger.Exit(1)`. It will exit even if the
	// logging level <= FatalLevel.
	FatalLevel
	// ErrorLevel level. Runtime errors. Used for errors that should definitely be noted.
	// Commonly used for hooks to send errors to an error tracking service.
	ErrorLevel
	// WarnLevel level. Non-critical entries that deserve eyes.
	WarnLevel
	// NoticeLevel level Uncommon events
	NoticeLevel
	// InfoLevel level. Examples: User logs in, SQL logs.
	InfoLevel
	// DebugLevel level. Usually only enabled when debugging. Very verbose logging.
	DebugLevel
)

var (
	ALLlevels = []Level{
		PanicLevel,
		FatalLevel,
		ErrorLevel,
		WarnLevel,
		NoticeLevel,
		InfoLevel,
		DebugLevel,
	}
	LevelName = map[Level]string{
		PanicLevel:  "PANIC",
		FatalLevel:  "FATAL",
		ErrorLevel:  "ERROR",
		NoticeLevel: "NOTICE",
		WarnLevel:   "WARNING",
		InfoLevel:   "INFO",
		DebugLevel:  "DEBUG",
	}
	lowerLevelName = func() map[Level]string {
		ret := make(map[Level]string, len(LevelName))
		for level, name := range LevelName {
			ret[level] = strings.ToLower(name)
		}
		return ret
	}()
)

const (
	CallerFlagFnlFn uint8 = iota
)

const (
	FieldKeyData = "data"

	FieldKeyTime     = "time"
	FieldKeyDate     = "date"
	FieldKeyDatetime = "datetime"

	FieldKeyLevel = "level"
	FieldKeyError = "error"
	FieldKeyExtra = "extra"

	// NOTICE: you must set `Logger.ReportCaller=true` for "caller"
	FieldKeyCaller  = "caller"
	FieldKeyMessage = "message"
)
