package structs_methods_and_interfaces

import "testing"

func TestPerimeter(t *testing.T) {
	r := Rectangle{10.0, 10.0}
	want := 40.0

	if got := r.Perimeter(); got != want {
		t.Errorf("%#v.Perimeter() = %g, want %g", r, got, want)
	}
}

func TestArea(t *testing.T) {
	tests := []struct {
		name string
		args Shape
		want float64
	}{
		{"rectangles",
			Rectangle{12.0, 6.0},
			72.0},

		{"circles",
			Circle{10.0},
			314.1592653589793},

		{"triangles",
			Triangle{12, 6}, 36.0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.Area(); got != tt.want {
				t.Errorf("%#v.Area() = %g, want %g", tt.args, got, tt.want)
			}
		})
	}
}
