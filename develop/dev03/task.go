package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func main() {
	flag := flags{
		reverseSort:    false,
		uniqueStrings:  false,
		ignoringBlanks: true,
		checkForSorted: false,
		numberSort:     true,
	}

	SortFile(flag, "file.txt")
}

// flags структура флагов
type flags struct {
	column                int  // - k — указание колонки для сортировки (слова в строке могут выступать в качестве колонок, по умолчанию разделитель — пробел)
	numberSort            bool // -n — сортировать по числовому значению
	reverseSort           bool // -r — сортировать в обратном порядке
	uniqueStrings         bool // -u - не выводить повторяющиеся строки
	month                 bool // -m - сортировать по названию месяца
	ignoringBlanks        bool // -b - игнорировать хвостовые пробелы
	checkForSorted        bool // -c — проверять отсортированы ли данные
	numericSortWithSuffix bool // -h — сортировать по числовому значению с учетом суффиксов
}

// SortFile Сортирует строки в файле, следуя указанным флагам
func SortFile(flag flags, fileName string) {
	file, err := os.OpenFile(fileName, os.O_RDWR, 0644)
	if err != nil {
		fmt.Println("Ошибка при открытии файла:", err)
		return
	}
	defer file.Close()

	var fileStrings []string

	if flag.ignoringBlanks == false {
		fileStrings, err = readFileByLines(file)
		if err != nil {
			fmt.Println("Ошибка при чтении файла построчно:", err)
			return
		}
	} else {
		fileStrings, err = readFileWithoutBlanks(file)
		if err != nil {
			fmt.Println("Ошибка при чтении файла построчно:", err)
			return
		}
	}

	// Проверяет отсортирован ли файл если этот флаг включен
	if flag.checkForSorted == true {
		if checkForSorted(fileStrings, flag) == true {
			fmt.Println("Файл уже отсортирован")
			return
		}
	}

	// Выбирает тип сортировки взависимости от установленных флагов
	if flag.reverseSort == true && flag.numberSort == true {
		sortByNumericDecrease(fileStrings)
	} else if flag.reverseSort == false && flag.numberSort == true {
		sortByNumericIncrease(fileStrings)
	} else if flag.reverseSort == true && flag.numberSort == false {
		decreaseStringsSort(fileStrings)
	} else {
		increaseStringsSort(fileStrings)
	}

	// Если флаг установлен, то удаляем дублирующиеся строки
	if flag.uniqueStrings == true {
		fileStrings = removeDuplicates(fileStrings)
	}

	err = writeAllStringsToFile(file, fileStrings)
	if err != nil {
		fmt.Println("Ошибка при записи данных в файл:", err)
		return
	}

	fmt.Println("Данные успешно записаны в файл.")
}

// readFileWithoutBlanks Читает все строки из файла игнорируя пробелы
func readFileWithoutBlanks(file *os.File) ([]string, error) {
	// Создаем новый сканер для чтения файла построчно
	scanner := bufio.NewScanner(file)

	// Считываем файл построчно
	fileStrings := make([]string, 0)
	for scanner.Scan() {
		stringWithoutBlanks := removeAllSpaces(scanner.Text())
		fileStrings = append(fileStrings, stringWithoutBlanks)
	}

	// Проверяем наличие ошибок во время сканирования
	if err := scanner.Err(); err != nil {
		fmt.Println("Ошибка при сканировании файла:", err)
		return nil, err
	}

	return fileStrings, nil
}

// readFileByLines Читает все строки из файла
func readFileByLines(file *os.File) ([]string, error) {
	// Создаем новый сканер для чтения файла построчно
	scanner := bufio.NewScanner(file)

	// Считываем файл построчно
	fileStrings := make([]string, 0)
	for scanner.Scan() {
		fileStrings = append(fileStrings, scanner.Text())
	}

	// Проверяем наличие ошибок во время сканирования
	if err := scanner.Err(); err != nil {
		fmt.Println("Ошибка при сканировании файла:", err)
		return nil, err
	}

	return fileStrings, nil
}

