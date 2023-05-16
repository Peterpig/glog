package glog_test

import (
	"fmt"
	"testing"

	"github.com/Peterpig/glog"
)

func TestXxx(t *testing.T) {
	fmt.Println("1111111111")
	fmt.Println(glog.InfoLevel)
	glog.Info("呵呵")
}
