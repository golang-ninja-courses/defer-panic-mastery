package recoverer

import (
	"runtime"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/goleak"
)

func TestMain(m *testing.M) {
	goleak.VerifyTestMain(m)
}

func TestRecoverer_Do(t *testing.T) {
	t.Run("no panic", func(t *testing.T) {
		r, s := newTestRecoverer(func(_ interface{}) {
			t.Fatal("unexpected call of recovery handler")
		})
		defer s.Wait()

		var fnCalls int
		for i := 0; i < 3; i++ {
			r.Do(func() { fnCalls++ })
		}
		assert.Equal(t, 3, fnCalls)
	})

	t.Run("panic", func(t *testing.T) {
		var handlerCalls int
		panicArg := new(runtime.TypeAssertionError)
		r, s := newTestRecoverer(func(err interface{}) {
			handlerCalls++
			assert.Equal(t, err, panicArg)
		})
		defer s.Wait()

		var fnCalls int
		for i := 0; i < 6; i++ {
			r.Do(func() {
				fnCalls++
				if i >= 3 {
					panic(panicArg)
				}
			})
		}
		assert.Equal(t, 6, fnCalls)
		assert.Equal(t, 3, handlerCalls)
	})
}

func TestRecoverer_DoWithRecoveryHandler(t *testing.T) {
	t.Run("no panic", func(t *testing.T) {
		r, s := newTestRecoverer(func(_ interface{}) {
			t.Fatal("unexpected call of default recovery handler")
		})
		defer s.Wait()

		h := func(_ interface{}) {
			t.Fatal("unexpected call of recovery handler")
		}

		var fnCalls int
		for i := 0; i < 3; i++ {
			r.DoWithRecoveryHandler(func() { fnCalls++ }, h)
		}
		assert.Equal(t, 3, fnCalls)
	})

	t.Run("panic", func(t *testing.T) {
		r, s := newTestRecoverer(func(_ interface{}) {
			t.Fatal("unexpected call of default recovery handler")
		})
		defer s.Wait()

		var handlerCalls int
		panicArg := new(runtime.TypeAssertionError)
		h := func(err interface{}) {
			handlerCalls++
			assert.Equal(t, err, panicArg)
		}

		var fnCalls int
		for i := 0; i < 6; i++ {
			r.DoWithRecoveryHandler(func() {
				fnCalls++
				if i >= 3 {
					panic(panicArg)
				}
			}, h)
		}
		assert.Equal(t, 6, fnCalls)
		assert.Equal(t, 3, handlerCalls)
	})
}

func TestRecoverer_Go(t *testing.T) {
	t.Run("no panic", func(t *testing.T) {
		r, s := newTestRecoverer(func(_ interface{}) {
			t.Fatal("unexpected call of recovery handler")
		})

		var fnCalls int
		var mu sync.Mutex
		for i := 0; i < 3; i++ {
			r.Go(func() {
				mu.Lock()
				fnCalls++
				mu.Unlock()
			})
		}
		s.Wait()
		assert.Equal(t, 3, fnCalls)
	})

	t.Run("panic", func(t *testing.T) {
		var mu sync.Mutex

		var handlerCalls int
		panicArg := new(runtime.TypeAssertionError)
		r, s := newTestRecoverer(func(err interface{}) {
			mu.Lock()
			defer mu.Unlock()
			handlerCalls++
			assert.Equal(t, err, panicArg)
		})

		var fnCalls int
		for i := 0; i < 6; i++ {
			i := i
			r.Go(func() {
				mu.Lock()
				fnCalls++
				mu.Unlock()
				if i >= 3 {
					panic(panicArg)
				}
			})
		}
		s.Wait()
		assert.Equal(t, 6, fnCalls)
		assert.Equal(t, 3, handlerCalls)
	})
}

func TestRecoverer_GoWithRecoveryHandler(t *testing.T) {
	t.Run("no panic", func(t *testing.T) {
		r, s := newTestRecoverer(func(_ interface{}) {
			t.Fatal("unexpected call of default recovery handler")
		})

		h := func(_ interface{}) {
			t.Fatal("unexpected call of recovery handler")
		}

		var fnCalls int
		var mu sync.Mutex
		for i := 0; i < 3; i++ {
			r.GoWithRecoveryHandler(func() {
				mu.Lock()
				fnCalls++
				mu.Unlock()
			}, h)
		}
		s.Wait()
		assert.Equal(t, 3, fnCalls)
	})

	t.Run("panic", func(t *testing.T) {
		var mu sync.Mutex

		r, s := newTestRecoverer(func(_ interface{}) {
			t.Fatal("unexpected call of default recovery handler")
		})

		var handlerCalls int
		panicArg := new(runtime.TypeAssertionError)
		h := func(err interface{}) {
			mu.Lock()
			defer mu.Unlock()
			handlerCalls++
			assert.Equal(t, err, panicArg)
		}

		var fnCalls int
		for i := 0; i < 6; i++ {
			i := i
			r.GoWithRecoveryHandler(func() {
				mu.Lock()
				fnCalls++
				mu.Unlock()
				if i >= 3 {
					panic(panicArg)
				}
			}, h)
		}
		s.Wait()
		assert.Equal(t, 6, fnCalls)
		assert.Equal(t, 3, handlerCalls)
	})
}

func newTestRecoverer(h RecoveryHandler) (*Recoverer, *goWaiter) {
	w := newGoWaiter()
	return NewRecoverer(h, w), w
}

var _ GoroutineStarter = (*goWaiter)(nil)

type goWaiter struct {
	wg sync.WaitGroup
}

func newGoWaiter() *goWaiter {
	return new(goWaiter)
}

func (g *goWaiter) Go(f func()) {
	g.wg.Add(1)
	go func() {
		defer g.wg.Done()
		f()
	}()
}

func (g *goWaiter) Wait() {
	g.wg.Wait()
}