// writeAllStringsToFile Записывает все отсортированные строки в файл
func writeAllStringsToFile(file *os.File, fileStrings []string) error {
	// Очищаем содержимое файла
	err := file.Truncate(0)
	if err != nil {
		fmt.Println("Ошибка при очистке файла:", err)
		return err
	}

	// Создаем новый писатель для записи в файл
	writer := bufio.NewWriter(file)

	// Записываем каждую строку в файл
	for _, line := range fileStrings {
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			fmt.Println("Ошибка при записи в файл:", err)
			return err
		}
	}

	// Сбрасываем буфер и убеждаемся, что все данные записаны в файл
	if err = writer.Flush(); err != nil {
		fmt.Println("Ошибка при сбросе буфера:", err)
		return err
	}

	return nil
}

// increaseStringsSort Сортирует массив строк в возрастающем порядке
func increaseStringsSort(fileStrings []string) {
	sort.Slice(fileStrings, func(i, j int) bool {
		if len(fileStrings[i]) == 0 {
		}
		return fileStrings[i] < fileStrings[j]
	})
}

// decreaseStringsSort Сортирует массив строк в убывающем порядке
func decreaseStringsSort(fileStrings []string) {
	sort.Slice(fileStrings, func(i, j int) bool {
		if len(fileStrings[i]) == 0 {
		}
		return fileStrings[i] > fileStrings[j]
	})
}

// sortByNumeric сортирует массив строк по числовому значению в возрастающем порядке
func sortByNumericIncrease(fileStrings []string) {
	sort.Slice(fileStrings, func(i, j int) bool {
		// Преобразуем строки в числа и сравниваем их
		return convertToInt(fileStrings[i]) < convertToInt(fileStrings[j])
	})
}

// sortByNumeric сортирует массив строк по числовому значению в убывающем порядке
func sortByNumericDecrease(fileStrings []string) {
	sort.Slice(fileStrings, func(i, j int) bool {
		// Преобразуем строки в числа и сравниваем их
		return convertToInt(fileStrings[i]) > convertToInt(fileStrings[j])
	})
}

// sortByNumericWithPrefix
func sortByNumericWithPrefix(fileStrings []string) {
	sort.Slice(fileStrings, func(i, j int) bool {
		// Преобразуем строки в числа и сравниваем их
		return convertToInt(fileStrings[i]) > convertToInt(fileStrings[j])
	})
}

// toBytes Функция для преобразования строки с единицей измерения в число байтов
func toBytes(s string) (int, error) {
	re := regexp.MustCompile(`(?i)^(\d+)([KMGTP]?)B?$`)
	matches := re.FindStringSubmatch(s)
	if matches == nil {
		return 0, fmt.Errorf("invalid format: %s", s)
	}

	base, err := strconv.Atoi(matches[1])
	if err != nil {
		return 0, err
	}

	switch strings.ToUpper(matches[2]) {
	case "K":
		return base * 1024, nil
	case "M":
		return base * 1024 * 1024, nil
	case "G":
		return base * 1024 * 1024 * 1024, nil
	case "T":
		return base * 1024 * 1024 * 1024 * 1024, nil
	case "P":
		return base * 1024 * 1024 * 1024 * 1024 * 1024, nil
	default:
		return base, nil
	}
}

// convertToInt Конвертирует строку в числовое значение
func convertToInt(str string) int {
	num, _ := strconv.Atoi(str)
	return num
}

// removeAllSpaces Удаляет все пробелы из строки
func removeAllSpaces(str string) string {
	return strings.TrimRight(str, " ")
}

// checkForSorted проверяет отсортированны ли строки
func checkForSorted(strArray []string, Flag flags) bool {

	if Flag.reverseSort == false {
		for j := 0; j < len(strArray); j++ {
			if j == 0 {
				continue
			} else {
				if strArray[j-1] > strArray[j] {
					return false
				}
			}
		}
	} else {
		for j := 0; j < len(strArray); j++ {
			if j == 0 {
				continue
			} else {
				if strArray[j-1] < strArray[j] {
					return false
				}
			}
		}
	}

	return true
}

// removeDuplicates Убирает дупликаты строк
func removeDuplicates(strings []string) []string {
	encountered := map[string]bool{} // Создаем карту для отслеживания встреченных строк
	result := []string{}             // Создаем пустой результат

	// Проходим по каждой строке в массиве строк
	for _, str := range strings {
		// Если текущая строка еще не встречалась, добавляем ее в результат и отмечаем как встреченную
		if !encountered[str] {
			encountered[str] = true
			result = append(result, str)
		}
	}

	return result
}
