package hw03frequencyanalysis

import (
	"regexp"
	"sort"
	"strings"
	"unicode"
)

var rgx = regexp.MustCompile("[a-zA-Z0-9А-Яа-я]*.[a-zA-Z0-9А-Яа-я]")

type wordsStruct struct {
	words []string
	count int
}

func Top10(s string) []string {
	// проверяем на пустую строку и выходим если это так
	if s == "" {
		return []string{}
	}

	lenFinalSlace := 10
	words := make(map[string]int)
	// формируем map  со словами и количеством вхождений
	for _, w := range strings.Fields(s) {
		tmpWord := w
		// получаем слово без спец символов или без спец символов и с тире
		if len(tmpWord) > 0 {
			if unicode.IsPunct(rune(tmpWord[0])) || unicode.IsPunct(rune(tmpWord[len(tmpWord)-1])) {
				tmpWord = strings.Join(rgx.FindStringSubmatch(tmpWord), "")
			}
		}
		// проверяем на пустоту и записываем в map
		if len(tmpWord) > 0 {
			words[strings.ToLower(tmpWord)]++
		}
	}

	// создаем мап где ключ это количество вхождений а занчение это слуйс из слов
	wordsKeyCount := make(map[int][]string)
	for k, v := range words {
		wordsKeyCount[v] = append(wordsKeyCount[v], k)
	}

	// создаем слайс из структур количесто вхождений и слайс слов
	// это позволит сразу брать слайсы слов и сортировать перед вставкой в финальный слайс
	sliceWords := make([]wordsStruct, 0, len(wordsKeyCount))
	for k, v := range wordsKeyCount {
		sliceWords = append(sliceWords, wordsStruct{words: v, count: k})
	}
	// сортируем по количеству вхождений
	sort.Slice(sliceWords, func(i, j int) bool {
		return sliceWords[i].count > sliceWords[j].count
	})

	// создаем финальный слайс по нужной длине
	if len(words) < lenFinalSlace {
		lenFinalSlace = len(words)
	}
	finalSlace := make([]string, 0, lenFinalSlace)

	flag := true
	i := 0
	// формируем финальный слайс, выходим когда он достигнет нужного размера
	for flag {
		sort.Strings(sliceWords[i].words)
		finalSlace = append(finalSlace, sliceWords[i].words...)
		if len(finalSlace) >= lenFinalSlace {
			flag = false
		}
		i++
	}
	return finalSlace[:lenFinalSlace]
}
