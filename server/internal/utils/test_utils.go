package utils

import (
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

type ErrorReader struct {
}

func (er *ErrorReader) Read(p []byte) (n int, err error) {
	return 0, errors.New("read error")
}

func ExecuteRequest(t *testing.T, method, path string, body io.Reader, handler http.HandlerFunc) *httptest.ResponseRecorder {
	req, err := http.NewRequest(method, path, body)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	handler.ServeHTTP(recorder, req)

	return recorder
}
