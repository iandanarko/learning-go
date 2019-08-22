package sync

import (
	"testing"
	"sync"
)
// A WaitGroup waits for a collection of goroutines to finish.
// The main goroutine calls Add to set the number of goroutines to wait for. 
// Then each of the goroutines runs and calls Done when finished.
// At the same time, Wait can be used to block until all goroutines have finished.
func TestCounter(t *testing.T) {
	t.Run("incrementing the counter 3 times leaves it at 3", func(t *testing.T) {
    counter := Counter{}
    counter.Inc()
    counter.Inc()
    counter.Inc()

    assertCounter(t, &counter, 3)
	})

	t.Run("it runs safely concurrently", func(t *testing.T) {
    wantedCount := 1000
    counter := Counter{}

    var wg sync.WaitGroup
    wg.Add(wantedCount)

    for i:=0; i<wantedCount; i++ {
        go func(w *sync.WaitGroup) {
            counter.Inc()
            w.Done()
        }(&wg)
    }
    wg.Wait()
		// A Mutex must not be copied after first use.
		// counter punya atrribute mutex jdi hrs pass by ref
    assertCounter(t, &counter, wantedCount)
	})
}
// func NewCounter() *Counter {
// 	return &Counter{}
// }
// There's only one way of passing arguments in Go and that is by value
// When you pass a pointer as an argument, what happens under the hood is that a copy of that pointer is created and passed to the underlying function
func assertCounter(t *testing.T, got *Counter, want int)  {
	t.Helper()
	if got.Value() != want {
			t.Errorf("got %d, want %d", got.Value(), want)
	}
}