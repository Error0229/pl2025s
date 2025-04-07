package shapes

type Shape interface {
	Area() float64
	Equal(other Shape) bool
	// DrawTemplateImpl
}

func SumOfShapes(shapes []Shape) float64 {
	sum := 0.0
	for _, shape := range shapes {
		sum += shape.Area()
	}
	return sum
}

type DrawTemplateImpl interface {
	CreateBackground()
	PreparePen()
	Draw()
}
