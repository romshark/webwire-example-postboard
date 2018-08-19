package logger

import (
	"io"
)

// Logger defines the interface of an API server logger
type Logger interface {
	// ErrorLogWriter returns the error log writer
	ErrorLogWriter() io.Writer

	// DebugLogWriter returns the debug log writer
	DebugLogWriter() io.Writer

	// Err prints an error log
	Err(v ...interface{})

	// Errf prints a formatted error log
	Errf(format string, v ...interface{})

	// Print prints a debug log
	Print(v ...interface{})

	// Printf prints a formatted debug log
	Printf(format string, v ...interface{})

	// Close closes the logger
	Close()
}
