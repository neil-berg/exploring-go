package iteration

// Repeat takes a character and repeats it N times
func Repeat(c string, n int) string {
	var repeated string
	for i := 0; i < n; i++ {
		repeated += c
	}
	return repeated
}
