package log

import (
	"os"
	"path"
	"time"

	rotateLogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	logger        *zap.Logger
	sugaredLogger *zap.SugaredLogger
}

const (
	encoderTimeKey   = "time"
	logTimeFormatter = "2006-01-02 15:04:05.000"
)

func NewLogger(dir, prefix string, level Level, stdout bool) *Logger {
	cfg := zap.NewProductionConfig()
	cfg.EncoderConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format(logTimeFormatter))
	}
	cfg.EncoderConfig.TimeKey = encoderTimeKey

	logf, _ := rotateLogs.New(
		path.Join(dir, prefix)+".%Y-%m-%d.log",
		rotateLogs.WithMaxAge(24*time.Hour*7),
		rotateLogs.WithRotationTime(1*time.Minute),
	)
	core := make([]zapcore.Core, 0, 2)
	core = append(core, zapcore.NewCore(
		zapcore.NewJSONEncoder(cfg.EncoderConfig),
		zapcore.AddSync(logf),
		level,
	))
	if stdout {
		core = append(core, zapcore.NewCore(
			zapcore.NewJSONEncoder(cfg.EncoderConfig),
			os.Stderr,
			level,
		))
	}

	logger := &Logger{
		logger: zap.New(zapcore.NewTee(core...), zap.WithCaller(true)),
	}
	logger.sugaredLogger = logger.logger.Sugar()
	return logger
}

func newLogger() *Logger {
	cfg := zap.NewProductionConfig()
	cfg.EncoderConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format(logTimeFormatter))
	}
	cfg.EncoderConfig.TimeKey = encoderTimeKey
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(cfg.EncoderConfig),
		os.Stderr,
		InfoLevel,
	)
	l := &Logger{
		logger: zap.New(core, zap.WithCaller(true)),
	}
	l.sugaredLogger = l.logger.Sugar()
	return l
}
