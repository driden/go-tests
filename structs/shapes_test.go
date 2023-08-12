package main

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{width: 10.0, height: 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f, wanted %2.f.", got, want)
	}
}

func TestArea(t *testing.T) {

	areaTests := []struct {
		shape   Shape
		hasArea float64
	}{
		{shape: Rectangle{width: 12.0, height: 6.0}, hasArea: 72.0},
		{shape: Circle{radius: 10}, hasArea: 314.1592653589793},
		{shape: Triangle{base: 12, height: 6}, hasArea: 36.0},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.hasArea {
			t.Errorf("%#v got %g, want %g", tt.shape, got, tt.hasArea)
		}
	}
}
