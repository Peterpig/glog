package glog

import (
	"bytes"
	"encoding/json"
)

type JSONFormatter struct {
	Fields []string

	Aliases StringMap

	PrettyPrint bool

	TimeFormat string
}

func NewJSONFormatter(fn ...func(*JSONFormatter)) *JSONFormatter {
	f := &JSONFormatter{
		Fields:     DefaultFields,
		TimeFormat: DefaultTimeFormat,
	}

	if len(fn) > 0 {
		fn[0](f)
	}

	return f
}

func (f *JSONFormatter) Format(r *Record) ([]byte, error) {
	data := make(M, len(f.Fields))

	for _, field := range f.Fields {
		outName, ok := f.Aliases[field]
		if !ok {
			outName = field
		}

		switch {
		case field == FieldKeyDatetime:
			data[outName] = r.Time.Format(f.TimeFormat)
		case field == FieldKeyLevel:
			data[outName] = r.LevelName
		case field == FieldKeyCaller && r.Caller != nil:
			data[outName] = formatCaller(r.Caller, true)
		case field == FieldKeyMessage:
			data[outName] = r.Message
		default:
			data[outName] = r.Fields[field]
		}
	}

	buffer := &bytes.Buffer{}
	encode := json.NewEncoder(buffer)

	if f.PrettyPrint {
		encode.SetIndent("", "    ")
	}

	err := encode.Encode(data)
	return buffer.Bytes(), err
}
