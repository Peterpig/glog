package glog

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/gookit/color"
)

var ColorTheme = map[Level]color.Color{
	FatalLevel:  color.FgRed,
	ErrorLevel:  color.FgMagenta,
	WarnLevel:   color.FgYellow,
	NoticeLevel: color.OpBold,
	InfoLevel:   color.FgGreen,
	DebugLevel:  color.FgCyan,
}

const DefaultTemplate = "[{{datetime}}] [{{channel}}] [{{level}}] [{{caller}}] {{message}} {{data}} {{extra}}\n"

type TextFormatter struct {
	template    string
	fieldMap    StringMap
	TimeFormat  string
	EnableColor bool
	ColorTheme  map[Level]color.Color
}

func (f *TextFormatter) Format(record *Record) ([]byte, error) {
	if f.EnableColor {
		return f.formatWithColor(record)
	}
	return f.formatWithColor(record)
}

func (f *TextFormatter) formatWithColor(record *Record) ([]byte, error) {

	// tplData := make(map[string]string, len(f.fieldMap))
	for field, tplVal := range f.fieldMap {
		fmt.Printf("field:%s,tplVal:%s\n", field, tplVal)
	}
	return nil, nil
}

func NewTextFormatter(template ...string) *TextFormatter {
	var fmtTpl string
	if len(template) > 0 {
		fmtTpl = template[0]
	} else {
		fmtTpl = DefaultTemplate
	}

	return &TextFormatter{
		template:   fmtTpl,
		fieldMap:   parseFieldMap(fmtTpl),
		TimeFormat: DefaultTimeFormat,
		ColorTheme: ColorTheme,
	}
}

// parse string "{{channel}}" to map { "channel": "{{channel}}" }
func parseFieldMap(format string) StringMap {
	reg := regexp.MustCompile(`{{\w+}}`)

	ss := reg.FindAllString(format, -1)
	fmt := make(StringMap)

	for _, tplVal := range ss {
		field := strings.Trim(tplVal, "{}")
		fmt[field] = tplVal
	}
	return fmt
}
