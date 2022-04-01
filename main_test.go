package main

import (
	"io"
	"net/http"
	"testing"
)

var (
	testInput = "http://www.blankwebsite.com/"
	//invalidInput            = "http://invalidurl.com/"
	testResponse, testError = http.Get(testInput)
	testBody                = testResponse.Body
)

func TestGetResponseBody(t *testing.T) {
	tests := []struct {
		input string
		body  io.Reader
		error error
	}{
		{input: testInput, body: testBody, error: nil},
		//{input: invalidInput, body: nil, error: testError},
	}

	for _, test := range tests {
		err := GetResponseBody(test.input)
		if err != test.body {
			t.Errorf("Input: %s was incorrect, got: %s, want: %s", test.input, err, test.body)
		}
	}
}
