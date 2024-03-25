package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

/*
Утилита grep


Реализовать утилиту фильтрации по аналогии с консольной утилитой (man grep — смотрим описание и основные параметры).


Реализовать поддержку утилитой следующих ключей:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", напечатать номер строки

*/

func main() {
	flag := flags{
		after:      0,
		before:     0,
		context:    0,
		count:      false,
		ignoreCase: true,
		invert:     false,
		fixed:      false,
		lineNum:    false,
	}

	pattern := "abc"
	fileName := "file.txt"
	answerFileName := "answer.txt"

	grep(flag, pattern, fileName, answerFileName)
}

type flags struct {
	after      int  //-A - "after" печатать +N строк после совпадения
	before     int  //-B - "before" печатать +N строк до совпадения
	context    int  //-C - "context" (A+B) печатать ±N строк вокруг совпадения
	count      bool //-c - "count" (количество строк)
	ignoreCase bool //-i - "ignore-case" (игнорировать регистр)
	invert     bool //-v - "invert" (вместо совпадения, исключать)
	fixed      bool //-F - "fixed", точное совпадение со строкой, не паттерн
	lineNum    bool //-n - "line num", напечатать номер строки
}

func grep(flag flags, pattern, fileName, answerFileName string) error {
	file, err := os.OpenFile(fileName, os.O_RDWR, 0666)
	if err != nil {
		fmt.Println("Ошибка при открытии файла:", err)
		return err
	}
	defer file.Close()

	var fileStrings []string

	fileStrings, err = readFileByLines(file)
	if err != nil {
		fmt.Println("Ошибка при чтении файла построчно:", err)
		return err
	}

	if flag.count == true {
		return countOfMatchedStrings(fileStrings, pattern, flag)
	} else if flag.after > 0 || flag.before > 0 || flag.context > 0 {
		fileStrings, err = flagABC(fileStrings, pattern, flag)
		if err != nil {
			return err
		}
	} else {
		fileStrings, err = justGrep(fileStrings, pattern, flag)
	}

	err = writeAllStringsToFile(answerFileName, fileStrings)
	if err != nil {
		fmt.Println("Ошибка при записи данных в файл:", err)
		return err
	}

	fmt.Println("Данные успешно записаны в файл.")
	return nil
}

// readFileByLines Читает все строки из файла
func readFileByLines(file *os.File) ([]string, error) {
	var fileStrings []string        // Массив для хранения строк файла
	var currentLine []byte          // Собираем текущую строку здесь
	reader := bufio.NewReader(file) // Создаем буферизованный reader для файла

	for {
		b, err := reader.ReadByte() // Считываем байт
		if err != nil {
			break // Если достигнут конец файла или произошла ошибка, выходим из цикла
		}

		if b == '\n' {
			// Когда встречаем символ новой строки, добавляем собранную строку в массив
			fileStrings = append(fileStrings, string(currentLine))
			currentLine = nil // Сбрасываем текущую строку
		} else {
			currentLine = append(currentLine, b) // Добавляем байт в текущую строку
		}
	}

	// Добавляем последнюю строку, если она не пустая
	if len(currentLine) > 0 {
		fileStrings = append(fileStrings, string(currentLine))
	}

	return fileStrings, nil
}

// writeAllStringsToFile Записывает все отсортированные строки в файл
func writeAllStringsToFile(fileName string, fileStrings []string) error {
	// Открываем файл для записи. Если файл существует, он будет перезаписан.
	file, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("ошибка при создании файла: %w", err)
	}
	defer file.Close()

	for _, line := range fileStrings {
		// Преобразуем строку в слайс байт и записываем её в файл.
		_, err := file.Write([]byte(line)) // Записываем каждый байт отдельно
		if err != nil {
			return fmt.Errorf("ошибка при заWписи в файл: %w", err)
		}
		// После каждой строки добавляем символ новой строки в файл.
		if _, err := file.Write([]byte("\n")); err != nil {
			return fmt.Errorf("ошибка при добавлении символа новой строки: %w", err)
		}
	}

	return nil
}

func countOfMatchedStrings(fileStrings []string, pattern string, flag flags) error {
	result := 0
	if flag.ignoreCase == true {
		pattern = "(?i)" + pattern
	}
	repattern, err := regexp.Compile(pattern)
	if err != nil {
		return fmt.Errorf("Некорректный паттерн: %w", err)
	}

	for _, xstr := range fileStrings {
		if repattern.MatchString(xstr) {
			result++
		}
	}

	fmt.Println("Количество строк содержащих совпадение равно:", result)

	return nil
}

func flagABC(fileStrings []string, pattern string, flag flags) ([]string, error) {
	result := make([]string, 0)
	alreadyAdded := make([]bool, len(fileStrings))

	if flag.ignoreCase == true {
		pattern = "(?i)" + pattern
	}
	if flag.fixed == true {
		pattern = "^" + pattern + "$"
	}

	repattern, err := regexp.Compile(pattern)
	if err != nil {
		return nil, fmt.Errorf("Некорректный паттерн: %w", err)
	}

	for i, xstr := range fileStrings {
		matchResult := repattern.MatchString(xstr)
		if flag.invert == true {
			matchResult = !matchResult
		}
		if matchResult {
			if flag.context > 0 || flag.before > 0 {
				xn1 := 0
				if flag.context > 0 {
					xn1 = flag.context
				} else {
					xn1 = flag.before
				}
				for i < xn1 {
					xn1--
				}
				if xn1 > 0 {
					for xn1 > 0 {
						if alreadyAdded[i-xn1] == false {
							result = append(result, fileStrings[i-xn1])
						}
						alreadyAdded[i-xn1] = true
						xn1--
					}
				}
			}

			if alreadyAdded[i] == false {
				result = append(result, xstr)
			}
			alreadyAdded[i] = true

			if flag.context > 0 || flag.after > 0 {
				xn2 := 0
				if flag.context > 0 {
					xn2 = flag.context
				} else {
					xn2 = flag.after
				}
				for xn2+i > len(fileStrings)-1 {
					xn2--
				}
				if xn2 > 0 {
					for xn2 > 0 {
						if alreadyAdded[i+xn2] == false {
							result = append(result, fileStrings[i+xn2])
						}
						alreadyAdded[i+xn2] = true
						xn2--
					}
				}
			}

		}
	}

	return result, nil
}

func justGrep(fileStrings []string, pattern string, flag flags) ([]string, error) {
	result := make([]string, 0)

	if flag.ignoreCase == true {
		pattern = "(?i)" + pattern
	}

	if flag.fixed == true {
		pattern = "^" + pattern + "$"
	}

	repattern, err := regexp.Compile(pattern)
	if err != nil {
		return nil, fmt.Errorf("Некорректный паттерн: %w", err)
	}

	for _, xstr := range fileStrings {
		matchResult := repattern.MatchString(xstr)
		if flag.invert == true {
			matchResult = !matchResult
		}
		if matchResult {
			result = append(result, xstr)
		}
	}

	return result, nil
}
