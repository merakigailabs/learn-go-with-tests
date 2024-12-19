package structsmethodsandinterfaces

import (
	"math"
	"testing"
)

const float64EqualityThreshold = 1e-9

func TestPerimeter(t *testing.T) {

	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		got := rectangle.Perimeter()
		want := 40.0

		assertFloatEqual(t, got, want)
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{10.0}
		got := circle.Perimeter()
		want := 62.83

		assertFloatAlmostEqual(t, got, want)
	})

}

func TestTableArea(t *testing.T) {

	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{12, 6}, 72.0},
		{Circle{10}, 314.1592653589793},
		{Triangle{12, 6}, 36.0},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("got %g want %g", got, tt.want)
		}
	}
}

func TestArea(t *testing.T) {

	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	}

	checkAreaNearlyEqual := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if math.Abs(got-want) <= float64EqualityThreshold {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	}

	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		checkArea(t, rectangle, 100.00)
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{10.0}
		got := circle.Area()
		want := 314.16

		checkAreaNearlyEqual(t, circle, 100.00)
		assertFloatAlmostEqual(t, got, want)
	})

}

func assertFloatEqual(t testing.TB, got, want float64) {
	t.Helper()

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func assertFloatAlmostEqual(t testing.TB, got, want float64) {
	t.Helper()
	if math.Abs(got-want) <= float64EqualityThreshold {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}
