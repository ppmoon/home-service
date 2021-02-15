package log

import "go.uber.org/zap"

func GetZapLogger() *zap.Logger {
	return logger
}
func GetZapLoggerSugar() *zap.SugaredLogger {
	return logSugar
}

func Debug(args ...interface{}) {
	logSugar.Debug(args...)
}

func Info(args ...interface{}) {
	logSugar.Info(args...)
}

func Warn(args ...interface{}) {
	logSugar.Warn(args...)
}

func Error(args ...interface{}) {
	logSugar.Error(args...)
}
func DPanic(args ...interface{}) {
	logSugar.DPanic(args...)
}
func Panic(args ...interface{}) {
	logSugar.Panic(args...)
}
func Fatal(args ...interface{}) {
	logSugar.Fatal(args...)
}
func Debugf(template string, args ...interface{}) {
	logSugar.Debugf(template, args...)
}

func Infof(template string, args ...interface{}) {
	logSugar.Infof(template, args...)
}

func Warnf(template string, args ...interface{}) {
	logSugar.Warnf(template, args...)
}

func Errorf(template string, args ...interface{}) {
	logSugar.Errorf(template, args...)
}

func DPanicf(template string, args ...interface{}) {
	logSugar.DPanicf(template, args...)
}

func Panicf(template string, args ...interface{}) {
	logSugar.Panicf(template, args...)
}
func Fatalf(template string, args ...interface{}) {
	logSugar.Fatalf(template, args...)
}
