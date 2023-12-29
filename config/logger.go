package config

import (
	"io"
	"log"
	"os"
)

type Logger struct {
	debug   *log.Logger
	info    *log.Logger
	warning *log.Logger
	err     *log.Logger
	writer  *io.Writer
}

func NewLogger(prefix string) *Logger {
	writer := io.Writer(os.Stdout)
	logger := log.New(writer, prefix, log.Ldate|log.Ltime)

	return &Logger{
		debug:   log.New(writer, "DEBUG: ", logger.Flags()),
		info:    log.New(writer, "INFO: ", logger.Flags()),
		warning: log.New(writer, "WARNING: ", logger.Flags()),
		err:     log.New(writer, "ERROR: ", logger.Flags()),
		writer:  &writer,
	}
}

func (l *Logger) Debug(value ...interface{}) {
	l.debug.Println(value...)
}

func (l *Logger) Info(value ...interface{}) {
	l.info.Println(value...)
}

func (l *Logger) Warnig(value ...interface{}) {
	l.warning.Println(value...)
}

func (l *Logger) Error(value ...interface{}) {
	l.err.Println(value...)
}

func (l *Logger) DebugF(format string, value ...interface{}) {
	l.debug.Printf(format, value...)
}

func (l *Logger) InfoF(format string, value ...interface{}) {
	l.info.Printf(format, value...)
}

func (l *Logger) WarningF(format string, value ...interface{}) {
	l.warning.Printf(format, value...)
}

func (l *Logger) ErrorF(format string, value ...interface{}) {
	l.err.Printf(format, value...)
}
