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

var (
	Info   = stdLogger.Info
	Warn   = stdLogger.Warn
	Error  = stdLogger.Error
	DPanic = stdLogger.DPanic
	Panic  = stdLogger.Panic
	Fatal  = stdLogger.Fatal
	Debug  = stdLogger.Debug

	SugaredInfo   = stdLogger.SugaredInfo
	SugaredInfof  = stdLogger.SugaredInfof
	SugaredWarn   = stdLogger.SugaredWarn
	SugaredWarnf  = stdLogger.SugaredWarnf
	SugaredError  = stdLogger.SugaredError
	SugaredErrorf = stdLogger.SugaredErrorf
	SugaredPanic  = stdLogger.SugaredPanic
	SugaredPanicf = stdLogger.SugaredPanicf
	SugaredFatal  = stdLogger.SugaredFatal
	SugaredFatalf = stdLogger.SugaredFatal
	SugaredDebug  = stdLogger.SugaredDebug
	SugaredDebugf = stdLogger.SugaredDebugf
)

// not safe for concurrent use
func resetStdLogger(l *Logger) {
	stdLogger = l
	Info = stdLogger.Info
	Warn = stdLogger.Warn
	Error = stdLogger.Error
	DPanic = stdLogger.DPanic
	Panic = stdLogger.Panic
	Fatal = stdLogger.Fatal
	Debug = stdLogger.Debug

	SugaredInfo = stdLogger.SugaredInfo
	SugaredInfof = stdLogger.SugaredInfof
	SugaredWarn = stdLogger.SugaredWarn
	SugaredWarnf = stdLogger.SugaredWarnf
	SugaredError = stdLogger.SugaredError
	SugaredErrorf = stdLogger.SugaredErrorf
	SugaredPanic = stdLogger.SugaredPanic
	SugaredPanicf = stdLogger.SugaredPanicf
	SugaredFatal = stdLogger.SugaredFatal
	SugaredFatalf = stdLogger.SugaredFatal
	SugaredDebug = stdLogger.SugaredDebug
	SugaredDebugf = stdLogger.SugaredDebugf
}

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
	return sync
}

// Init is same as InitWithPath but with a default dir "./"
func Init(prefix string, stdout bool) SyncFunc {
	initLog("./", prefix, InfoLevel, stdout)
	return sync
}

func initLog(dir, prefix string, level Level, stdout bool) {
	resetStdLogger(NewLogger(dir, prefix, level, stdout))
}

func sync() {
	if stdLogger != nil {
		_ = stdLogger.logger.Sync()
	}
}

func StdLogger() *Logger {
	return stdLogger
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
