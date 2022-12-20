package hw03frequencyanalysis

import (
	"regexp"
	"sort"
	"strings"
	"unicode"
)

const (
	lenFinalSlace = 10
)

var rgx = regexp.MustCompile("[a-zA-Z0-9А-Яа-я]*.[a-zA-Z0-9А-Яа-я]")

type wordsStruct struct {
	name  string
	count int
}

func search(s []int, v int) bool {
	for _, i := range s {
		tmp := i
		if tmp == v {
			return true
		}
	}
	return false
}

func Top10(s string) []string {
	words := make(map[string]int)
	// формируем map  со словами и количеством вхождений
	for _, w := range strings.Fields(s) {
		tmpWord := w
		//получаем слово без спец. символов или без спец. символов и с тире
		if len(tmpWord) > 0 && (unicode.IsPunct(rune(tmpWord[len(tmpWord)-1])) || unicode.IsPunct(rune(tmpWord[len(tmpWord)-1]))) {
			tmpWord = strings.Join(rgx.FindStringSubmatch(tmpWord), "")
		}
		// проверяем и записываем в map
		_, ok := words[strings.ToLower(tmpWord)]
		if len(tmpWord) > 0 && ok {
			words[strings.ToLower(tmpWord)] = words[strings.ToLower(tmpWord)] + 1
		} else {
			words[strings.ToLower(tmpWord)] = 1
		}
	}

	// фирмируем слайс с уникальным количеством вхождений
	countSlice := []int{}
	for _, v := range words {
		// написал функцию поиска, но потом увидел что такая есть в стандартном пакете, решил оставить
		if !search(countSlice, v) {
			countSlice = append(countSlice, v)
		}
	}
	// сортируем и переварачиваем, первым элементом будет самое частое слово (не смог найти сортировку по убываю, знаете такую?)
	sort.Ints(countSlice)
	for i := len(countSlice)/2 - 1; i >= 0; i-- {
		opp := len(countSlice) - 1 - i
		countSlice[i], countSlice[opp] = countSlice[opp], countSlice[i]
	}

	// формируем финальный слайс
	finalSlace := []string{}
	for i := 0; i < lenFinalSlace; i++ {
		tmpSliceWords := []string{}
		// нужно только lenFinalSlace элементов
		if len(finalSlace) < lenFinalSlace {
			for k, v := range words {
				if i < len(countSlice) && v == countSlice[i] {
					tmpSliceWords = append(tmpSliceWords, k)
				}
			}
		}
		// для этого количества повторений сортируем вывод по алфавиту
		sort.Strings(tmpSliceWords)
		for _, v := range tmpSliceWords {
			finalSlace = append(finalSlace, v)
		}
	}
	// если мало слов в тексте было, выводим что есть
	if len(finalSlace) > lenFinalSlace {
		finalSlace = finalSlace[:lenFinalSlace]
	}
	return finalSlace
}
