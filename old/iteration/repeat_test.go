package iteration

import (
	"fmt"
	"testing"
)

const character = "a"
const number = 6

func TestRepeat(t *testing.T) {
	repeated := Repeat(character, number)
	expected := "aaaaaa"

	if repeated != expected {
		t.Errorf("Expected %q but got %q ", expected, repeated)
	}
}

// func samp(str string, num int) string {
// 	str1 := ""
// 	for i := 0; i < num; i++ {
// 		str1 += str
// 	}

// 	return str1
// }

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat(character, number)
	}
}

func ExampleRepeat() {

	repeated := Repeat("a", 8)
	fmt.Println(repeated)
	// Output: aaaaaaaa

}