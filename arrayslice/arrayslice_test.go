package arrayslice

import (
	"reflect"
	"testing"
)

func TestSumOneArray(t *testing.T) {
	t.Run("sums items in one array", func(t *testing.T) {
		got := SumOneArray([]int{1, 2, 3})
		want := 6
		if got != want {
			t.Errorf("want %v, got %v", want, got)
		}
	})

	t.Run("returns 0 if passed an empty array", func(t *testing.T) {
		got := SumOneArray([]int{})
		want := 0
		if want != got {
			t.Errorf("want %v, got %v", want, got)
		}
	})
}

func TestSumNArrays(t *testing.T) {
	t.Run("sums items in N arrays", func(t *testing.T) {
		got := SumNArrays([]int{1, 2, 3}, []int{4, 5, 6}, []int{})
		want := []int{6, 15, 0}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("want %v, got %v", want, got)
		}
	})
}
