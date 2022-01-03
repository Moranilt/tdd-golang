package structs

import (
	"fmt"
	"testing"
)

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10, 10}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{name: "Rectangle", shape: Rectangle{12, 6}, want: 72.0},
		{name: "Circle", shape: Circle{10}, want: 314.1592653589793},
		{name: "Triangle", shape: Triangle{12, 6}, want: 36.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.want {
				t.Errorf("%#v got %g want %g", tt.shape, got, tt.want)
			}
		})
	}
}

func ExampleTriangle_Area() {
	triangle := Triangle{Base: 12, Height: 6}
	fmt.Print(triangle.Area())
	// Output: 36
}

func ExampleRectangle_Area() {
	rectangle := Rectangle{12, 6}
	fmt.Print(rectangle.Area())
	// Output: 72
}

func ExampleCircle_Area() {
	circle := Circle{10}
	fmt.Print(circle.Area())
	// Output: 314.1592653589793
}
func BenchmarkPerimeter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Perimeter(Rectangle{873, 2356})
	}
}

func BenchmarkRectangleArea(b *testing.B) {
	rectangle := Rectangle{873, 2356}
	for i := 0; i < b.N; i++ {
		rectangle.Area()
	}
}
