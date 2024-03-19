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
		reverseSort:         false,
		uniqueStrings:       false,
		ignoringRightBlanks: true,
		checkForSorted:      false,
		numberSort:          false,
	}

	SortFile(flag, "file.txt", "C:/Users/eblan  elite/GolandProjects/WB_L2/develop/dev03/answer.txt")
}

// flags структура флагов
type flags struct {
	column                int  // - k — указание колонки для сортировки (слова в строке могут выступать в качестве колонок, по умолчанию разделитель — пробел)
	numberSort            bool // -n — сортировать по числовому значению
	reverseSort           bool // -r — сортировать в обратном порядке
	uniqueStrings         bool // -u - не выводить повторяющиеся строки
	month                 bool // -m - сортировать по названию месяца
	ignoringRightBlanks   bool // -b - игнорировать хвостовые пробелы
	checkForSorted        bool // -c — проверять отсортированы ли данные
	numericSortWithSuffix bool // -h — сортировать по числовому значению с учетом суффиксов
}

// SortFile Сортирует строки в файле, следуя указанным флагам
func SortFile(flag flags, fileName, answerFileDestination string) {
	file, err := os.OpenFile(fileName, os.O_RDWR, 0666)
	if err != nil {
		fmt.Println("Ошибка при открытии файла:", err)
		return
	}
	defer file.Close()

	var fileStrings []string

	if flag.ignoringRightBlanks == false {
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

	err = writeAllStringsToFile(answerFileDestination, fileStrings)
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
	//// Очищаем содержимое файла
	//err := file.Truncate(0)
	//if err != nil {
	//	fmt.Println("Ошибка при очистке файла:", err)
	//	return err
	//}

	// Открываем файл для записи. Если файл существует, он будет перезаписан.
	file, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("ошибка при создании файла: %w", err)
	}
	defer file.Close()

	for _, line := range fileStrings {
		// Преобразуем строку в слайс байт и записываем её в файл.
		fmt.Println([]byte(line))
		_, err := file.Write([]byte(line)) // Записываем каждый байт отдельно
		if err != nil {
			return fmt.Errorf("ошибка при записи в файл: %w", err)
		}
		// После каждой строки добавляем символ новой строки в файл.
		if _, err := file.Write([]byte("\n")); err != nil {
			return fmt.Errorf("ошибка при добавлении символа новой строки: %w", err)
		}
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
