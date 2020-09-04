package strings

import "strings"

// CountChar counts the number of chars in a string
// and returns a repeated string of those chars N times
func CountChar(s, char string) string {
	count := strings.Count(s, char)
	repeat := strings.Repeat(char, count)
	return repeat
}
