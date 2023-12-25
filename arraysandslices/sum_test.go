package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	// t.Run("Collection of 5 numbers", func(t *testing.T) {
	// 	numbers := []int{1, 2, 3, 4, 5}

	// 	got := Sum(numbers)
	// 	want := 15

	// 	if got != want {
	// 		t.Errorf("Got %d want %d given, %v", got, want, numbers)
	// 	}
	// })
	// Coverage is 100% even with this task deleted
	// -----------------------------------------------------------
	// aryan@xyz:~/learning-go/arraysandslices$ go test -cover
	// PASS
	// coverage: 100.0% of statements
	// ok      main/arraysandslices    0.002s

	t.Run("Collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("Got %d want %d given, %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {

	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Got %v want %v", got, want)
	}

}

func TestSumAllTails(t *testing.T) {
	// got := SumAllTails([]int{1, 2}, []int{0, 9})
	// want := []int{2, 9}

	// if !reflect.DeepEqual(got, want) {
	// 	t.Errorf("Got %v want %v", got, want)
	// }

	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Got %v want %v", got, want)
		}
	}


	t.Run("Make the sum of some slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}

		checkSums(t, got, want)
	})

	t.Run("Safety sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}

		checkSums(t, got, want)
	})

}
