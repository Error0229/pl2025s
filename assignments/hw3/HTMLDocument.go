package main

type HTMLDocument struct{}

func (hd *HTMLDocument) PrepareData() string {
	return "<html><body>This is raw HTML data.</body></html>"
}

func (hd *HTMLDocument) FormatContent(data string) string {
	return "<div>" + data + "</div>"
}

func (hd *HTMLDocument) Save(formattedContent string) string {
	return "Saving HTML document: " + formattedContent
}
