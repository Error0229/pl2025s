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
	content := "Ch1: I hate the world as much as I love the world🎉"
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

	if result.Info() != "StopWordManager" {
		t.Errorf("Info() = %v; want StopWordManager", result.Info())
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

	if result.Info() != "WordFrequencyManager" {
		t.Errorf("Info() = %v; want WordFrequencyManager", result.Info())
	}
}
