package logger

import (
	"io"
	"log"
)

type LogLevel int

const (
	// debugLog indicates Debug severity.
	debugLog LogLevel = iota
	// infoLog indicates Info severity.
	infoLog
	// warningLog indicates Warning severity.
	warningLog
	// errorLog indicates Error severity.
	errorLog
)

// severityName contains the string representation of each severity.
var severityName = []string{
	debugLog:   "DEBUG",
	infoLog:    "INFO",
	warningLog: "WARNING",
	errorLog:   "ERROR",
}

// Logger does underlying logging work for scaleway-sdk-go.
type Logger interface {
	// Info logs to INFO log. Arguments are handled in the manner of fmt.Print.
	Info(args ...interface{})
	// Infoln logs to INFO log. Arguments are handled in the manner of fmt.Println.
	Infoln(args ...interface{})
	// Infof logs to INFO log. Arguments are handled in the manner of fmt.Printf.
	Infof(format string, args ...interface{})
	// Warning logs to WARNING log. Arguments are handled in the manner of fmt.Print.
	Warning(args ...interface{})
	// Warningln logs to WARNING log. Arguments are handled in the manner of fmt.Println.
	Warningln(args ...interface{})
	// Warningf logs to WARNING log. Arguments are handled in the manner of fmt.Printf.
	Warningf(format string, args ...interface{})
	// Error logs to ERROR log. Arguments are handled in the manner of fmt.Print.
	Error(args ...interface{})
	// Errorln logs to ERROR log. Arguments are handled in the manner of fmt.Println.
	Errorln(args ...interface{})
	// Errorf logs to ERROR log. Arguments are handled in the manner of fmt.Printf.
	Errorf(format string, args ...interface{})
}

// SetLogger sets logger that is used in by the SDK.
// Not mutex-protected, should be called before any scaleway-sdk-go functions.
func SetLogger(l Logger) {
	logger = l
}

// SetLevel create a new default logger with the given log level.
// Not mutex-protected, should be called before any scaleway-sdk-go functions.
func SetLevel(level LogLevel) {
	logger = newLogger(level)
}

// NewLogger creates a logger with the provided writers.
// Error logs will be written to errorW, warningW, infoW and debugW.
// Warning logs will be written to warningW, infoW and debugW.
// Info logs will be written to infoW and debugW.
// Debug logs will be written to debugW.
func NewLogger(debugW, infoW, warningW, errorW io.Writer) Logger {
	var m []*log.Logger
	m = append(m, log.New(debugW, severityName[debugLog]+": ", log.LstdFlags))
	m = append(m, log.New(io.MultiWriter(debugW, infoW), severityName[infoLog]+": ", log.LstdFlags))
	m = append(m, log.New(io.MultiWriter(debugW, infoW, warningW), severityName[warningLog]+": ", log.LstdFlags))
	m = append(m, log.New(io.MultiWriter(debugW, infoW, warningW, errorW), severityName[errorLog]+": ", log.LstdFlags))
	return &loggerT{m: m}
}
