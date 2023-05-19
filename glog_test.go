package glog_test

import (
	"testing"

	"github.com/Peterpig/glog"
)

func TestCommon(t *testing.T) {
	glog.Info("呵呵")
	glog.Debug("测试一下 %s", "xxx")
	glog.Error("我是错误")
}

func TestJsonFormatter(t *testing.T) {
	glog.SetFormate(glog.NewJSONFormatter())
	glog.Info("Json格式测试")
	glog.Error("Josn错误测试")
}

func TestJosnFormatterIndent(t *testing.T) {

	f := func(f *glog.JSONFormatter) {
		f.PrettyPrint = true
	}

	glog.SetFormate(glog.NewJSONFormatter(f))

	glog.Info("Josn格式缩进测试")
	glog.Error("Json %s", "格式")
}
