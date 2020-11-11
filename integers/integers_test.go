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
	tests := []struct {
		a    int
		b    int
		want int
	}{
		{2, 3, 5},
		{5, 10, 15},
	}

	for _, test := range tests {
		name := fmt.Sprintf("adding %d + %d equals %d", test.a, test.b, test.want)
		t.Run(name, func(t *testing.T) {
			got := Add(test.a, test.b)
			if got != test.want {
				t.Errorf("got %d, want %d", got, test.want)
			}
		})
	}
}
