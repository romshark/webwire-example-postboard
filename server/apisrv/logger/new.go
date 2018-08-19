package logger

import (
	"log"
	"os"

	"github.com/pkg/errors"
	"github.com/qbeon/webwire-messenger/server/apisrv/config"
)

// New creates a new API server logger
func New(conf *config.Config) (Logger, error) {
	// Use standard output buffers by default
	errorLogWriter := os.Stderr
	debugLogWriter := os.Stdout

	var errorLogFile, debugLogFile *os.File

	if conf.Log != nil {
		var err error
		logFileFlags := os.O_RDWR | os.O_CREATE | os.O_APPEND
		logFileMode := os.FileMode(0666)

		if conf.Log.ErrorLogFilePath != "" {
			// Write error logs to file
			errorLogFile, err = os.OpenFile(
				conf.Log.ErrorLogFilePath,
				logFileFlags,
				logFileMode,
			)
			if err != nil {
				return nil, errors.Wrapf(
					err, "couldn't open error log file (%s)",
					conf.Log.ErrorLogFilePath,
				)
			}
			errorLogWriter = errorLogFile
		}

		if conf.Log.DebugLogFilePath != "" {
			// Write debug logs to file
			debugLogFile, err = os.OpenFile(
				conf.Log.ErrorLogFilePath,
				logFileFlags,
				logFileMode,
			)
			if err != nil {
				return nil, errors.Wrapf(
					err,
					"couldn't open debug log file (%s)",
					conf.Log.DebugLogFilePath,
				)
			}
			debugLogWriter = debugLogFile
		}
	}

	// Initialize loggers
	return &logger{
		errorLogWriter: errorLogWriter,
		debugLogWriter: debugLogWriter,
		errorLogFile:   errorLogFile,
		debugLogFile:   debugLogFile,
		errorLogger: log.New(
			errorLogWriter,
			"ERROR: ",
			log.Ldate|log.Ltime|log.Lshortfile,
		),
		debugLogger: log.New(
			debugLogWriter,
			"DEBUG: ",
			log.Ldate|log.Ltime|log.Lshortfile,
		),
	}, nil
}
