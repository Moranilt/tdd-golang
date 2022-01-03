package arrays

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2, 3}, []int{7, 9, 0})
	want := []int{6, 16}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSumAllTails(t *testing.T) {
	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}
	t.Run("make the sums of some slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 3}, []int{0, 9, 1})
		want := []int{5, 10}

		checkSums(t, got, want)
	})
	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{0, 9, 1})
		want := []int{0, 10}

		checkSums(t, got, want)
	})
}

func ExampleSum() {
	sum := Sum([]int{1, 2, 3, 4})
	fmt.Print(sum)
	// Output: 10
}

func ExampleSumAll() {
	sums := SumAll([]int{3, 4, 5}, []int{5, 6, 7})
	fmt.Print(sums)
	// Output: [12 18]
}

func ExampleSumAllTails() {
	sumsTails := SumAllTails([]int{3, 4, 5}, []int{5, 6, 7})
	fmt.Print(sumsTails)
	// Output: [9 13]
}

func BenchmarkSumAll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumAll([]int{78, 231, 5, 76, 324, 43, 4, 62, 2412, 1, 786, 8979067, 0, 2139984, 24, 213}, []int{873721, 321, 55, 7, 98, 0, 7345, 28423, 8324, 238, 34, 54})
	}
}

func BenchmarkSumAllTails(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumAllTails([]int{78, 231, 5, 76, 324, 43, 4, 62, 2412, 1, 786, 8979067, 0, 2139984, 24, 213}, []int{873721, 321, 55, 7, 98, 0, 7345, 28423, 8324, 238, 34, 54})
	}
}
