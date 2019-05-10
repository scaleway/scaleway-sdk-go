package logger

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

var logger = newLogger(LogLevelWarning)

// Info logs to the INFO log.
func Info(args ...interface{}) { logger.Info(args...) }
func (g *loggerT) Info(args ...interface{}) {
	g.m[LogLevelInfo].Print(args...)
}

// Infof logs to the INFO log. Arguments are handled in the manner of fmt.Printf.
func Infof(format string, args ...interface{}) { logger.Infof(format, args...) }
func (g *loggerT) Infof(format string, args ...interface{}) {
	g.m[LogLevelInfo].Printf(format, args...)
}

// Infoln logs to the INFO log. Arguments are handled in the manner of fmt.Println.
func Infoln(args ...interface{}) { logger.Infoln(args...) }
func (g *loggerT) Infoln(args ...interface{}) {
	g.m[LogLevelInfo].Println(args...)
}

// Warning logs to the WARNING log.
func Warning(args ...interface{}) { logger.Warning(args...) }
func (g *loggerT) Warning(args ...interface{}) {
	g.m[LogLevelWarning].Print(args...)
}

// Warningf logs to the WARNING log. Arguments are handled in the manner of fmt.Printf.
func Warningf(format string, args ...interface{}) { logger.Warningf(format, args...) }
func (g *loggerT) Warningf(format string, args ...interface{}) {
	g.m[LogLevelWarning].Printf(format, args...)
}

// Warningln logs to the WARNING log. Arguments are handled in the manner of fmt.Println.
func Warningln(args ...interface{}) { logger.Warningln(args...) }
func (g *loggerT) Warningln(args ...interface{}) {
	g.m[LogLevelWarning].Println(args...)
}

// Error logs to the ERROR log.
func Error(args ...interface{}) { logger.Error(args...) }
func (g *loggerT) Error(args ...interface{}) {
	g.m[LogLevelError].Print(args...)
}

// Errorf logs to the ERROR log. Arguments are handled in the manner of fmt.Printf.
func Errorf(format string, args ...interface{}) { logger.Errorf(format, args...) }
func (g *loggerT) Errorf(format string, args ...interface{}) {
	g.m[LogLevelError].Printf(format, args...)
}

// Errorln logs to the ERROR log. Arguments are handled in the manner of fmt.Println.
func Errorln(args ...interface{}) { logger.Errorln(args...) }
func (g *loggerT) Errorln(args ...interface{}) {
	g.m[LogLevelError].Println(args...)
}

// loggerT is the default logger used by scaleway-sdk-go.
type loggerT struct {
	m []*log.Logger
}

// newLogger creates a logger to be used as default logger.
// All logs are written to stderr.
func newLogger(level LogLevel) Logger {
	errorW := ioutil.Discard
	warningW := ioutil.Discard
	infoW := ioutil.Discard
	debugW := ioutil.Discard
	if isEnabled("SCW_DEBUG") {
		level = LogLevelDebug
	}
	switch level {
	case LogLevelDebug:
		debugW = os.Stderr
	case LogLevelInfo:
		infoW = os.Stderr
	case LogLevelWarning:
		warningW = os.Stderr
	case LogLevelError:
		errorW = os.Stderr
	}

	return NewLogger(debugW, infoW, warningW, errorW)
}

func isEnabled(envKey string) bool {
	env, exist := os.LookupEnv(envKey)
	if !exist {
		return false
	}

	value, err := strconv.ParseBool(env)
	if err != nil {
		fmt.Fprintf(os.Stderr, "environment variable %s has invalid boolean value\n", envKey)
	}

	return value
}
