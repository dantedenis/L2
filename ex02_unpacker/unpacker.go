package main

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

func Unpacked(str string) (string, error) {
	// если удается строку без ошибок преобразовать в инт, это некорректная строка
	if _, err := strconv.Atoi(str); err == nil {
		return "", errors.New("некорректная строка")
	}

	var buf strings.Builder

	// переменная для предыдущего символа
	var prev rune
	// флаг на эскейп последовательность
	var isEscaped bool
	for _, char := range str {
		if unicode.IsDigit(char) && !isEscaped {
			// переводим чар в инт
			num := int(char - '0')
			// и делаем репит символа и пишем в буфер
			repeat := strings.Repeat(string(prev), num-1)
			buf.WriteString(repeat)
			// иначе проверяет на возможность эскейп последовательности и обработываем и так же пишем в буфер
		} else {
			isEscaped = string(char) == "\\" && string(prev) != "\\"
			if !isEscaped {
				buf.WriteRune(char)
			}
			prev = char
		}
	}
	// возвращаем строку
	return buf.String(), nil
}
