package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Level = zapcore.Level

const (
	InfoLevel   Level = zap.InfoLevel   // 0, default level
	WarnLevel   Level = zap.WarnLevel   // 1
	ErrorLevel  Level = zap.ErrorLevel  // 2
	DPanicLevel Level = zap.DPanicLevel // 3, used in development log
	// PanicLevel logs a message, then panics
	PanicLevel Level = zap.PanicLevel // 4
	// FatalLevel logs a message, then calls os.Exit(1).
	FatalLevel Level = zap.FatalLevel // 5
	DebugLevel Level = zap.DebugLevel // -1
)

var stdLogger = newLogger()

// SyncFunc calls the underlying Core's Sync method, flushing any buffered log
// entries. Applications should take care to call Sync before exiting.
type SyncFunc func()

// InitWithPath init the log module, this should be called at the very beginning of the whole program.
// dir is the dir path to store log file, prefix is log file prefix
// for example InitWithPath("/var", "prod") will create "/var/prod.2021-02-02.log"
// log will sync to std out if stdout is true
func InitWithPath(dir, prefix string, stdout bool) SyncFunc {
	initLog(dir, prefix, InfoLevel, stdout)
	return loggerSync
}

// Init is same as InitWithPath but with a default dir "./"
func Init(prefix string, stdout bool) SyncFunc {
	initLog("./", prefix, InfoLevel, stdout)
	return loggerSync
}

func initLog(dir, prefix string, level Level, stdout bool) {
	stdLogger = NewLogger(dir, prefix, level, stdout)
}

func loggerSync() {
	if stdLogger != nil {
		_ = stdLogger.logger.Sync()
		_ = stdLogger.sugaredLogger.Sync()
	}
}

func StdLogger() *Logger {
	return stdLogger
}

// warp stdLogger.logger
func DebugWithFields(msg string, fields ...Field) {
	stdLogger.logger.Debug(msg, fields...)
}

func InfoWithFields(msg string, fields ...Field) {
	stdLogger.logger.Info(msg, fields...)
}

func WarnWithFields(msg string, fields ...Field) {
	stdLogger.logger.Warn(msg, fields...)
}

func ErrorWithFields(msg string, fields ...Field) {
	stdLogger.logger.Error(msg, fields...)
}
func DPanicWithFields(msg string, fields ...Field) {
	stdLogger.logger.DPanic(msg, fields...)
}
func PanicWithFields(msg string, fields ...Field) {
	stdLogger.logger.Panic(msg, fields...)
}
func FatalWithFields(msg string, fields ...Field) {
	stdLogger.logger.Fatal(msg, fields...)
}

// warp stdLogger.sugaredLogger
func Debug(args ...interface{}) {
	stdLogger.sugaredLogger.Debug(args...)
}

func Debugf(format string, args ...interface{}) {
	stdLogger.sugaredLogger.Debugf(format, args...)
}

func Info(args ...interface{}) {
	stdLogger.sugaredLogger.Info(args...)
}

func Infof(format string, args ...interface{}) {
	stdLogger.sugaredLogger.Infof(format, args...)
}

func Warn(args ...interface{}) {
	stdLogger.sugaredLogger.Warn(args...)
}

func Warnf(format string, args ...interface{}) {
	stdLogger.sugaredLogger.Warnf(format, args...)
}

func Error(args ...interface{}) {
	stdLogger.sugaredLogger.Error(args...)
}

func Errorf(format string, args ...interface{}) {
	stdLogger.sugaredLogger.Errorf(format, args...)
}

// Fatal uses fmt.Sprint to construct and log a message, then calls os.Exit.
func Fatal(args ...interface{}) {
	stdLogger.sugaredLogger.Fatal(args...)
}

// Fatalf uses fmt.Sprintf to log a templated message, then calls os.Exit.
func Fatalf(format string, args ...interface{}) {
	stdLogger.sugaredLogger.Fatalf(format, args...)
}

// Panic uses fmt.Sprint to construct and log a message, then panics.
func Panic(args ...interface{}) {
	stdLogger.sugaredLogger.Panic(args...)
}

// Panicf uses fmt.Sprintf to log a templated message, then panics.
func Panicf(format string, args ...interface{}) {
	stdLogger.sugaredLogger.Panicf(format, args...)
}
