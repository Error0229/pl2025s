// main_test.go
package main

import (
	"os"
	"reflect"
	"testing"
)

func setup(content string) string {
	// Create temporary test file
	tmpfile, err := os.CreateTemp("", "test.*.txt")
	if err != nil {
		panic(err)
	}
	if _, err := tmpfile.Write([]byte(content)); err != nil {
		panic(err)
	}
	tmpfile.Close()
	return tmpfile.Name()
}

func cleanup(filename string) {
	os.Remove(filename)
}

func TestReadFile(t *testing.T) {
	content := "Hello World"
	filename := setup(content)
	defer cleanup(filename)

	result := readFile(filename)
	if string(result) != content {
		t.Errorf("readFile(%s) = %s; want %s", filename, string(result), content)
	}
}

func TestFilterCharsAndNormalize(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"Hello123😎❤️😁😂🤣", "hello123 "},
		{"Hello, World!", "hello world "},
		{"UPPER lower", "upper lower"},
	}

	for _, tt := range tests {
		data := []rune(tt.input)
		if got := string(filterCharsAndNormalize(data)); got != tt.expected {
			t.Errorf("filterCharsAndNormalize(%s) = %s; want %s", tt.input, got, tt.expected)
		}
	}
}

func TestScan(t *testing.T) {
	data := []rune("hello world test")

	expected := []string{"hello", "world", "test"}
	result := scan(data)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("scan() got %v; want %v", result, expected)
	}
}

func TestRemoveStopWords(t *testing.T) {

	words := []string{"the", "quick", "brown", "fox", "is", "running"}
	result := removeStopWords(words)

	expected := []string{"quick", "brown", "fox", "running"}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("removeStopWords() got %v; want %v", words, expected)
	}
}

func TestFrequencies(t *testing.T) {
	words := []string{"hello", "world", "hello", "test", "world", "hello"}
	result := frequencies(words)

	// Create map for easier comparison
	freqMap := make(map[string]int)
	for word, count := range result {
		freqMap[word] = count
	}

	expected := map[string]int{
		"hello": 3,
		"world": 2,
		"test":  1,
	}

	if !reflect.DeepEqual(freqMap, expected) {
		t.Errorf("frequencies() got %v; want %v", freqMap, expected)
	}
}

func TestSortFreqs(t *testing.T) {
	wordFreqs := map[string]int{
		"hello": 3,
		"world": 2,
		"test":  4,
	}
	result := sortFreqs(wordFreqs)

	expected := [][2]interface{}{
		{"test", 4},
		{"hello", 3},
		{"world", 2},
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("sortFreqs() got %v; want %v", result, expected)
	}
}
