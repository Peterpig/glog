package glog

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/gookit/color"
	"github.com/valyala/bytebufferpool"
)

var ColorTheme = map[Level]color.Color{
	FatalLevel:  color.FgRed,
	ErrorLevel:  color.FgMagenta,
	WarnLevel:   color.FgYellow,
	NoticeLevel: color.OpBold,
	InfoLevel:   color.FgGreen,
	DebugLevel:  color.FgCyan,
}

const DefaultTemplate = "[{{datetime}}] [{{level}}] [{{caller}}] {{message}}\n"

type TextFormatter struct {
	template    string
	fields      []string
	TimeFormat  string
	EnableColor bool
	ColorTheme  map[Level]color.Color
}

var textPool bytebufferpool.Pool

func (f *TextFormatter) Format(record *Record) ([]byte, error) {
	buf := textPool.Get()
	defer textPool.Put(buf)

	fieldLen := len(f.fields)
	for i, field := range f.fields {
		fmt.Printf("field:%s\n", field)

		switch {
		case field == FieldKeyDatetime:
			buf.B = record.Time.AppendFormat(buf.B, f.TimeFormat)
		case field == FieldKeyCaller && record.Caller != nil:
			buf.WriteString(formatCaller(record.Caller, false))
		case field == FieldKeyLevel:
			if f.EnableColor {
				buf.WriteString(f.renderColorByLevel(record.Level.Name(), record.Level))
			} else {
				buf.WriteString(record.Level.Name())
			}
		case field == FieldKeyMessage:
			if f.EnableColor {
				buf.WriteString(f.renderColorByLevel(record.Message, record.Level))
			} else {
				buf.WriteString(record.Level.Name())
			}
		default:
			if _, ok := record.Fields[field]; ok {
				buf.WriteString(fmt.Sprintf("%v", record.Fields[field]))
			} else {
				buf.WriteString(field)
			}
		}
		if i <= fieldLen-1 {
			buf.WriteString(" ")
		}

	}
	return buf.B, nil
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
		fields:     parseTemplateToFeilds(fmtTpl),
		TimeFormat: DefaultTimeFormat,
		ColorTheme: ColorTheme,
	}
}

func parseTemplateToFeilds(tplStr string) []string {
	re := regexp.MustCompile(`{{\w+}}`)
	ss := re.FindAllString(tplStr, -1)
	fields := make([]string, 0, len(ss)*2)
	for _, tplVar := range ss {
		fields = append(fields, strings.Trim(tplVar, "{}"))
	}
	return fields
}

func (f *TextFormatter) renderColorByLevel(text string, level Level) string {
	if them, ok := f.ColorTheme[level]; ok {
		return them.Render(text)
	}
	return text
}
