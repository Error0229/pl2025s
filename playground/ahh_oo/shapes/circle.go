package shapes

import "math"

type Circle struct {
	Center *Point
	Radius float64
}

func NewCircle(x float64, y float64, r float64) *Circle {
	return &Circle{
		Center: NewPoint(x, y),
		Radius: r,
	}
}

func (c *Circle) Area() float64 {
	return c.Radius * c.Radius * math.Pi
}

func (c1 *Circle) Equal(s Shape) bool {
	other, ok := s.(*Circle)
	if !ok {
		return false
	}
	return c1.Center.Equal(*other.Center) && c1.Radius == other.Radius
}
