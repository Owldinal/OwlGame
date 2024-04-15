package log

import (
	"github.com/sirupsen/logrus"
	"os"
	"owl-backend/internal/config"
)

var Logger = logrus.New()

func init() {
	level, err := logrus.ParseLevel(config.C.LogLevel)
	if err != nil {
		Logger.Warn("Invalid log level. Using default 'info'")
		level = logrus.InfoLevel
	}
	Logger.SetOutput(os.Stdout)
	Logger.SetLevel(level)
	Logger.SetFormatter(&logrus.TextFormatter{FullTimestamp: true})

	switch config.C.LogOutput {
	case "file":
		logFile, err := os.OpenFile(config.C.LogFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			Logger.Warn("Failed to log to file, using default stderr")
		} else {
			Logger.SetOutput(logFile)
		}
	case "console":
		// 默认情况下 logrus 会输出到 stderr，这里不需要做额外配置
	default:
		Logger.Warn("Unknown log output. Using default stderr.")
	}
}

func Trace(args ...interface{}) {
	Logger.Log(logrus.TraceLevel, args...)
}

func Debug(args ...interface{}) {
	Logger.Log(logrus.DebugLevel, args...)
}

func Info(args ...interface{}) {
	Logger.Log(logrus.InfoLevel, args...)
}

func Warn(args ...interface{}) {
	Logger.Log(logrus.WarnLevel, args...)
}

func Error(args ...interface{}) {
	Logger.Log(logrus.ErrorLevel, args...)
}

func Fatal(args ...interface{}) {
	Logger.Log(logrus.FatalLevel, args...)
	Logger.Exit(1)
}

func Panic(args ...interface{}) {
	Logger.Log(logrus.PanicLevel, args...)
}

func Tracef(format string, args ...interface{}) {
	Logger.Logf(logrus.TraceLevel, format, args...)
}

func Debugf(format string, args ...interface{}) {
	Logger.Logf(logrus.DebugLevel, format, args...)
}

func Infof(format string, args ...interface{}) {
	Logger.Logf(logrus.InfoLevel, format, args...)
}

func Warnf(format string, args ...interface{}) {
	Logger.Logf(logrus.WarnLevel, format, args...)
}

func Errorf(format string, args ...interface{}) {
	Logger.Logf(logrus.ErrorLevel, format, args...)
}

func Fatalf(format string, args ...interface{}) {
	Logger.Logf(logrus.FatalLevel, format, args...)
	Logger.Exit(1)
}

func Panicf(format string, args ...interface{}) {
	Logger.Logf(logrus.PanicLevel, format, args...)
}
