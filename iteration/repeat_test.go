package iteration

import "testing"

func TestRepeat(t *testing.T) {

	t.Run("should repeat 10 time", func(t *testing.T) {
		repeated := Repeat("a", 10)
		expected := "aaaaaaaaaa"

		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	})

	t.Run("should repeat 5 time", func(t *testing.T) {
		repeated := Repeat("a", 0)
		expected := "aaaaa"

		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	})

}

func ExampleRepeat() {
	Repeat("b", 3)
	// Output "bbb"
}

// go test -bench=.
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 10)
	}
}
