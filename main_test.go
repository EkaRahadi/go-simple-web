package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
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

func TestMainQuery(t *testing.T) {
	tests := []struct {
		name     string
		request  []string
		expected string
	}{
		{
			name:     "Hello, Eka Rahadi",
			request:  []string{"Eka", "Rahadi"},
			expected: "Hello, Eka Rahadi",
		},
		{
			name:     "Hello, Elba Ayu",
			request:  []string{"Elba", "Ayu"},
			expected: "Hello, Elba Ayu",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodGet,
				fmt.Sprintf("http://localhost:9000/say-hello?first_name=%s&last_name=%s", test.request[0], test.request[1]), nil)
			recorder := httptest.NewRecorder()

			QueryHandler(recorder, request)

			response := recorder.Result()
			body, _ := io.ReadAll(response.Body)
			bodyStr := string(body)
			assert.Equal(t, test.expected, bodyStr, fmt.Sprintf("Should Match 'Hello, %s %s'", test.request[0], test.request[1]))
		})
	}
}

func TestMainQueryArray(t *testing.T) {
	tests := []struct {
		name     string
		request  []string
		expected string
	}{
		{
			name:     "Hello, Eka Rahadi",
			request:  []string{"Eka", "Rahadi"},
			expected: "Hello, Eka Rahadi",
		},
		{
			name:     "Hello, Elba Ayu",
			request:  []string{"Elba", "Ayu"},
			expected: "Hello, Elba Ayu",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodGet,
				fmt.Sprintf("http://localhost:9000/say-hello-array?name=%s&name=%s", test.request[0], test.request[1]), nil)
			recorder := httptest.NewRecorder()

			QueryHandlerArray(recorder, request)

			response := recorder.Result()
			body, _ := io.ReadAll(response.Body)
			bodyStr := string(body)
			assert.Equal(t, test.expected, bodyStr, fmt.Sprintf("Should Match 'Hello, %s %s'", test.request[0], test.request[1]))
		})
	}
}

func TestFormPostHandler(t *testing.T) {
	tests := []struct {
		name     string
		request  []string
		expected string
	}{
		{
			name:     "Hello, Eka Rahadi",
			request:  []string{"Eka", "Rahadi"},
			expected: "Hello, Eka Rahadi",
		},
		{
			name:     "Hello, Elba Ayu",
			request:  []string{"Elba", "Ayu"},
			expected: "Hello, Elba Ayu",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			requestBody := strings.NewReader(fmt.Sprintf("first_name=%s&last_name=%s", test.request[0], test.request[1]))
			request := httptest.NewRequest(http.MethodPost, "http://localhost:9000/say-hello-post", requestBody)
			request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
			recorder := httptest.NewRecorder()

			FormPostHandler(recorder, request)

			response := recorder.Result()
			body, _ := io.ReadAll(response.Body)
			bodyStr := string(body)
			assert.Equal(t, test.expected, bodyStr, fmt.Sprintf("Should Match 'Hello, %s %s'", test.request[0], test.request[1]))
		})
	}
}
