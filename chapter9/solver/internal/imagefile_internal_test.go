package solver 

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOpenFile_Errors(t *testing.T) {
	testCases := map[string]struct {
		input string
		err   string
	}{
		"No such file": {
			input: "nosuchfile.png",
			err:   "no such file or directory",
		},
		"not a rgba image": {
			input: "../testdata/rgb.png",
			err:   "Expected RGBA image, got *image.Paletted",
		},
		"can't open existing file": {
			input: "../testdata/norights.png",
			err: "permission denied",
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			img, err := openMaze(tc.input)

			assert.Nil(t, img)
			assert.Error(t, err)
			assert.ErrorContains(t, err, tc.err)
		})
	}
}
