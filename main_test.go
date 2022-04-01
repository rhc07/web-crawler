package main

import (
	"io"
	"net/http"
	"testing"
)

var (
	testInput       = "http://www.blankwebsite.com/"
	testResponse, _ = http.Get(testInput)
	testBody        = testResponse.Body
)

func TestGetResponseBody(t *testing.T) {
	tests := []struct {
		input string
		body  io.Reader
	}{
		{input: testInput, body: io.Reader(testBody)},
	}

	for _, test := range tests {
		err := GetResponseBody(test.input)
		if err != test.body {
			t.Errorf("Input: %s was incorrect, got: %s, want: %s", test.input, err, test.body)
		}
	}
}

func TestGetAnchorLinks(t *testing.T) {
	tests := []struct {
		input  io.Reader
		output []string
	}{
		{input: testResponse.Body, output: []string{"http://www.pointlesssites.com/"}},
	}

	for _, test := range tests {
		actual := GetAnchorLinks(testBody)
		if actual != nil {
			t.Errorf("Input: %s was incorrect, got: %s, want: %s", test.input, actual, test.output)
		}
	}
}
