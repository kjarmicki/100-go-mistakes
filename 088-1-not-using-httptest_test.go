package main

import (
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// testing an http handler

func ApiHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("X-API-VERSION", "1.0")
	b, _ := io.ReadAll(r.Body)
	_, _ = w.Write(append([]byte("hello "), b...))
	w.WriteHeader(http.StatusOK)
}

func TestApiHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://localhost", strings.NewReader("foo"))
	w := httptest.NewRecorder()
	ApiHandler(w, req)
	body, _ := ioutil.ReadAll(w.Result().Body)

	assert.Equal(t, http.StatusOK, w.Result().StatusCode)
	assert.Equal(t, "1.0", w.Result().Header.Get("X-API-VERSION"))
	assert.Equal(t, "hello foo", string(body))
}

// testing an http client

func GetDuration(client *http.Client, url string, body string) (time.Duration, error) {
	resp, err := client.Post(url, "text/plain", strings.NewReader(body))
	if err != nil {
		return 0, err
	}
	respBody, _ := ioutil.ReadAll(resp.Body)
	parsed, _ := strconv.Atoi(string(respBody))

	return time.Duration(int64(parsed)), nil
}

func TestGetDuration(t *testing.T) {
	srv := httptest.NewServer(
		http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				_, _ = w.Write([]byte("42"))
			},
		),
	)
	defer srv.Close()

	client := &http.Client{}
	duration, err := GetDuration(client, srv.URL, "knock knock")

	assert.NoError(t, err)
	assert.Equal(t, time.Duration(int64(42)), duration)
}
