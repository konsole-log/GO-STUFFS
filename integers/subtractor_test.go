package integers

import (
	"fmt"
	"testing"
)

func TestSubtract(t *testing.T){
	got := Subtract(5,2)
	want := 3
	if got != want{
		t.Errorf("required %d but got %d.",want, got)
	}
}

func ExampleSubtract(){
	result := Subtract(5,4)
	fmt.Println(result)
	// Output: 1
}
