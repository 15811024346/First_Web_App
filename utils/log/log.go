package log

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path"
	"strings"
	"time"
)

//定义BUG级别的常量
const (
	DebugLevel = iota
	ReleaseLevel
	WaringLevel
	ErrorLevel
	FatalLevel
)
const (
	printDebugLevel   = "[debug  ] "
	printReleaseLevel = "[release] "
	printWaringLevel  = "[waring]"
	printErrorLevel   = "[error  ] "
	printFatalLevel   = "[fatal  ] "
)

//定义一个logger结构体

type Logger struct {
	level      int
	baseLogger *log.Logger
	baseFile   *os.File
}

//给日志创建一个初始化方法
func New(strLevel string, pathName string, flag int) (*Logger, error) {
	//level
	var level int
	switch strings.ToLower(strLevel) {
	case "debug":
		level = DebugLevel
	case "release":
		level = ReleaseLevel
	case "waring":
		level = WaringLevel
	case "error":
		level = ErrorLevel
	case "fatal":
		level = FatalLevel
	default:
		return nil, errors.New("unknown level :" + strLevel)
	}
	//logger  重新声明两个变量 ，用来接收，log的指针类型值，和文件路径。
	var baseLogger *log.Logger
	var baseFile *os.File

	if pathName != "" {
		now := time.Now()
		filename := fmt.Sprintf("%d%02d%02d_%02d_%02d_%02d.log", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
		file, err := os.Create(path.Join(pathName, filename))
		if err != nil {
			return nil, err
		}
		baseLogger = log.New(file, "", flag)
		baseFile = file
	} else {
		baseLogger = log.New(os.Stdout, "", flag)
	}

	logger := new(Logger)
	logger.level = level
	logger.baseFile = baseFile
	logger.baseLogger = baseLogger
	return logger, nil
}

//log日志的close方法

func (l *Logger) Close() {
	if l.baseFile != nil {
		l.baseFile.Close()
	}
	l.baseLogger = nil
	l.baseFile = nil
}

//写一个log吐出的方法
func (l *Logger) doPrintf(level int, printLevel string, format string, a ...interface{}) {
	if level < l.level {
		return
	}
	if l.baseLogger == nil {
		panic("log closed")
	}
	format = printLevel + format
	l.baseLogger.Output(3, fmt.Sprintf(format, a...))
	if level == FatalLevel {
		os.Exit(1)
	}
}

//给Logger结构体制造方法
func (l *Logger) Debug(format string, a ...interface{}) {
	l.doPrintf(DebugLevel, printDebugLevel, format, a...)
}
func (l *Logger) Release(format string, a ...interface{}) {
	l.doPrintf(ReleaseLevel, printReleaseLevel, format, a...)
}
func (l *Logger) Waring(format string, a ...interface{}) {
	l.doPrintf(WaringLevel, printWaringLevel, format, a...)
}
func (l *Logger) Error(format string, a ...interface{}) {
	l.doPrintf(ErrorLevel, printErrorLevel, format, a...)
}
func (l *Logger) Fatal(format string, a ...interface{}) {
	l.doPrintf(FatalLevel, printFatalLevel, format, a...)
}

var gLogger, _ = New("debug", "", log.LstdFlags)

func Export(l *Logger) {
	if l != nil {
		gLogger = l
	}
}

func Debug(format string, a ...interface{}) {
	gLogger.doPrintf(DebugLevel, printDebugLevel, format, a...)
}

func Release(format string, a ...interface{}) {
	gLogger.doPrintf(ReleaseLevel, printReleaseLevel, format, a...)
}

func Waring(format string, a ...interface{}) {
	gLogger.doPrintf(WaringLevel, printWaringLevel, format, a...)
}

func Error(format string, a ...interface{}) {
	gLogger.doPrintf(ErrorLevel, printErrorLevel, format, a...)
}

func Fatal(format string, a ...interface{}) {
	gLogger.doPrintf(FatalLevel, printFatalLevel, format, a...)
}
func Close() {
	gLogger.Close()
}
