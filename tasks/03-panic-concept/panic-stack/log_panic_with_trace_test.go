package panichelpers

import (
	"bytes"
	"errors"
	"fmt"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLogPanicWithTrace_Example(t *testing.T) {
	defer LogPanicWithTrace(t)
	panic("sky is falling")
}

func TestLogPanicWithTrace_NoPanic(t *testing.T) {
	b := newBufLogger()

	defer func() {
		assert.Empty(t, b.String(), "no panic but something was logged")
	}()
	defer LogPanicWithTrace(&b)
}

func TestLogPanicWithTrace_Smoke(t *testing.T) {
	b := newBufLogger()

	defer func() {
		s, err := removeFilePathsFromStackTrace(b.Bytes())
		require.NoError(t, err)
		assert.True(t, expectedStackWithoutFilesRe.Match(s), "received:\n%v", string(s))
	}()

	defer LogPanicWithTrace(&b)
	aaa()
}

func aaa() {
	bbb()
}

func bbb() {
	ccc()
}

func ccc() {
	panic("sky is falling")
}

var expectedStackWithoutFilesRe = regexp.MustCompile(`sky is falling
goroutine \d{1,3} \[running]:
runtime/debug.Stack\(\)
github\.com/golang-ninja-courses/defer-panic-mastery/tasks/03-panic-concept/panic-stack\.LogPanicWithTrace\({0x[0-9a-f]{4,16}\??, 0x[0-9a-f]{4,16}\??}\)
panic\({0x[0-9a-f]{4,16}\??, 0x[0-9a-f]{4,16}\??}\)
github\.com/golang-ninja-courses/defer-panic-mastery/tasks/03-panic-concept/panic-stack\.ccc\((\.{3})?\)
github\.com/golang-ninja-courses/defer-panic-mastery/tasks/03-panic-concept/panic-stack\.bbb\((\.{3})?\)
github\.com/golang-ninja-courses/defer-panic-mastery/tasks/03-panic-concept/panic-stack\.aaa\((\.{3})?\)
github\.com/golang-ninja-courses/defer-panic-mastery/tasks/03-panic-concept/panic-stack\.TestLogPanicWithTrace_Smoke\(0x[0-9a-f]{4,16}\??\)
testing\.tRunner\(0x[0-9a-f]{4,16}\??, 0x[0-9a-f]{4,16}\??\)
created by testing\.\(\*T\)\.Run`)

func removeFilePathsFromStackTrace(s []byte) ([]byte, error) {
	var res bytes.Buffer

	lines := bytes.Split(s, []byte("\n"))
	if len(lines) < 5 {
		return nil, errors.New("too few frames")
	}

	res.Write(lines[0]) // panic value.
	res.WriteRune('\n')
	res.Write(lines[1])
	res.WriteRune('\n') // Header "goroutine 19 [running]:".

	for i := 2; i < len(lines); i += 2 {
		res.Write(lines[i])
		res.WriteRune('\n')
	}

	return res.Bytes(), nil
}

type bufLogger struct {
	*bytes.Buffer
}

func newBufLogger() bufLogger {
	return bufLogger{bytes.NewBuffer(nil)}
}

func (b bufLogger) Logf(f string, args ...any) {
	_, _ = b.WriteString(fmt.Sprintf(f, args...))
}
