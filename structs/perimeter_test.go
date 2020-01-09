package structs

import "testing"

func TestPerimeter(t *testing.T) {
	e := 40.0
	a := Perimeter(Rectangle{10.0, 10.0})

	if e != a {
		t.Errorf("expected: %.2f actual: %.2f", e, a)
	}
}

func TestArea(t *testing.T) {

	areaTests := []struct {
		name     string
		shape    Shape
		expected float64
	}{
		{name: "Rectangle", shape: Rectangle{12, 6}, expected: 72.0},
		{name: "Circle", shape: Circle{10}, expected: 314.1592653589793},
		{name: "Triangle", shape: Triangle{12, 6}, expected: 36.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.shape.area()
			if actual != tt.expected {
				t.Errorf("%#v expected: %g actual: %g", tt.shape, tt.expected, actual)
			}
		})

	}
}
