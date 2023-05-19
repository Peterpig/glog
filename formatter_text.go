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

const DefaultTemplate = "[{{datetime}}] [{{level}}] [{{caller}}] {{message}}\n"

type TextFormatter struct {
	template    string
	fields      []string
	TimeFormat  string
	EnableColor bool
	ColorTheme  map[Level]color.Color
}

func (f *TextFormatter) Format(record *Record) ([]byte, error) {

	fieldLen := len(f.fields)
	pars := make([]string, fieldLen*2)

	var tmpVal string
	for _, field := range f.fields {
		tempKey := fmt.Sprintf("{{%s}}", field)

		switch {
		case field == FieldKeyDatetime:
			tmpVal = record.Time.Format(f.TimeFormat)
		case field == FieldKeyCaller && record.Caller != nil:
			tmpVal = formatCaller(record.Caller, false)
		case field == FieldKeyLevel:
			if f.EnableColor {
				tmpVal = f.renderColorByLevel(record.Level.Name(), record.Level)
			} else {
				tmpVal = record.Level.Name()
			}
		case field == FieldKeyMessage:
			if f.EnableColor {
				tmpVal = f.renderColorByLevel(record.Message, record.Level)
			} else {
				tmpVal = record.Message
			}
		default:
			if _, ok := record.Fields[field]; ok {
				tmpVal = fmt.Sprintf("%v", record.Fields[field])
			} else {
				tmpVal = field
			}
		}

		pars = append(pars, tempKey, tmpVal)
	}

	str := strings.NewReplacer(pars...).Replace(f.template)
	return []byte(str), nil
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
