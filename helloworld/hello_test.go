package helloworld

import (
	"fmt"
	"testing"
)

func TestHello(t *testing.T) {
	t.Run("In Spanish", func(t *testing.T) {
		got := Hello("Anush", "Spanish")
		want := "Hola Anush"
		assertCorrectMessage(t, got, want)
	})

	t.Run("In Nepali", func(t *testing.T) {
		got := Hello("Anush", "Nepali")
		want := "Namaste Anush"
		assertCorrectMessage(t, got, want)
	})
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Anush", "")
		want := "Hello Anush"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello World"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

}

func ExampleHello() {
	name := "Anush"
	fmt.Println(Hello(name, spanish))
	// Output: Hola Anush
}

func BenchmarkHello(b *testing.B) {
	for b.Loop() {
		Hello("Anush", "nepali")
	}
}
