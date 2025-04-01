package main

import (
	"bytes"
	"io"
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

func TestNewDataStorageManager(t *testing.T) {
	content := "Hello World"
	filename := setup(content)
	defer cleanup(filename)

	result, err := NewDataStorageManager(filename)
	if err != nil {
		t.Errorf("NewDataStorageManager(%s) = %v; want nil", filename, err)
	}
	if result.Info() != "DataStorageManager: My major data structure is a string" {
		t.Errorf("Info() = %v; want DataStorageManager: My major data structure is a string", result.Info())
	}
}

func TestDataStorageManagerWords(t *testing.T) {
	content := "Ch1: I hate the world as much as I love the world."
	filename := setup(content)
	defer cleanup(filename)

	result, _ := NewDataStorageManager(filename)
	expected := []string{"ch1", "i", "hate", "the", "world", "as", "much", "as", "i", "love", "the", "world"}
	if got := result.Words(); !reflect.DeepEqual(got, expected) {
		t.Errorf("Words() = %v; want %v", got, expected)
	}
}

func TestNewStopWordManager(t *testing.T) {
	result, err := NewStopWordManager()
	if err != nil {
		t.Errorf("NewStopWordManager() = %v; want nil", err)
	}
	if result.Info() != "StopWordManager: My major data structure is a map[string]struct{}" {
		t.Errorf("Info() = %v; want StopWordManager: My major data structure is a map[string]struct{}", result.Info())
	}

}

func TestStopWordManagerIsStopWord(t *testing.T) {
	result, _ := NewStopWordManager()
	tests := []struct {
		input    string
		expected bool
	}{
		{"hello", false},
		{"the", true},
		{"is", true},
		{"on", true},
	}
	for _, tt := range tests {
		if got := result.IsStopWord(tt.input); got != tt.expected {
			t.Errorf("IsStopWord(%s) = %v; want %v", tt.input, got, tt.expected)
		}
	}
}

func TestNewWordFrequencyManager(t *testing.T) {
	result, err := NewWordFrequencyManager()
	if err != nil {
		t.Errorf("NewWordFrequencyManager() = %v; want nil", err)
	}
	if result.Info() != "WordFrequencyManager: My major data structure is a map[string]int" {
		t.Errorf("Info() = %v; want WordFrequencyManager: My major data structure is a map[string]int", result.Info())
	}

}

func TestWordFrequencyManagerIncrementCount(t *testing.T) {
	result, _ := NewWordFrequencyManager()
	tests := []struct {
		word  string
		count int
	}{
		{"hello", 4},
		{"world", 10},
		{"hate", 3},
		{"like", 1},
	}
	for _, tt := range tests {
		for i := 0; i < tt.count; i++ {
			result.IncrementCount(tt.word)
		}
	}
	for _, tt := range tests {
		if got := result.wordFreqs[tt.word]; got != tt.count {
			t.Errorf("IncrementCount(%s) = %v; want %v", tt.word, got, tt.count)
		}
	}
}

func TestWordFrequencyManagerSorted(t *testing.T) {
	result, _ := NewWordFrequencyManager()
	tests := []struct {
		word  string
		count int
	}{
		{"hello", 4},
		{"world", 10},
		{"hate", 3},
		{"like", 1},
	}
	for _, tt := range tests {
		for i := 0; i < tt.count; i++ {
			result.IncrementCount(tt.word)
		}
	}

	expected := [][2]any{
		{"world", 10},
		{"hello", 4},
		{"hate", 3},
		{"like", 1},
	}
	if got := result.Sorted(); !reflect.DeepEqual(got, expected) {
		t.Errorf("Sorted() = %v; want %v", got, expected)
	}
}

func TestNewWordFrequencyController(t *testing.T) {
	content := "Ch1: I hate the world as much as I love the world🎉"
	filename := setup(content)
	defer cleanup(filename)

	result, err := NewWordFrequencyController(filename)
	if err != nil {
		t.Errorf("NewWordFrequencyController(%s) = %v; want nil", filename, err)
	}
	if result.Info() != "WordFrequencyController" {
		t.Errorf("Info() = %v; want WordFrequencyController", result.Info())
	}
}

func TestWordFrequencyControllerRun(t *testing.T) {
	content := "Ch1: I hate the world as much as I love the world🎉"
	filename := setup(content)
	defer cleanup(filename)

	// Capture stdout
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	result, _ := NewWordFrequencyController(filename)
	result.Run()

	// Restore stdout
	w.Close()
	os.Stdout = oldStdout

	// Read captured output
	var buf bytes.Buffer
	io.Copy(&buf, r)
	got := buf.String()

	want := "world - 2\nch1 - 1\nhate - 1\nlove - 1\nmuch - 1\n"
	if got != want {
		t.Errorf("Run() output = %q; want %q", got, want)
	}
}
