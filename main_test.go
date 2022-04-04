package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	testInput    = "http://www.blankwebsite.com/"
	testInput2   = "https://en.wikipedia.org/wiki/Ken_Thompson"
	invalidInput = "http://invalidurl.com/"
)

func TestGetAnchorLinks(t *testing.T) {
	tests := []struct {
		input  string
		output string
		error  string
	}{
		{input: testInput, output: "http://www.pointlesssites.com/", error: ""},
		{input: testInput2, output: "https://vi.wikipedia.org/wiki/Ken_Thompson", error: ""},
		{input: invalidInput, output: "", error: "no body in the response"},
	}

	for _, test := range tests {
		actual, err := GetAnchorLinks(test.input)
		if test.error != "" {
			assert.ErrorContains(t, err, test.error)
		}
		if test.output != "" {
			assert.Contains(t, actual, test.output)
		}
	}
}
