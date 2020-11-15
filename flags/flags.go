package main

import (
	"flag"
	"fmt"
)

func main() {
	wordPtr := flag.String("word", "foo", "a string")

	var svar string
	flag.StringVar(&svar, "svar", "bar", "Enter in a string var")

	flag.Parse()
	fmt.Println("The word is", *wordPtr)
	fmt.Println("The svar is", svar)
}
