package main

import (
	"fmt"
	"sort"
	"strings"
)

/*
Поиск анаграмм по словарю

Написать функцию поиска всех множеств анаграмм по словарю.


Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.


Требования:
1.Входные данные для функции: ссылка на массив, каждый элемент которого - слово на русском языке в кодировке utf8

2.Выходные данные: ссылка на мапу множеств анаграмм

3.Ключ - первое встретившееся в словаре слово из множества. Значение - ссылка на массив, каждый элемент которого,
слово из множества.

4.Массив должен быть отсортирован по возрастанию.

5.Множества из одного элемента не должны попасть в результат.

6.Все слова должны быть приведены к нижнему регистру.

7.В результате каждое слово должно встречаться только один раз.
*/

func main() {
	arr := []string{"тяпка", "ПЯТАК", "Пятка", "бетон", "СЛИТОК", "столик", "листок", "вБа", "Авб", "абВ", "листок", "вБа"}

	annagramSets := findAnagrammsSets(arr)

	for _, xstrArray := range annagramSets {
		fmt.Println(xstrArray)
	}
}

// findAnagrammsSets Производит поиск анаграмм, и возвращает массив множеств слов анаграмм
// Каждое множество анаграмм отсортировано
func findAnagrammsSets(stringArray []string) map[string][]string {
	annagramSet := make(map[string][]string)
	keysMap := make(map[string]string)

	for _, xstr := range stringArray {
		xstr = strings.ToLower(xstr)
		arrRune := []rune(xstr)

		sort.Slice(arrRune, func(i, j int) bool {
			return arrRune[i] < arrRune[j]
		})

		sortedStr := string(arrRune)

		if keysMap[sortedStr] == "" {
			keysMap[sortedStr] = xstr
			annagramSet[xstr] = append(annagramSet[xstr], xstr)
		} else {
			if !wordAlreadyInclude(annagramSet, keysMap, sortedStr, xstr) {
				annagramSet[keysMap[sortedStr]] = append(annagramSet[keysMap[sortedStr]], xstr)
			}
		}

	}

	for _, xStrArray := range annagramSet {
		sort.Slice(xStrArray, func(i, j int) bool {
			return xStrArray[i] < xStrArray[j]
		})
	}

	return annagramSet
}

// wordAlreadyInclude Проверяет включено ли данное слово в список анаграмм или нет,
// возвращает true если это слово уже есть в списке анаграмм и false если нет.
func wordAlreadyInclude(annagramSet map[string][]string, keysMap map[string]string, key, word string) bool {
	for _, xstr := range annagramSet[keysMap[key]] {
		if xstr == word {
			return true
		}
	}

	return false
}
