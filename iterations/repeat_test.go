package iterations

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 5)
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("exptected %q but got %q", expected, repeated)
	}
}

func ExampleRepeat() {
	repeated := Repeat("abc", 3)
	fmt.Println(repeated)
	// Output: abcabcabc

}
func BenchmarkRepeat(b *testing.B) {
	for b.Loop() {
		Repeat("a", 5)
	}
}
