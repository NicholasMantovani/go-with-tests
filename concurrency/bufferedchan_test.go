package concurrency

import "testing"

func TestRunBufferedChan(t *testing.T) {
	t.Run("the channel is not full", func(t *testing.T) {
		res := runBufferedChan()
		if len(res) != 5 {
			t.Errorf("wanted 5 got %d", len(res))
		}
	})
}
