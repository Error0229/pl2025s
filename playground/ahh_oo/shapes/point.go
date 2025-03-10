package shapes

import (
	"math"
)

type Point struct {
	x, y float64
}

func NewPoint(x, y float64) Point {
	return Point{x, y}
}

func (self Point) Distance(other Point) float64 {
	return math.Sqrt((other.x-self.x)*(other.y-self.y) + (other.y-self.y)*(other.y-self.y))
}

func (self *Point) Translate(dx, dy float64) {
	self.x += dx
	self.y += dy
}

func (self Point) Area() float64 {
	return 0.0
}

func (self Point) Equal(other Point) bool {
	return self.x == other.x && self.y == other.y
}
