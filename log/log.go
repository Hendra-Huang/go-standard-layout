package log

import (
	"flag"
)

// Fields of log, this is to hack the logger interface so we don't have to call logger
type Fields map[string]interface{}

type logFlag struct {
	LogLevel string
}

type fieldsGetter interface {
	GetFields() map[string]interface{}
}

func (f *logFlag) Parse(fs *flag.FlagSet, args []string) error {
	fs.StringVar(&f.LogLevel, "log_level", "", "define log level")
	return fs.Parse(args)
}

var defaultLogger *Logger

func init() {
	defaultLogger = NewDefaultLogger()
}

func SetLevel(level interface{}) {
	defaultLogger.SetLevel(level)
}

func Debug(msg interface{}) {
	defaultLogger.Debug(msg)
}

func Debugf(format string, v ...interface{}) {
	defaultLogger.Debugf(format, v...)
}

func Info(msg interface{}) {
	defaultLogger.Info(msg)
}

func Infof(format string, v ...interface{}) {
	defaultLogger.Infof(format, v...)
}

func Warn(msg interface{}) {
	defaultLogger.Warn(msg)
}

func Warnf(format string, v ...interface{}) {
	defaultLogger.Warnf(format, v...)
}

func Error(msg interface{}) {
	defaultLogger.Error(msg)
}

func Errorf(format string, v ...interface{}) {
	defaultLogger.Errorf(format, v...)
}

func Errors(err error) {
	defaultLogger.Errors(err)
}

func Fatal(msg interface{}) {
	defaultLogger.Fatal(msg)
}

func Fatalf(format string, v ...interface{}) {
	defaultLogger.Fatalf(format, v...)
}

func WithFields(f Fields) *Logger {
	return defaultLogger.WithFields(f)
}
