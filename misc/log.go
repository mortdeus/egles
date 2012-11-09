package glog

import (
	"log"
	"os"
)

type l struct {
	*log.Logger
}

var glog *l

func Init(flag int) {
	if flag == 0 {
		flag = log.LstdFlags
	}
	glog = &l{log.New(os.Stdout, "glog:\t", flag)}
	glog.Println("Init")
}
func Println(v interface{}) {
	glog.Println(v)
}
func Printf(format string, v ...interface{}) {
	glog.Printf(format, v)
}
func Panicln(v interface{}) {
	glog.Panicln(v)
}
func Panicf(format string, v ...interface{}) {
	glog.Panicf(format, v)
}
