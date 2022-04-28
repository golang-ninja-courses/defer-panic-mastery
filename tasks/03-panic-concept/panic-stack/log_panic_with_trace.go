package panichelpers

type Logger interface {
	Logf(format string, args ...any)
}

func LogPanicWithTrace(l Logger) {
	// Реализуй меня.
}
