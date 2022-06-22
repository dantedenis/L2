package main

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

func Unpacked(str string) (string, error) {
	if _, err := strconv.Atoi(str); err == nil {
		return "", errors.New("некорректная строка")
	}

	var buf strings.Builder

	var prev rune
	var isEscaped bool
	for _, char := range str {
		if unicode.IsDigit(char) && !isEscaped {
			num := int(char - '0')
			repeat := strings.Repeat(string(prev), num-1)
			buf.WriteString(repeat)
		} else {
			isEscaped = string(char) == "\\" && string(prev) != "\\"
			if !isEscaped {
				buf.WriteRune(char)
			}
			prev = char
		}
	}

	return buf.String(), nil
}
