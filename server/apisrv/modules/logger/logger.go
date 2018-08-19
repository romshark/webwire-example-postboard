package logger

import (
	"io"
	"log"
	"os"
)

// logger represents a Logger interface implementation
type logger struct {
	errorLogWriter io.Writer
	debugLogWriter io.Writer
	errorLogger    *log.Logger
	debugLogger    *log.Logger
	errorLogFile   *os.File
	debugLogFile   *os.File
}

// ErrorLogWriter implements the Logger interface
func (l *logger) ErrorLogWriter() io.Writer {
	return l.errorLogWriter
}

// DebugLogWriter implements the Logger interface
func (l *logger) DebugLogWriter() io.Writer {
	return l.debugLogWriter
}

// Err implements the Logger interface
func (l *logger) Err(v ...interface{}) {
	l.errorLogger.Print(v...)
}

// Errf implements the Logger interface
func (l *logger) Errf(format string, v ...interface{}) {
	l.errorLogger.Printf(format, v...)
}

// Print implements the Logger interface
func (l *logger) Print(v ...interface{}) {
	l.debugLogger.Print(v...)
}

// Printf implements the Logger interface
func (l *logger) Printf(format string, v ...interface{}) {
	l.debugLogger.Printf(format, v...)
}

// Close implements the Logger interface
func (l *logger) Close() {
	if l.errorLogFile != nil {
		l.errorLogFile.Close()
	}
	if l.debugLogFile != nil {
		l.debugLogFile.Close()
	}
}
