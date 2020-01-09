package arrays

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("expected: %d actual: %d given: %v", want, got, numbers)
		}
	})
}

func assertSlicedSums(t *testing.T, a, e []int) {
	t.Helper()

	if !reflect.DeepEqual(a, e) {
		t.Errorf("expected: %d actual: %d", e, a)
	}
}

func TestSumAll(t *testing.T) {
	a := SumAll([]int{1, 2}, []int{0, 9})
	e := []int{3, 9}

	assertSlicedSums(t, a, e)
}

func TestSumAllTails(t *testing.T) {

	t.Run("make the sum of some slices", func(t *testing.T) {
		a := SumAllTails([]int{1, 2}, []int{0, 9})
		e := []int{2, 9}

		assertSlicedSums(t, a, e)
	})

	t.Run("safely some empty slices", func(t *testing.T) {
		a := SumAllTails([]int{}, []int{0, 9})
		e := []int{0, 9}

		assertSlicedSums(t, a, e)
	})

}
