package morse

import (
	"errors"
	"strings"
)

// Valid returns error when text has availble characters
func Valid(text string) error {
	text = strings.ToUpper(text)
	bs := []byte(text)
	for _, c := range bs {
		if _, ok := dict[c]; !ok {
			return errors.New("no-valid character in text")
		}
	}
	return nil
}

// ToMorse translate text to morse_code
func ToMorse(text string) string {
	text = strings.ToUpper(text)
	bs := []byte(text)
	result := []string{}
	for _, c := range bs {
		result = append(result, dict[c])
	}
	return strings.Join(result, " ")
}

// FromMorse translate morse_code to text
func FromMorse(morseCode string) string {
	return ""
}
