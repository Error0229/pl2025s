package test

import (
	"math"
	"shape_app/shapes"
	"testing"
)

func TestSumOfArea(t *testing.T) {
	p1 := shapes.NewPoint(3.0, 4.0)
	c1 := shapes.NewCircle(7.0, 8.0, 5.0)
	c2 := shapes.NewCircle(7.0, 10.0, 5.0)

	sumOfArea := shapes.SumOfShapes([]shapes.Shape{p1, c1, c2})
	if math.Abs(sumOfArea-157) > 0.1 {
		t.Error("Expected something close to 157, got ", sumOfArea)
	}
}
