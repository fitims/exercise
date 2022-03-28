package log

import (
	"io"
	golog "log"
)

var (
	prefix = ""
)

func SetPrefix(p string) {
	prefix = p
}

func getPrefix(p string) string {
	if len(prefix) == 0 {
		return "[" + p + "] : "
	}
	return "[" + prefix + ":" + p + "] : "
}

func getInfoPrefix() string {
	return getPrefix("INFO")
}

func getWarnPrefix() string {
	return getPrefix("WARN")
}

func getErrorPrefix() string {
	return getPrefix("ERROR")
}

func getFatalPrefix() string {
	return getPrefix("FATAL")
}

func getPanicPrefix() string {
	return getPrefix("PANIC")
}

func Info(v ...interface{}) {
	golog.SetPrefix(getInfoPrefix())
	golog.Print(v...)
}

func Infoln(v ...interface{}) {
	golog.SetPrefix(getInfoPrefix())
	golog.Println(v...)
}

func Infof(format string, v ...interface{}) {
	golog.SetPrefix(getInfoPrefix())
	golog.Printf(format, v...)
}

func Warn(v ...interface{}) {
	golog.SetPrefix(getWarnPrefix())
	golog.Print(v...)
}

func Warnln(v ...interface{}) {
	golog.SetPrefix(getWarnPrefix())
	golog.Println(v...)
}

func Warnf(format string, v ...interface{}) {
	golog.SetPrefix(getWarnPrefix())
	golog.Printf(format, v...)
}

func Error(v ...interface{}) {
	golog.SetPrefix(getErrorPrefix())
	golog.Print(v...)
}

func Errorln(v ...interface{}) {
	golog.SetPrefix(getErrorPrefix())
	golog.Println(v...)
}

func Errorf(format string, v ...interface{}) {
	golog.SetPrefix(getErrorPrefix())
	golog.Printf(format, v...)
}

func Fatal(v ...interface{}) {
	golog.SetPrefix(getFatalPrefix())
	golog.Fatal(v...)
}

func Fatalf(format string, v ...interface{}) {
	golog.SetPrefix(getFatalPrefix())
	golog.Fatalf(format, v...)
}

func Fatalln(v ...interface{}) {
	golog.SetPrefix(getFatalPrefix())
	golog.Fatalln(v...)
}

func SetOutput(w io.Writer) {
	golog.SetOutput(w)
}

func Panic(v ...interface{}) {
	golog.SetPrefix(getPanicPrefix())
	golog.Panic(v)
}

func Panicf(format string, v ...interface{}) {
	golog.SetPrefix(getPanicPrefix())
	golog.Panicf(format, v)
}

func Panicln(v ...interface{}) {
	golog.SetPrefix(getPanicPrefix())
	golog.Panicln(v)
}
