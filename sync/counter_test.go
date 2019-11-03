package sync

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("incrementing the counter 3 times leaves it at 3", func(t *testing.T) {
		counter := NewCounter()

		counter.Inc()
		counter.Inc()
		counter.Inc()

		assertCounter(t, counter, 3)
	})

	t.Run("the counter is safe to use concurrently", func(t *testing.T) {
		want := 1000
		counter := NewCounter()

		var wg sync.WaitGroup
		wg.Add(want)

		for i := 0; i < want; i++ {
			go func(w *sync.WaitGroup) {
				counter.Inc()
				w.Done()
			}(&wg)
		}

		wg.Wait()

		assertCounter(t, counter, want)
	})
}

func assertCounter(t *testing.T, counter *Counter, want int) {
	t.Helper()

	if counter.Value() != want {
		t.Errorf("value = %d, want %d", counter.Value(), want)
	}
}

func BenchmarkCounter(b *testing.B) {
	counter := new(Counter)
	for i := 0; i < b.N; i++ {
		counter.Inc()
	}
}
