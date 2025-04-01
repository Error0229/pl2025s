package main

type TextDocument struct{}

func (td *TextDocument) PrepareData() string {
	return "This is the raw text data"
}

func (td *TextDocument) FormatContent(data string) string {
	return "Formatted Text: " + data
}

func (td *TextDocument) Save(formattedContent string) string {
	return "Saving text document: " + formattedContent
}
