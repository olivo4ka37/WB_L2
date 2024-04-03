package main

import (
	"strconv"
	"strings"
)

/*
=== Задача на распаковку ===
Создать Go-функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы/руны, например:
"a4bc2d5e" => "aaaabccddddde"
"abcd" => "abcd"
"45" => "" (некорректная строка)
"" => ""

=== Дополнительно ===
	Реализовать поддержку escape-последовательностей.
	Например:
	qwe\4\5 => qwe45 (*)
	qwe\45 => qwe44444 (*)
	qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка, функция должна возвращать ошибку. Написать unit-тесты.
*/

// Unpack распаковывает строку, возвращает "некорректная строка" если
// была передана неккоректная строка соответственно
// Для эффективной конкатеннации строк использует тип strings.Builder{}
// Также если в строке содержится escape последовательность, вызывает unpackWithEscape
// тем самым может работать с escape последовательностями
func Unpack(str string) string {
	resultBuffer := strings.Builder{}
	lastSymbol := ""
	repeatTimes := ""

	if len(str) == 0 {
		return ""
	}

	if isThatNumber(rune(str[0])) {
		return "некорректная строка"
	}

	for i, x := range str {
		if isThatEscape(x) {
			result := resultBuffer.String()
			return unpackWithEscape(result, str, lastSymbol, i)
		} else if !isThatNumber(x) {
			for j, _ := strconv.Atoi(repeatTimes); j > 1; j-- {
				resultBuffer.WriteString(lastSymbol)
			}

			resultBuffer.WriteString(string(x))
			lastSymbol = string(x)

			repeatTimes = ""
		} else {
			repeatTimes += string(x)
		}

	}

	if repeatTimes != "" {
		for j, _ := strconv.Atoi(repeatTimes); j > 1; j-- {
			resultBuffer.WriteString(lastSymbol)
		}
	}

	result := resultBuffer.String()

	return result
}

// unpackWithEscape распаковывает строку содержащую escape последовательность
func unpackWithEscape(result, str, lastSymbol string, n int) string {
	resultBuffer := strings.Builder{}

	for i := n; i < len(str); i++ {
		if isThatEscape(rune(str[i])) && !isThatEscape(rune(str[i-1])) {
			lastSymbol = ""
			continue
		} else if isThatNumber(rune(str[i])) && lastSymbol != "" {
			for j, _ := strconv.Atoi(string(str[i])); j > 1; j-- {
				resultBuffer.WriteString(lastSymbol)
			}
			lastSymbol = string(str[i])
		} else {
			resultBuffer.WriteString(string(str[i]))
			lastSymbol = string(str[i])
		}
	}
	result += resultBuffer.String()

	return result
}

// isThatNumber возвращает true если руна является цифрой
func isThatNumber(run rune) bool {
	if run >= 48 && run <= 57 {
		return true
	}

	return false
}

// isThatEscape возвращает true если руна является бэкслешом
func isThatEscape(run rune) bool {
	if run == '\\' {
		return true
	}

	return false
}
