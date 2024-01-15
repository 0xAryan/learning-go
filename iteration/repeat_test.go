package iteration

import (
	"fmt"
	"testing"
)

var character string = "a"
var no int = 6

func TestRepeat(t *testing.T) {

	repeated := Repeat(character, no)
	expected := "aaaaaa"

	if repeated != expected {
		t.Errorf("Expected %q but got %q", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	fmt.Println(Repeat("A", 5))
	// Output: AAAAA
}