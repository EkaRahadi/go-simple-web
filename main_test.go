package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Hello, world!",
			request:  "Hello, world!",
			expected: "Hello, world!",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodGet, "http://localhost:9000", nil)
			recorder := httptest.NewRecorder()

			IndexHandler(recorder, request)

			response := recorder.Result()
			body, _ := io.ReadAll(response.Body)
			bodyStr := string(body)
			assert.Equal(t, test.expected, bodyStr, "Should match 'Hello, world!'")
		})
	}
}
