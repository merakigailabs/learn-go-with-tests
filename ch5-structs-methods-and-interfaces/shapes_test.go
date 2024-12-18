package structsmethodsandinterfaces

import (
	"math"
	"testing"
)

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

func TestArea(t *testing.T) {
	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		got := rectangle.Area()
		want := 100.00

		assertFloatEqual(t, got, want)
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{10.0}
		got := circle.Area()
		want := 314.16

		assertFloatAlmostEqual(t, got, want)
	})

}

func assertFloatEqual(t testing.TB, got, want float64) {
	t.Helper()

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

const float64EqualityThreshold = 1e-9

func assertFloatAlmostEqual(t testing.TB, got, want float64) {
	t.Helper()
	if math.Abs(got-want) <= float64EqualityThreshold {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}
