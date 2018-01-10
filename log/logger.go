package log

/*
Log format is JSON format by default. But we can change it dynamically.
Inspired by and a subset copy of upspin/log. Using go-kit/log as its default logger

This log library is created because I want a simple JSON logger for my application.
Instead of importing a big log library, this is a more simple log library.
*/

import (
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	kitlog "github.com/go-kit/kit/log"
)

type Level int

const (
	DebugLevel Level = iota
	InfoLevel
	WarnLevel
	ErrorLevel
	FatalLevel
	DisableLevel
)

func stringToLevel(s string) Level {
	switch strings.ToLower(s) {
	case "debug":
		return DebugLevel
	case "info":
		return InfoLevel
	case "warn":
		return WarnLevel
	case "error":
		return ErrorLevel
	case "fatal":
		return FatalLevel
	default:
		return InfoLevel
	}
}

// conver level to string
func levelToString(l Level) string {
	switch l {
	case DebugLevel:
		return "debug"
	case InfoLevel:
		return "info"
	case WarnLevel:
		return "warn"
	case ErrorLevel:
		return "error"
	case FatalLevel:
		return "fatal"
	default:
		return "info"
	}
}

type Format int8

const (
	FmtFormat Format = iota
	JSONFormat
)

type Logger struct {
	// state properties of logger
	level       Level
	levelString string

	// logFormat used to save the current logformat being used
	// this need to tracked as SetOutput will used current logFormat
	// and SetFormat is altering the current logFormat
	logFormat Format

	// defaultLogger will always go to stderr
	// this used to show immediate error when program is running
	defaultLogger kitlog.Logger

	// external logger is used for other use case
	// but usually used to write the log to a file
	externalLogger kitlog.Logger
	externalExists bool
	externalWriter io.Writer

	// fields for withfields
	// this should be used by copying the object of logger
	fields Fields
}

func NewDefaultLogger() *Logger {
	logger := &Logger{
		level:         InfoLevel,
		levelString:   levelToString(InfoLevel),
		defaultLogger: kitlog.NewJSONLogger(os.Stderr),
		logFormat:     JSONFormat,
	}
	return logger
}

// SetLevel to tokologger
// If level is not defined, then level is InfoLevel
func (l *Logger) SetLevel(level interface{}) {
	var lvl Level
	switch level.(type) {
	case Level:
		lvl = level.(Level)
	case string:
		lvl = stringToLevel(level.(string))
	default:
		lvl = InfoLevel
	}
	l.level = lvl
	l.levelString = levelToString(lvl)
}

// SetOutput define where we want to point externalLogger, usually is used for saving log to file
// Double logging is expected if externalLogger/SetOutput is pointed to stderr
func (l *Logger) SetOutput(writer io.Writer) error {
	l.externalExists = true
	l.externalWriter = writer
	l.externalLogger = createNewKitLogger(l.logFormat, writer)
	return nil
}

// SetFormat output of logger
func (l *Logger) SetFormat(format Format) {
	l.logFormat = format
	l.defaultLogger = createNewKitLogger(format, os.Stderr)
	if l.externalExists {
		l.externalLogger = createNewKitLogger(format, l.externalWriter)
	}
}

func createNewKitLogger(format Format, writer io.Writer) kitlog.Logger {
	switch format {
	case JSONFormat:
		return kitlog.NewJSONLogger(writer)
	case FmtFormat:
		return kitlog.NewLogfmtLogger(writer)
	default:
		return kitlog.NewJSONLogger(writer)
	}
}

func (l *Logger) Debug(msg interface{}) {
	l.print(DebugLevel, msg, fieldsToArrayInterface(l.fields)...)
}

func (l *Logger) Debugf(format string, v ...interface{}) {
	l.print(DebugLevel, fmt.Sprintf(format, v...), fieldsToArrayInterface(l.fields)...)
}

func (l *Logger) Print(msg interface{}) {
	l.print(InfoLevel, msg, fieldsToArrayInterface(l.fields)...)
}

func (l *Logger) Printf(format string, v ...interface{}) {
	l.print(InfoLevel, fmt.Sprintf(format, v...), fieldsToArrayInterface(l.fields)...)
}

func (l *Logger) Info(msg interface{}) {
	l.print(InfoLevel, msg, fieldsToArrayInterface(l.fields)...)
}

func (l *Logger) Infof(format string, v ...interface{}) {
	l.print(InfoLevel, fmt.Sprintf(format, v...), fieldsToArrayInterface(l.fields)...)
}

func (l *Logger) Warn(msg interface{}) {
	l.print(WarnLevel, msg, fieldsToArrayInterface(l.fields)...)
}

func (l *Logger) Warnf(format string, v ...interface{}) {
	l.print(WarnLevel, fmt.Sprintf(format, v...), fieldsToArrayInterface(l.fields)...)
}

func (l *Logger) Error(msg interface{}) {
	l.print(ErrorLevel, msg, fieldsToArrayInterface(l.fields)...)
}

func (l *Logger) Errorf(format string, v ...interface{}) {
	l.print(ErrorLevel, fmt.Sprintf(format, v...), fieldsToArrayInterface(l.fields)...)
}

// Errors should be called by using errors package
// errors package have special error fields to add more context in error
func (l *Logger) Errors(err error) {
	var fields map[string]interface{}
	switch err.(type) {
	case fieldsGetter:
		fields = err.(fieldsGetter).GetFields()
	}
	// transform error fields into log fields
	logFields := fieldsToArrayInterface(Fields(fields))
	l.print(ErrorLevel, err.Error(), logFields...)
}

func (l *Logger) Fatal(msg interface{}, Fields ...Fields) {
	l.print(FatalLevel, msg, fieldsToArrayInterface(l.fields)...)
}

func (l *Logger) Fatalf(format string, v ...interface{}) {
	l.print(FatalLevel, fmt.Sprintf(format, v...), fieldsToArrayInterface(l.fields)...)
}

// print will print the actual log, all printer is pointing to this print
// several params is added in this function, like msg, level and time
// os exit is called when its called via FatalLevel
func (l *Logger) print(logLevel Level, msg interface{}, v ...interface{}) {
	if logLevel < l.level {
		return
	}
	v = append(v, "msg", msg)
	v = append(v, "level", levelToString(logLevel))
	v = append(v, "time", time.Now().String())
	l.defaultLogger.Log(v...)
	if l.externalExists {
		l.externalLogger.Log(v...)
	}
	// make sure to exit when fatal
	if logLevel == FatalLevel {
		os.Exit(1)
	}
}

// WithFields provide a functionality to log fields passed to the function
// the functionality is 100% same with logrus.Fields and logrus.WithFields
// the Logger object will be copied and returned as *Logger for further use
func (l *Logger) WithFields(f Fields) *Logger {
	// add fields to copied logger object
	l.fields = f
	return l
}

// fieldsToArrayInterface used to tranfrom fields to []interface
// this is because the go-kit/log receive []interface as parameters
func fieldsToArrayInterface(fields Fields) []interface{} {
	// return if Fields is not exists
	if len(fields) <= 0 {
		return nil
	}

	// always get Fields 0
	fieldsLength := len(fields)
	v := make([]interface{}, fieldsLength*2)
	counter := 0
	for key, value := range fields {
		v[counter] = key
		counter++
		v[counter] = value
		counter++
	}
	return v
}
