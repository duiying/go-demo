package logger

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"path"
	"time"
)

var levelMap = map[logrus.Level]string{
	logrus.DebugLevel: "debug",
	logrus.InfoLevel: "info",
	logrus.WarnLevel: "warning",
	logrus.FatalLevel: "fatal",
	logrus.ErrorLevel: "error",
	logrus.PanicLevel: "panic",
	logrus.TraceLevel: "trace",
}

// Debug debug
func Debug(info interface{}) {
	logToFile(logrus.DebugLevel, info)
	return
}

// Info info
func Info(info interface{}) {
	logToFile(logrus.InfoLevel, info)
	return
}

// Warn warn
func Warn(info interface{}) {
	logToFile(logrus.WarnLevel, info)
	return
}

// Fatal fatal
func Fatal(info interface{}) {
	logToFile(logrus.FatalLevel, info)
	return
}

// Error error
func Error(info interface{}) {
	logToFile(logrus.ErrorLevel, info)
	return
}

// Panic panic
func Panic(info interface{}) {
	logToFile(logrus.PanicLevel, info)
	return
}

// 记录日志到文件
func logToFile(level logrus.Level, something interface{}) {
	var info string
	switch value := something.(type) {
	case string:
		info = value
	case map[string]interface{}:
		jsonStr, err := json.Marshal(value)
		if err != nil {
			return
		}
		info = string(jsonStr)
	default:
		return
	}
	// 日志文件
	logFilePath := os.Getenv("LOG_PATH")
	today := time.Now().Format("20060102")
	logFileName := fmt.Sprintf("%s.log.%s", levelMap[level], today)
	fileName := path.Join(logFilePath, logFileName)
	// 日志文件不存在时创建
	isExist, _ := pathExists(fileName)
	if !isExist {
		f, err := os.Create(fileName)
		if err != nil {
			return
		}
		defer f.Close()
	}
	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Printf("Error while logger: %v \n", err)
		return
	}

	// 实例
	logger := logrus.New()
	logger.Out = src
	logger.Formatter = &logrus.JSONFormatter{}
	logger.SetLevel(logrus.DebugLevel)

	// 记录日志
	app := os.Getenv("APP_NAME")
	field := logrus.Fields{
		"app": app,
	}
	switch level {
	case logrus.DebugLevel:
		logger.WithFields(field).Debug(info)
	case logrus.InfoLevel:
		logger.WithFields(field).Info(info)
	case logrus.WarnLevel:
		logger.WithFields(field).Warn(info)
	case logrus.FatalLevel:
		logger.WithFields(field).Fatal(info)
	case logrus.ErrorLevel:
		logger.WithFields(field).Error(info)
	case logrus.PanicLevel:
		logger.WithFields(field).Panic(info)
	}
	return
}

// 判断文件是否存在
func pathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}