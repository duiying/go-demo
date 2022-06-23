package logger

import (
	"fmt"
	"github.com/duiying/go-demo/pkg/config"
	jsoniter "github.com/json-iterator/go"
	"github.com/sirupsen/logrus"
	"github.com/xiam/to"
	"io"
	"os"
	"path"
	"time"
)

var levelMap = map[logrus.Level]string{
	logrus.DebugLevel: "debug",
	logrus.InfoLevel:  "info",
	logrus.WarnLevel:  "warning",
	logrus.FatalLevel: "fatal",
	logrus.ErrorLevel: "error",
	logrus.PanicLevel: "panic",
	logrus.TraceLevel: "trace",
}

func Debug(msg string, keysAndValues ...interface{}) {
	logToFile(logrus.DebugLevel, msg, keysAndValues...)
	return
}

func Info(msg string, keysAndValues ...interface{}) {
	logToFile(logrus.InfoLevel, msg, keysAndValues...)
	return
}

func Warn(msg string, keysAndValues ...interface{}) {
	logToFile(logrus.WarnLevel, msg, keysAndValues...)
	return
}

func Fatal(msg string, keysAndValues ...interface{}) {
	logToFile(logrus.FatalLevel, msg, keysAndValues...)
	return
}

func Error(msg string, keysAndValues ...interface{}) {
	logToFile(logrus.ErrorLevel, msg, keysAndValues...)
	return
}

func Panic(msg string, keysAndValues ...interface{}) {
	logToFile(logrus.PanicLevel, msg, keysAndValues...)
	return
}

// 记录日志到文件
func logToFile(level logrus.Level, msg string, keysAndValues ...interface{}) {
	var info string
	var err error
	var out io.Writer

	if config.Debug { // Debug 模式下输出到控制台
		out = os.Stdout
	} else { // 其他模式下输出到日志文件
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
			defer func() {
				_ = f.Close()
			}()
		}
		out, err = os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		if err != nil {
			fmt.Printf("Error OpenFile logger: %v \n", err)
			return
		}
	}

	// 实例
	logger := logrus.New()
	logger.Out = out
	logger.Formatter = &logrus.JSONFormatter{}
	logger.SetLevel(logrus.DebugLevel)

	// 组装日志内容
	field := logrus.Fields{
		"app": os.Getenv("APP_NAME"),
		"tag": msg,
	}
	var lastKey string
	for k, v := range keysAndValues {
		if k%2 == 0 {
			lastKey = toString(v)
			if len(lastKey) < 1 {
				return
			}
		} else {
			field[lastKey] = toString(v)
		}
	}

	switch level {
	case logrus.DebugLevel:
		logger.WithFields(field).Debug()
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

func toString(something interface{}) string {
	switch value := something.(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64, complex128, complex64, bool, string, []byte:
		return to.String(value)
	case error:
		return value.Error()
	default:
		info, err := jsoniter.MarshalToString(value)
		if err != nil {
			return ""
		}
		return info
	}
}
