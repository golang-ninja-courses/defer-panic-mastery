package processor

import (
	"math/rand"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func TestProcessFiles(t *testing.T) {
	const firstFileLine = "test"

	fToSkip, err := os.CreateTemp("", "TestProcessFiles_*_skip")
	require.NoError(t, err)
	require.NoError(t, fToSkip.Close())

	f, err := os.CreateTemp("", "TestProcessFiles_*")
	require.NoError(t, err)
	defer func() { require.NoError(t, f.Close()) }()

	_, err = f.WriteString(firstFileLine)
	require.NoError(t, err)

	var (
		n         = 100_000
		skipEvery = rand.Intn(1000)
		skip      = n / skipEvery
	)
	paths := strings.Split(strings.TrimRight(strings.Repeat(f.Name()+"|", n), "|"), "|")
	for i := 0; i < n; i += skipEvery {
		paths[i] = fToSkip.Name()
	}

	res, err := ProcessFiles(paths)
	require.NoError(t, err)

	assert.InDelta(t, len(res), n-skip, 1)
	for i, r := range res {
		assert.Equal(t, firstFileLine, r, "index=%d", i)
	}
}

func TestProcessFiles_SkipAll(t *testing.T) {
	fToSkip, err := os.CreateTemp("", "TestProcessFiles_*_skip")
	require.NoError(t, err)
	require.NoError(t, fToSkip.Close())

	const n = 100_000
	paths := strings.Split(strings.TrimRight(strings.Repeat(fToSkip.Name()+"|", n), "|"), "|")

	res, err := ProcessFiles(paths)
	require.NoError(t, err)
	assert.Empty(t, res)
}
