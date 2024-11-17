package utils

import (
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

func DecodeUTF8(input string) (string, error) {
	decoder := unicode.UTF8.NewDecoder()
	decoded, _, err := transform.String(decoder, input)
	if err != nil {
		return "", err
	}
	return decoded, nil
}
