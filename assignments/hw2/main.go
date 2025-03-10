package main

import (
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
)

type TFExercise struct {
	TypeName      string
	DataStructure string
}

func (tfe *TFExercise) Info() string {
	return fmt.Sprintf("%s: My major data structure is a %s",
		tfe.TypeName,
		tfe.DataStructure)
}

type DataStorageManager struct {
	TFExercise
	data string
}

func NewDataStorageManager(pathToFile string) (*DataStorageManager, error) {
	content, err := os.ReadFile(pathToFile)
	if err != nil {
		return nil, err
	}
	reg, err := regexp.Compile(`[\W_]+`)
	cleanedData := reg.ReplaceAllString(strings.ToLower(string(content)), " ")
	if err != nil {
		return nil, err
	}
	return &DataStorageManager{
		TFExercise: TFExercise{
			TypeName:      "DataStorageManager",
			DataStructure: "string",
		},
		data: cleanedData,
	}, nil
}

func (dsm *DataStorageManager) Words() []string {
	return strings.Fields(string(dsm.data))
}

type StopWordManager struct {
	TFExercise
	stopWords map[string]struct{}
}

func NewStopWordManager() (*StopWordManager, error) {
	swm := &StopWordManager{
		TFExercise: TFExercise{
			TypeName:      "StopWordManager",
			DataStructure: "map[string]struct{}",
		},
		stopWords: make(map[string]struct{}),
	}
	stopWords, err := os.ReadFile("../stop_words.txt")
	if err != nil {
		return nil, err
	}
	words := strings.Split(string(stopWords), ",")
	for _, word := range words {
		swm.stopWords[strings.TrimSpace(word)] = struct{}{}
	}
	for ch := 'a'; ch <= 'z'; ch++ {
		swm.stopWords[string(ch)] = struct{}{}
	}
	return swm, nil
}

func (swm *StopWordManager) IsStopWord(word string) bool {
	_, ok := swm.stopWords[word]
	return ok
}

type WordFrequencyManager struct {
	TFExercise
	wordFreqs map[string]int
}

func NewWordFrequencyManager() (*WordFrequencyManager, error) {
	return &WordFrequencyManager{
		TFExercise: TFExercise{
			TypeName:      "WordFrequencyManager",
			DataStructure: "map[string]int",
		},
		wordFreqs: make(map[string]int),
	}, nil
}

func (wfm *WordFrequencyManager) IncrementCount(word string) {
	wfm.wordFreqs[word]++
}

func (wfm *WordFrequencyManager) Sorted() [][2]any {
	var freqPairs [][2]any
	for word, freq := range wfm.wordFreqs {
		freqPairs = append(freqPairs, [2]any{word, freq})
	}
	sort.Slice(freqPairs, func(i, j int) bool {
		return freqPairs[i][1].(int) > freqPairs[j][1].(int)
	})
	return freqPairs
}

type WordFrequencyController struct {
	storageManager  *DataStorageManager
	stopWordManager *StopWordManager
	wordFreqManager *WordFrequencyManager
}

func NewWordFrequencyController(pathToFile string) (*WordFrequencyController, error) {
	storageManager, err := NewDataStorageManager(pathToFile)
	if err != nil {
		return nil, err
	}
	stopWordManager, err := NewStopWordManager()
	if err != nil {
		return nil, err
	}
	wordFreqManager, err := NewWordFrequencyManager()
	if err != nil {
		return nil, err
	}
	return &WordFrequencyController{
		storageManager:  storageManager,
		stopWordManager: stopWordManager,
		wordFreqManager: wordFreqManager,
	}, nil
}

func (wfc *WordFrequencyController) Run() {
	for _, word := range wfc.storageManager.Words() {
		if !wfc.stopWordManager.IsStopWord(word) {
			wfc.wordFreqManager.IncrementCount(word)
		}
	}
	wordFreqs := wfc.wordFreqManager.Sorted()
	for _, pair := range wordFreqs[:25] {
		fmt.Printf("%s  -  %d\n", pair[0].(string), pair[1].(int))
	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: main.go <path_to_file>")
		return
	}
	wfc, err := NewWordFrequencyController(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	wfc.Run()
}
