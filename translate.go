package gomorse

import (
	"errors"
	"strings"
)

var errNoValid error = errors.New("no-valid character in text")

// Valid returns error when text has availble characters
func Valid(text string) error {
	text = strings.ToUpper(text)
	bs := []byte(text)
	for _, c := range bs {
		if _, ok := dict[c]; !ok {
			return errNoValid
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
	words := strings.Split(morseCode, "   ")
	var text []string
	for _, word := range words {
		codes := strings.Split(word, " ")
		var chars []byte
		for _, code := range codes {
			chars = append(chars, redict[code])
		}
		text = append(text, string(chars))
	}
	return strings.Join(text, " ")
}
