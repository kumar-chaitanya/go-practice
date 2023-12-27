package array

import "testing"

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
}