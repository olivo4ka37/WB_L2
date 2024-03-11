package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	flag := flags{
		reverseSort:   true,
		uniqueStrings: false,
	}

	Unpack(flag, "example.txt")
}

type flags struct {
	column                int  // - k — указание колонки для сортировки (слова в строке могут выступать в качестве колонок, по умолчанию разделитель — пробел)
	numberSort            bool // -n — сортировать по числовому значению
	reverseSort           bool // -r — сортировать в обратном порядке
	uniqueStrings         bool // -u - не выводить повторяющиеся строки
	month                 bool // -m - сортировать по названию месяца
	ignoringBlanks        bool // -b -игнорировать хвостовые пробелы
	checkForSorted        bool // -c — проверять отсортированы ли данные
	numericSortWithSuffix bool // -h — сортировать по числовому значению с учетом суффиксов
}

func Unpack(flag flags, fileName string) {
	// Открываем файл
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Ошибка при открытии файла:", err)
		return
	}
	defer file.Close()

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
		return
	}

	// Выбираем тип сортировки сортировка по умолчанию или по убыванию
	if flag.reverseSort == true {
		decreaseStringsSort(fileStrings)
	} else {
		increaseStringsSort(fileStrings)
	}

	// Если флаг установлен, то удаляем дублирующиеся строки
	if flag.uniqueStrings == true {
		fileStrings = removeDuplicates(fileStrings)
	}

	// Открываем файл для чтения
	file, err = os.OpenFile(fileName, os.O_RDWR, 0644)
	if err != nil {
		fmt.Println("Ошибка при открытии файла:", err)
		return
	}
	defer file.Close()

	// Очищаем содержимое файла
	err = file.Truncate(0)
	if err != nil {
		fmt.Println("Ошибка при очистке файла:", err)
		return
	}

	// Создаем новый писатель для записи в файл
	writer := bufio.NewWriter(file)

	// Записываем каждую строку в файл
	for _, line := range fileStrings {
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			fmt.Println("Ошибка при записи в файл:", err)
			return
		}
	}

	// Сбрасываем буфер и убеждаемся, что все данные записаны в файл
	if err = writer.Flush(); err != nil {
		fmt.Println("Ошибка при сбросе буфера:", err)
		return
	}

	fmt.Println("Данные успешно записаны в файл.")
}

func increaseStringsSort(fileStrings []string) {
	sort.Slice(fileStrings, func(i, j int) bool {
		if len(fileStrings[i]) == 0 {
		}
		return fileStrings[i] < fileStrings[j]
	})
}

func decreaseStringsSort(fileStrings []string) {
	sort.Slice(fileStrings, func(i, j int) bool {
		if len(fileStrings[i]) == 0 {
		}
		return fileStrings[i] > fileStrings[j]
	})
}

func sortByNumeric(fileStrings []string) {
	sort.Slice(fileStrings, func(i, j int) bool {
		// Преобразуем строки в числа и сравниваем их
		return convertToInt(fileStrings[i]) < convertToInt(fileStrings[j])
	})
}

func convertToInt(str string) int {
	num, _ := strconv.Atoi(str)
	return num
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
