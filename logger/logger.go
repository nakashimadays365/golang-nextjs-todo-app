package logger

import (
	"fmt"
	"log"
	"os"
)

type Logger struct {
	lggr *log.Logger
}

var defLogger = NewLogger()

func NewLogger() Logger {
	return Logger{
		lggr: log.New(os.Stdout, "", log.Lshortfile|log.LstdFlags),
	}
}

func (l Logger) Info(v ...interface{}) {
	l.output("[info] ", v...)
}

func (l Logger) Infof(text string, v ...interface{}) {
	l.outputf("[info] "+text, v...)
}

func (l Logger) Warn(v ...interface{}) {
	l.output("[warn] ", v...)
}

func (l Logger) Warnf(text string, v ...interface{}) {
	l.outputf("[warn] "+text, v...)
}

func (l Logger) Error(v ...interface{}) {
	l.output("[error] ", v...)
}

func (l Logger) Errorf(text string, v ...interface{}) {
	l.outputf("[error] "+text, v...)
}

func (l Logger) output(label string, v ...interface{}) {
	l.lggr.Output(4, label+fmt.Sprintln(v...))
}

func (l Logger) outputf(text string, v ...interface{}) {
	l.lggr.Output(4, fmt.Sprintf(text, v...))
}

func Info(v ...interface{}) {
	defLogger.Info(v...)
}

func Infof(text string, v ...interface{}) {
	defLogger.Infof(text, v...)
}

func Warn(v ...interface{}) {
	defLogger.Warn(v...)
}

func Warnf(text string, v ...interface{}) {
	defLogger.Warnf(text, v...)
}

func Error(v ...interface{}) {
	defLogger.Error(v...)
}

func Errorf(text string, v ...interface{}) {
	defLogger.Errorf(text, v...)
}
