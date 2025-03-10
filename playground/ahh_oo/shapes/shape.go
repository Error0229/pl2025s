package shapes

type Shape interface {
	Area() float64
	Equal(other Shape) bool
}
