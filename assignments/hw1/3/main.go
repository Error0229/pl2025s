package main

import (
	"os"
)

func readFile(pathToFile string) []rune {
	return nil
}

func filterCharsAndNormalize(strData []rune) []rune {
	return nil
}

func scan(strData []rune) []string {
	return nil
}

func removeStopWords(wordList []string) []string {
	return nil
}

func frequencies(wordList []string) map[string]int {
	return nil
}

// let the sort function return the the pair array with the highest frequency first
func sort(wordFreqs map[string]int) [][2]interface{} {
	return nil
}
func printAll(wordFreqs [][2]interface{}) {

}

func main() {
	if len(os.Args) < 2 {
		panic("Please provide the file for word frequency counting")
	}
	printAll(sort(frequencies(removeStopWords(scan(filterCharsAndNormalize(readFile(os.Args[1])))))))
}
