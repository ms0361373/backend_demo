package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func NewRolling(name string) *zap.Logger {
	wFile := zapcore.AddSync(&lumberjack.Logger{
		Filename:   name,
		MaxSize:    10, // megabytes
		MaxBackups: 10,
		MaxAge:     28, // days
	})
	wStdout := zapcore.AddSync(os.Stdout)

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(zapcore.EncoderConfig{
			TimeKey:        "ts",
			LevelKey:       "level",
			NameKey:        "logger",
			CallerKey:      "caller",
			MessageKey:     "msg",
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.LowercaseLevelEncoder,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		}),
		zapcore.NewMultiWriteSyncer(wFile, wStdout),
		zap.NewAtomicLevel(),
	)
	return zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel))
}
