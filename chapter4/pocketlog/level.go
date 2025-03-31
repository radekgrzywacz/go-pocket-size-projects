package pocketlog

// Level represents an available logging level.
type Level byte

const (
	// LevelDebug represents the lowest level of log, mostly used for debugging purposes.
	LevelDebug Level = iota
	// LevelInfo represents a logging level that contains information deemed valuable.
	LevelInfo
	// LevelWarn represents a logging level that contains a warning about potential problems.
	LevelWarn
	// LevelError represents the highest logging level, only to be used to trace errors.
	LevelError
	// LevelFatal represents logging level that contains an information about fatal error.
	LevelFatal
)

func (lvl Level) String() string {
	switch lvl {
	case LevelDebug:
		return "[DEBUG]"
	case LevelInfo:
		return "[INFO]"
	case LevelError:
		return "[ERROR]"
	default:
		// Should not happen.
		return ""
	}
}
