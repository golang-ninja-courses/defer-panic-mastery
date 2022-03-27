package timekeeper

import (
	"fmt"
	"time"
)

func ExampleTimekeeper() {
	t := NewTimekeeper(newLoggingStorage(time.Second))

	defer t.MeasureExecutionTime("ExampleTimekeeper")()
	time.Sleep(time.Second)

	// Output:
	// ExampleTimekeeper: elapsed 1s
}

type loggingStorage struct {
	r time.Duration
}

func newLoggingStorage(rounding time.Duration) loggingStorage {
	return loggingStorage{r: rounding}
}

func (s loggingStorage) Record(operation string, d time.Duration) {
	fmt.Printf("%v: elapsed %v", operation, d.Round(s.r))
}
