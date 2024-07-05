package sync

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("incrementing the counter 3 times leaves it at 3", func(t *testing.T) {
		counter := newCounter()
		counter.Inc()
		counter.Inc()
		counter.Inc()

		assertCounter(t, counter, 3)
	})

	t.Run("it runs safely concurrently", func(t *testing.T) {
		wantedCount := 1000
		counter := newCounter()
		var wg sync.WaitGroup

		wg.Add(wantedCount)
		for i := 0; i < wantedCount; i++ {
			go func(w *sync.WaitGroup) {
				counter.Inc()
				wg.Done()
			}(&wg)
		}

		wg.Wait()
		assertCounter(t, counter, wantedCount)
	})
}

func assertCounter(t *testing.T, got *Counter, want int) {
	if got.Value() != want {
		t.Errorf("got %d, want %d", got.Value(), want)
	}
}
