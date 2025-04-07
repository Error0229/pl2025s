package test

import (
	"shape_app/shapes"
	"testing"
)

func TestPoint(t *testing.T) {
	p1 := shapes.NewPoint(1.0, 2.0)
	p2 := shapes.NewPoint(1.0, 2.0)
	if !p1.Equal(*p2) {
		t.Error("Expected true, got ", p1.Equal(*p2))
	}

}

func TestDistance(t *testing.T) {
	p1 := shapes.NewPoint(1.0, 2.0)
	p2 := shapes.NewPoint(4.0, 6.0)
	if p1.Distance(*p2) != 5.0 {
		t.Error("Expected 5.0, got ", p1.Distance(*p2))
	}
}

func TestTranslate(t *testing.T) {
	p := shapes.NewPoint(1.0, 2.0)
	p.Translate(2.0, 3.0)
	expected := shapes.NewPoint(3.0, 5.0)
	if !p.Equal(*expected) {
		t.Error("Expected true, got ", p.Equal(*expected))
	}
}

func TestArea(t *testing.T) {
	p := shapes.NewPoint(1.0, 2.0)
	if p.Area() != 0.0 {
		t.Error("Expected 0.0, got ", p.Area())
	}
}
