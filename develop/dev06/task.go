package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	result, err := cut()
	if err != nil {
		fmt.Println(err)
	}
	for _, str := range result {
		fmt.Println(str)
	}
}

// cut Выполняет реализацию утилиты консольной утилиты cut с поддержкой флагов -f, -d, -s. Флаг -f обязателен.
func cut() ([]string, error) {
	fields := flag.String("f", "", "выбрать поля (колонки)")
	delimiter := flag.String("d", "\t", "использовать другой разделитель")
	suppressEmpty := flag.Bool("s", false, "только строки с разделителем")

	flag.Parse()

	if *fields == "" {
		return nil, fmt.Errorf("Ошибка: флаг -f обязателен!")
	}

	fieldIndices := parseFieldIndices(*fields)
	fmt.Println(fieldIndices)

	scanner := bufio.NewScanner(os.Stdin)
	var result []string

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, *delimiter)

		if *suppressEmpty && len(parts) == 1 {
			continue
		}

		var extractedFields []string
		for _, index := range fieldIndices {
			if index >= 0 && index < len(parts) {
				extractedFields = append(extractedFields, parts[index])
			}
		}

		result = append(result, strings.Join(extractedFields, *delimiter))
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading input:", err)
		os.Exit(1)
	}

	return result, nil
}

// parseFieldIndices Парсирует флаг -f и возвращает массив индексов полей, которые будут подлежать выводу.
func parseFieldIndices(fieldsStr string) []int {
	var fieldIndices []int
	for _, fieldStr := range strings.Split(fieldsStr, ",") {
		if strings.Contains(fieldStr, "-") {
			rangeParts := strings.Split(fieldStr, "-")
			start, err := strconv.Atoi(rangeParts[0])
			if err != nil {
				fmt.Fprintln(os.Stderr, "Error parsing field range:", err)
				os.Exit(1)
			}
			end, err := strconv.Atoi(rangeParts[1])
			if err != nil {
				fmt.Fprintln(os.Stderr, "Error parsing field range:", err)
				os.Exit(1)
			}
			for i := start; i <= end; i++ {
				fieldIndices = append(fieldIndices, i-1) // Adjust for 0-based indexing
			}
		} else {
			index, err := strconv.Atoi(fieldStr)
			if err != nil {
				fmt.Fprintln(os.Stderr, "Error parsing field index:", err)
				os.Exit(1)
			}
			fieldIndices = append(fieldIndices, index-1) // Adjust for 0-based indexing
		}
	}
	return fieldIndices
}
