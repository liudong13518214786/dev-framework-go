package logging

import (
	"fmt"
	"go-gin-example/pkg/file"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

//日志模块
var (
	f         *os.File
	logger    *log.Logger
	logPrefix = ""
	loglevel  = []string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL"}
)

type Level int

const (
	DEBUG = iota
	INFO
	WARN
	ERROR
	FATAL
)

func LogoInit() {
	var err error
	f, err = file.MustOpen("log", "log/")
	if err != nil {
		log.Fatal("log init err:", err)
	}
	logger = log.New(f, "", 1)
}

func Debug(v interface{}) {
	SetPrefix(DEBUG)
	logger.Println(v)
}

func Info(v interface{}) {
	SetPrefix(INFO)
	logger.Println(v)
}

func Warn(v interface{}) {
	SetPrefix(WARN)
	logger.Println(v)
}

func Error(v interface{}) {
	SetPrefix(ERROR)
	logger.Println(v)
}

func Fatal(v interface{}) {
	SetPrefix(FATAL)
	logger.Println(v)
}

func SetPrefix(level Level) {
	_, fileNmae, line, ok := runtime.Caller(2)
	if ok {
		logPrefix = fmt.Sprintf("[%s][%s:%d]", loglevel[level], filepath.Base(fileNmae), line)
	} else {
		logPrefix = fmt.Sprintf("[%s]", loglevel[level])
	}
	logger.SetPrefix(logPrefix)
}
