package sync

import (
	"sync"
	"testing"
)

func TestIncrement(t *testing.T) {
	t.Run("test counter in sequental program", func(t *testing.T) {
		counter := Counter{}
		counter.Inc()
		counter.Inc()

		want := 2
		assertCouter(t, &counter, want)
	})

	t.Run("test couter in concurrent environment", func(t *testing.T) {
		conter := Counter{}
		var wg sync.WaitGroup
		for range 1000 {
			wg.Add(1)
			go func() {
				conter.Inc()
				wg.Done()
			}()
		}
		wg.Wait()

		want := 1000
		assertCouter(t, &conter, want)
	})
}

func assertCouter(t testing.TB, c *Counter, want int) {
	t.Helper()
	got := c.Get()
	if got != want {
		t.Errorf("expected %d but got %d", want, got)
	}
}
