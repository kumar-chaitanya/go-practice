package smi

import "testing"

func TestPerimeter(t *testing.T) {
	/* Define the testing data */
	testCases := []struct {
		name string
		shape Shape
		wantArea float64
		wantPerimeter float64
	}{
		{ name: "Rectangle", shape: Rectangle{Width: 12.0, Height: 8.0}, wantArea: 96.0, wantPerimeter: 40.0 },
		{ name: "Circle", shape: Circle{Radius: 10.0}, wantArea: 314.1592653589793, wantPerimeter: 62.83185307179586 },
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			gotArea := tc.shape.Area()
			gotPerimeter := tc.shape.Perimeter()

			if gotArea != tc.wantArea {
				t.Errorf("%#v got %g want %g", tc.shape, gotArea, tc.wantArea)
			}

			if gotPerimeter != tc.wantPerimeter {
				t.Errorf("%#v got %g want %g", tc.shape, gotPerimeter, tc.wantPerimeter)
			}
		})
	}
}
