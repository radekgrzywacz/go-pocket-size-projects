package pocketlog

import "io"

type Option func(*Logger)

// WithOutput return a function that sets up the output of the logger.
func WithOutput(output io.Writer) Option {
	return func(l *Logger) {
		l.output = output
	}
}
