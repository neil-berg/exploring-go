package main

import (
	"fmt"
	"strings"
)

func main() {
	ch1 := make(chan string)
	// ch2 := make(chan string)

	data := []string{"Apple", "Pear", "Orange", "Kiwi", "baana", "Papaya"}
	// data2 := []string{"Bottom", "Puck", "Penix", "Kiwi", "Foop", "Paddddd"}

	go findFruits(data, ch1)
	// go findFruits(data2, ch2)

	var items []string
	for f1 := range ch1 {
		items = append(items, f1)
		// fmt.Println(f1)
	}
	fmt.Println(items)
}

func findFruits(fruits []string, ch chan<- string) {
	for _, f := range fruits {
		if strings.HasPrefix(strings.ToLower(f), "p") {
			ch <- f
		}
	}
	close(ch)
}
