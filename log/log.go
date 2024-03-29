package log

import (
	"fmt"
	"log"

	//"path/filepath"
	_ "reflect"
	"runtime"
	//"strings"
)

type Level int

const (
	DEBUG Level = iota
	INFO
	WARNING
	ERROR
	FATAL
)

var gLevel Level

func SetLevel(l Level) {
	gLevel = l
}

func Fatalf(formating string, args ...interface{}) {
	if gLevel <= FATAL {
		LOG("FATAL", formating, args...)
	}
}

func Errorf(formating string, args ...interface{}) {
	if gLevel <= ERROR {
		LOG("ERROR", formating, args...)
	}
}

func Warningf(formating string, args ...interface{}) {
	if gLevel <= WARNING {
		LOG("WARNING", formating, args...)
	}
}

func Infof(formating string, args ...interface{}) {
	if gLevel <= INFO {
		LOG("INFO", formating, args...)
	}
}

func Debugf(formating string, args ...interface{}) {
	if gLevel <= DEBUG {
		LOG("DEBUG", formating, args...)
	}
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
		//filename := filepath.Base(filename)
	}

	//log.Printf("[%s]\t%s:%s() %d:%s\n", level, filename, funcName, line, fmt.Sprintf(formating, args...))
	log.Printf("[%s] %s:%d:%s\n", level, filename, line, fmt.Sprintf(formating, args...))
}
