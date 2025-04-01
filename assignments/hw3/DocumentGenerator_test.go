package main

import (
	"testing"
)

func TestHTMLDocumentGenerate(t *testing.T) {
	hd := &HTMLDocument{}
	bg := &BaseGenerator{DocumentGenerator: hd}
	result := bg.Generate()

	expected := "Saving HTML document: <div><html><body>This is raw HTML data.</body></html></div>"
	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}

func TestTextDocumentGenerate(t *testing.T) {
	td := &TextDocument{}
	bg := &BaseGenerator{DocumentGenerator: td}
	result := bg.Generate()

	expected := "Saving text document: Formatted Text: This is the raw text data"
	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}
