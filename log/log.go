package log

import (
	"fmt"
	"log"
	"path/filepath"
	_ "reflect"
	"runtime"
	//"strings"
)

func Fatalf(formating string, args ...interface{}) {
	LOG("FATAL", formating, args...)
}

func Errorf(formating string, args ...interface{}) {
	LOG("ERROR", formating, args...)
}

func Warningf(formating string, args ...interface{}) {
	LOG("WARNING", formating, args...)
}

func Infof(formating string, args ...interface{}) {
	LOG("INFO", formating, args...)
}

func Debugf(formating string, args ...interface{}) {
	LOG("DEBUG", formating, args...)
}

func LOG(level string, formating string, args ...interface{}) {
	filename, line, _ := "???", 0, "???"
	_, filename, line, ok := runtime.Caller(2)
	if ok {
		/* funcName = runtime.FuncForPC(pc).Name()
		funcName = filepath.Ext(funcName)
		funcName = strings.TrimPrefix(funcName, ".")
		if funcName == "0" {
			funcName = "init"
		}
		*/
		filename = filepath.Base(filename)
	}

	//log.Printf("[%s]\t%s:%s() %d:%s\n", level, filename, funcName, line, fmt.Sprintf(formating, args...))
	log.Printf("[%s] %s:%d:%s\n", level, filename, line, fmt.Sprintf(formating, args...))
}
