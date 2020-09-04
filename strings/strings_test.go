package strings

import "testing"

func TestStrings(t *testing.T) {
	t.Run("Counts and repeats a character", func(t *testing.T) {
		got := CountChar("California", "i")
		want := "ii"
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
