package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(input string) (string, error) {
	var b strings.Builder
	asRune := []rune(input)

	for i, v := range asRune {
		repeat := 1
		isDigit := unicode.IsDigit(v)

		if isDigit && i == 0 {
			return "", ErrInvalidString
		}

		if i+1 < len(asRune) {
			next := asRune[i+1]
			isNextDigit := unicode.IsDigit(next)

			if isDigit && isNextDigit {
				return "", ErrInvalidString
			}

			if isDigit {
				continue
			}

			if isNextDigit {
				repeat, _ = strconv.Atoi(string(next))
			}
		}

		vString := string(v)
		b.WriteString(strings.Repeat(vString, repeat))
	}

	return b.String(), nil
}
