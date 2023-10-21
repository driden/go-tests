package sync

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("Incrementing counter 3 times leaves it at 3", func(t *testing.T) {
		counter := Counter{}
		counter.Inc()
		counter.Inc()
		counter.Inc()

		assertCounter(t, &counter, 3)
	})

	t.Run("It runs safe concurrently", func(t *testing.T) {
		wantedCount := 1000
		counter := Counter{}

		var wg sync.WaitGroup
		wg.Add(wantedCount)

		for i := 0; i < wantedCount; i++ {
			go func() {
				counter.Inc()
				wg.Done()
			}()
		}

		wg.Wait()

		assertCounter(t, &counter, wantedCount)
	})
}

func assertCounter(t testing.TB, got *Counter, expected int) {
	t.Helper()
	if got.Value() != expected {
		t.Errorf("got %d, wanted %d", got.Value(), expected)
	}
}
