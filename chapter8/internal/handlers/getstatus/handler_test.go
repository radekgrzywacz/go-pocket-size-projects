package getstatus

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"learngo/httpgordle/internal/api"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandle(t *testing.T) {
	req, err := http.NewRequest(http.MethodPost, "/games/", nil)
	require.NoError(t, err)

	req.SetPathValue(api.GameID, "123456")

	recorder := httptest.NewRecorder()
	Handle(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.Equal(t, "application/json", recorder.Header().Get("Content-Type"))
	assert.JSONEq(t, `{"id":"123456","attempts_left":0,"guesses":[],"word_length":0,"status":""}`, recorder.Body.String())
}
