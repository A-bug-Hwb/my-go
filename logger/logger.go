// mylog

package logger

import (
	"fmt"
	"os"
	"path"
	"runtime"

	"github.com/sirupsen/logrus"
)

type mylogger struct {
	fp      string
	funcn   string
	logInfo *logrus.Logger
}

var log = &mylogger{
	logInfo: logrus.New(),
}

func init() {
	log.logInfo.SetOutput(os.Stdout)
	log.logInfo.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02T15:04:05",
	})
	//Info("初始化日志配置......")
}

// getCallerInfo 获取调用者的函数名，调用行
func (logger *mylogger) getCallerInfo() {
	// 调用链往上翻三层，找到调用函数的信息
	pc, file, line, ok := runtime.Caller(3)
	if !ok {
		return
	}
	funcName := runtime.FuncForPC(pc).Name()
	logger.funcn = path.Base(funcName)

	_, fileName := path.Split(file)
	logger.fp = fmt.Sprintf("%s:%d", fileName, line)
}

// printer 日志处理函数
func (logger *mylogger) printer(level logrus.Level, msg ...interface{}) {
	log.getCallerInfo()
	logger.logInfo.WithFields(logrus.Fields{"filePath": logger.fp, "func": logger.funcn}).Log(level, msg...)
}

// printerf 格式化日志处理
func (logger *mylogger) printerf(level logrus.Level, format string, msg ...interface{}) {
	log.getCallerInfo()
	logger.logInfo.WithFields(logrus.Fields{"filePath": logger.fp, "func": logger.funcn}).Logf(level, format, msg...)
}

func Info(msg ...interface{}) {
	log.printer(logrus.InfoLevel, msg...)
}

func Infof(format string, msg ...interface{}) {
	log.printerf(logrus.InfoLevel, format, msg...)
}

func Debug(msg ...interface{}) {
	log.printer(logrus.DebugLevel, msg...)
}

func Error(msg ...interface{}) {
	log.printer(logrus.ErrorLevel, msg...)
}

func Errorf(format string, msg ...interface{}) {
	log.printerf(logrus.ErrorLevel, format, msg...)
}
