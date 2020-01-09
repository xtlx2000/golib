package log

import (
	"github.com/xtlx2000/golib/rotateLog/mlog"
)

func SetupLog(filename string, maxBytes, number int) {
	mlog.StartEx(mlog.LevelInfo, filename, maxBytes, number)
}

func main() {

	mlog.Info("Hello World !")

	ipsum := "ipsum"
	mlog.Warning("Lorem %s", ipsum)
}

func Infof(format string, a ...interface{}) {
	mlog.Info(format, a...)
}

func Warningf(format string, a ...interface{}) {
	mlog.Warning(format, a...)
}

func Errorf(format string, a ...interface{}) {
	mlog.Errorf(format, a...)
}

func Fatalf(format string, a ...interface{}) {
	mlog.Fatalf(format, a...)
}
