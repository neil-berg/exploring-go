package arrayslice

// SumOneArray (really a slice) sums the items in one array and returns the sum
func SumOneArray(arr []int) int {
	var sum int = 0
	for _, number := range arr {
		sum += number
	}
	return sum
}

// SumNArrays sums the items in each array (actually slices) and return a slice of the sums
func SumNArrays(arrs ...[]int) []int {
	nArrs := len(arrs)
	var sums = make([]int, nArrs)
	for i, numbers := range arrs {
		sums[i] = SumOneArray(numbers)
	}
	return sums
}
