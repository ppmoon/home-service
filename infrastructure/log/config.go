package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var logSugar *zap.SugaredLogger
var logger *zap.Logger

func Init() {
	callerSkip := zap.AddCallerSkip(1)
	caller := zap.AddCaller()
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder, // ISO8601 UTC 时间格式
	}
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout)),
		zap.NewAtomicLevel(),
	)
	//实例化zap
	logger = zap.New(core, caller, callerSkip)
	defer logger.Sync() //
	logSugar = logger.Sugar()
}
