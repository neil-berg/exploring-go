package main

import (
	"fmt"
	"strings"
	// "github.com/neil-berg/exploring-go/morestrings"
)

// Hello is a silly test function to print a name and language prefix
func Hello(language, name string) string {
	m := map[string]string{
		"english": "hello",
		"spansih": "hola",
		"french":  "bonjour",
	}

	return m[strings.ToLower(language)] + ", " + name
}

func main() {
	fmt.Println(Hello("ENglish", "Neil"))
	// fmt.Println(morestrings.ReverseRunes("!oG ,olleH"))
}
