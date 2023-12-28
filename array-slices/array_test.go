package array

import "testing"
import "reflect"

func TestSum(t *testing.T) {
	t.Run("sum of integer array", func(t *testing.T) {
		numbers := [...]int{1, 2, 3, 4, 5, 6}
		got := SumArray(numbers)
		want := 21
	
		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})

	t.Run("sum of integer slice", func(t *testing.T) {
		numbers := []int{1, 2, 4}
		got := SumSlice(numbers)
		want := 7

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})

	t.Run("sum of all slices", func(t * testing.T) {
		numbers := [][]int{{1, 5}, {0, 9}, {-1, -3}}
		got := SumAll(numbers)
		want := []int{6, 9, -4}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
}