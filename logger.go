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
	DevEnv = "dev"

	encoderTimeKey   = "time"
	logTimeFormatter = "2006-01-02 15:04:05.000"
)

func NewLogger(dir, prefix string, level Level, stdout bool) *Logger {
	var encoder zapcore.Encoder
	showCaller := true
	if os.Getenv("APP_ENV") == DevEnv {
		encoder = customDevEncoder()
	} else {
		encoder = customProdEncoder()
		showCaller = false
	}
	rotateCfg, _ := rotateLogs.New(
		path.Join(dir, prefix)+".%Y-%m-%d.log",
		rotateLogs.WithMaxAge(24*time.Hour*7),
		rotateLogs.WithRotationTime(1*time.Minute),
	)
	opts := []zapcore.WriteSyncer{
		zapcore.AddSync(rotateCfg),
	}
	if stdout {
		opts = append(opts, zapcore.AddSync(os.Stderr))
	}
	core := zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(opts...), level)

	logger := &Logger{
		logger: zap.New(core, zap.WithCaller(showCaller)),
	}
	logger.sugaredLogger = logger.logger.Sugar()
	return logger
}

func customDevEncoder() zapcore.Encoder {
	customTimeEncoder := func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString("[" + t.Format(logTimeFormatter) + "]")
	}
	customLevelEncoder := func(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString("[" + level.CapitalString() + "]")
	}
	encoderConf := zapcore.EncoderConfig{
		CallerKey:        "caller",
		LevelKey:         "level",
		MessageKey:       "msg",
		TimeKey:          encoderTimeKey,
		StacktraceKey:    "stacktrace",
		LineEnding:       zapcore.DefaultLineEnding,
		EncodeTime:       customTimeEncoder,
		EncodeLevel:      customLevelEncoder,
		EncodeCaller:     zapcore.FullCallerEncoder,
		EncodeDuration:   zapcore.SecondsDurationEncoder,
		ConsoleSeparator: " ",
	}
	return zapcore.NewConsoleEncoder(encoderConf)
}

func customProdEncoder() zapcore.Encoder {
	customTimeEncoder := func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format(logTimeFormatter))
	}
	encoderConf := zapcore.EncoderConfig{
		CallerKey:      "caller",
		LevelKey:       "level",
		MessageKey:     "msg",
		TimeKey:        encoderTimeKey,
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeTime:     customTimeEncoder,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
	}
	return zapcore.NewJSONEncoder(encoderConf)
}

func newLogger() *Logger {
	core := zapcore.NewCore(customProdEncoder(), os.Stderr, InfoLevel)
	l := &Logger{
		logger: zap.New(core, zap.WithCaller(true)),
	}
	l.sugaredLogger = l.logger.Sugar()
	return l
}

func (l *Logger) withFields(fields []Field, fieldsInterface []interface{}) *Logger {
	return &Logger{
		l.logger.With(fields...),
		l.sugaredLogger.With(fieldsInterface...),
	}
}

func (l *Logger) Debug(msg string, fields ...Field) {
	l.logger.Debug(msg, fields...)
}

func (l *Logger) Info(msg string, fields ...Field) {
	l.logger.Info(msg, fields...)
}

func (l *Logger) Warn(msg string, fields ...Field) {
	l.logger.Warn(msg, fields...)
}

func (l *Logger) Error(msg string, fields ...Field) {
	l.logger.Error(msg, fields...)
}
func (l *Logger) DPanic(msg string, fields ...Field) {
	l.logger.DPanic(msg, fields...)
}
func (l *Logger) Panic(msg string, fields ...Field) {
	l.logger.Panic(msg, fields...)
}
func (l *Logger) Fatal(msg string, fields ...Field) {
	l.logger.Fatal(msg, fields...)
}

func (l *Logger) SugaredDebug(args ...interface{}) {
	l.sugaredLogger.Debug(args...)
}

func (l *Logger) SugaredDebugf(format string, args ...interface{}) {
	l.sugaredLogger.Debugf(format, args...)
}

func (l *Logger) SugaredInfo(args ...interface{}) {
	l.sugaredLogger.Info(args...)
}

func (l *Logger) SugaredInfof(format string, args ...interface{}) {
	l.sugaredLogger.Infof(format, args...)
}

func (l *Logger) SugaredWarn(args ...interface{}) {
	l.sugaredLogger.Warn(args...)
}

func (l *Logger) SugaredWarnf(format string, args ...interface{}) {
	l.sugaredLogger.Warnf(format, args...)
}

func (l *Logger) SugaredError(args ...interface{}) {
	l.sugaredLogger.Error(args...)
}

func (l *Logger) SugaredErrorf(format string, args ...interface{}) {
	l.sugaredLogger.Errorf(format, args...)
}

func (l *Logger) SugaredFatal(args ...interface{}) {
	l.sugaredLogger.Fatal(args...)
}

func (l *Logger) SugaredFatalf(format string, args ...interface{}) {
	l.sugaredLogger.Fatalf(format, args...)
}

func (l *Logger) SugaredPanic(args ...interface{}) {
	l.sugaredLogger.Panic(args...)
}

func (l *Logger) SugaredPanicf(format string, args ...interface{}) {
	l.sugaredLogger.Panicf(format, args...)
}
