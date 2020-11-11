package main

import (
	"fmt"
)

// PairAddsToMax take a slice of ints and returns the pair of values that sum
// to the max value in the slice or [] if no pair exists.
// E.g. [-1, 3, 4, 8, 12] // [4, 8]
func PairAddsToMax(v []int) []int {
	var out []int
	max := getMax(v)

	for i, n := range v {
		for _, n2 := range v[i+1:] {
			if n+n2 == max {
				out = append(out, n, n2)
				break
			}
		}
	}
	// fmt.Println(reflect.DeepEqual(out, []int{1, 12}))
	return out
}

func getMax(v []int) int {
	var max int
	for _, n := range v {
		if n > max {
			max = n
		}
	}
	return max
}

func main() {
	vals := []int{8, 6, 1, 3, 13, -5, 12}
	res := PairAddsToMax(vals)
	fmt.Println(res)
}
