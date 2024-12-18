package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 5)
	expected := "aaaaa"

	assertCorrectMessage(t, repeated, expected)
}

func assertCorrectMessage(t testing.TB, got, want string) {
	if got != want {
		t.Errorf("expected %q but got %q", got, want)
	}
}

func ExampleRepeat() {
	res := Repeat("a", 5)
	fmt.Println(res)
	// Output: aaaaa
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}
