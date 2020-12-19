package log

import (
	"github.com/sirupsen/logrus"
	"os"
)

func init() {
	logrus.SetFormatter(&logrus.TextFormatter{})
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.WarnLevel)
	logrus.SetReportCaller(true)
}

func Info(args ...interface{}) {
	logrus.Info(args)
}

func Error(args ...interface{}) {
	logrus.Error(args)
}

func Warn(args ...interface{}) {
	logrus.Warn(args)
}
