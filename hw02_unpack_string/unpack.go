package hw02unpackstring

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func isDigit09(r rune) bool {
	if r >= 48 && r <= 57 {
		return true
	}
	return false
}

func Unpack(s string) (string, error) {
	var finalString strings.Builder
	sliceRune := []rune(s)

	if len(sliceRune) == 0 {
		return "", nil
	}

	if isDigit09(sliceRune[0]) {
		return "", fmt.Errorf("first rune is digit. %w", ErrInvalidString)
	}

	for i := 0; i < len(sliceRune); i++ {
		tmpRune := sliceRune[i]

		if i+1 < len(sliceRune) {
			if string(sliceRune[i]) == `\` && (isDigit09(sliceRune[i+1]) || string(sliceRune[i+1]) == `\`) {
				tmpRune = sliceRune[i+1]
				i++
			} else if i+1 < len(sliceRune) && string(sliceRune[i]) == `\` && unicode.IsLetter(sliceRune[i+1]) {
				return "", fmt.Errorf("after / rune. ERROR. %w", ErrInvalidString)
			}
		}

		if i+1 < len(sliceRune) && isDigit09(sliceRune[i+1]) {
			if i+2 < len(sliceRune) && isDigit09(sliceRune[i+2]) {
				return "", fmt.Errorf("double digit. %w", ErrInvalidString)
			}
			t, _ := strconv.Atoi(string(sliceRune[i+1]))
			finalString.WriteString(strings.Repeat(string(tmpRune), t))
			i++
		} else {
			finalString.WriteRune(tmpRune)
		}
	}

	return finalString.String(), nil
}
