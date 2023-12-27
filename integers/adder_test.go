package integers

import "testing"

func TestHello(t *testing.T) {
	t.Run("add positive ints", func(t *testing.T) {
		got := Add(2, 3);
		want := 5

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("add mixed ints", func(t *testing.T) {
		got := Add(-2, 3);
		want := 1

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("add negative ints", func(t *testing.T) {
		got := Add(-2, -1);
		want := -3

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}