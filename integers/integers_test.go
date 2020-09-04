package integers

import (
	"fmt"
	"testing"
)

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}

func TestAdd(t *testing.T) {

	t.Run("adds to positive numbers", func(t *testing.T) {
		got := Add(5, 5)
		want := 10
		if got != want {
			t.Errorf("got '%d' want '%d'", got, want)
		}
	})

	t.Run("adds two negative numbers", func(t *testing.T) {
		got := Add(-5, -5)
		want := -10
		if got != want {
			t.Errorf("got '%d' want '%d'", got, want)
		}
	})
}
