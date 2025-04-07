package test

import (
	"math"
	"shape_app/shapes"
	"testing"
)

func TestCircle(t *testing.T) {
	c := shapes.NewCircle(10, 10, 5)
	if c.Center.X != 10.0 {
		t.Error("Expected 10.0, got ", c.Center.X)
	}
	if c.Center.Y != 10.0 {
		t.Error("Expected 10.0, got ", c.Center.Y)
	}
	if c.Radius != 5.0 {
		t.Error("Expected 5.0, got ", c.Radius)
	}
}

func TestCircleArea(t *testing.T) {
	c := shapes.NewCircle(10, 10, 5)
	area := c.Area()
	if math.Abs(area-78.5) >= 0.1 {
		t.Error("Expected something close to 78.5, got ", area)
	}
}
