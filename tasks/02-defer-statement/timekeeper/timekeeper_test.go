package timekeeper

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func TestTimekeeper(t *testing.T) {
	s := newInMemoryStorage()
	tk := NewTimekeeper(s)

	fastOperation := func() {
		defer tk.MeasureExecutionTime("fastOperation")()
		time.Sleep(10 * time.Millisecond)
	}

	mediumOperation := func() {
		defer tk.MeasureExecutionTime("mediumOperation")()
		time.Sleep(500 * time.Millisecond)
	}

	slowOperation := func() {
		defer tk.MeasureExecutionTime("slowOperation")()
		time.Sleep(time.Second)
	}

	randomDuration := time.Duration(rand.Intn(1001)+100) * time.Millisecond // [100, 1000] ms
	randomOperation := func() {
		defer tk.MeasureExecutionTime("randomOperation")()
		time.Sleep(randomDuration)
	}

	fastOperation()
	mediumOperation()
	slowOperation()
	randomOperation()

	assertOperationDuration(t, s, "fastOperation", 10*time.Millisecond)
	assertOperationDuration(t, s, "mediumOperation", 500*time.Millisecond)
	assertOperationDuration(t, s, "slowOperation", time.Second)
	assertOperationDuration(t, s, "randomOperation", randomDuration)
}

func assertOperationDuration(t *testing.T, s inMemoryStorage, op string, expected time.Duration) {
	t.Helper()
	require.Contains(t, s, op, "record for operation %q is not found", op)
	assert.Equal(t, expected, s[op].Round(expected), "invalid duration of %q", op)
}

type inMemoryStorage map[string]time.Duration

func newInMemoryStorage() inMemoryStorage {
	return make(map[string]time.Duration)
}

func (s inMemoryStorage) Record(operation string, d time.Duration) {
	s[operation] = d
}
