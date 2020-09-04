package main

import (
	"fmt"
	// "github.com/neil-berg/exploring-go/morestrings"
)

// Hello is a silly test function to print a name and language prefix
func Hello(language, name string) string {
	const englishPrefix = "Hello"
	const spanishPrefix = "Hola"
	const frenchPrefix = "Bonjour"

	prefix := englishPrefix
	switch language {
	case "spanish":
		prefix = spanishPrefix
	case "french":
		prefix = frenchPrefix
	}
	return prefix + ", " + name
}

func main() {
	fmt.Println(Hello("english", "Neil"))
	// fmt.Println(morestrings.ReverseRunes("!oG ,olleH"))
}
