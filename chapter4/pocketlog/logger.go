package pocketlog

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

// Logger is used to log information.
type Logger struct {
	threshold        Level
	output           io.Writer
	maxMessageLength uint
}

type Log struct {
	Level   string `json:"level"`
	Message string `json:"message"`
}

// New returns you a logger, ready to log at the required threshold.
// Give it a list of configuration functions to tune it at your will.
// The default output is Stdout.
func New(threshold Level, opts ...Option) *Logger {
	lgr := &Logger{threshold: threshold, output: os.Stdout}

	for _, configFunc := range opts {
		configFunc(lgr)
	}

	return lgr
}

// Debugf formats and prints a message if the log level is debug or higher.
func (l *Logger) Debugf(format string, args ...any) {
	if l.threshold > LevelDebug {
		return
	}

	l.logf(LevelDebug, format, args...)
}

// Infof formats and prints a message if the log level is info or higher.
func (l *Logger) Infof(format string, args ...any) {
	if l.threshold > LevelInfo {
		return
	}

	l.logf(LevelInfo, format, args...)
}

// Errorf formats and prints a message if the log level is error or higher.
func (l *Logger) Errorf(format string, args ...any) {
	if l.threshold > LevelError {
		return
	}

	l.logf(LevelError, format, args...)
}

// Logf formats and prints a message if the log level is high enough
func (l *Logger) Logf(lvl Level, format string, args ...any) {
	if l.threshold > lvl {
		return
	}

	l.logf(lvl, format, args...)
}

// logf prints the message to the output.
// Add decorations here, if any.
func (l *Logger) logf(lvl Level, format string, args ...any) {
	message := fmt.Sprintf(format, args...)

	if l.maxMessageLength != 0 && uint(len([]rune(message))) > l.maxMessageLength {
		message = string([]rune(message)[:l.maxMessageLength]) + "[TRIMMED]"
	}
	log := Log{
		Level:   lvl.String(),
		Message: message,
	}

	logJson, err := json.Marshal(log)
	if err != nil {
		fmt.Println(err)
	}
	_, _ = fmt.Fprintf(l.output, "%s\n", logJson)
}
