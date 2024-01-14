package appLogs

import (
	"log"
	"os"
	"runtime"
)

type Logger struct {
	InfoLog  *log.Logger
	WarnLog  *log.Logger
	ErrorLog *log.Logger
}

func getFile() *os.File {
	if _, err := os.Stat("/logs"); os.IsNotExist(err) {
		if err := os.MkdirAll("logs", os.ModePerm); err != nil {
			log.Println(err)
			return nil
		}
	}
	f, err := os.OpenFile("logs/log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
		return nil
	}
	return f
}

func (l *Logger) Init() {
	flags := log.LstdFlags
	l.InfoLog = log.New(os.Stdout, "INFO: ", flags)
	l.WarnLog = log.New(os.Stdout, "WARNING: ", flags)
	l.ErrorLog = log.New(os.Stdout, "ERROR: ", flags)
}

func (l *Logger) InfoLogPrint(v ...interface{}) {
	pc, fn, line, _ := runtime.Caller(1)
	l.InfoLog.SetOutput(os.Stdout)
	l.InfoLog.Printf("%v [%s:%s:%d]", v, runtime.FuncForPC(pc).Name(), fn, line)
	f := getFile()
	defer f.Close()
	l.InfoLog.SetOutput(f)
	l.InfoLog.Printf("%v [%s:%s:%d]", v, runtime.FuncForPC(pc).Name(), fn, line)
}

func (l *Logger) WarnLogPrint(v ...interface{}) {
	pc, fn, line, _ := runtime.Caller(1)
	l.WarnLog.SetOutput(os.Stdout)
	l.WarnLog.Printf("%v [%s:%s:%d]", v, runtime.FuncForPC(pc).Name(), fn, line)
	f := getFile()
	defer f.Close()
	l.WarnLog.SetOutput(f)
	l.WarnLog.Printf("%v [%s:%s:%d]", v, runtime.FuncForPC(pc).Name(), fn, line)
}
func (l *Logger) ErrorLogPrint(v ...interface{}) {
	pc, fn, line, _ := runtime.Caller(1)
	l.ErrorLog.SetOutput(os.Stdout)
	l.ErrorLog.Printf("%v [%s:%s:%d]", v, runtime.FuncForPC(pc).Name(), fn, line)
	f := getFile()
	defer f.Close()
	l.ErrorLog.SetOutput(f)
	l.ErrorLog.Printf("%v [%s:%s:%d]", v, runtime.FuncForPC(pc).Name(), fn, line)
}
