package iteration

import "testing"

func TestIteration(t *testing.T) {
	t.Run("Repeats a character N times", func(t *testing.T) {
		got := Repeat("a", 3)
		want := "aaa"
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
