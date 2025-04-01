package main

type BaseGenerator struct{ DocumentGenerator }

func (g *BaseGenerator) Generate() string {
	data := g.PrepareData()
	formatted := g.FormatContent(data)
	return g.Save(formatted)
}
