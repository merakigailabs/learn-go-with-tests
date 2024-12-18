package main

import "testing"

func TestHello(t *testing.T) {

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Amine")
		want := "Hello, Amine"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello to world when name is empty", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World!"

		assertCorrectMessage(t, got, want)
	})

}

//	refactoring assert message function, testing.TB is an interface that *testing.T and * testing.B both
//
// satisfy, so you can call helper functions from a test or a benchmark.
func assertCorrectMessage(t testing.TB, got, want string) {
	// This is needed to tell the test suite that this method is a helper.
	t.Helper()
	if got != want {
		t.Errorf("git %q want %q", got, want)
	}
}
