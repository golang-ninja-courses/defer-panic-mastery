package recoverer

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDummyRecoverer_Do(t *testing.T) {
	t.Run("no panic", func(t *testing.T) {
		r := NewDummyRecoverer()

		var fnCalls int
		for i := 0; i < 3; i++ {
			r.Do(func() { fnCalls++ })
		}
		assert.Equal(t, 3, fnCalls)
	})

	t.Run("panic", func(t *testing.T) {
		r := NewDummyRecoverer()

		var fnCalls int
		assert.Panics(t, func() {
			for i := 0; i < 6; i++ {
				r.Do(func() {
					fnCalls++
					if i >= 3 {
						panic("sky is falling")
					}
				})
			}
		})
		assert.Equal(t, 3+1, fnCalls)
	})
}

func TestDummyRecoverer_DoWithRecoveryHandler(t *testing.T) {
	t.Run("no panic", func(t *testing.T) {
		r := NewDummyRecoverer()

		h := func(_ any) {
			t.Fatal("unexpected call of recovery handler")
		}

		var fnCalls int
		for i := 0; i < 3; i++ {
			r.DoWithRecoveryHandler(func() { fnCalls++ }, h)
		}
		assert.Equal(t, 3, fnCalls)
	})

	t.Run("panic", func(t *testing.T) {
		r := NewDummyRecoverer()

		h := func(_ any) {
			t.Fatal("unexpected call of recovery handler")
		}

		var fnCalls int
		assert.Panics(t, func() {
			for i := 0; i < 6; i++ {
				r.DoWithRecoveryHandler(func() {
					fnCalls++
					if i >= 3 {
						panic("sky is falling")
					}
				}, h)
			}
		})
		assert.Equal(t, 3+1, fnCalls)
	})
}

func TestDummyRecoverer_Go(t *testing.T) {
	t.Run("no panic", func(t *testing.T) {
		r := NewDummyRecoverer()

		var fnCalls int
		var mu sync.Mutex

		var wg sync.WaitGroup
		wg.Add(3)

		for i := 0; i < 3; i++ {
			r.Go(func() {
				defer wg.Done()

				mu.Lock()
				fnCalls++
				mu.Unlock()
			})
		}
		wg.Wait()
		assert.Equal(t, 3, fnCalls)
	})

	t.Run("panic", func(t *testing.T) {
		// Can be tested, but it will be nasty code.
	})
}

func TestDummyRecoverer_GoWithRecoveryHandler(t *testing.T) {
	t.Run("no panic", func(t *testing.T) {
		r := NewDummyRecoverer()

		h := func(_ any) {
			t.Fatal("unexpected call of recovery handler")
		}

		var wg sync.WaitGroup
		wg.Add(3)

		var fnCalls int
		var mu sync.Mutex
		for i := 0; i < 3; i++ {
			r.GoWithRecoveryHandler(func() {
				defer wg.Done()

				mu.Lock()
				fnCalls++
				mu.Unlock()
			}, h)
		}
		wg.Wait()
		assert.Equal(t, 3, fnCalls)
	})

	t.Run("panic", func(t *testing.T) {
		// Can be tested, but it will be nasty code.
	})
}

func TestDummyRecoverer_GoMethodsConcurrency(t *testing.T) {
	r := NewDummyRecoverer()

	ch1, ch2 := make(chan int), make(chan int)
	val := 42

	r.Go(func() { ch1 <- val })
	r.GoWithRecoveryHandler(func() { ch2 <- <-ch1 }, func(_ any) {})

	received := <-ch2
	assert.Equal(t, val, received)
}
