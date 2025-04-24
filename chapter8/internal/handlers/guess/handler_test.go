package guess

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"learngo/httpgordle/internal/api"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandle(t *testing.T) {
	req, err := http.NewRequest(http.MethodPut, "/games/", strings.NewReader(`{"guess":"pocket"}`))
	require.NoError(t, err)

	req.SetPathValue(api.GameID, "123456")

	recorder := httptest.NewRecorder()

	Handle(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.Equal(t, "application/json", recorder.Header().Get("Content-Type"))
	assert.JSONEq(t, `{"id":"123456","attempts_left":0,"guesses":null,"word_length":0,"status":""}`, recorder.Body.String())
}
