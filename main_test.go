package main

import (
	"bytes"
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
		{input: testInput, body: testBody},
	}

	for _, test := range tests {
		err := GetResponseBody(test.input)
		res := bytes.Equal(StreamToByte(err), StreamToByte(test.body))
		if res != true {
			t.Errorf("Input: %s was incorrect, got: %s, want: %s", test.input, err, test.body)
		}
	}
}

func TestGetAnchorLinks(t *testing.T) {
	tests := []struct {
		input  io.Reader
		output []string
	}{
		{input: testBody, output: []string{"http://www.pointlesssites.com/"}},
	}

	for _, test := range tests {
		actual := GetAnchorLinks(test.input)
		for _, testValue := range test.output {
			for _, actualValue := range actual {
				if testValue != actualValue {
					t.Errorf("Input: %s was incorrect, got: %s, want: %s", test.input, actual, test.output)
				}
			}
		}
	}
}

func StreamToByte(stream io.Reader) []byte {
	buf := new(bytes.Buffer)
	buf.ReadFrom(stream)
	return buf.Bytes()
}
