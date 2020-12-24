package logger

import (
	"github.com/apsdehal/go-logger"
	"os"
)

type Logger interface {
 	Fatal(message string)
	FatalF(format string, a ...interface{})
	Fatalf(format string, a ...interface{})
	Panic(message string)
	PanicF(format string, a ...interface{})
	Panicf(format string, a ...interface{})
	Critical(message string)
	CriticalF(format string, a ...interface{})
	Criticalf(format string, a ...interface{})
	Error(message string)
	ErrorF(format string, a ...interface{})
	Errorf(format string, a ...interface{})
	Warning(message string)
	WarningF(format string, a ...interface{})
	Warningf(format string, a ...interface{})
	Notice(message string)
	NoticeF(format string, a ...interface{})
	Noticef(format string, a ...interface{})
	Info(message string) 
	InfoF(format string, a ...interface{})
	Infof(format string, a ...interface{})
	Debug(message string) 
	DebugF(format string, a ...interface{})
	Debugf(format string, a ...interface{})
}

func Create_logger(name string) Logger{

	log, err := logger.New(name, 1, os.Stdout)
	if err != nil {
		panic(err) // Check for error
	}
	return log
}