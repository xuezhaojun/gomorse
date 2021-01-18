package morse

import (
	"testing"
)

func TestToMorse(t *testing.T) {
	testcases := []struct {
		text      string
		morseCode string
	}{
		{
			"SOS",
			"... --- ...",
		},
		{
			"sos",
			"... --- ...",
		},
	}
	for _, testcase := range testcases {
		if ToMorse(testcase.text) != testcase.morseCode {
			t.Error(testcase, "not pass")
		}
	}
}

func TestFromMorse(t *testing.T) {
	testcases := []struct {
		text      string
		morseCode string
	}{
		{
			"SOS",
			"... --- ...",
		},
	}
	for _, testcase := range testcases {
		if FromMorse(testcase.morseCode) != testcase.text {
			t.Error(testcase, "not pass")
		}
	}
}
