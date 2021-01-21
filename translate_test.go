package morse

import (
	"testing"
)

func TestValid(t *testing.T) {
	testcases := []struct {
		text  string
		valid error
	}{
		{
			text:  "SOS",
			valid: nil,
		},
		{
			text:  "警告",
			valid: errNoValid,
		},
	}

	for index, testcase := range testcases {
		if Valid(testcase.text) != testcase.valid {
			t.Error(testcase, index, "not pass")
		}
	}
}

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
		{
			"I love you!",
			"..   .-.. --- ...- .   -.-- --- ..- -.-.--",
		},
	}
	for index, testcase := range testcases {
		get := ToMorse(testcase.text)
		if get != testcase.morseCode {
			t.Error(testcase, index, "not pass;", "target is:", testcase.morseCode, ",get:", get)
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
		{
			"I LOVE YOU!",
			"..   .-.. --- ...- .   -.-- --- ..- -.-.--",
		},
	}
	for _, testcase := range testcases {
		get := FromMorse(testcase.morseCode)
		if get != testcase.text {
			t.Error(testcase, "not pass", "except:", testcase.text, ",get:", get)
		}
	}
}
