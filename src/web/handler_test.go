package web

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"go-rest-project/logger"
)

func TestSendResponse(t *testing.T) {
	respRec := httptest.NewRecorder()
	testBody := "lorem ipsum"

	logEntry := logger.LogEntry("test")
	SendResponse(respRec, http.StatusOK, testBody, logEntry)

	assert.Equal(t, http.StatusOK, respRec.Code)

	body, err := ioutil.ReadAll(respRec.Body)
	assert.NoError(t, err)
	assert.Contains(t, string(body), testBody)
}

func TestSendError(t *testing.T) {
	respRec := httptest.NewRecorder()
	testErr := errors.New("lorem ipsum")

	logEntry := logger.LogEntry("test")
	sendError(respRec, http.StatusInternalServerError, testErr, logEntry)

	assert.Equal(t, http.StatusInternalServerError, respRec.Code)

	body, err := ioutil.ReadAll(respRec.Body)
	assert.NoError(t, err)
	assert.Contains(t, string(body), testErr.Error())
}
