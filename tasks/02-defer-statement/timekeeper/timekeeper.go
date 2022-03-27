package timekeeper

import "time"

type MetricsStorage interface {
	Record(operation string, d time.Duration)
}

type Timekeeper struct {
	storage MetricsStorage
}

func NewTimekeeper(s MetricsStorage) *Timekeeper {
	return &Timekeeper{storage: s}
}

func (t *Timekeeper) MeasureExecutionTime( /* Реализуй меня */ ) /* Реализуй меня */ {
	// Реализуй меня.
}
