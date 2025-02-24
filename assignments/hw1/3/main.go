package main

import (
	"os"
	"regexp"
	"sort"
	"strings"
)

func readFile(pathToFile string) []rune {
	content, err := os.ReadFile(pathToFile)
	if err != nil {
		panic(err)
	}
	return []rune(string(content))
}

func filterCharsAndNormalize(strData []rune) []rune {
	// use regex to remove all non-alphanumeric characters
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		panic(err)
	}
	return []rune(reg.ReplaceAllString(strings.ToLower(string(strData)), " "))
}

func scan(strData []rune) []string {
	return strings.Fields(string(strData))
}

func removeStopWords(wordList []string) []string {
	content, err := os.ReadFile("../stop_words.txt")
	if err != nil {
		panic(err)
	}
	stopWords := strings.Split(string(content), ",")
	stopWords = append(stopWords, make([]string, 26)...)
	for c := 'a'; c <= 'z'; c++ {
		stopWords = append(stopWords, string(c))
	}
	stopWordsSet := make(map[string]struct{})
	for _, word := range stopWords {
		stopWordsSet[word] = struct{}{}
	}
	var words []string
	for _, word := range wordList {
		if _, ok := stopWordsSet[word]; !ok {
			words = append(words, word)
		}
	}
	return words
}

func frequencies(wordList []string) map[string]int {

	wordFreqs := make(map[string]int)
	for _, w := range wordList {
		if _, ok := wordFreqs[w]; ok {
			wordFreqs[w]++
		} else {
			wordFreqs[w] = 1
		}
	}
	return wordFreqs
}

// let the sort function return the the pair array with the highest frequency first
func sortFreqs(wordFreqs map[string]int) [][2]interface{} {
	var freqPairs [][2]interface{}
	for word, freq := range wordFreqs {
		freqPairs = append(freqPairs, [2]interface{}{word, freq})
	}
	// sort the array by frequency
	sort.Slice(freqPairs, func(i, j int) bool {
		return freqPairs[i][1].(int) > freqPairs[j][1].(int)
	})
	return freqPairs
}
func printAll(wordFreqs [][2]interface{}) {
	if len(wordFreqs) > 0 {
		println(wordFreqs[0][0].(string), " - ", wordFreqs[0][1].(int))
		printAll(wordFreqs[1:])
	}
}

func main() {
	if len(os.Args) < 2 {
		panic("Please provide the file for word frequency counting")
	}
	printAll(sortFreqs(frequencies(removeStopWords(scan(filterCharsAndNormalize(readFile(os.Args[1]))))))[:25])
}
