package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(ls string) (string, error) {
	var finalString strings.Builder
	s := []rune(ls)

	if len(s) == 0 {
		return "", nil
	}

	if unicode.IsDigit(s[0]) {
		return "First rune is digit.", ErrInvalidString
	}

	for i := 0; i < len(s); i++ {
		tmpRune := s[i]

		if strings.EqualFold(string(s[i]), `\`) && (unicode.IsDigit(s[i+1]) || strings.EqualFold(string(s[i+1]), `\`)) {
			tmpRune = s[i+1]
			i++
		} else if strings.EqualFold(string(s[i]), `\`) && unicode.IsLetter(s[i+1]) {
			return "After / rune. ERROR.", ErrInvalidString
		}

		if i+1 < len(s) && unicode.IsDigit(s[i+1]) {
			if i+2 < len(s) && unicode.IsDigit(s[i+2]) {
				return "Double digit.", ErrInvalidString
			}
			t, _ := strconv.Atoi(string(s[i+1]))
			finalString.WriteString(strings.Repeat(string(tmpRune), t))
			i++
		} else {
			finalString.WriteRune(tmpRune)
		}
	}

	return finalString.String(), nil
}
