package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {

	t.Run("should be 4", func(t *testing.T) {
		sum := Add(2, 2)
		expected := 4

		if sum != expected {
			t.Errorf("expected '%d' but got '%d'", expected, sum)
		}
	})

	t.Run("should be 6", func(t *testing.T) {
		sum := Add(2, 4)
		expected := 6

		if sum != expected {
			t.Errorf("expected '%d' but got '%d'", expected, sum)
		}
	})

}

// This will add an example to the godoc
func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
