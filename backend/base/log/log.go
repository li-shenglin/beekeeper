package log

import (
	"os"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

var (
	lg = logrus.New()
)

func GetLog() *logrus.Logger {
	return lg
}

func Config(logFile string, level logrus.Level) {
	lg.SetReportCaller(true)
	if logFile != "" {
		file, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644) //创建一个log日志文件
		if err != nil {
			panic(err)
		}
		fileName := file.Name()
		lg.Out = file
		logWriter, err := rotatelogs.New(
			fileName+".%Y%m%d.log",
			rotatelogs.WithLinkName(fileName),
			rotatelogs.WithMaxAge(7*24*time.Hour),
			rotatelogs.WithRotationTime(1*time.Hour),
		)
		if err != nil {
			panic(err)
		}
		writerMap := lfshook.WriterMap{
			logrus.InfoLevel:  logWriter,
			logrus.FatalLevel: logWriter,
			logrus.DebugLevel: logWriter,
			logrus.WarnLevel:  logWriter,
			logrus.ErrorLevel: logWriter,
			logrus.PanicLevel: logWriter,
		}
		lg.AddHook(lfshook.NewHook(writerMap, &logrus.JSONFormatter{
			TimestampFormat: "2006-01-02 15:04:05",
		}))
	} else {
		lg.Out = os.Stdout
		lg.SetFormatter(&logrus.TextFormatter{
			FullTimestamp:   true,
			TimestampFormat: "2006-01-02 15:04:05",
		})
	}

	lg.SetLevel(level)
	lg.Infoln("log config success")
}
