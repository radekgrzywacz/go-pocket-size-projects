package pocketlog_test

import (
	"logger/pocketlog"
	"testing"
)

const (
	debugMessage = "Why write I still all one, ever the same,"
	infoMessage  = "And keep invention in a noted weed,"
	errorMessage = "That every word doth almost tell my name,"
)

func TestLogger_DebugInfoError(t *testing.T) {
	tt := map[string]struct {
		level    pocketlog.Level
		expected string
	}{
		"debug": {
			level: pocketlog.LevelDebug,
			expected: `{"level":"[DEBUG]","message":"` + debugMessage + "\"}\n" +
				`{"level":"[INFO]","message":"` + infoMessage + "\"}\n" +
				`{"level":"[ERROR]","message":"` + errorMessage + "\"}\n",
		},
		"info": {
			level: pocketlog.LevelInfo,
			expected: `{"level":"[INFO]","message":"` + infoMessage + "\"}\n" +
				`{"level":"[ERROR]","message":"` + errorMessage + "\"}\n",
		},
		"error": {
			level:    pocketlog.LevelError,
			expected: `{"level":"[ERROR]","message":"` + errorMessage + "\"}\n",
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			tw := &testWriter{}

			testedLogger := pocketlog.New(tc.level, pocketlog.WithOutput(tw))

			testedLogger.Debugf(debugMessage)
			testedLogger.Infof(infoMessage)
			testedLogger.Logf(pocketlog.LevelError, errorMessage)

			if tw.contents != tc.expected {
				t.Errorf("invalid contents, expected %q, got %q", tc.expected, tw.contents)
			}
		})
	}
}

func ExampleLogger_Debugf() {
	debugLogger := pocketlog.New(pocketlog.LevelDebug)
	debugLogger.Debugf("Hello, %s", "world")
	// Output: Hello, world
}

// testWriter isa struct that implements io.Writer.
// We use it to validate that we can write to the specific output.
type testWriter struct {
	contents string
}

// Write implements the io.Writer interface
func (tw *testWriter) Write(p []byte) (n int, err error) {
	tw.contents = tw.contents + string(p)
	return len(p), nil
}
