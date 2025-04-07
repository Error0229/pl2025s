package shapes

import (
	"math"
)

type Point struct {
	X, Y float64
}

func NewPoint(x, y float64) *Point {
	return &Point{x, y}
}

func (self Point) Distance(other Point) float64 {
	return math.Sqrt((other.X-self.X)*(other.Y-self.Y) + (other.Y-self.Y)*(other.Y-self.Y))
}

func (self *Point) Translate(dx, dy float64) {
	self.X += dx
	self.Y += dy
}

func (self Point) Area() float64 {
	return 0.0
}

func (self Point) Equal(s Shape) bool {
	other, ok := s.(*Point)
	if !ok {
		return false
	}
	return self.X == other.X && self.Y == other.Y
}
